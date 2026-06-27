package design

type PaymentStrategy interface {
	Pay(amount int) error
}

type CreditCardPayment struct {
	cardNumber   string
	expiryDate   string
	securityCode string
}

func (c *CreditCardPayment) Pay(amount int) error {
	return nil
}

type PayPalPayment struct {
	email    string
	password string
}

func (p *PayPalPayment) Pay(amount int) error {
	return nil
}

type BankTransferPayment struct {
	bankName      string
	accountNumber string
}

func (b *BankTransferPayment) Pay(amount int) error {
	return nil
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) ProcessPayment(amount int) error {
	return p.strategy.Pay(amount)
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: strategy}
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}
