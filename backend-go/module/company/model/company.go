package companymodel

import (
	"cinema/common"
	"errors"
	"strings"
)

const EntityName = "Company"
const TableName = "companies"

type Company struct {
	common.SQLModel `json:",inline"`
	OwnerID         int                `json:"owner_id" gorm:"column:owner_id;"`
	Owner           *common.SimpleUser `json:"owner" gorm:"preload:false;foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LogoPath        string             `json:"logo_path" gorm:"column:logo_path;"`
	Name            string             `json:"name" gorm:"column:name;"`
}

func (Company) TableName() string { return TableName }

func (c *Company) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeCompany)
	if owner := c.Owner; owner != nil {
		owner.Mask(isAdminOrOwner)
	}
}

type CompanyCreate struct {
	common.SQLModel `json:",inline" swaggerignore:"true"`
	OwnerID         int    `json:"owner_id" gorm:"column:user_id;" swaggerignore:"true"`
	LogoPath        string `json:"logo_path" gorm:"column:logo_path;"`
	Name            string `json:"name" gorm:"column:name;"`
}

func (CompanyCreate) TableName() string { return TableName }

func (data *CompanyCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCompany)
}

func (data *CompanyCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
