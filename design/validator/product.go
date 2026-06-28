package validator

import "strings"

type CreateProductRequest struct {
	ID    string
	Name  string
	Price int64
}

type FieldError struct {
	Field   string
	Message string
}

func ValidateCreateProductRequest(req CreateProductRequest) []FieldError {
	var errs []FieldError

	if strings.TrimSpace(req.ID) == "" {
		errs = append(errs, FieldError{Field: "id", Message: "ID cannot be empty"})
	}

	if strings.TrimSpace(req.Name) == "" {
		errs = append(errs, FieldError{Field: "name", Message: "Name cannot be empty"})
	}

	if req.Price < 0 {
		errs = append(errs, FieldError{Field: "price", Message: "Price must be greater than or equal to 0"})
	}

	return errs
}
