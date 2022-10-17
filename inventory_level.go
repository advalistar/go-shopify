package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const inventoryLevelBasePath = "inventory_levels"

type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	ListWithPagination(interface{}) ([]InventoryLevel, *Pagination, error)
	GetOrderList() []string
}

type InventoryLevelServiceOp struct {
	client *Client
}

type InventoryLevel struct {
	ID           int64      `json:"id"`
	SubjectID    int64      `json:"subject_id"`
	CreatedAt    *time.Time `json:"created_at"`
	Subject_type string     `json:"subject_type"`
	Verb         string     `json:"verb"`
	Arguments    []string   `json:"arguments"`
	Body         string     `json:"body"`
	Message      string     `json:"message"`
	Author       string     `json:"author"`
	Description  string     `json:"description"`
	Path         string     `json:"path"`
}

// InventoryLevelsResource represents the result from the
// admin/inventoryLevels.json endpoint.
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

type InventoryLevelsOptions struct {
	// PageInfo is used with new pagination search.
	PageInfo string `url:"page_info,omitempty"`

	Page             int       `url:"page,omitempty"`
	Limit            int       `url:"limit,omitempty"`
	SinceID          int64     `url:"since_id,omitempty"`
	CreatedAtMin     time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax     time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin     time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax     time.Time `url:"updated_at_max,omitempty"`
	Order            string    `url:"order,omitempty"`
	Fields           string    `url:"fields,omitempty"`
	Vendor           string    `url:"vendor,omitempty"`
	IDs              []int64   `url:"ids,omitempty,comma"`
	InventoryItemIDs []int64   `url:"inventory_item_ids,omitempty,comma"`
	LocationIDs      []int64   `url:"location_ids,omitempty,comma"`
}

// List gets all application inventoryLevels.
func (s InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelBasePath)
	resource := &InventoryLevelsResource{}
	return resource.InventoryLevels, s.client.Get(path, resource, options)
}

func (s *InventoryLevelServiceOp) ListWithPagination(options interface{}) ([]InventoryLevel, *Pagination, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelBasePath)
	resource := new(InventoryLevelsResource)
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

	return resource.InventoryLevels, pagination, nil
}

func (s *InventoryLevelServiceOp) GetOrderList() []string {
	str := new(InventoryLevel)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
