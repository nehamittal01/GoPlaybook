package service

import (
	"fmt"
	"goplaybook/model"
)

var (
	movieTittleToTweetPostMap = make(map[string]*model.TweetPost)
	tweet                     []string
)

func CreateTweet(movies []model.Movie, reviews []model.Reviews) {
	for _, n := range reviews {
		rating := GetRating(n.Score)
		tweetVal := model.CreateNewTweetPost(n.Review, rating, n.Title, 0)
		movieTittleToTweetPostMap[n.Title] = tweetVal
	}

	for _, n := range movie {
		val, ok := movieTittleToTweetPostMap[n.Title]
		if ok {
			model.UpdateTweetPost(val, n.Year)

		}
	}
	for k := range movieTittleToTweetPostMap {
		finalPost := validateFinalAndUpdatePost(movieTittleToTweetPostMap[k])
		fmt.Println(finalPost)
		tweet = append(tweet, finalPost)
	}
}

func GetRating(score int) string {
	var rating string
	if score >= 0 && score <= 10 {
		rating = "1/2"
	}
	if score > 10 && score <= 20 {
		rating = "★"
	}
	if score > 20 && score <= 30 {
		rating = "* 1/2"
	}
	if score > 30 && score <= 40 {
		rating = "**"
	}
	if score > 40 && score <= 50 {
		rating = "**1/2"
	}
	if score > 50 && score <= 60 {
		rating = "***"
	}
	if score > 60 && score <= 70 {
		rating = "***1/2"
	}
	if score > 70 && score <= 80 {
		rating = "****"
	}
	if score > 80 && score <= 90 {
		rating = "****1/2"
	}
	if score > 90 && score <= 100 {
		rating = "*****"
	}
	return rating
}
