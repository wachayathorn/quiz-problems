package design

import "fmt"

type Product struct {
	ID    string
	Name  string
	Price int
}

type ProductRepositoryInterface interface {
	Create(product Product) error
	GetByID(id string) (Product, bool)
	List(page int, pageSize int) []Product
	Update(product Product) error
	Delete(id string) bool
}

type ProductRepository struct {
	products map[string]Product
}

func NewProductRepository() ProductRepositoryInterface {
	return &ProductRepository{
		products: make(map[string]Product),
	}
}

func (r *ProductRepository) Create(product Product) error {
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

func (r *ProductRepository) GetByID(id string) (Product, bool) {
	product, exists := r.products[id]
	return product, exists
}

func (r *ProductRepository) List(page int, pageSize int) []Product {
	var products []Product
	for _, product := range r.products {
		products = append(products, product)
	}
	return products
}

func (r *ProductRepository) Update(product Product) error {
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

func (r *ProductRepository) Delete(id string) bool {
	if id == "" {
		return false
	}

	delete(r.products, id)
	return true
}
