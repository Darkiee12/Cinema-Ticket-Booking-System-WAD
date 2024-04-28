package ticketstore

import (
	"context"
	"gorm.io/gorm"
	"log"
)

func (store *sqlStore) model(ctx context.Context, model ...interface{}) *gorm.DB {
	tx := extractTx(ctx)

	if tx != nil {
		return tx
	}
	return store.db
}

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func (store *sqlStore) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	// begin transaction
	tx := store.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// run callback
	err := tFunc(injectTx(ctx, tx))
	if err != nil {
		// if error, rollback
		tx.Rollback()
		return err
	}
	// if no error, commit
	if errCommit := tx.Commit().Error; errCommit != nil {
		tx.Rollback()
		log.Printf("commit transaction: %v", errCommit)
		return errCommit
	}
	return nil
}
