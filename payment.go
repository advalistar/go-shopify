package goshopify

import (
	"fmt"
)

const (
	paymentsBasePath = "checkouts/%s/payments"
)

type PaymentService interface {
	List(string, interface{}) ([]Payment, error)
}

type PaymentServiceOp struct {
	client *Client
}

type Payment struct {
	ID                            int64        `json:"id"`
	UniqueToken                   string       `json:"unique_token"`
	PaymentProcessingErrorMessage string       `json:"payment_processing_error_message"`
	Transaction                   *Transaction `json:"transaction"`
	CreditCard                    *CreditCard  `json:"credit_card"`
	Checkout                      *Checkout    `json:"checkout"`
}

type CreditCard struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	FirstDigits int64  `json:"first_digits"`
	LastDigits  int64  `json:"last_digits"`
	Brand       string `json:"brand"`
	ExpiryMonth int64  `json:"expiry_month"`
	ExpiryYear  int64  `json:"expiry_year"`
	CustomerID  int64  `json:"customer_id"`
}

type PaymentsResource struct {
	Payments []Payment `json:"payments"`
}

func (s *PaymentServiceOp) List(token string, options interface{}) ([]Payment, error) {
	path := fmt.Sprintf(paymentsBasePath+".json", token)
	resource := &PaymentsResource{}
	return resource.Payments, s.client.Get(path, resource, options)
}
