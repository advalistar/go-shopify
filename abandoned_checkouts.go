package goshopify

import (
	"fmt"
)

const (
	abandonedCheckoutsBasePath = "checkouts"
)

type AbandonedCheckoutService interface {
	List(interface{}) ([]Checkout, error)
}

type AbandonedCheckoutServiceOp struct {
	client *Client
}

type AbandonedCheckoutResource struct {
	Checkouts []Checkout `json:"checkouts"`
}

// List checkouts
func (s *AbandonedCheckoutServiceOp) List(options interface{}) ([]Checkout, error) {
	path := fmt.Sprintf("%s.json", abandonedCheckoutsBasePath)
	resource := &AbandonedCheckoutResource{}
	return resource.Checkouts, s.client.Get(path, resource, options)
}
