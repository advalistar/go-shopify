package goshopify

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const payoutsBasePath = "shopify_payments/payouts"

type PayoutService interface {
	List(interface{}) ([]Payout, error)
	ListWithPagination(interface{}) ([]Payout, *Pagination, error)
}

type PayoutServiceOp struct {
	client *Client
}

// Payout represents a Shopify blog
type Payout struct {
	ID       int64            `json:"id"`
	Status   string           `json:"status"`
	Date     *time.Time       `json:"date"`
	Currency string           `json:"currency"`
	Amount   *decimal.Decimal `json:"amount"`
	Summary  *Summary         `json:"summary"`
}

type Summary struct {
	AdjustmentsFeeAmount      *decimal.Decimal `json:"adjustments_fee_amount"`
	AdjustmentsGrossAmount    *decimal.Decimal `json:"adjustments_gross_amount"`
	ChargesFeeAmount          *decimal.Decimal `json:"charges_fee_amount"`
	ChargesGrossAmount        *decimal.Decimal `json:"charges_gross_amount"`
	RefundsFeeAmount          *decimal.Decimal `json:"refunds_fee_amount"`
	RefundsGrossAmount        *decimal.Decimal `json:"refunds_gross_amount"`
	ReservedFundsFeeAmount    *decimal.Decimal `json:"reserved_funds_fee_amount"`
	ReservedFundsGrossAmount  *decimal.Decimal `json:"reserved_funds_gross_amount"`
	RetriedPayoutsFeeAmount   *decimal.Decimal `json:"retried_payouts_fee_amount"`
	RetriedPayoutsGrossAmount *decimal.Decimal `json:"retried_payouts_gross_amount"`
}

// PayoutsResource is the result from the payouts.json endpoint
type PayoutsResource struct {
	Payouts []Payout `json:"payouts"`
}

// List all payouts
func (s *PayoutServiceOp) List(options interface{}) ([]Payout, error) {
	path := fmt.Sprintf("%s.json", payoutsBasePath)
	resource := new(PayoutsResource)
	err := s.client.Get(path, resource, options)
	return resource.Payouts, err
}

func (s *PayoutServiceOp) ListWithPagination(options interface{}) ([]Payout, *Pagination, error) {
	path := fmt.Sprintf("%s.json", payoutsBasePath)
	resource := new(PayoutsResource)
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

	return resource.Payouts, pagination, nil
}
