package ticketmodel

type Filter struct {
	UserID int   `json:"user_id,omitempty"`
	ShowID int   `json:"show_id,omitempty" form:"showID"`
	Status []int `json:"status,omitempty" form:"status"`
}

func (f *Filter) Validate() {
	if f.ShowID <= 0 {
		f.ShowID = 1
	}
}
