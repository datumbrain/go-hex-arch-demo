package rq

import "hexashop/internal/core/domain"

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func DomainProduct(p Product) domain.Product {
	return domain.Product{
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
}
