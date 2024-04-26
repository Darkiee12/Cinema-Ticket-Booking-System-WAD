package moviemodel

type Filter struct {
	OwnerID int   `json:"owner_id,omitempty" form:"-"`
	Status  []int `json:"-"`
}
