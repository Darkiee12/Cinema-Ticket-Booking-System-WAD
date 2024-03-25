package auditoriummodel

import (
	"cinema/common"
	"errors"
	"strings"
)

const EntityName = "Auditorium"
const TableName = "auditoriums"

type Auditorium struct {
	common.SQLModel `json:",inline"`
	Name            string               `json:"name" gorm:"column:name;"`
	Seats           int                  `json:"seats" gorm:"column:seats;"`
	CinemaID        int                  `json:"-" gorm:"column:cinema_id;"`
	Cinema          *common.SimpleCinema `json:"cinema" gorm:"preload:false;foreignKey:CinemaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Auditorium) TableName() string { return TableName }

func (c *Auditorium) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeAuditorium)
	if cinema := c.Cinema; cinema != nil {
		cinema.Mask(isAdminOrOwner)
	}
}

type AuditoriumCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Seats           int    `json:"seats" gorm:"column:seats;"`
	CinemaID        int    `json:"cinema_id" gorm:"column:cinema_id;"`
}

func (AuditoriumCreate) TableName() string { return TableName }

func (data *AuditoriumCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCinema)
}

func (data *AuditoriumCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type UpdateCinema struct {
	Name     string `json:"name" gorm:"column:name;"`
	Seats    int    `json:"seats" gorm:"column:seats;"`
	CinemaID int    `json:"cinema_id" gorm:"column:cinema_id;"`
}

func (UpdateCinema) TableName() string { return TableName }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
