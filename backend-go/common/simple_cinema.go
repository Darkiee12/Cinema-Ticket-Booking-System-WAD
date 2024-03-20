package common

type SimpleCinema struct {
	SQLModel `json:",inline"`
	Name     string `json:"name" gorm:"column:name;"`
	Address  string `json:"address"  gorm:"column:address;"`
	Email    string `json:"email" gorm:"column:email;"`
}

func (SimpleCinema) TableName() string {
	return "cinemas"
}

func (u *SimpleCinema) Mask(isAdmin bool) {
	u.GenUID(DbTypeCinema)
}
