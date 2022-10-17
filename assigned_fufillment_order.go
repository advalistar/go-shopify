package goshopify

import (
	"fmt"
	"reflect"
)

const (
	assignedFulfillmentOrderBasePath = "assigned_fulfillment_orders"
)

type AssignedFulfillmentOrderService interface {
	List(interface{}) ([]FulfillmentOrder, error)
	GetOrderList() []string
}

// AssignedFulfillmentOrderServiceOp handles communication with the order related methods of the
// Shopify API.
type AssignedFulfillmentOrderServiceOp struct {
	client *Client
}

// Represents the result from the assignedFulfillmentOrder.json endpoint
type AssignedFulfillmentOrderResource struct {
	FulfillmentOrders []FulfillmentOrder `json:"fulfillment_orders"`
}

// List assignedFulfillmentOrder
func (s *AssignedFulfillmentOrderServiceOp) List(options interface{}) ([]FulfillmentOrder, error) {
	path := fmt.Sprintf("%s.json", assignedFulfillmentOrderBasePath)
	resource := &AssignedFulfillmentOrderResource{}
	return resource.FulfillmentOrders, s.client.Get(path, resource, options)
}

func (s *AssignedFulfillmentOrderServiceOp) GetOrderList() []string {
	str := new(FulfillmentOrder)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
