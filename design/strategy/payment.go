package strategy

type PaymentMethod interface {
	Pay(amount int) error
}

type CreditCardMethod struct {
	CardNumber   string
	ExpiryDate   string
	SecurityCode string
}

func (m *CreditCardMethod) Pay(amount int) error {
	return nil
}

type PayPalMethod struct {
	Email    string
	Password string
}

func (m *PayPalMethod) Pay(amount int) error {
	return nil
}

type BankTransferMethod struct {
	BankName      string
	AccountNumber string
}

func (m *BankTransferMethod) Pay(amount int) error {
	return nil
}

type PaymentProcessor struct {
	method PaymentMethod
}

func NewPaymentProcessor(method PaymentMethod) *PaymentProcessor {
	return &PaymentProcessor{method: method}
}

func (p *PaymentProcessor) SetMethod(method PaymentMethod) {
	p.method = method
}

func (p *PaymentProcessor) ProcessPayment(amount int) error {
	return p.method.Pay(amount)
}
