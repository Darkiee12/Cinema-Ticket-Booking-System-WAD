package countrymodel

import (
	"errors"
	"strings"
)

const EntityName = "Country"
const TableName = "countries"

type Country struct {
	Iso31661 string `json:"iso_3166_1" gorm:"column:iso_3166_1;"`
	Name     string `json:"name" gorm:"column:name;"`
}

func (Country) TableName() string { return TableName }

func (data *Country) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	data.Iso31661 = strings.TrimSpace(data.Iso31661)
	if data.Name == "" || data.Iso31661 == "" {
		return ErrNameIsEmpty
	}
	return nil
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
