package goshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const (
	ordersBasePath     = "orders"
	ordersResourceName = "orders"
)

// OrderService is an interface for interfacing with the orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/order
type OrderService interface {
	List(interface{}) ([]Order, error)
	ListWithPagination(interface{}) ([]Order, *Pagination, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*Order, error)
	Create(Order) (*Order, error)
	Update(Order) (*Order, error)
	Cancel(int64, interface{}) (*Order, error)
	Close(int64) (*Order, error)
	Open(int64) (*Order, error)

	// MetafieldsService used for Order resource to communicate with Metafields resource
	MetafieldsService

	// FulfillmentsService used for Order resource to communicate with Fulfillments resource
	FulfillmentsService
}

// OrderServiceOp handles communication with the order related methods of the
// Shopify API.
type OrderServiceOp struct {
	client *Client
}

// A struct for all available order count options
type OrderCountOptions struct {
	Page              int       `url:"page,omitempty"`
	Limit             int       `url:"limit,omitempty"`
	SinceID           int64     `url:"since_id,omitempty"`
	CreatedAtMin      time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax      time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin      time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax      time.Time `url:"updated_at_max,omitempty"`
	Order             string    `url:"order,omitempty"`
	Fields            string    `url:"fields,omitempty"`
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
}

// A struct for all available order list options.
// See: https://help.shopify.com/api/reference/order#index
type OrderListOptions struct {
	ListOptions
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
	ProcessedAtMin    time.Time `url:"processed_at_min,omitempty"`
	ProcessedAtMax    time.Time `url:"processed_at_max,omitempty"`
	Order             string    `url:"order,omitempty"`
}

// A struct of all available order cancel options.
// See: https://help.shopify.com/api/reference/order#index
type OrderCancelOptions struct {
	Amount   *decimal.Decimal `json:"amount,omitempty"`
	Currency string           `json:"currency,omitempty"`
	Restock  bool             `json:"restock,omitempty"`
	Reason   string           `json:"reason,omitempty"`
	Email    bool             `json:"email,omitempty"`
	Refund   *Refund          `json:"refund,omitempty"`
}

// Order represents a Shopify order
type Order struct {
	ID                     int64                   `json:"id"`
	AdminGraphqlAPIID      string                  `json:"admin_graphql_api_id"`
	AppID                  int                     `json:"app_id"`
	BrowserIp              string                  `json:"browser_ip"`
	BuyerAcceptsMarketing  bool                    `json:"buyer_accepts_marketing"`
	CancelReason           string                  `json:"cancel_reason"`
	CancelledAt            *time.Time              `json:"cancelled_at"`
	CartToken              string                  `json:"cart_token"`
	CheckoutID             int64                   `json:"checkout_id"`
	CheckoutToken          string                  `json:"checkout_token"`
	ClientDetails          *ClientDetails          `json:"client_details"`
	ClosedAt               *time.Time              `json:"closed_at"`
	Confirmed              bool                    `json:"confirmed"`
	ContactEmail           string                  `json:"contact_email"`
	CreatedAt              *time.Time              `json:"created_at"`
	Currency               string                  `json:"currency"`
	CustomerLocale         string                  `json:"customer_locale"`
	DeviceID               int64                   `json:"device_id"`
	DiscountCodes          []*DiscountCode         `json:"discount_codes"`
	Email                  string                  `json:"email"`
	FinancialStatus        string                  `json:"financial_status"`
	FulfillmentStatus      string                  `json:"fulfillment_status"`
	Gateway                string                  `json:"gateway"`
	LandingSite            string                  `json:"landing_site"`
	LandingSiteRef         string                  `json:"landing_site_ref"`
	LocationID             int64                   `json:"location_id"`
	Name                   string                  `json:"name"`
	Note                   string                  `json:"note"`
	NoteAttributes         []*NoteAttribute        `json:"note_attributes"`
	Number                 int                     `json:"number"`
	OrderNumber            int                     `json:"order_number"`
	OrderStatusUrl         string                  `json:"order_status_url"`
	PaymentGatewayNames    []string                `json:"payment_gateway_names"`
	Phone                  string                  `json:"phone"`
	PresentmentCurrency    string                  `json:"presentment_currency"`
	ProcessedAt            *time.Time              `json:"processed_at"`
	ProcessingMethod       string                  `json:"processing_method"`
	Reference              string                  `json:"reference"`
	ReferringSite          string                  `json:"referring_site"`
	SourceIdentifier       string                  `json:"source_identifier"`
	SourceName             string                  `json:"source_name"`
	SourceURL              string                  `json:"source_url"`
	SubtotalPrice          *decimal.Decimal        `json:"subtotal_price"`
	SubtotalPriceSet       *AmountSet              `json:"subtotal_price_set"`
	Tags                   string                  `json:"tags"`
	TaxLines               []*TaxLine              `json:"tax_lines"`
	TaxesIncluded          bool                    `json:"taxes_included"`
	Test                   bool                    `json:"test"`
	Token                  string                  `json:"token"`
	TotalDiscounts         *decimal.Decimal        `json:"total_discounts"`
	TotalDiscountsSet      *AmountSet              `json:"total_discounts_set"`
	TotalLineItemsPrice    *decimal.Decimal        `json:"total_line_items_price"`
	TotalLineItemsPriceSet *AmountSet              `json:"total_line_items_price_set"`
	TotalPrice             *decimal.Decimal        `json:"total_price"`
	TotalPriceSet          *AmountSet              `json:"total_price_set"`
	TotalPriceUSD          *decimal.Decimal        `json:"total_price_usd"`
	TotalShippingPriceSet  *AmountSet              `json:"total_shipping_price_set"`
	TotalTax               *decimal.Decimal        `json:"total_tax"`
	TotalTaxSet            *AmountSet              `json:"total_tax_set"`
	TotalTipReceived       *decimal.Decimal        `json:"total_tip_received"`
	TotalWeight            int                     `json:"total_weight"`
	UpdatedAt              *time.Time              `json:"updated_at"`
	UserID                 int64                   `json:"user_id"`
	BillingAddress         *Address                `json:"billing_address"`
	Customer               *Customer               `json:"customer"`
	DiscountApplications   []*DiscountApplications `json:"discount_applications"`
	Fulfillments           []*Fulfillment          `json:"fulfillments"`
	LineItems              []*LineItem             `json:"line_items"`
	PaymentDetails         *PaymentDetails         `json:"payment_details"`
	Refunds                []*Refund               `json:"refunds"`
	ShippingAddress        *Address                `json:"shipping_address"`
	ShippingLines          []*ShippingLines        `json:"shipping_lines"`
}

type Address struct {
	FirstName    string  `json:"first_name"`
	Address1     string  `json:"address1"`
	Phone        string  `json:"phone"`
	City         string  `json:"city"`
	Zip          string  `json:"zip"`
	Province     string  `json:"province"`
	Country      string  `json:"country"`
	LastName     string  `json:"last_name"`
	Address2     string  `json:"address2"`
	Company      string  `json:"company"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"country_code"`
	ProvinceCode string  `json:"province_code"`
	ID           int64   `json:"id"`
}

type DiscountCode struct {
	Amount *decimal.Decimal `json:"amount"`
	Code   string           `json:"code"`
	Type   string           `json:"type"`
}

type LineItem struct {
	ID                         int64                  `json:"id"`
	ProductID                  int64                  `json:"product_id"`
	VariantID                  int64                  `json:"variant_id"`
	Quantity                   int                    `json:"quantity"`
	Price                      *decimal.Decimal       `json:"price"`
	TotalDiscount              *decimal.Decimal       `json:"total_discount"`
	Title                      string                 `json:"title"`
	VariantTitle               string                 `json:"variant_title"`
	Name                       string                 `json:"name"`
	SKU                        string                 `json:"sku"`
	Vendor                     string                 `json:"vendor"`
	GiftCard                   bool                   `json:"gift_card"`
	Taxable                    bool                   `json:"taxable"`
	FulfillmentService         string                 `json:"fulfillment_service"`
	RequiresShipping           bool                   `json:"requires_shipping"`
	VariantInventoryManagement string                 `json:"variant_inventory_management"`
	PreTaxPrice                *decimal.Decimal       `json:"pre_tax_price"`
	Properties                 []*NoteAttribute       `json:"properties"`
	ProductExists              bool                   `json:"product_exists"`
	FulfillableQuantity        int                    `json:"fulfillable_quantity"`
	Grams                      int                    `json:"grams"`
	FulfillmentStatus          string                 `json:"fulfillment_status"`
	TaxLines                   []*TaxLine             `json:"tax_lines"`
	OriginLocation             *Address               `json:"origin_location"`
	DestinationLocation        *Address               `json:"destination_location"`
	AppliedDiscount            *AppliedDiscount       `json:"applied_discount"`
	DiscountAllocations        []*DiscountAllocations `json:"discount_allocations"`
}

type DiscountAllocations struct {
	Amount                   *decimal.Decimal `json:"amount"`
	AmountSet                *AmountSet       `json:"amount_set"`
	DiscountApplicationIndex int              `json:"discount_application_index"`
}

type AmountSet struct {
	ShopMoney        *AmountSetEntry `json:"shop_money"`
	PresentmentMoney *AmountSetEntry `json:"presentment_money"`
}

type AmountSetEntry struct {
	Amount       *decimal.Decimal `json:"amount"`
	CurrencyCode string           `json:"currency_code"`
}

type DiscountApplications struct {
	TargetType               string      `json:"target_type"`
	DiscountApplicationsType string      `json:"type"`
	Value                    interface{} `json:"value"`
	ValueType                string      `json:"value_type"`
	AllocationMethod         string      `json:"allocation_method"`
	TargetSelection          string      `json:"target_selection"`
	Code                     string      `json:"code"`
}

// UnmarshalJSON custom unmarsaller for LineItem required to mitigate some older orders having LineItem.Properies
// which are empty JSON objects rather than the expected array.
func (li *LineItem) UnmarshalJSON(data []byte) error {
	type alias LineItem
	aux := &struct {
		Properties json.RawMessage `json:"properties"`
		*alias
	}{alias: (*alias)(li)}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	if len(aux.Properties) == 0 {
		return nil
	} else if aux.Properties[0] == '[' { // if the first character is a '[' we unmarshal into an array
		var p []*NoteAttribute
		err = json.Unmarshal(aux.Properties, &p)
		if err != nil {
			return err
		}
		li.Properties = p
	} else { // else we unmarshal it into a struct
		var p *NoteAttribute
		err = json.Unmarshal(aux.Properties, &p)
		if err != nil {
			return err
		}
		if p.Name == "" && p.Value == nil { // if the struct is empty we set properties to nil
			li.Properties = nil
		} else {
			li.Properties = []*NoteAttribute{p} // else we set them to an array with the property nested
		}
	}

	return nil
}

type LineItemProperty struct {
	Message string `json:"message"`
}

type NoteAttribute struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// Represents the result from the orders/X.json endpoint
type OrderResource struct {
	Order *Order `json:"order"`
}

// Represents the result from the orders.json endpoint
type OrdersResource struct {
	Orders []Order `json:"orders"`
}

type PaymentDetails struct {
	AVSResultCode     string `json:"avs_result_code"`
	CreditCardBin     string `json:"credit_card_bin"`
	CVVResultCode     string `json:"cvv_result_code"`
	CreditCardNumber  string `json:"credit_card_number"`
	CreditCardCompany string `json:"credit_card_company"`
}

type ShippingLines struct {
	ID                            int64            `json:"id"`
	Title                         string           `json:"title"`
	Price                         *decimal.Decimal `json:"price"`
	Code                          string           `json:"code"`
	Source                        string           `json:"source"`
	Phone                         string           `json:"phone"`
	RequestedFulfillmentServiceID string           `json:"requested_fulfillment_service_id"`
	DeliveryCategory              string           `json:"delivery_category"`
	CarrierIdentifier             string           `json:"carrier_identifier"`
	TaxLines                      []*TaxLine       `json:"tax_lines"`
	Handle                        string           `json:"handle"`
}

// UnmarshalJSON custom unmarshaller for ShippingLines implemented to handle requested_fulfillment_service_id being
// returned as json numbers or json nulls instead of json strings
func (sl *ShippingLines) UnmarshalJSON(data []byte) error {
	type alias ShippingLines
	aux := &struct {
		*alias
		RequestedFulfillmentServiceID interface{} `json:"requested_fulfillment_service_id"`
	}{alias: (*alias)(sl)}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch aux.RequestedFulfillmentServiceID.(type) {
	case nil:
		sl.RequestedFulfillmentServiceID = ""
	default:
		sl.RequestedFulfillmentServiceID = fmt.Sprintf("%v", aux.RequestedFulfillmentServiceID)
	}

	return nil
}

type TaxLine struct {
	Price    *decimal.Decimal `json:"price"`
	Rate     *decimal.Decimal `json:"rate"`
	Title    string           `json:"title"`
	PriceSet *AmountSet       `json:"price_set"`
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

type ClientDetails struct {
	AcceptLanguage string `json:"accept_language"`
	BrowserHeight  int    `json:"browser_height"`
	BrowserIp      string `json:"browser_ip"`
	BrowserWidth   int    `json:"browser_width"`
	SessionHash    string `json:"session_hash"`
	UserAgent      string `json:"user_agent"`
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

// List orders
func (s *OrderServiceOp) List(options interface{}) ([]Order, error) {
	orders, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderServiceOp) ListWithPagination(options interface{}) ([]Order, *Pagination, error) {
	path := fmt.Sprintf("%s.json", ordersBasePath)
	resource := new(OrdersResource)
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

	return resource.Orders, pagination, nil
}

// Count orders
func (s *OrderServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", ordersBasePath)
	return s.client.Count(path, options)
}

// Get individual order
func (s *OrderServiceOp) Get(orderID int64, options interface{}) (*Order, error) {
	path := fmt.Sprintf("%s/%d.json", ordersBasePath, orderID)
	resource := new(OrderResource)
	err := s.client.Get(path, resource, options)
	return resource.Order, err
}

// Create order
func (s *OrderServiceOp) Create(order Order) (*Order, error) {
	path := fmt.Sprintf("%s.json", ordersBasePath)
	wrappedData := OrderResource{Order: &order}
	resource := new(OrderResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Order, err
}

// Update order
func (s *OrderServiceOp) Update(order Order) (*Order, error) {
	path := fmt.Sprintf("%s/%d.json", ordersBasePath, order.ID)
	wrappedData := OrderResource{Order: &order}
	resource := new(OrderResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Order, err
}

// Cancel order
func (s *OrderServiceOp) Cancel(orderID int64, options interface{}) (*Order, error) {
	path := fmt.Sprintf("%s/%d/cancel.json", ordersBasePath, orderID)
	resource := new(OrderResource)
	err := s.client.Post(path, options, resource)
	return resource.Order, err
}

// Close order
func (s *OrderServiceOp) Close(orderID int64) (*Order, error) {
	path := fmt.Sprintf("%s/%d/close.json", ordersBasePath, orderID)
	resource := new(OrderResource)
	err := s.client.Post(path, nil, resource)
	return resource.Order, err
}

// Open order
func (s *OrderServiceOp) Open(orderID int64) (*Order, error) {
	path := fmt.Sprintf("%s/%d/open.json", ordersBasePath, orderID)
	resource := new(OrderResource)
	err := s.client.Post(path, nil, resource)
	return resource.Order, err
}

// List metafields for an order
func (s *OrderServiceOp) ListMetafields(orderID int64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldService.List(options)
}

// Count metafields for an order
func (s *OrderServiceOp) CountMetafields(orderID int64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldService.Count(options)
}

// Get individual metafield for an order
func (s *OrderServiceOp) GetMetafield(orderID int64, metafieldID int64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldService.Get(metafieldID, options)
}

// Create a new metafield for an order
func (s *OrderServiceOp) CreateMetafield(orderID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldService.Create(metafield)
}

// Update an existing metafield for an order
func (s *OrderServiceOp) UpdateMetafield(orderID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldService.Update(metafield)
}

// Delete an existing metafield for an order
func (s *OrderServiceOp) DeleteMetafield(orderID int64, metafieldID int64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldService.Delete(metafieldID)
}

// List fulfillments for an order
func (s *OrderServiceOp) ListFulfillments(orderID int64, options interface{}) ([]Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.List(options)
}

// Count fulfillments for an order
func (s *OrderServiceOp) CountFulfillments(orderID int64, options interface{}) (int, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Count(options)
}

// Get individual fulfillment for an order
func (s *OrderServiceOp) GetFulfillment(orderID int64, fulfillmentID int64, options interface{}) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Get(fulfillmentID, options)
}

// Create a new fulfillment for an order
func (s *OrderServiceOp) CreateFulfillment(orderID int64, fulfillment Fulfillment) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Create(fulfillment)
}

// Update an existing fulfillment for an order
func (s *OrderServiceOp) UpdateFulfillment(orderID int64, fulfillment Fulfillment) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Update(fulfillment)
}

// Complete an existing fulfillment for an order
func (s *OrderServiceOp) CompleteFulfillment(orderID int64, fulfillmentID int64) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Complete(fulfillmentID)
}

// Transition an existing fulfillment for an order
func (s *OrderServiceOp) TransitionFulfillment(orderID int64, fulfillmentID int64) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Transition(fulfillmentID)
}

// Cancel an existing fulfillment for an order
func (s *OrderServiceOp) CancelFulfillment(orderID int64, fulfillmentID int64) (*Fulfillment, error) {
	fulfillmentService := &FulfillmentServiceOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentService.Cancel(fulfillmentID)
}
