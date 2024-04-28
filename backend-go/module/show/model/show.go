package showmodel

import (
	"time"
)

const EntityName = "Show"
const TableName = "shows"

type Show struct {
	ID           int        `gorm:"column:id;primary_key" json:"-"`
	Date         *time.Time `gorm:"column:date" json:"date"`
	EndTime      *time.Time `gorm:"column:end_time" json:"endTime"`
	StartTime    *time.Time `gorm:"column:start_time" json:"startTime"`
	AuditoriumID int64      `gorm:"column:auditorium_id" json:"-"`
	ImdbID       string     `gorm:"column:imdb_id" json:"imdbID"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"-"`
}

func (Show) TableName() string { return TableName }
