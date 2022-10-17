package goshopify

import (
	"fmt"
	"reflect"
)

const (
	balanceBasePath = "shopify_payments/balance"
)

type BalanceService interface {
	List(interface{}) ([]Balance, error)
	GetOrderList() []string
}

// BalanceServiceOp handles communication with the order related methods of the
// Shopify API.
type BalanceServiceOp struct {
	client *Client
}

// Balance represents a Shopify order
type Balance struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// Represents the result from the balance.json endpoint
type BalanceResource struct {
	Balance []Balance `json:"balance"`
}

// List balance
func (s *BalanceServiceOp) List(options interface{}) ([]Balance, error) {
	path := fmt.Sprintf("%s.json", balanceBasePath)
	resource := &BalanceResource{}
	return resource.Balance, s.client.Get(path, resource, options)
}

func (s *BalanceServiceOp) GetOrderList() []string {
	str := new(Balance)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
