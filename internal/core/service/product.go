package service

import (
	"context"
	"hexashop/internal/core/domain"
	"hexashop/internal/port"
	"log"
)

type Product struct {
	productRepo port.ProductRepo
}

func NewProduct(repo port.ProductRepo) *Product {
	return &Product{productRepo: repo}
}

func (svc *Product) CreateProduct(ctx context.Context, req *domain.Product) error {
	if err := svc.productRepo.Create(ctx, req); err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (svc *Product) GetProduct(ctx context.Context, id int) (*domain.Product, error) {
	product, err := svc.productRepo.Get(ctx, id)
	if err != nil {
		log.Println(ctx, err)
		return nil, err
	}

	return product, nil
}

func (svc *Product) UpdateProduct(ctx context.Context, id int, req *domain.Product) error {
	req.ID = id
	if err := svc.productRepo.Update(ctx, req); err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (svc *Product) DeleteProduct(ctx context.Context, id int) error {
	if err := svc.productRepo.Delete(ctx, id); err != nil {
		log.Println(ctx, err)
		return err
	}
	return nil
}
