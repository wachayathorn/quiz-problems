package design

import "errors"

type CheckoutRequest struct {
	PaymentType PaymentType
	Amount      int
	OrderID     string
	UserEmail   string
}

type CheckoutResult struct {
	OrderID     string
	Transaction string
	Provider    PaymentType
}

type orderValidator struct{}

func (v *orderValidator) validate(req CheckoutRequest) error {
	if req.OrderID == "" {
		return errors.New("order id is required")
	}
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}

type transactionRecorder struct{}

func (r *transactionRecorder) record(orderID string, provider PaymentType, amount int) string {
	return "txn_" + orderID + "_" + string(provider)
}

type paymentNotifier struct{}

func (n *paymentNotifier) notify(email, orderID string) error {
	if email == "" {
		return nil
	}
	return nil
}

type PaymentFacade struct {
	factory   *PaymentFactory
	validator *orderValidator
	recorder  *transactionRecorder
	notifier  *paymentNotifier
}

func NewPaymentFacade(config PaymentConfig) *PaymentFacade {
	return &PaymentFacade{
		factory:   NewPaymentFactory(config),
		validator: &orderValidator{},
		recorder:  &transactionRecorder{},
		notifier:  &paymentNotifier{},
	}
}

func (f *PaymentFacade) Checkout(req CheckoutRequest) (*CheckoutResult, error) {
	if err := f.validator.validate(req); err != nil {
		return nil, err
	}

	payment, err := f.factory.CreatePayment(req.PaymentType)
	if err != nil {
		return nil, err
	}

	if err := payment.Pay(req.Amount); err != nil {
		return nil, err
	}

	txnID := f.recorder.record(req.OrderID, payment.Provider(), req.Amount)

	if err := f.notifier.notify(req.UserEmail, req.OrderID); err != nil {
		return nil, err
	}

	return &CheckoutResult{
		OrderID:     req.OrderID,
		Transaction: txnID,
		Provider:    payment.Provider(),
	}, nil
}
