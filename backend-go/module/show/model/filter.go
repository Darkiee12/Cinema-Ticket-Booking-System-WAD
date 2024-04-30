package showmodel

import (
	"cinema/common"
)

type Filter struct {
	Date *common.Date `json:"date" form:"date"`
	//StartTime *common.Time `json:"startTime" form:"startTime"`
	ImdbID string `json:"imdbID" form:"imdbID"`
}
