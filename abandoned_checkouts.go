package goshopify

import (
	"fmt"
	"reflect"
)

const (
	abandonedCheckoutsBasePath = "checkouts"
)

type AbandonedCheckoutService interface {
	List(interface{}) ([]Checkout, error)
	GetOrderList() []string
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

func (s *AbandonedCheckoutServiceOp) GetOrderList() []string {
	str := new(Checkout)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
