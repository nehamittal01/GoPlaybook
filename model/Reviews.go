package model

type Reviews struct {
	Title  string `json:"title,omitempty"`
	Review string `json:"review,omitempty"`
	Score  int    `json:"score,omitempty"`
}
