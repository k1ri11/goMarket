package api

import (
	"github.com/gorilla/mux"
	"goMarket/internal/api/handlers"
	"goMarket/internal/api/middleware"
	"goMarket/internal/app/product"
	"net/http"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Инициализация хранилища и сервисов
	productRepo := product.NewRepository()
	productService := product.NewService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// CRUD маршруты для продуктов
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// Добавляем логирование ко всем маршрутам.
	return middleware.RequestLogger(r)
}
