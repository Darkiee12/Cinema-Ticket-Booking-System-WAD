package auditoriumseatsstore

import (
	"cinema/common"
	"context"
)

func (store *sqlStore) OpenTransaction(context context.Context) (*sqlStore, error) {
	tx := store.db.Begin()
	if tx.Error != nil {
		return nil, common.ErrDB(tx.Error)
	}
	return &sqlStore{db: tx}, nil
}
func (store *sqlStore) RollbackTransaction(context context.Context) {
	store.db.Rollback()
}
func (store *sqlStore) CloseTransaction(context context.Context) error {
	if err := store.db.Commit().Error; err != nil {
		store.db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
