package goshopify

import (
	"fmt"

	"github.com/shopspring/decimal"
)

const countriesBasePath = "countries"

type CountryService interface {
	List(interface{}) ([]Country, error)
}

type CountryServiceOp struct {
	client *Client
}

// ShippingCountry represents a Shopify shipping country
type Country struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Tax       *decimal.Decimal `json:"tax"`
	Code      string           `json:"code"`
	TaxName   string           `json:"tax_name"`
	CountryID int64            `json:"shipping_zone_id"`
	Provinces []*Province      `json:"provinces"`
}

type Province struct {
	ID             int64            `json:"id"`
	CountryID      int64            `json:"country_id"`
	Name           string           `json:"name"`
	Code           string           `json:"code"`
	Tax            *decimal.Decimal `json:"tax"`
	TaxName        string           `json:"tax_name"`
	TaxType        string           `json:"tax_type"`
	TaxPercentage  *decimal.Decimal `json:"tax_percentage"`
	ShippingZoneID int64            `json:"shipping_zone_id"`
}

// Represents the result from the shipping_zones.json endpoint
type CountriesResource struct {
	Countries []Country `json:"countries"`
}

// List shipping zones
func (s *CountryServiceOp) List(options interface{}) ([]Country, error) {
	path := fmt.Sprintf("%s.json", countriesBasePath)
	resource := new(CountriesResource)
	err := s.client.Get(path, resource, options)
	return resource.Countries, err
}
