package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const refundBasePath = "orders/%d/refund"

type RefundService interface {
	List(int64, interface{}) ([]Refund, error)
	ListWithPagination(int64, interface{}) ([]Refund, *Pagination, error)
	GetOrderList() []string
}

type RefundServiceOp struct {
	client *Client
}

type Refund struct {
	ID                int64             `json:"id"`
	AdminGraphqlAPIID string            `json:"admin_graphql_api_id"`
	CreatedAt         *time.Time        `json:"created_at"`
	Note              string            `json:"note"`
	OrderID           int64             `json:"order_id"`
	ProcessedAt       *time.Time        `json:"processed_at"`
	Restock           bool              `json:"restock"`
	UserID            int64             `json:"user_id"`
	OrderAdjustments  *OrderAdjustments `json:"order_adjustments"`
	Transactions      []*Transaction    `json:"transactions"`
	RefundLineItems   []*RefundLineItem `json:"refund_line_items"`
}

type OrderAdjustments struct {
	ID       int64 `json:"id"`
	OrderID  int64 `json:"order_id"`
	RefundID int64 `json:"refund_id"`
}

type RefundLineItem struct {
	ID          int64            `json:"id"`
	LineItemID  int64            `json:"line_item_id"`
	LocationID  int64            `json:"location_id"`
	Quantity    int              `json:"quantity"`
	RestockType string           `json:"restock_type"`
	Subtotal    *decimal.Decimal `json:"subtotal"`
	SubtotalSet *AmountSet       `json:"subtotal_set"`
	TotalTax    *decimal.Decimal `json:"total_tax"`
	TotalTaxSet *AmountSet       `json:"total_tax_set"`
	LineItem    *LineItem        `json:"line_item"`
}

type RefundsResource struct {
	Refunds []Refund `json:"refunds"`
}

// List of discount codes
func (s *RefundServiceOp) List(orderID int64, options interface{}) ([]Refund, error) {
	path := fmt.Sprintf(refundBasePath+".json", orderID)
	resource := new(RefundsResource)
	err := s.client.Get(path, resource, options)
	return resource.Refunds, err
}

func (s *RefundServiceOp) ListWithPagination(orderID int64, options interface{}) ([]Refund, *Pagination, error) {
	path := fmt.Sprintf(refundBasePath+".json", orderID)
	resource := new(RefundsResource)
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

	return resource.Refunds, pagination, nil
}

func (s *RefundServiceOp) GetOrderList() []string {
	str := new(Refund)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
