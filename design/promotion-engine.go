package design

type CartItem struct {
	SKU      string
	Price    int
	Quantity int
}

type Promotion struct {
	Type        string // "percentage" or "fixed"
	Value       int
	MinSpend    int
	MaxDiscount int
}

func calculateFinalPrice(items []CartItem, promo Promotion) int {
	totalPrice := 0
	for _, item := range items {
		if item.Price <= 0 || item.Quantity <= 0 {
			continue
		}

		totalPrice += item.Price * item.Quantity
	}

	if totalPrice <= 0 {
		return 0
	}

	if totalPrice < promo.MinSpend {
		return totalPrice
	}

	if promo.Value <= 0 {
		return totalPrice
	}

	discount := 0
	switch promo.Type {
	case "percentage":
		discount = (totalPrice * promo.Value) / 100
	case "fixed":
		discount = promo.Value
	default:
		return totalPrice
	}

	if promo.MaxDiscount > 0 && discount > promo.MaxDiscount {
		discount = promo.MaxDiscount
	}

	finalPrice := totalPrice - discount
	if finalPrice < 0 {
		return 0
	}

	return finalPrice
}
