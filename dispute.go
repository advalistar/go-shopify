package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const disputesBasePath = "shopify_payments/disputes"

type DisputeService interface {
	List(interface{}) ([]Dispute, error)
	ListWithPagination(interface{}) ([]Dispute, *Pagination, error)
	GetOrderList() []string
}

type DisputeServiceOp struct {
	client *Client
}

// Dispute represents a Shopify blog
type Dispute struct {
	ID                int64            `json:"id"`
	OrderID           int64            `json:"order_id"`
	Type              string           `json:"type"`
	Amount            *decimal.Decimal `json:"amount"`
	Currency          string           `json:"currency"`
	Reason            string           `json:"reason"`
	NetworkReasonCode int64            `json:"network_reason_code"`
	Status            string           `json:"status"`
	EvidenceDueBy     *time.Time       `json:"evidence_due_by"`
	EvidenceSentOn    *time.Time       `json:"evidence_sent_on"`
	FinalizedOn       *time.Time       `json:"finalized_on"`
	InitiatedAt       *time.Time       `json:"initiated_at"`
}

// DisputesResource is the result from the disputes.json endpoint
type DisputesResource struct {
	Disputes []Dispute `json:"disputes"`
}

// List all disputes
func (s *DisputeServiceOp) List(options interface{}) ([]Dispute, error) {
	path := fmt.Sprintf("%s.json", disputesBasePath)
	resource := new(DisputesResource)
	err := s.client.Get(path, resource, options)
	return resource.Disputes, err
}

func (s *DisputeServiceOp) ListWithPagination(options interface{}) ([]Dispute, *Pagination, error) {
	path := fmt.Sprintf("%s.json", disputesBasePath)
	resource := new(DisputesResource)
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

	return resource.Disputes, pagination, nil
}

func (s *DisputeServiceOp) GetOrderList() []string {
	str := new(Dispute)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
