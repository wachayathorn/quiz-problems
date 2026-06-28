package product

import "fmt"

type InMemoryStore interface {
	Create(product Product) error
	GetByID(id string) (Product, bool)
	List(page, pageSize int) []Product
	Update(product Product) error
	Delete(id string) bool
}

type InMemoryRepository struct {
	products map[string]Product
}

func NewInMemoryRepository() InMemoryStore {
	return &InMemoryRepository{
		products: make(map[string]Product),
	}
}

func (r *InMemoryRepository) Create(product Product) error {
	if product.ID == "" {
		return fmt.Errorf("product ID cannot be empty")
	}
	if product.Name == "" {
		return fmt.Errorf("product name cannot be empty")
	}
	if product.Price <= 0 {
		return fmt.Errorf("product price must be greater than 0")
	}

	r.products[product.ID] = product
	return nil
}

func (r *InMemoryRepository) GetByID(id string) (Product, bool) {
	product, exists := r.products[id]
	return product, exists
}

func (r *InMemoryRepository) List(page, pageSize int) []Product {
	products := make([]Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products
}

func (r *InMemoryRepository) Update(product Product) error {
	if product.ID == "" {
		return fmt.Errorf("product ID cannot be empty")
	}
	if product.Name == "" {
		return fmt.Errorf("product name cannot be empty")
	}
	if product.Price <= 0 {
		return fmt.Errorf("product price must be greater than 0")
	}

	r.products[product.ID] = product
	return nil
}

func (r *InMemoryRepository) Delete(id string) bool {
	if id == "" {
		return false
	}
	delete(r.products, id)
	return true
}
