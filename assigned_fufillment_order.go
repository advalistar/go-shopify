package goshopify

import (
	"fmt"
)

const (
	assignedFufillmentOrderBasePath = "assigned_fufillment_orders"
)

type AssignedFufillmentOrderService interface {
	List(interface{}) ([]AssignedFufillmentOrder, error)
}

// AssignedFufillmentOrderServiceOp handles communication with the order related methods of the
// Shopify API.
type AssignedFufillmentOrderServiceOp struct {
	client *Client
}

// AssignedFufillmentOrder represents a Shopify order
type AssignedFufillmentOrder struct {
	ID                       int64         `json:"id"`
	ShopID                   int64         `json:"shop_id"`
	OrderID                  int64         `json:"order_id"`
	AssignedLocationID       int64         `json:"assigned_location_id"`
	RequestStatus            string        `json:"request_status"`
	Status                   string        `json:"status"`
	SupportedActions         []string      `json:"supported_actions"`
	Destination              *Destination  `json:"destination"`
	LineItems                []*LineItem   `json:"line_items"`
	OutgoingRequests         []interface{} `json:"outgoing_requests"`
	FulfillmentAerviceHandle string        `json:"fulfillment_service_handle"`
	AssignedLocation         *Location     `json:"assigned_location"`
}

type Destination struct {
	ID        int64  `json:"id"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	Company   string `json:"company"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	Zip       string `json:"zip"`
}

// Represents the result from the assignedFufillmentOrder.json endpoint
type AssignedFufillmentOrderResource struct {
	AssignedFufillmentOrder []AssignedFufillmentOrder `json:"assignedFufillmentOrders"`
}

// List assignedFufillmentOrder
func (s *AssignedFufillmentOrderServiceOp) List(options interface{}) ([]AssignedFufillmentOrder, error) {
	path := fmt.Sprintf("%s.json", assignedFufillmentOrderBasePath)
	resource := &AssignedFufillmentOrderResource{}
	return resource.AssignedFufillmentOrder, s.client.Get(path, resource, options)
}
