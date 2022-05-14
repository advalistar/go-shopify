package goshopify

import (
	"fmt"
)

const fulfillmentServicesBasePath = "fulfillment_services"

type FulfillmentSvcService interface {
	List(interface{}) ([]FulfillmentSvc, error)
}

type FulfillmentSvcServiceOp struct {
	client *Client
}

type FulfillmentSvc struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	Email                  string `json:"email"`
	ServiceName            string `json:"service_name"`
	Handle                 string `json:"handle"`
	FulfillmentOrdersOptIn bool   `json:"fulfillment_orders_opt_in"`
	IncludePendingStock    bool   `json:"include_pending_stock"`
	ProviderID             int64  `json:"provider_id"`
	LocationID             int64  `json:"location_id"`
	CallbackURL            string `json:"callback_url"`
	TrackingSupport        bool   `json:"tracking_support"`
	InventoryManagement    bool   `json:"inventory_management"`
	AdminGraphqlAPIID      string `json:"admin_graphql_api_id"`
}

type FulfillmentSvcsResource struct {
	FulfillmentSvcs []FulfillmentSvc `json:"fulfillment_services"`
}

// List shipping zones
func (s *FulfillmentSvcServiceOp) List(options interface{}) ([]FulfillmentSvc, error) {
	path := fmt.Sprintf("%s.json", fulfillmentServicesBasePath)
	resource := new(FulfillmentSvcsResource)
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentSvcs, err
}
