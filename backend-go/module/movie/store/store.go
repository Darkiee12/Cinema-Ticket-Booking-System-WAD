package moviestore

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db.Omit("created_at", "updated_at")}
}
