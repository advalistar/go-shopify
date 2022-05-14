package goshopify

import (
	"fmt"
	"time"
)

const currenciesBasePath = "currencies"

type CurrencyService interface {
	List(interface{}) ([]Currency, error)
}

type CurrencyServiceOp struct {
	client *Client
}

// ShippingCurrency represents a Shopify shipping country
type Currency struct {
	Currency      string     `json:"currency"`
	RateUpdatedAt *time.Time `json:"rate_updated_at"`
	Enabled       bool       `json:"enabled"`
}

// Represents the result from the shipping_zones.json endpoint
type CurrenciesResource struct {
	Currencies []Currency `json:"currencies"`
}

// List shipping zones
func (s *CurrencyServiceOp) List(options interface{}) ([]Currency, error) {
	path := fmt.Sprintf("%s.json", currenciesBasePath)
	resource := new(CurrenciesResource)
	err := s.client.Get(path, resource, options)
	return resource.Currencies, err
}
