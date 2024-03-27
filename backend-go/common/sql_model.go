package common

import "time"

type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id"`
	FakeID    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;-"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.ID), dbType, 1)
	m.FakeID = &uid
}
