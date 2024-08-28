package dbmanager

import (
	"context"
	"github.com/nilroad/kateb"
)

type DB interface {
	Begin() (any, error)
	Commit() error
	Rollback() error
}

type TrxContextKey string

type TrxManager struct {
	DB
	ContextKeyName TrxContextKey
	logger         *kateb.Logger
}

func (r *TrxManager) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	d, err := r.Begin()
	db, ok := d.(DB)
	if !ok {
		panic("implement SQLDb interface")
	}

	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			if err := db.Rollback(); err != nil {
				r.logger.Error("failed to rollback", map[string]interface{}{
					"err": err,
				})
			}
			panic(err)
		}
	}()

	ctx = context.WithValue(ctx, r.ContextKeyName, db)
	err = fn(ctx)
	if err != nil {
		if err := db.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err := db.Commit(); err != nil {
		return err
	}

	return nil
}
