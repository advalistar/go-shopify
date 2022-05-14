package goshopify

import (
	"fmt"
	"time"
)

const (
	fulfillmentOrdersBasePath       = "fulfillment_orders/%d"
	ordersFulfillmentOrdersBasePath = "orders/%d/fulfillment_orders"
)

type FulfillmentOrderService interface {
	Get(int64, interface{}) (*FulfillmentOrder, error)
	List(int64, interface{}) ([]FulfillmentOrder, error)
}

type FulfillmentOrderServiceOp struct {
	client *Client
}

type FulfillmentOrder struct {
	ID                       int64              `json:"id"`
	ShopID                   int64              `json:"shop_id"`
	OrderID                  int64              `json:"order_id"`
	AssignedLocationID       int64              `json:"assigned_location_id"`
	RequestStatus            string             `json:"request_status"`
	Status                   string             `json:"status"`
	SupportedActions         []string           `json:"supported_actions"`
	Destination              *Destination       `json:"destination"`
	LineItems                []*LineItem        `json:"line_items"`
	FulfillmentAerviceHandle string             `json:"fulfillment_service_handle"`
	AssignedLocation         *Location          `json:"assigned_location"`
	MerchantRequests         []*MerchantRequest `json:"merchant_requests"`
}

type MerchantRequest struct {
	Message        string           `json:"message"`
	RequestOptions []*RequestOption `json:"request_options"`
	Kind           string           `json:"kind"`
}

type RequestOption struct {
	ShippingMethod string     `json:"shipping_method"`
	Note           string     `json:"note"`
	Date           *time.Time `json:"date"`
}

type FulfillmentOrdersResource struct {
	FulfillmentOrders []FulfillmentOrder `json:"fulfillment_orders"`
}

type FulfillmentOrderResource struct {
	FulfillmentOrder *FulfillmentOrder `json:"fulfillment_order"`
}

func (s FulfillmentOrderServiceOp) Get(fulfillmentOrderID int64, options interface{}) (*FulfillmentOrder, error) {
	path := fmt.Sprintf(fulfillmentOrdersBasePath+".json", fulfillmentOrderID)
	resource := &FulfillmentOrderResource{}
	return resource.FulfillmentOrder, s.client.Get(path, resource, options)
}

// List of discount codes
func (s *FulfillmentOrderServiceOp) List(orderID int64, options interface{}) ([]FulfillmentOrder, error) {
	path := fmt.Sprintf(ordersFulfillmentOrdersBasePath+".json", orderID)
	resource := new(FulfillmentOrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentOrders, err
}
