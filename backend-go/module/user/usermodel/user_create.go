package usermodel

import (
	"cinema/common"
)

const EntityName = "UserID"

type UserCreate struct {
	common.SQLModel `json:",inline" swaggerignore:"true"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"password" gorm:"column:password;"`
	Name            string `json:"name" gorm:"column:name;"`
	Role            string `json:"-" gorm:"column:tier;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (UserCreate) TableName() string {
	return TableName
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return TableName
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}
