package usermodel

import (
	"cinema/common"
	"errors"
	"regexp"
)

const (
	TableName = "users"
)

type User struct {
	common.SQLModel `json:",inline"`
	DateOfBirth     *common.Date `json:"date_of_birth" gorm:"column:date_of_birth;"`
	Email           string       `json:"email" gorm:"column:email;"`
	Password        string       `json:"-" gorm:"column:password;"`
	Gender          *string      `json:"gender,omitempty" gorm:"gender"`
	Name            *string      `json:"name,omitempty" gorm:"column:name;"`
	Role            string       `json:"role,omitempty" gorm:"column:tier;"`
	Salt            string       `json:"-" gorm:"column:salt;"`
	Phone           *string      `json:"phone,omitempty" gorm:"column:phone_number;"`
}

func (u *User) GetUserId() int {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (User) TableName() string {
	return TableName
}

type UserUpdate struct {
	DateOfBirth *common.Date `json:"date_of_birth,omitempty" gorm:"column:date_of_birth;"`
	Gender      *string      `json:"gender,omitempty" gorm:"gender"`
	Name        *string      `json:"name,omitempty" gorm:"column:name;"`
	Phone       *string      `json:"phone,omitempty" gorm:"column:phone_number;"`
}

func (UserUpdate) TableName() string {
	return TableName
}

func (u *UserUpdate) Validate() error {
	if u.Phone != nil && len(*u.Phone) > 0 {
		phoneRegex := `^\d{10}$`
		re := regexp.MustCompile(phoneRegex)

		if !re.MatchString(*u.Phone) {
			return ErrPhoneInvalid
		}
	}
	return nil
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrPhoneInvalid = common.NewCustomError(
		errors.New("phone number invalid"),
		"phone number invalid",
		"ErrPhoneInvalid",
	)
)
