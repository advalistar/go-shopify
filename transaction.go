package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

// TransactionService is an interface for interfacing with the transactions endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/transaction
type TransactionService interface {
	List(int64, interface{}) ([]Transaction, error)
	ListWithPagination(int64, interface{}) ([]Transaction, *Pagination, error)
	Count(int64, interface{}) (int, error)
	Get(int64, int64, interface{}) (*Transaction, error)
	Create(int64, Transaction) (*Transaction, error)
	GetOrderList() []string
}

// TransactionServiceOp handles communication with the transaction related methods of the
// Shopify API.
type TransactionServiceOp struct {
	client *Client
}

type Transaction struct {
	ID                int64            `json:"id"`
	AdminGraphqlAPIID string           `json:"admin_graphql_api_id"`
	Amount            *decimal.Decimal `json:"amount"`
	Authorization     string           `json:"authorization"`
	CreatedAt         *time.Time       `json:"created_at"`
	Currency          string           `json:"currency"`
	DeviceID          *int64           `json:"device_id"`
	ErrorCode         string           `json:"error_code"`
	Gateway           string           `json:"gateway"`
	Kind              string           `json:"kind"`
	LocationID        *int64           `json:"location_id"`
	Message           string           `json:"message"`
	OrderID           int64            `json:"order_id"`
	ParentID          *int64           `json:"parent_id"`
	ProcessedAt       *time.Time       `json:"processed_at"`
	Receipt           *Receipt         `json:"receipt"`
	SourceName        string           `json:"source_name"`
	Status            string           `json:"status"`
	Test              bool             `json:"test"`
	UserID            *int64           `json:"user_id"`

	PaymentDetails *PaymentDetails `json:"payment_details,omitempty"`
}

// TransactionResource represents the result from the orders/X/transactions/Y.json endpoint
type TransactionResource struct {
	Transaction *Transaction `json:"transaction"`
}

// TransactionsResource represents the result from the orders/X/transactions.json endpoint
type TransactionsResource struct {
	Transactions []Transaction `json:"transactions"`
}

// List transactions
func (s *TransactionServiceOp) List(orderID int64, options interface{}) ([]Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	resource := new(TransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Transactions, err
}

func (s *TransactionServiceOp) ListWithPagination(orderID int64, options interface{}) ([]Transaction, *Pagination, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	resource := new(TransactionsResource)
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

	return resource.Transactions, pagination, nil
}

// Count transactions
func (s *TransactionServiceOp) Count(orderID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/transactions/count.json", ordersBasePath, orderID)
	return s.client.Count(path, options)
}

// Get individual transaction
func (s *TransactionServiceOp) Get(orderID int64, transactionID int64, options interface{}) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions/%d.json", ordersBasePath, orderID, transactionID)
	resource := new(TransactionResource)
	err := s.client.Get(path, resource, options)
	return resource.Transaction, err
}

// Create a new transaction
func (s *TransactionServiceOp) Create(orderID int64, transaction Transaction) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	wrappedData := TransactionResource{Transaction: &transaction}
	resource := new(TransactionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Transaction, err
}

func (s *TransactionServiceOp) GetOrderList() []string {
	str := new(Transaction)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
