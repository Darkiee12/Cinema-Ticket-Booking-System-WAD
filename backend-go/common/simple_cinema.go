package common

type Cinema struct {
	SQLModel    `json:",inline"`
	OwnerID     int         `json:"owner_id" gorm:"column:owner_id;"`
	Owner       *SimpleUser `json:"owner" gorm:"preload:false;foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string      `json:"name" gorm:"column:name;"`
	Address     string      `json:"address"  gorm:"column:address;"`
	Capacity    int         `json:"capacity" gorm:"column:capacity;"`
	Email       string      `json:"email" gorm:"column:email;"`
	PhoneNumber string      `json:"phone_number" gorm:"column:phone_number;"`
}

func (Cinema) TableName() string {
	return "cinemas"
}

func (u *Cinema) Mask(isAdmin bool) {
	u.GenUID(DbTypeCinema)
}
