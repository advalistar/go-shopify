package goshopify

import (
	"github.com/shopspring/decimal"
)

// ShippingZoneService is an interface for interfacing with the shipping zones endpoint
// of the Shopify API.
// See: https://help.shopify.com/api/reference/store-properties/shippingzone
type ShippingZoneService interface {
	List() ([]ShippingZone, error)
}

// ShippingZoneServiceOp handles communication with the shipping zone related methods
// of the Shopify API.
type ShippingZoneServiceOp struct {
	client *Client
}

// ShippingZone represents a Shopify shipping zone
type ShippingZone struct {
	ID                           int64                          `json:"id"`
	Name                         string                         `json:"name"`
	ProfileID                    string                         `json:"profile_id"`
	LocationGroupID              string                         `json:"location_group_id"`
	AdminGraphqlAPIID            string                         `json:"admin_graphql_api_id"`
	Countries                    []*Country                     `json:"countries"`
	WeightBasedShippingRates     []*WeightBasedShippingRate     `json:"weight_based_shipping_rates"`
	PriceBasedShippingRates      []*PriceBasedShippingRate      `json:"price_based_shipping_rates"`
	CarrierShippingRateProviders []*CarrierShippingRateProvider `json:"carrier_shipping_rate_providers"`
}

// WeightBasedShippingRate represents a Shopify weight-constrained shipping rate
type WeightBasedShippingRate struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	Price          *decimal.Decimal `json:"price"`
	ShippingZoneID int64            `json:"shipping_zone_id"`
	WeightLow      *decimal.Decimal `json:"weight_low"`
	WeightHigh     *decimal.Decimal `json:"weight_high"`
}

// PriceBasedShippingRate represents a Shopify subtotal-constrained shipping rate
type PriceBasedShippingRate struct {
	ID               int64            `json:"id"`
	Name             string           `json:"name"`
	Price            *decimal.Decimal `json:"price"`
	ShippingZoneID   int64            `json:"shipping_zone_id"`
	MinOrderSubtotal *decimal.Decimal `json:"min_order_subtotal"`
	MaxOrderSubtotal *decimal.Decimal `json:"max_order_subtotal"`
}

// CarrierShippingRateProvider represents a Shopify carrier-constrained shipping rate
type CarrierShippingRateProvider struct {
	ID               int64             `json:"id"`
	CarrierServiceID int64             `json:"carrier_service_id"`
	FlatModifier     *decimal.Decimal  `json:"flat_modifier"`
	ServiceFilter    map[string]string `json:"service_filter"`
	PercentModifier  *decimal.Decimal  `json:"percent_modifier"`
	ShippingZoneID   int64             `json:"shipping_zone_id"`
}

// Represents the result from the shipping_zones.json endpoint
type ShippingZonesResource struct {
	ShippingZones []ShippingZone `json:"shipping_zones"`
}

// List shipping zones
func (s *ShippingZoneServiceOp) List() ([]ShippingZone, error) {
	resource := new(ShippingZonesResource)
	err := s.client.Get("shipping_zones.json", resource, nil)
	return resource.ShippingZones, err
}
