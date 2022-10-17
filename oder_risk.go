package goshopify

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/shopspring/decimal"
)

const orderRiskBasePath = "orders/%d/risks"

type OrderRiskService interface {
	List(int64) ([]Risk, error)
	ListWithPagination(int64, interface{}) ([]Risk, *Pagination, error)
	GetOrderList() []string
}

type OrderRiskServiceOp struct {
	client *Client
}

type Risk struct {
	ID              int64            `json:"id"`
	OrderID         int64            `json:"order_id"`
	CheckoutID      int64            `json:"checkout_id"`
	Source          string           `json:"source"`
	Score           *decimal.Decimal `json:"score"`
	Recommendation  string           `json:"recommendation"`
	Display         bool             `json:"display"`
	CauseCancel     bool             `json:"cause_cancel"`
	Message         string           `json:"message"`
	MerchantMessage string           `json:"merchant_message"`
}

type RisksResource struct {
	Risks []Risk `json:"risks"`
}

// List of discount codes
func (s *OrderRiskServiceOp) List(orderID int64) ([]Risk, error) {
	path := fmt.Sprintf(orderRiskBasePath+".json", orderID)
	resource := new(RisksResource)
	err := s.client.Get(path, resource, nil)
	return resource.Risks, err
}

func (s *OrderRiskServiceOp) ListWithPagination(orderID int64, options interface{}) ([]Risk, *Pagination, error) {
	path := fmt.Sprintf(orderRiskBasePath+".json", orderID)
	resource := new(RisksResource)
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

	return resource.Risks, pagination, nil
}

func (s *OrderRiskServiceOp) GetOrderList() []string {
	str := new(Risk)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
