package userrepo

import (
	"context"
	"sotoon/internal/adapter/storage/mysql"
	"sotoon/internal/core/entity"
)

type Repository struct {
	*mysql.SQLDb
}

func New(db *mysql.SQLDb) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.DB(ctx).WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
