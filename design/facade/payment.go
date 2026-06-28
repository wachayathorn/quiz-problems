package facade

import (
	"errors"

	"github.com/wachayathorn/quiz-problems/design/factory"
)

type CheckoutRequest struct {
	Provider  factory.ProviderType
	Amount    int
	OrderID   string
	UserEmail string
}

type CheckoutResult struct {
	OrderID     string
	Transaction string
	Provider    factory.ProviderType
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

func (r *transactionRecorder) record(orderID string, provider factory.ProviderType) string {
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
	factory   *factory.PaymentFactory
	validator *orderValidator
	recorder  *transactionRecorder
	notifier  *paymentNotifier
}

func NewPaymentFacade(config factory.GatewayConfig) *PaymentFacade {
	return &PaymentFacade{
		factory:   factory.NewPaymentFactory(config),
		validator: &orderValidator{},
		recorder:  &transactionRecorder{},
		notifier:  &paymentNotifier{},
	}
}

func (f *PaymentFacade) Checkout(req CheckoutRequest) (*CheckoutResult, error) {
	if err := f.validator.validate(req); err != nil {
		return nil, err
	}

	gateway, err := f.factory.Create(req.Provider)
	if err != nil {
		return nil, err
	}

	if err := gateway.Pay(req.Amount); err != nil {
		return nil, err
	}

	txnID := f.recorder.record(req.OrderID, gateway.Provider())

	if err := f.notifier.notify(req.UserEmail, req.OrderID); err != nil {
		return nil, err
	}

	return &CheckoutResult{
		OrderID:     req.OrderID,
		Transaction: txnID,
		Provider:    gateway.Provider(),
	}, nil
}
