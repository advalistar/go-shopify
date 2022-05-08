package goshopify

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const inventoryItemsBasePath = "inventory_items"

// InventoryItemService is an interface for interacting with the
// inventory items endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventoryitem
type InventoryItemService interface {
	List(interface{}) ([]InventoryItem, error)
	ListWithPagination(interface{}) ([]InventoryItem, *Pagination, error)
	Get(int64, interface{}) (*InventoryItem, error)
	Update(InventoryItem) (*InventoryItem, error)
}

// InventoryItemServiceOp is the default implementation of the InventoryItemService interface
type InventoryItemServiceOp struct {
	client *Client
}

// InventoryItem represents a Shopify inventory item
type InventoryItem struct {
	ID                           int64                           `json:"id"`
	SKU                          string                          `json:"sku"`
	CreatedAt                    *time.Time                      `json:"created_at"`
	UpdatedAt                    *time.Time                      `json:"updated_at"`
	RequiresShipping             bool                            `json:"requires_shipping"`
	Cost                         *decimal.Decimal                `json:"cost"`
	CountryCodeOfOrigin          string                          `json:"country_code_of_origin"`
	ProvinceCodeOfOrigin         string                          `json:"province_code_of_origin"`
	HarmonizedSystemCode         string                          `json:"harmonized_system_code"`
	Tracked                      bool                            `json:"tracked"`
	CountryHarmonizedSystemCodes []*CountryHarmonizedSystemCodes `json:"country_harmonized_system_codes"`
	AdminGraphqlAPIID            string                          `json:"admin_graphql_api_id"`
}

type CountryHarmonizedSystemCodes struct {
	HarmonizedSystemCode string `json:"harmonized_system_code"`
	CountryCode          string `json:"country_code"`
}

// InventoryItemResource is used for handling single item requests and responses
type InventoryItemResource struct {
	InventoryItem *InventoryItem `json:"inventory_item"`
}

// InventoryItemsResource is used for handling multiple item responsees
type InventoryItemsResource struct {
	InventoryItems []InventoryItem `json:"inventory_items"`
}

// List inventory items
func (s *InventoryItemServiceOp) List(options interface{}) ([]InventoryItem, error) {
	path := fmt.Sprintf("%s.json", inventoryItemsBasePath)
	resource := new(InventoryItemsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItems, err
}

func (s *InventoryItemServiceOp) ListWithPagination(options interface{}) ([]InventoryItem, *Pagination, error) {
	path := fmt.Sprintf("%s.json", inventoryItemsBasePath)
	resource := new(InventoryItemsResource)
	headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")

	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.InventoryItems, pagination, nil
}

// Get a inventory item
func (s *InventoryItemServiceOp) Get(id int64, options interface{}) (*InventoryItem, error) {
	path := fmt.Sprintf("%s/%d.json", inventoryItemsBasePath, id)
	resource := new(InventoryItemResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItem, err
}

// Update a inventory item
func (s *InventoryItemServiceOp) Update(item InventoryItem) (*InventoryItem, error) {
	path := fmt.Sprintf("%s/%d.json", inventoryItemsBasePath, item.ID)
	wrappedData := InventoryItemResource{InventoryItem: &item}
	resource := new(InventoryItemResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.InventoryItem, err
}
