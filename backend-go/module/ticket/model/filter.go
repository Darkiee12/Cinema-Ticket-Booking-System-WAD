package ticketmodel

type Filter struct {
	ShowID int   `json:"show_id,omitempty" form:"-"`
	Status []int `json:"-"`
}
