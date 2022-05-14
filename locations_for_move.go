package goshopify

import (
	"fmt"
)

const (
	locationsForMoveBasePath = "fulfillment_orders/%d/locations_for_move"
)

type LocationsForMoveService interface {
	List(int64, interface{}) ([]LocationsForMove, error)
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
