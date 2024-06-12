package service

import (
	"goplaybook/model"
	"unicode/utf8"
)

func validateAndUpdatePost(tweetPost *model.TweetPost) string {
	finalPost := model.BuildPost(tweetPost.Title, tweetPost.Year, tweetPost.Review, tweetPost.Rating)

	if utf8.RuneCountInString(finalPost) > 140 && utf8.RuneCountInString(tweetPost.Title) > 25 {

		tweetPost.Title = tweetPost.Title[:25]
	}
	finalPost = model.BuildPost(tweetPost.Title, tweetPost.Year, tweetPost.Review, tweetPost.Rating)

	//fmt.Println(finalPost)
	//fmt.Println(utf8.RuneCountInString(finalPost))

	if utf8.RuneCountInString(finalPost) > 140 {

		finalPost = model.BuildPost(tweetPost.Title, tweetPost.Year, "", tweetPost.Rating)
		finalPostLength := utf8.RuneCountInString(finalPost)
		//fmt.Println(finalPost)
		//fmt.Println(finalPostLength)

		var trimUpto int
		if finalPostLength >= 140 {
			trimUpto = finalPostLength - 140
		} else {
			trimUpto = 140 - finalPostLength
		}

		if utf8.RuneCountInString(tweetPost.Review) > trimUpto {
			tweetPost.Review = tweetPost.Review[:trimUpto]
		}
		finalPost = model.BuildPost(tweetPost.Title, tweetPost.Year, tweetPost.Review, tweetPost.Rating)
	}

	return finalPost
}
