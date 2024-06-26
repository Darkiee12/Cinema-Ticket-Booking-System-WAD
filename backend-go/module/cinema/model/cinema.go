package cinemamodel

import (
	"cinema/common"
	"errors"
	"strings"
)

const EntityName = "Cinema"
const TableName = "cinemas"

type Cinema common.Cinema

func (Cinema) TableName() string { return TableName }

func (c *Cinema) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeCinema)

	if owner := c.Owner; owner != nil {
		owner.Mask(isAdminOrOwner)
	}
}

type CinemaCreate struct {
	common.SQLModel `json:",inline" swaggerignore:"true"`
	OwnerID         int    `json:"owner_id" gorm:"column:owner_id;"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address"  gorm:"column:address;"`
	Email           string `json:"email" gorm:"column:email;"`
	PhoneNumber     string `json:"phone_number" gorm:"column:phone_number;"`
	Banner          string `json:"banner" gorm:"column:banner;"`
}

func (CinemaCreate) TableName() string { return TableName }

func (data *CinemaCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCinema)
}

func (data *CinemaCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}

	data.Email = strings.TrimSpace(data.Email)
	if data.Email == "" {
		return ErrEmailIsEmpty
	}

	data.PhoneNumber = strings.TrimSpace(data.PhoneNumber)
	if data.PhoneNumber == "" {
		return ErrPhoneNumberIsEmpty

	}

	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return ErrAddressIsEmpty
	}

	data.Banner = strings.TrimSpace(data.Banner)
	if data.Banner == "" {
		return ErrAddressIsEmpty
	}
	return nil
}

type UpdateCinema struct {
	Name        string `json:"name" gorm:"column:name;"`
	Address     string `json:"address"  gorm:"column:address;"`
	Email       string `json:"email" gorm:"column:email;"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;"`
	Banner      string `json:"banner" gorm:"column:banner;"`
}

func (UpdateCinema) TableName() string { return TableName }

var (
	ErrNameIsEmpty        = errors.New("name can not be empty")
	ErrEmailIsEmpty       = errors.New("email can not be empty")
	ErrPhoneNumberIsEmpty = errors.New("phone number can not be empty")
	ErrAddressIsEmpty     = errors.New("address can not be empty")
	ErrBannerIsEmpty      = errors.New("banner can not be empty")
)
