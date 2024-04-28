package ticketmodel

import (
	"errors"
	"time"
)

const EntityName = "Ticket"
const TableName = "tickets"

type Ticket struct {
	ID         int       `gorm:"column:id;primary_key"`
	SeatNumber int       `gorm:"column:seat_number"`
	Status     int16     `gorm:"column:status"`
	Timestamp  time.Time `gorm:"column:timestamp"`
	ShowID     int64     `gorm:"column:show_id"`
	UserID     int64     `gorm:"column:user_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (Ticket) TableName() string { return TableName }

type TicketCreate struct {
	SeatNumber int       `gorm:"column:seat_number"`
	Status     int16     `gorm:"column:status"`
	Timestamp  time.Time `gorm:"column:timestamp"`
	ShowID     int64     `gorm:"column:show_id"`
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
