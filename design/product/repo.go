package product

import "context"

type ProductRepository interface {
	ExistsBySKU(ctx context.Context, sku string) (bool, error)
	Create(ctx context.Context, product Product) error
}

type productRepositoryImpl struct {
	// Add any dependencies here such DB instance
}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}

func (r *productRepositoryImpl) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	// Implementation would go here
	return false, nil
}

func (r *productRepositoryImpl) Create(ctx context.Context, product Product) error {
	// Implementation would go here
	return nil
}
