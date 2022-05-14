package goshopify

import (
	"fmt"
)

const (
	assignedFufillmentOrderBasePath = "assigned_fufillment_orders"
)

type AssignedFufillmentOrderService interface {
	List(interface{}) ([]FulfillmentOrder, error)
}

// AssignedFufillmentOrderServiceOp handles communication with the order related methods of the
// Shopify API.
type AssignedFufillmentOrderServiceOp struct {
	client *Client
}

// Represents the result from the assignedFufillmentOrder.json endpoint
type AssignedFufillmentOrderResource struct {
	FulfillmentOrders []FulfillmentOrder `json:"fulfillment_orders"`
}

// List assignedFufillmentOrder
func (s *AssignedFufillmentOrderServiceOp) List(options interface{}) ([]FulfillmentOrder, error) {
	path := fmt.Sprintf("%s.json", assignedFufillmentOrderBasePath)
	resource := &AssignedFufillmentOrderResource{}
	return resource.FulfillmentOrders, s.client.Get(path, resource, options)
}
