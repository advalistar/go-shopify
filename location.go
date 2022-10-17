package goshopify

import (
	"fmt"
	"reflect"
	"time"
)

const locationsBasePath = "locations"

// LocationService is an interface for interfacing with the location endpoints
// of the Shopify API.
// See: https://help.shopify.com/en/api/reference/inventory/location
type LocationService interface {
	// Retrieves a list of locations
	List(options interface{}) ([]Location, error)
	// Retrieves a single location by its ID
	Get(ID int64, options interface{}) (*Location, error)
	// Retrieves a count of locations
	Count(options interface{}) (int, error)
	GetOrderList() []string
}

type Location struct {
	ID                    int64     `json:"id"`
	Name                  string    `json:"name"`
	Address1              string    `json:"address1"`
	Address2              string    `json:"address2"`
	City                  string    `json:"city"`
	Zip                   string    `json:"zip"`
	Province              string    `json:"province"`
	Country               string    `json:"country"`
	Phone                 string    `json:"phone"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	CountryCode           string    `json:"country_code"`
	CountryName           string    `json:"country_name"`
	ProvinceCode          string    `json:"province_code"`
	Legacy                bool      `json:"legacy"`
	Active                bool      `json:"active"`
	AdminGraphqlAPIID     string    `json:"admin_graphql_api_id"`
	LocalizedCountryName  string    `json:"localized_country_name"`
	LocalizedProvinceName string    `json:"localized_province_name"`
}

// LocationServiceOp handles communication with the location related methods of
// the Shopify API.
type LocationServiceOp struct {
	client *Client
}

func (s *LocationServiceOp) List(options interface{}) ([]Location, error) {
	path := fmt.Sprintf("%s.json", locationsBasePath)
	resource := new(LocationsResource)
	err := s.client.Get(path, resource, options)
	return resource.Locations, err
}

func (s *LocationServiceOp) Get(ID int64, options interface{}) (*Location, error) {
	path := fmt.Sprintf("%s/%d.json", locationsBasePath, ID)
	resource := new(LocationResource)
	err := s.client.Get(path, resource, options)
	return resource.Location, err
}

func (s *LocationServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", locationsBasePath)
	return s.client.Count(path, options)
}

// Represents the result from the locations/X.json endpoint
type LocationResource struct {
	Location *Location `json:"location"`
}

// Represents the result from the locations.json endpoint
type LocationsResource struct {
	Locations []Location `json:"locations"`
}

func (s *LocationServiceOp) GetOrderList() []string {
	str := new(Location)

	var orderList []string
	for i := 0; i < reflect.TypeOf(&str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
