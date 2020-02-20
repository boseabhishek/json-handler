# json-handler
merges jsons into one (WIP) - some constraints could be improvised!

### How to use:

#### Step 1:
> install Go and set GOPATH etc 
(Just one time pain!)

#### Step 2:
In the example, [jobs](jobs/) is the name of the directory which would have all the sub-directories(11, 12, 13 etc. - DOES NOT care about dir names BUT cares about file name TestStatus.json).

[jobs](jobs/) also the file job-summary.json which is needed and strict on name.

> once you clone this project, you could replace the jobs directory with the one you want and execute step 3

#### Step 3:

> Usage: go run main.go -dir _directory_

This will hopefully produce a file called final.json.



### ToDo (for Author)

Package this Go code into an app so that no Go installation is needed at client.


