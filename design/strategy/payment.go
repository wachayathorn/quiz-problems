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

func main() {
	paymentProcessor := NewPaymentProcessor(&CreditCardMethod{
		CardNumber:   "1234567890",
		ExpiryDate:   "12/2026",
		SecurityCode: "123",
	})
	paymentProcessor.ProcessPayment(100)

	paymentProcessor.SetMethod(&PayPalMethod{
		Email:    "test@example.com",
		Password: "password",
	})
	paymentProcessor.ProcessPayment(100)

	paymentProcessor.SetMethod(&BankTransferMethod{
		BankName:      "Bank of America",
		AccountNumber: "1234567890",
	})
	paymentProcessor.ProcessPayment(100)
}
