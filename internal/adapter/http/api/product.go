package api

import (
	"encoding/json"
	"hexashop/internal/adapter/http/rq"
	"hexashop/internal/port"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	productSvc port.ProductSvc
}

func NewProduct(productSvc port.ProductSvc) *Product {
	return &Product{
		productSvc: productSvc,
	}
}

// CreateProduct handler
func (api *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var Product rq.Product
	if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	req := rq.DomainProduct(Product)
	if err := api.productSvc.CreateProduct(r.Context(), &req); err != nil {
		http.Error(w, "Failed to create Product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Product)
}

// GetProduct handler
func (api *Product) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	Product, err := api.productSvc.GetProduct(r.Context(), idInt)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Product)
}

// UpdateProduct handler
func (api *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	var Product rq.Product
	if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	req := rq.DomainProduct(Product)
	if err := api.productSvc.UpdateProduct(r.Context(), idInt, &req); err != nil {
		http.Error(w, "Failed to update Product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Product)
}

// DeleteProduct handler
func (api *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)
	if err := api.productSvc.DeleteProduct(r.Context(), idInt); err != nil {
		http.Error(w, "Failed to delete Product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
