package utils

import "strings"

func GetRating(score int) string {
	var rating string
	if score >= 0 && score <= 10 {
		rating = "½"
	}
	if score > 10 && score <= 20 {
		rating = "★"
	}
	if score > 20 && score <= 30 {
		rating = "*½"
	}
	if score > 30 && score <= 40 {
		rating = "**"
	}
	if score > 40 && score <= 50 {
		rating = "**½"
	}
	if score > 50 && score <= 60 {
		rating = "***"
	}
	if score > 60 && score <= 70 {
		rating = "***½"
	}
	if score > 70 && score <= 80 {
		rating = "****"
	}
	if score > 80 && score <= 90 {
		rating = "****½"
	}
	if score > 90 && score <= 100 {
		rating = "*****"
	}
	return rating
}

func TrimSpace(st string) string {
	return strings.TrimSpace(st)
}

func TrimLeft(st string, regex string) string {
	return strings.TrimLeft(st, regex)
}

func TrimRight(st string, regex string) string {
	return strings.TrimLeft(st, regex)
}

func TrimLeftAndRight(st string, regex string) string {
	val := TrimLeft(st, regex)
	val = TrimRight(val, regex)
	return val
}
