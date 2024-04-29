package showmodel

import (
	"cinema/common"
	"time"
)

const EntityName = "Show"
const TableName = "shows"

type Show struct {
	ID               int        `json:"id" gorm:"column:id;primary_key"`
	Date             *time.Time `json:"date" gorm:"column:date" `
	EndTime          *time.Time `json:"endTime" gorm:"column:end_time"`
	StartTime        *time.Time `json:"startTime" gorm:"column:start_time"`
	AuditoriumID     int64      `json:"-" gorm:"column:auditorium_id"`
	AuditoriumFakeID string     `json:"auditoriumID" gorm:"-"`
	ImdbID           string     `json:"imdbID" gorm:"column:imdb_id"`
	CreatedAt        *time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt        *time.Time `json:"-" gorm:"column:updated_at"`
}

func (Show) TableName() string { return TableName }

func (s *Show) Mask(isAdminOrOwner bool) {
	uid := common.NewUID(uint32(s.AuditoriumID), common.DbTypeAuditorium, 1)
	s.AuditoriumFakeID = uid.String()
}

type ShowUpdate struct {
	Date         *time.Time `gorm:"column:date" json:"date"`
	EndTime      *time.Time `gorm:"column:end_time" json:"endTime"`
	StartTime    *time.Time `gorm:"column:start_time" json:"startTime"`
	AuditoriumID int64      `gorm:"column:auditorium_id" json:"-"`
	ImdbID       string     `gorm:"column:imdb_id" json:"imdbID"`
}

func (ShowUpdate) TableName() string { return TableName }
