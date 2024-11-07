package repo

import (
	"context"
	"hexashop/internal/adapter/storage/postgres"
	"hexashop/internal/core/domain"
)

type Product struct {
	repo *postgres.Client
}

func NewProduct(db *postgres.Client) *Product {
	return &Product{repo: db}
}

func (u *Product) Create(ctx context.Context, product *domain.Product) error {
	return u.repo.DB().Create(product).Error
}

func (u *Product) Get(ctx context.Context, id int) (*domain.Product, error) {
	product := &domain.Product{}
	err := u.repo.DB().First(product, id).Error
	return product, err
}

func (u *Product) Update(ctx context.Context, product *domain.Product) error {
	return u.repo.DB().Save(product).Error
}

func (u *Product) Delete(ctx context.Context, id int) error {
	return u.repo.DB().Delete(&domain.Product{}, id).Error
}
