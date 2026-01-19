package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type txKey struct{}

// TxManager handles database transactions
type TxManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type txManager struct {
	db *sqlx.DB
}

// NewTxManager creates a new transaction manager
func NewTxManager(db *sqlx.DB) TxManager {
	return &txManager{db: db}
}

// WithTransaction executes fn within a transaction
func (tm *txManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	// Inject transaction into context
	ctx = context.WithValue(ctx, txKey{}, tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetTx extracts transaction from context
func GetTx(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}

// Queryable interface for both DB and Tx
type Queryable interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// GetQueryable returns tx if exists in context, otherwise db
func GetQueryable(ctx context.Context, db *sqlx.DB) Queryable {
	if tx := GetTx(ctx); tx != nil {
		return tx
	}
	return db
}
