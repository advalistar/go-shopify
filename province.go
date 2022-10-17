package goshopify

import (
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
)

const provincesBasePath = "countries/%d/provinces"

type ProvinceService interface {
	List(int64, interface{}) ([]Province, error)
	GetOrderList() []string
}

type ProvinceServiceOp struct {
	client *Client
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

type ProvincesResource struct {
	Provinces []Province `json:"provinces"`
}

// List of discount codes
func (s *ProvinceServiceOp) List(countryID int64, options interface{}) ([]Province, error) {
	path := fmt.Sprintf(provincesBasePath+".json", countryID)
	resource := new(ProvincesResource)
	err := s.client.Get(path, resource, options)
	return resource.Provinces, err
}

func (s *ProvinceServiceOp) GetOrderList() []string {
	str := new(Province)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
