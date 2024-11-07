package repo

import (
	"context"
	"hexashop/internal/adapter/storage/postgres"
	"hexashop/internal/core/domain"
)

type User struct {
	repo *postgres.Client
}

func NewUser(db *postgres.Client) *User {
	return &User{repo: db}
}

func (u *User) Create(ctx context.Context, user *domain.User) error {
	return u.repo.DB().Create(user).Error
}

func (u *User) Get(ctx context.Context, id int) (*domain.User, error) {
	user := &domain.User{}
	err := u.repo.DB().First(user, id).Error
	return user, err
}

func (u *User) Update(ctx context.Context, user *domain.User) error {
	return u.repo.DB().Save(user).Error
}

func (u *User) Delete(ctx context.Context, id int) error {
	return u.repo.DB().Delete(&domain.User{}, id).Error
}
