package design

import "errors"

type PaymentType string

const (
	PaymentStripe PaymentType = "stripe"
	PaymentOmise  PaymentType = "omise"
	PaymentPayPal PaymentType = "paypal"
)

type Payment interface {
	Pay(amount int) error
	Provider() PaymentType
}

type StripePayment struct {
	apiKey string
}

func (p *StripePayment) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (p *StripePayment) Provider() PaymentType {
	return PaymentStripe
}

type OmisePayment struct {
	secretKey string
}

func (p *OmisePayment) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (p *OmisePayment) Provider() PaymentType {
	return PaymentOmise
}

type PayPalPayment struct {
	clientID string
}

func (p *PayPalPayment) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (p *PayPalPayment) Provider() PaymentType {
	return PaymentPayPal
}

type PaymentConfig struct {
	StripeAPIKey string
	OmiseKey     string
	PayPalClient string
}

type PaymentFactory struct {
	config PaymentConfig
}

func NewPaymentFactory(config PaymentConfig) *PaymentFactory {
	return &PaymentFactory{config: config}
}

func (f *PaymentFactory) CreatePayment(paymentType PaymentType) (Payment, error) {
	switch paymentType {
	case PaymentStripe:
		return &StripePayment{apiKey: f.config.StripeAPIKey}, nil
	case PaymentOmise:
		return &OmisePayment{secretKey: f.config.OmiseKey}, nil
	case PaymentPayPal:
		return &PayPalPayment{clientID: f.config.PayPalClient}, nil
	default:
		return nil, errors.New("unsupported payment type")
	}
}
