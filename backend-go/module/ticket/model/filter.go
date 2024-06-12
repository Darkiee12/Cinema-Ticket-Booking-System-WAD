package ticketmodel

type Filter struct {
	UserID int   `json:"user_id,omitempty"`
	ShowID int   `json:"show_id,omitempty" form:"showID"`
	Status []int `json:"status,omitempty" form:"status"`
}

func (f *Filter) Validate(isAdmin bool) {
	if f.ShowID <= 0 {
		f.ShowID = 1
	}
	if (f.Status == nil || len(f.Status) == 0) && !isAdmin {
		f.Status = []int{1}
	}
}
