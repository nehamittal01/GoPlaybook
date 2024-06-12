package model

import (
	"goplaybook/utils"
	"strconv"
	"strings"
)

type TweetPost struct {
	Title  string `json:"title,omitempty"`
	Review string `json:"review,omitempty"`
	Rating string `json:"score,omitempty"`
	Year   string `json:"year"`
}

func CreateNewTweetPost(review string, rating string, title string, year int) *TweetPost {
	tweetPost := new(TweetPost)
	if title != "" && len(utils.TrimSpace(title)) > 0 {

		tweetPost.Title = utils.TrimLeftAndRight(title, " ")
	}
	if rating != "" && len(utils.TrimSpace(rating)) > 0 {

		tweetPost.Rating = utils.TrimLeftAndRight(rating, " ")
	}
	if review != "" && len(utils.TrimSpace(review)) > 0 {

		tweetPost.Review = utils.TrimLeftAndRight(review, " ")
	}
	if year > 0 {

		tweetPost.Year = strconv.Itoa(year)
		tweetPost.Year = utils.TrimLeftAndRight(tweetPost.Year, " ")
	} else {
		tweetPost.Year = ""
	}
	return tweetPost
}

func UpdateTweetPost(tweetPost *TweetPost, year int) {
	if year > 0 {
		tweetPost.Year = strconv.Itoa(year)
	} else {
		tweetPost.Year = ""
	}

}

func BuildPost(title string, year string, review string, rating string) string {
	var finalPost strings.Builder

	if title != "" {
		finalPost.WriteString(title)
	}

	if year != "" {
		finalPost.WriteString(" (")
		finalPost.WriteString(year)
		finalPost.WriteString(")")
	}

	if review != "" {
		finalPost.WriteString(": ")
		finalPost.WriteString(review)
	} else {
		finalPost.WriteString(": ")
	}

	if rating != "" {
		finalPost.WriteString(" ")
		finalPost.WriteString(rating)
	}

	return finalPost.String()

}
