package goshopify

import (
	"fmt"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const (
	checkoutsBasePath             = "checkouts/%s"
	checkoutsShoppingRateBasePath = "checkouts/%s/shipping_rates"
)

type CheckoutService interface {
	Get(string, interface{}) (*Checkout, error)
	ShoppingRateList(string, interface{}) ([]ShoppingRate, error)
	GetOrderList() []string
}

type CheckoutServiceOp struct {
	client *Client
}

type Checkout struct {
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

type ShoppingRate struct {
	ID                     int64                 `json:"id"`
	Price                  string                `json:"price"`
	Title                  string                `json:"title"`
	Checkout               *ShoppingRateCheckout `json:"checkout"`
	PhoneRequired          bool                  `json:"phone_required"`
	DeliveryRange          interface{}           `json:"delivery_range"`
	EstimatedTimeInTransit interface{}           `json:"estimated_time_in_transit"`
	Handle                 string                `json:"handle"`
}

type ShoppingRateCheckout struct {
	TotalTax      *decimal.Decimal `json:"total_tax"`
	TotalPrice    *decimal.Decimal `json:"total_price"`
	SubtotalPrice *decimal.Decimal `json:"subtotal_price"`
}

type CheckoutsShoppingRateResource struct {
	ShoppingRates []ShoppingRate `json:"shipping_rates"`
}

type CheckoutResource struct {
	Checkout *Checkout `json:"checkout"`
}

func (s *CheckoutServiceOp) Get(token string, options interface{}) (*Checkout, error) {
	path := fmt.Sprintf(checkoutsBasePath+".json", token)
	resource := &CheckoutResource{}
	return resource.Checkout, s.client.Get(path, resource, options)
}

func (s *CheckoutServiceOp) ShoppingRateList(token string, options interface{}) ([]ShoppingRate, error) {
	path := fmt.Sprintf(checkoutsShoppingRateBasePath+".json", token)
	resource := &CheckoutsShoppingRateResource{}
	return resource.ShoppingRates, s.client.Get(path, resource, options)
}

func (s *CheckoutServiceOp) GetOrderList() []string {
	str := new(Checkout)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
