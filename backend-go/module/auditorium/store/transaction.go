package auditoriumstore

import (
	"cinema/common"
	"context"
)

func (store *sqlStore) OpenTransaction(context context.Context) {
	tx := store.db.Begin()
	store.tx = store.db
	store.db = tx
}
func (store *sqlStore) RollbackTransaction(context context.Context) {
	store.db.Rollback()
}
func (store *sqlStore) CloseTransaction(context context.Context) error {
	defer func() {
		store.db = store.tx
		store.tx = nil
	}()
	if err := store.db.Commit().Error; err != nil {
		store.db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
