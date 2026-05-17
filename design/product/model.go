package product

type Product struct {
	ID    string
	Name  string
	Price int
}

type CreateProductInput struct {
	SKU   string
	Name  string
	Price int
}
