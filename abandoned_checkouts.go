package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	checkoutsBasePath = "checkouts"
)

// CheckoutsService is an interface for interfacing with the checkouts endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/order
type CheckoutsService interface {
	List(interface{}) ([]Checkouts, error)
}

// CheckoutsServiceOp handles communication with the order related methods of the
// Shopify API.
type CheckoutsServiceOp struct {
	client *Client
}

// Checkouts represents a Shopify order
type Checkouts struct {
	ID                       int64            `json:"id"`
	Token                    string           `json:"token"`
	CartToken                string           `json:"cart_token"`
	Email                    string           `json:"email"`
	Gateway                  string           `json:"gateway"`
	BuyerAcceptsMarketing    bool             `json:"buyer_accepts_marketing"`
	CreatedAt                *time.Time       `json:"created_at"`
	UpdatedAt                *time.Time       `json:"updated_at"`
	LandingSite              string           `json:"landing_site"`
	Note                     string           `json:"note"`
	NoteAttributes           []*NoteAttribute `json:"note_attributes"`
	ReferringSite            string           `json:"referring_site"`
	ShippingLines            []*ShippingLines `json:"shipping_lines"`
	TaxesIncluded            bool             `json:"taxes_included"`
	TotalWeight              int              `json:"total_weight"`
	Currency                 string           `json:"currency"`
	CompletedAt              *time.Time       `json:"completed_at"`
	ClosedAt                 *time.Time       `json:"closed_at"`
	UserID                   int64            `json:"user_id"`
	LocationID               int64            `json:"location_id"`
	SourceIdentifier         string           `json:"source_identifier"`
	SourceURL                string           `json:"source_url"`
	DeviceID                 int64            `json:"device_id"`
	Phone                    string           `json:"phone"`
	CustomerLocale           string           `json:"customer_locale"`
	LineItems                []*LineItem      `json:"line_items"`
	Name                     string           `json:"name"`
	Source                   string           `json:"source"`
	AbandonedCheckoutURL     string           `json:"abandoned_checkout_url"`
	DiscountCodes            []*DiscountCode  `json:"discount_codes"`
	TaxLines                 []*TaxLine       `json:"tax_lines"`
	SourceName               string           `json:"source_name"`
	PresentmentCurrency      string           `json:"presentment_currency"`
	BuyerAcceptsSmsMarketing bool             `json:"buyer_accepts_sms_marketing"`
	SmsMarketingPhone        string           `json:"sms_marketing_phone"`
	TotalDiscounts           *decimal.Decimal `json:"total_discounts"`
	TotalLineItemsPrice      *decimal.Decimal `json:"total_line_items_price"`
	TotalPrice               *decimal.Decimal `json:"total_price"`
	TotalTax                 *decimal.Decimal `json:"total_tax"`
	SubtotalPrice            *decimal.Decimal `json:"subtotal_price"`
	TotalDuties              *decimal.Decimal `json:"total_duties"`
	BillingAddress           *Address         `json:"billing_address"`
	ShippingAddress          *Address         `json:"shipping_address"`
	Customer                 *Customer        `json:"customer"`
}

// Represents the result from the checkouts.json endpoint
type CheckoutsResource struct {
	Checkouts []Checkouts `json:"checkouts"`
}

// List checkouts
func (s *CheckoutsServiceOp) List(options interface{}) ([]Checkouts, error) {
	path := fmt.Sprintf("%s.json", checkoutsBasePath)
	resource := &CheckoutsResource{}
	return resource.Checkouts, s.client.Get(path, resource, options)
}
