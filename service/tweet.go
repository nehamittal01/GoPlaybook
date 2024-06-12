package service

import (
	"fmt"
	"goplaybook/model"
	"goplaybook/utils"
)

var (
	movieTittleToTweetPostMap = make(map[string]*model.TweetPost)
	tweet                     []string
)

func CreateTweet(movies []model.Movie, reviews []model.Reviews) {
	for _, n := range reviews {

		rating := utils.GetRating(n.Score)
		tweetVal := model.CreateNewTweetPost(n.Review, rating, n.Title, 0)

		if len(utils.TrimSpace(n.Title)) > 0 && len(utils.TrimSpace(n.Review)) > 0 && len(utils.TrimSpace(rating)) > 0 {

			movieTittleToTweetPostMap[n.Title] = tweetVal
		}

	}

	for _, n := range movie {
		val, ok := movieTittleToTweetPostMap[n.Title]
		if ok {
			model.UpdateTweetPost(val, n.Year)

		}
	}

	for k := range movieTittleToTweetPostMap {
		finalPost := validateAndUpdatePost(movieTittleToTweetPostMap[k])
		fmt.Println(finalPost)
		//fmt.Println(utf8.RuneCountInString(finalPost))
		tweet = append(tweet, finalPost)
	}
}
