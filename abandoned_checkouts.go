package goshopify

import (
	"fmt"
)

const (
	adandonedCheckoutsBasePath = "checkouts"
)

type AdandonedCheckoutService interface {
	List(interface{}) ([]Checkout, error)
}

type AdandonedCheckoutServiceOp struct {
	client *Client
}

type AdandonedCheckoutResource struct {
	Checkouts []Checkout `json:"checkouts"`
}

// List checkouts
func (s *AdandonedCheckoutServiceOp) List(options interface{}) ([]Checkout, error) {
	path := fmt.Sprintf("%s.json", adandonedCheckoutsBasePath)
	resource := &AdandonedCheckoutResource{}
	return resource.Checkouts, s.client.Get(path, resource, options)
}
