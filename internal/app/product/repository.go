package product

import (
	"errors"
	"sync"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

type Repository struct {
	products []Product
	nextID   int
	mu       sync.Mutex // для потокобезопасности при доступе к данным
}

// NewRepository создает новое локальное хранилище продуктов.
func NewRepository() *Repository {
	return &Repository{
		products: []Product{
			{ID: 1, Name: "iPhone 14", Description: "Apple smartphone", Price: 999.99, Stock: 10},
			{ID: 2, Name: "Samsung Galaxy S23", Description: "Samsung smartphone", Price: 899.99, Stock: 15},
			{ID: 3, Name: "Google Pixel 7", Description: "Google smartphone", Price: 799.99, Stock: 20},
		},
		nextID: 4,
	}
}

// GetAll возвращает все продукты из локального хранилища.
func (r *Repository) GetAll() []Product {
	return r.products
}

// GetByID возвращает продукт по ID.
func (r *Repository) GetByID(id int) (*Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

// Create добавляет новый продукт в хранилище.
func (r *Repository) Create(product Product) Product {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	r.nextID++
	r.products = append(r.products, product)
	return product
}

// Update обновляет существующий продукт по ID.
func (r *Repository) Update(id int, updated Product) (*Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, product := range r.products {
		if product.ID == id {
			updated.ID = id
			r.products[i] = updated
			return &r.products[i], nil
		}
	}
	return nil, errors.New("product not found")
}

// Delete удаляет продукт по ID.
func (r *Repository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
