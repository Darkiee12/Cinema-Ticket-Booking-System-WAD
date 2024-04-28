package auditoriumseatsmodel

type Filter struct {
	Type         []int `json:"types"`
	AuditoriumID int   `json:"-" gorm:"column:auditorium_id;"`
	Status       []int `json:"-"`
}
