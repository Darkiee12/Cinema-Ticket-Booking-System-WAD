package ticketmodel

import (
	"cinema/common"
	"errors"
	"time"
)

const EntityName = "Ticket"
const TableName = "tickets"

type Ticket struct {
	common.SQLModel `json:",inline"`
	SeatNumber      int                `gorm:"column:seat_number" json:"seat_number"`
	Timestamp       time.Time          `gorm:"column:timestamp" json:"timestamp"`
	ShowID          int64              `gorm:"column:show_id" json:"show_id"`
	UserID          int64              `gorm:"column:user_id" json:"-"`
	User            *common.SimpleUser `gorm:"preload:false;foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}

func (Ticket) TableName() string { return TableName }

func (t *Ticket) Mask(isAdminOrOwner bool) {
	if user := t.User; user != nil {
		user.Mask(isAdminOrOwner)
	}
}

type TicketCreate struct {
	SeatNumber int       `gorm:"column:seat_number" json:"seat_number"`
	Status     int16     `gorm:"column:status" json:"status"`
	Timestamp  time.Time `gorm:"column:timestamp" json:"timestamp"`
	ShowID     int64     `gorm:"column:show_id" json:"showID"`
}

func (TicketCreate) TableName() string { return TableName }

type TicketUpdate struct {
	Status int16 `gorm:"column:status"`
	UserID int64 `gorm:"column:user_id"`
}

func (TicketUpdate) TableName() string { return TableName }

var (
	ErrTicketHasBeenSold = errors.New("ticket has been sold")
	ErrTicketUnavailable = errors.New("ticket is unavailable")
)

const (
	TicketStatusSold      = 0
	TicketStatusAvailable = 1
)
