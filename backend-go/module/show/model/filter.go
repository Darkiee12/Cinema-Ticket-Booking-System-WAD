package showmodel

import "time"

type Filter struct {
	Date      *time.Time `json:"date"`
	StartTime *time.Time `json:"startTime"`
	ImdbID    string     `json:"imdbID"`
}
