package factory

import "errors"

type ProviderType string

const (
	ProviderStripe ProviderType = "stripe"
	ProviderOmise  ProviderType = "omise"
	ProviderPayPal ProviderType = "paypal"
)

type Gateway interface {
	Pay(amount int) error
	Provider() ProviderType
}

type StripeGateway struct {
	apiKey string
}

func (g *StripeGateway) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (g *StripeGateway) Provider() ProviderType {
	return ProviderStripe
}

type OmiseGateway struct {
	secretKey string
}

func (g *OmiseGateway) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (g *OmiseGateway) Provider() ProviderType {
	return ProviderOmise
}

type PayPalGateway struct {
	clientID string
}

func (g *PayPalGateway) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

func (g *PayPalGateway) Provider() ProviderType {
	return ProviderPayPal
}

type GatewayConfig struct {
	StripeAPIKey string
	OmiseKey     string
	PayPalClient string
}

type PaymentFactory struct {
	config GatewayConfig
}

func NewPaymentFactory(config GatewayConfig) *PaymentFactory {
	return &PaymentFactory{config: config}
}

func (f *PaymentFactory) Create(provider ProviderType) (Gateway, error) {
	switch provider {
	case ProviderStripe:
		return &StripeGateway{apiKey: f.config.StripeAPIKey}, nil
	case ProviderOmise:
		return &OmiseGateway{secretKey: f.config.OmiseKey}, nil
	case ProviderPayPal:
		return &PayPalGateway{clientID: f.config.PayPalClient}, nil
	default:
		return nil, errors.New("unsupported payment provider")
	}
}
