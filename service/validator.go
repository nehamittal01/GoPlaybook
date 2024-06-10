package service

import (
	"goplaybook/model"
	"strings"
)

func validateFinalAndUpdatePost(tweetPost *model.TweetPost) string {
	finalPost := buildPost(tweetPost.Title, tweetPost.Year, tweetPost.Review, tweetPost.Rating)

	if len(finalPost) > 140 && len(tweetPost.Title) > 25 {
		tweetPost.Title = tweetPost.Title[:25]
	}

	finalPost = buildPost(tweetPost.Title, tweetPost.Year, tweetPost.Review, tweetPost.Rating)

	if len(finalPost) > 140 {
		finalPost = buildPost(tweetPost.Title, tweetPost.Year, "", tweetPost.Rating)
		finalPostLength := len(strings.TrimSpace(finalPost))

		var trimUpto int

		if finalPostLength >= 140 {
			trimUpto = finalPostLength - 140
		} else {
			trimUpto = 140 - finalPostLength
		}
		if len(tweetPost.Review) > trimUpto {
			tweetPost.Review = tweetPost.Review[:trimUpto]
		}

	}
	finalPost = buildPost(tweetPost.Title, tweetPost.Year, tweetPost.Review, tweetPost.Rating)
	return finalPost
}

func buildPost(title string, year string, review string, rating string) string {
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
	}

	if rating != "" {
		finalPost.WriteString(" ")
		finalPost.WriteString(rating)
	}

	return finalPost.String()

}
