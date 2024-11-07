package port

import (
	"context"
	"hexashop/internal/core/domain"
)

type ProductSvc interface {
	CreateProduct(ctx context.Context, req *domain.Product) error
	GetProduct(ctx context.Context, id int) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id int, req *domain.Product) error
	DeleteProduct(ctx context.Context, id int) error
}

type ProductRepo interface {
	Create(ctx context.Context, product *domain.Product) error
	Get(ctx context.Context, id int) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id int) error
}
