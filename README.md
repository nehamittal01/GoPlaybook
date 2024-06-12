# GoPlayBook
You need to have Golang installed

The project's goal is to generate tweets from Reviews.json and Movies.json, two input JSON files.

## Getting Started
Two JSON files in the project include several test scenarios.

#### Logic Implemented
1. If a review, score, or title is blank, it won't be considered for tweet.
2. Ratings are computed using the score input, with 0-10 bracket as 1/2 rating, 10-20 as *, 20-30 as *1/2 and so on.
3. Trimmed title string if tweet length exceeds 140. If still length is greater than 140 then also Trimmed review string.
4. All length are determined after eliminating white spaces fom beginning and end.

### Installing
#### TAR File
* Create a directory using command: mkdir directoryName
* cd directoryName
* Download 99designs-coding-test.tar.gz and extract it using command tar -xzpf 99designs-coding-test.tar.gz
* cd goplaybook

#### Github
* Create a directory using command: mkdir directoryName
* cd directoryName
* run https://github.com/nehamittal01/goplaybook.git
* cd goplaybook


### Executing program
* Run below commands inside goplaybook directory
* go build main.go
* go run main.go  Movies.json Reviews.json
* Final output will get print in command line