package product

import (
	"context"

	"github.com/wachayathorn/go-errcode"
)

type ProductService interface {
	CreateProduct(ctx context.Context, input CreateProductInput) error
}

type ProductServiceImpl struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &ProductServiceImpl{
		repo: repo,
	}
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, input CreateProductInput) error {
	if input.SKU == "" {
		return errcode.BadRequest.WithMessage("sku is required")
	}
	if input.Name == "" {
		return errcode.BadRequest.WithMessage("name is required")
	}
	if input.Price <= 0 {
		return errcode.BadRequest.WithMessage("price must be greater than 0")
	}

	exists, err := s.repo.ExistsBySKU(ctx, input.SKU)
	if err != nil {
		return errcode.InternalServerError.
			WithMessage("failed to check sku existence").
			WithCause(err)
	}

	if exists {
		return errcode.BadRequest.WithMessage("sku already exists")
	}

	product := Product{
		ID:    input.SKU,
		Name:  input.Name,
		Price: input.Price,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return errcode.InternalServerError.
			WithMessage("failed to create product").
			WithCause(err)
	}

	return nil
}
