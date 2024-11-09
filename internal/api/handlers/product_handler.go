package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"goMarket/internal/app/product"
)

type ProductHandler struct {
	service *product.Service
}

func NewProductHandler(service *product.Service) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetProducts возвращает список всех продуктов.
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByID возвращает продукт по ID.
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// CreateProduct добавляет новый продукт.
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product product.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdProduct := h.service.CreateProduct(product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProduct)
}

// UpdateProduct обновляет существующий продукт.
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updatedProduct product.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.service.UpdateProduct(id, updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct удаляет продукт по ID.
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.service.DeleteProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
