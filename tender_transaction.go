package goshopify

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const tenderTransactionsBasePath = "tender_transactions"

type TenderTransactionService interface {
	List(interface{}) ([]TenderTransaction, error)
	ListWithPagination(interface{}) ([]TenderTransaction, *Pagination, error)
}

type TenderTransactionServiceOp struct {
	client *Client
}

// TenderTransaction represents a Shopify tenderTransaction
type TenderTransaction struct {
	ID              int64            `json:"id"`
	OrderID         int64            `json:"order_id"`
	Amount          *decimal.Decimal `json:"amount"`
	Currency        string           `json:"currency"`
	UserID          *int64           `json:"user_id"`
	Test            bool             `json:"test"`
	ProcessedAt     *time.Time       `json:"processed_at"`
	RemoteReference string           `json:"remote_reference"`
	PaymentDetails  *PaymentDetails  `json:"payment_details"`
	PaymentMethod   string           `json:"payment_method"`
}

type TenderTransactionResource struct {
	TenderTransaction *TenderTransaction `json:"tender_transaction"`
}

type TenderTransactionsResource struct {
	TenderTransactions []TenderTransaction `json:"tender_transactions"`
}

// List tenderTransactions
func (s *TenderTransactionServiceOp) List(options interface{}) ([]TenderTransaction, error) {
	path := fmt.Sprintf("%s.json", tenderTransactionsBasePath)
	resource := new(TenderTransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.TenderTransactions, err
}

func (s *TenderTransactionServiceOp) ListWithPagination(options interface{}) ([]TenderTransaction, *Pagination, error) {
	path := fmt.Sprintf("%s.json", tenderTransactionsBasePath)
	resource := new(TenderTransactionsResource)
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

	return resource.TenderTransactions, pagination, nil
}
