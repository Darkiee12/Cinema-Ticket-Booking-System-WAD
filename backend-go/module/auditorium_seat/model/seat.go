package auditoriumseatsmodel

import "cinema/common"

const EntityName = "Auditorium Seats"
const TableName = "auditorium_seats"

type AuditoriumSeat struct {
	common.SQLModel `json:",inline" gorm:"embedded"`
	SeatNumber      int `json:"seat_number" gorm:"column:seat_number;"`
	Type            int `json:"type" gorm:"column:type;"`
	AuditoriumID    int `json:"-" gorm:"column:auditorium_id;"`
}

func (AuditoriumSeat) TableName() string { return TableName }
