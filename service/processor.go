package service

import (
	"encoding/json"
	"fmt"
	"goplaybook/model"
	"io/ioutil"
)

var (
	movie   []model.Movie
	reviews []model.Reviews
)

func Process(args []string) {

	movieJsonFile, movieError := readJsonFile(args[1])
	if movieError != nil {
		fmt.Println(movieError)
		return
	}
	movieError = json.Unmarshal(movieJsonFile, &movie)
	if movieError != nil {
		fmt.Println(movieError)
		return
	}

	reviewJsonFile, reviewError := readJsonFile(args[2])
	if reviewError != nil {
		fmt.Println(reviewError)
		return
	}
	reviewError = json.Unmarshal(reviewJsonFile, &reviews)
	if reviewError != nil {
		fmt.Println(reviewError)
		return
	}

	CreateTweet(movie, reviews)

}

func readJsonFile(fileData string) ([]byte, error) {
	return ioutil.ReadFile(fileData)
}
