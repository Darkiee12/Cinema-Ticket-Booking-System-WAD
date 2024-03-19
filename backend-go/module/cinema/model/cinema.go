package cinemamodel

import "cinema/common"

const EntityName = "Cinema"

type Cinema struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"-" gorm:"column:cinema_id"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address"  gorm:"column:address;"`
	Capacity        int    `json:"capacity" gorm:"column:capacity;"`
	Email           string `json:"email" gorm:"column:email;"`
	PhoneNumber     string `json:"phone_number" gorm:"column:phone_number;"`
}

func (Cinema) TableName() string { return "cinemas" }

func (r *Cinema) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeCinema)
}
