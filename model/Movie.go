package model

type Movie struct {
	Title string `json:"title,omitempty"`
	Year  int    `json:"year"`
}
