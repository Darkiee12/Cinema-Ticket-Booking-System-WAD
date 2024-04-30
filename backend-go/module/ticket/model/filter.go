package ticketmodel

type Filter struct {
	ShowID int   `json:"show_id,omitempty" form:"showID"`
	Status []int `json:"status,omitempty" form:"status"`
}

func (f *Filter) Validate() {
	if f.ShowID <= 0 {
		f.ShowID = 1
	}
	if f.Status == nil || len(f.Status) == 0 {
		f.Status = []int{1}
	}
}
