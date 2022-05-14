package goshopify

import (
	"fmt"
)

const (
	locationsForMoveBasePath = "order/%d/fulfillment_orders"
)

type LocationsForMoveService interface {
	List(int64, interface{}) ([]FulfillmentOrder, error)
}

type LocationsForMoveServiceOp struct {
	client *Client
}

type LocationsForMovesResource struct {
	FulfillmentOrders []FulfillmentOrder `json:"fulfillment_orders"`
}

// List of discount codes
func (s *LocationsForMoveServiceOp) List(orderID int64, options interface{}) ([]FulfillmentOrder, error) {
	path := fmt.Sprintf(locationsForMoveBasePath+".json", orderID)
	resource := new(LocationsForMovesResource)
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentOrders, err
}
