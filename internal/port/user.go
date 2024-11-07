package port

import (
	"context"
	"hexashop/internal/core/domain"
)

type UserSvc interface {
	CreateUser(ctx context.Context, req *domain.User) error
	GetUser(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, id int, req *domain.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserRepo interface {
	Create(ctx context.Context, product *domain.User) error
	Get(ctx context.Context, id int) (*domain.User, error)
	Update(ctx context.Context, product *domain.User) error
	Delete(ctx context.Context, id int) error
}
