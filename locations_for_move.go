package goshopify

import (
	"fmt"
	"reflect"
)

const (
	locationsForMoveBasePath = "fulfillment_orders/%d/locations_for_move"
)

type LocationsForMoveService interface {
	List(int64, interface{}) ([]LocationsForMove, error)
	GetOrderList() []string
}

type LocationsForMoveServiceOp struct {
	client *Client
}

type LocationsForMove struct {
	Location *LocationsForMoveLocation `json:"destination"`
	Message  string                    `json:"message"`
	Movable  bool                      `json:"movable"`
}

type LocationsForMoveLocation struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type LocationsForMovesResource struct {
	LocationsForMoves []LocationsForMove `json:"locations_for_move"`
}

// List of discount codes
func (s *LocationsForMoveServiceOp) List(fulfillmentOrderID int64, options interface{}) ([]LocationsForMove, error) {
	path := fmt.Sprintf(locationsForMoveBasePath+".json", fulfillmentOrderID)
	resource := new(LocationsForMovesResource)
	err := s.client.Get(path, resource, options)
	return resource.LocationsForMoves, err
}

func (s *LocationsForMoveServiceOp) GetOrderList() []string {
	str := new(LocationsForMove)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
