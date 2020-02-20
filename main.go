package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// SubJob struct defined
type SubJob struct {
	AirFlowDagID       string `json:"AirFlowDagId"`
	AirFlowDagRunID    string `json:"AirFlowDagRunId"`
	ApplicationName    string `json:"ApplicationName"`
	ApplicationVersion string `json:"ApplicationVersion"`
	DigestUsed         string `json:"DigestUsed"`
	Elapsed            string `json:"Elapsed"`
	EndTime            string `json:"EndTime"`
	ReturnCode         string `json:"ReturnCode"`
	SimBranch          string `json:"SimBranch"`
	SimBuildSet        string `json:"SimBuildSet"`
	SimJenkinsBuildNo  string `json:"SimJenkinsBuildNo"`
	StartTime          string `json:"StartTime"`
	TestName           string `json:"TestName"`
	TestNumber         string `json:"TestNumber"`
	WorkingDirectory   string `json:"WorkingDirectory"`
}

// JobSummary struct defined
type JobSummary struct {
	OverallStartTime      string   `json:"OverallStartTime"`
	OverallEndTime        string   `json:"OverallEndTime"`
	OverallElapsedSeconds string   `json:"OverallElapsedSeconds"`
	OverallReturnCode     string   `json:"OverallReturnCode"`
	WorkingDirectory      string   `json:"WorkingDirectory"`
	ContainerDigest       string   `json:"ContainerDigest"`
	SimJenkinsBuildNo     string   `json:"SimJenkinsBuildNo"`
	SimBranch             string   `json:"SimBranch"`
	SimBuildSet           string   `json:"SimBuildSet"`
	AirFlowDagID          string   `json:"AirFlowDagId"`
	AirFlowDagRunID       string   `json:"AirFlowDagRunId"`
	ApplicationName       string   `json:"ApplicationName"`
	ApplicationVersion    string   `json:"ApplicationVersion"`
	SubJobs               []SubJob `json:"SubJobs"`
}

func main() {

	var subjobs []SubJob
	var subjob SubJob

	var jobsummary JobSummary

	mainDir := "jobs"

	//for job-summary
	byteValue, err := readFile(mainDir + "/" + "job-summary.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(byteValue, &jobsummary)

	//for all subdirectories
	subDirNames, err := getSubDirNames(mainDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, subDirName := range subDirNames {

		absPath := mainDir + "/" + subDirName

		fileName, err := file(absPath)
		if err != nil {
			fmt.Println(err)
		}

		byteValue, _ := readFile(fileName)

		json.Unmarshal(byteValue, &subjob)

		subjobs = append(subjobs, subjob)
	}

	jobsummary.SubJobs = subjobs

	// Convert structs to JSON.
	data, err := json.Marshal(jobsummary)
	if err != nil {
		log.Fatal(err)
	}

	writeToFile(data)
}

func getSubDirNames(mainDirPath string) ([]string, error) {
	var dirNames []string

	dirs, err := ioutil.ReadDir(mainDirPath)
	if err != nil {
		return dirNames, fmt.Errorf("error reading directory name %s provided", mainDirPath)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			dirNames = append(dirNames, dir.Name())
		}
	}

	return dirNames, nil
}

func file(dir string) (string, error) {

	var files []string

	fileInfo, err := ioutil.ReadDir(dir)

	if err != nil {
		return "", fmt.Errorf("error reading subdir name %s provided", dir)
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	if len(files) != 1 {
		return "", fmt.Errorf("contains more than 1 file inside %s", dir)
	}

	if files[0] != "users.json" {
		return "", fmt.Errorf("users.json file not found or incorrect name of file inside %s", dir)
	}

	return files[0], nil
}

func readFile(fileName string) ([]byte, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		return []byte{}, fmt.Errorf("error opening file %s", fileName)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading json to byte array from file %s", jsonFile)
	}

	return byteValue, nil
}

func writeToFile(byte []byte) {

	fileName := "final.json"

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file %#v", err)
		return
	}
	l, err := f.WriteString(string(byte))
	if err != nil {
		fmt.Printf("error writing to file %#v", err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
