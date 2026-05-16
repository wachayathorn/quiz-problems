package design

import "strings"

type CreateProductRequest struct {
	ID    string
	Name  string
	Price int64
}

type ValidationError struct {
	Field   string
	Message string
}

func ValidateCreateProductRequest(req CreateProductRequest) []ValidationError {
	var errors []ValidationError

	if strings.TrimSpace(req.ID) == "" {
		errors = append(errors, ValidationError{
			Field:   "id",
			Message: "ID cannot be empty",
		})
	}

	if strings.TrimSpace(req.Name) == "" {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name cannot be empty",
		})
	}

	if req.Price < 0 {
		errors = append(errors, ValidationError{
			Field:   "price",
			Message: "Price must be greater than or equal to 0",
		})
	}

	return errors
}
