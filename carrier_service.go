package goshopify

import (
	"fmt"
)

const (
	carrierServicerBasePath = "carrier_services"
)

type CarrierServiceService interface {
	List(interface{}) ([]CarrierService, error)
}

// CarrierServiceServiceOp handles communication with the order related methods of the
// Shopify API.
type CarrierServiceServiceOp struct {
	client *Client
}

// CarrierService represents a Shopify order
type CarrierService struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	Active             bool   `json:"active"`
	ServiceDiscovery   bool   `json:"service_discovery"`
	CarrierServiceType string `json:"carrier_service_type"`
	AdminGraphqlAPIID  string `json:"admin_graphql_api_id"`
	Format             string `json:"format"`
	CallbackURL        string `json:"callback_url"`
}

// Represents the result from the carrierServicer.json endpoint
type CarrierServiceResource struct {
	CarrierServices []CarrierService `json:"carrier_services"`
}

// List carrierServicer
func (s *CarrierServiceServiceOp) List(options interface{}) ([]CarrierService, error) {
	path := fmt.Sprintf("%s.json", carrierServicerBasePath)
	resource := &CarrierServiceResource{}
	return resource.CarrierServices, s.client.Get(path, resource, options)
}
