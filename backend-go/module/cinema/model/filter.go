package cinemamodel

type Filter struct {
	FakeCinemaID string `json:"-" form:"cinema_id"`
	CinemaID     int    `json:"cinema_id,omitempty" form:"-"`
	Status       []int  `json:"-"`
}
