package auditoriummodel

type Filter struct {
	CinemaID   int    `json:"cinema_id,omitempty" form:"-"`
	CinemaName string `json:"cinema_name,omitempty" form:"cinema_name"`
	Status     []int  `json:"-"`
}
