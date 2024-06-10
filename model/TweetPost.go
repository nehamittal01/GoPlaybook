package model

import "strconv"

type TweetPost struct {
	Title  string `json:"title,omitempty"`
	Review string `json:"review,omitempty"`
	Rating string `json:"score,omitempty"`
	Year   string `json:"year"`
}

func CreateNewTweetPost(review string, rating string, title string, year int) *TweetPost {
	tweetPost := new(TweetPost)
	if title != "" {
		tweetPost.Title = title
	}

	if rating != "" {
		tweetPost.Rating = rating
	}
	if review != "" {
		tweetPost.Review = review
	}
	if year > 0 {
		tweetPost.Year = strconv.Itoa(year)
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
