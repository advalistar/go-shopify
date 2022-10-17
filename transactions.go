package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const transactionsBasePath = "shopify_payments/balance/transactions"

type ShopifyPaymentsTransactionsService interface {
	List(interface{}) ([]ShopifyPaymentsTransactions, error)
	ListWithPagination(interface{}) ([]ShopifyPaymentsTransactions, *Pagination, error)
	GetOrderList() []string
}

type ShopifyPaymentsTransactionsServiceOp struct {
	client *Client
}

type ShopifyPaymentsTransactions struct {
	ID                       int64            `json:"id"`
	Type                     string           `json:"type"`
	Test                     bool             `json:"test"`
	PayoutID                 int64            `json:"payout_id"`
	PayoutStatus             string           `json:"payout_status"`
	Currency                 string           `json:"currency"`
	Amount                   *decimal.Decimal `json:"amount"`
	Fee                      *decimal.Decimal `json:"fee"`
	Net                      *decimal.Decimal `json:"net"`
	SourceID                 int64            `json:"source_id"`
	SourceType               string           `json:"source_type"`
	SourceOrderID            int64            `json:"source_order_id"`
	SourceOrderTransactionID int64            `json:"source_order_transaction_id"`
	ProcessedAt              *time.Time       `json:"processed_at"`
}

type ShopifyPaymentsTransactionsResource struct {
	ShopifyPaymentsTransactions []ShopifyPaymentsTransactions `json:"transactions"`
}

// List all transactions
func (s *ShopifyPaymentsTransactionsServiceOp) List(options interface{}) ([]ShopifyPaymentsTransactions, error) {
	path := fmt.Sprintf("%s.json", transactionsBasePath)
	resource := new(ShopifyPaymentsTransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.ShopifyPaymentsTransactions, err
}

func (s *ShopifyPaymentsTransactionsServiceOp) ListWithPagination(options interface{}) ([]ShopifyPaymentsTransactions, *Pagination, error) {
	path := fmt.Sprintf("%s.json", transactionsBasePath)
	resource := new(ShopifyPaymentsTransactionsResource)
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

	return resource.ShopifyPaymentsTransactions, pagination, nil
}

func (s *ShopifyPaymentsTransactionsServiceOp) GetOrderList() []string {
	str := new(ShopifyPaymentsTransactions)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
