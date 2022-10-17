package goshopify

import (
	"fmt"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const (
	fulfillmentEventsBasePath = "orders/%d/fulfillments/%d/events"
)

type FulfillmentEventService interface {
	List(int64, int64, interface{}) ([]FulfillmentEvent, error)
	GetOrderList() []string
}

type FulfillmentEventServiceOp struct {
	client *Client
}

type FulfillmentEvent struct {
	ID                  int64            `json:"id"`
	FulfillmentID       int64            `json:"fulfillment_id"`
	Status              string           `json:"status"`
	Message             string           `json:"message"`
	HappenedAt          *time.Time       `json:"happened_at"`
	City                string           `json:"city"`
	Province            string           `json:"province"`
	Country             string           `json:"country"`
	Zip                 string           `json:"zip"`
	Address1            string           `json:"address1"`
	Latitude            *decimal.Decimal `json:"latitude"`
	Longitude           *decimal.Decimal `json:"longitude"`
	ShopID              int64            `json:"shop_id"`
	CreatedAt           *time.Time       `json:"created_at"`
	UpdatedAt           *time.Time       `json:"updated_at"`
	EstimatedDeliveryAt *time.Time       `json:"estimated_delivery_at"`
	OrderID             int64            `json:"order_id"`
	AdminGraphqlAPIID   string           `json:"admin_graphql_api_id"`
}

type FulfillmentEventsResource struct {
	FulfillmentEvents []FulfillmentEvent `json:"fulfillment_events"`
}

// List of discount codes
func (s *FulfillmentEventServiceOp) List(fulfillmentID, orderID int64, options interface{}) ([]FulfillmentEvent, error) {
	path := fmt.Sprintf(fulfillmentEventsBasePath+".json", fulfillmentID, orderID)
	resource := new(FulfillmentEventsResource)
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentEvents, err
}

func (s *FulfillmentEventServiceOp) GetOrderList() []string {
	str := new(FulfillmentEvent)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
