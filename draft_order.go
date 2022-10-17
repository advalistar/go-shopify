package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const (
	draftOrdersBasePath     = "draft_orders"
	draftOrdersResourceName = "draft_orders"
)

// DraftOrderService is an interface for interfacing with the draft orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/orders/draftorder
type DraftOrderService interface {
	List(interface{}) ([]DraftOrder, error)
	ListWithPagination(interface{}) ([]DraftOrder, *Pagination, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*DraftOrder, error)
	Create(DraftOrder) (*DraftOrder, error)
	Update(DraftOrder) (*DraftOrder, error)
	Delete(int64) error
	Invoice(int64, DraftOrderInvoice) (*DraftOrderInvoice, error)
	Complete(int64, bool) (*DraftOrder, error)
	GetOrderList() []string

	// MetafieldsService used for DrafT Order resource to communicate with Metafields resource
	MetafieldsService
}

// DraftOrderServiceOp handles communication with the draft order related methods of the
// Shopify API.
type DraftOrderServiceOp struct {
	client *Client
}

// DraftOrder represents a shopify draft order
type DraftOrder struct {
	ID                int64            `json:"id"`
	Note              string           `json:"note"`
	Email             string           `json:"email"`
	TaxesIncluded     bool             `json:"taxes_included"`
	Currency          string           `json:"currency"`
	InvoiceSentAt     *time.Time       `json:"invoice_sent_at"`
	CreatedAt         *time.Time       `json:"created_at"`
	UpdatedAt         *time.Time       `json:"updated_at"`
	TaxExempt         bool             `json:"tax_exempt"`
	CompletedAt       *time.Time       `json:"completed_at"`
	Name              string           `json:"name"`
	Status            string           `json:"status"`
	LineItems         []*LineItem      `json:"line_items"`
	ShippingAddress   *Address         `json:"shipping_address"`
	BillingAddress    *Address         `json:"billing_address"`
	InvoiceURL        string           `json:"invoice_url"`
	AppliedDiscount   *AppliedDiscount `json:"applied_discount"`
	OrderID           int64            `json:"order_id"`
	ShippingLine      *ShippingLines   `json:"shipping_line"`
	TaxLines          []*TaxLine       `json:"tax_lines"`
	Tags              string           `json:"tags"`
	NoteAttributes    []*NoteAttribute `json:"note_attribute"`
	TotalPrice        string           `json:"total_price"`
	SubtotalPrice     *decimal.Decimal `json:"subtotal_price"`
	TotalTax          string           `json:"total_tax"`
	AdminGraphqlAPIID string           `json:"admin_graphql_api_id"`
	Customer          *Customer        `json:"customer"`

	UseCustomerDefaultAddress bool `json:"use_customer_default_address,omitempty"`
}

// AppliedDiscount is the discount applied to the line item or the draft order object.
type AppliedDiscount struct {
	Title       string `json:"applied_discount"`
	Description string `json:"description"`
	Value       string `json:"value"`
	ValueType   string `json:"value_type"`
	Amount      string `json:"amount"`
}

// DraftOrderInvoice is the struct used to create an invoice for a draft order
type DraftOrderInvoice struct {
	To            string   `json:"to,omitempty"`
	From          string   `json:"from,omitempty"`
	Subject       string   `json:"subject,omitempty"`
	CustomMessage string   `json:"custom_message,omitempty"`
	Bcc           []string `json:"bcc,omitempty"`
}

type DraftOrdersResource struct {
	DraftOrders []DraftOrder `json:"draft_orders"`
}

type DraftOrderResource struct {
	DraftOrder *DraftOrder `json:"draft_order"`
}

type DraftOrderInvoiceResource struct {
	DraftOrderInvoice *DraftOrderInvoice `json:"draft_order_invoice,omitempty"`
}

// DraftOrderListOptions represents the possible options that can be used
// to further query the list draft orders endpoint
type DraftOrderListOptions struct {
	Fields       string     `url:"fields,omitempty"`
	Limit        int        `url:"limit,omitempty"`
	SinceID      int64      `url:"since_id,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax *time.Time `url:"updated_at_max,omitempty"`
	IDs          string     `url:"ids,omitempty"`
	Status       string     `url:"status,omitempty"`
}

// DraftOrderCountOptions represents the possible options to the count draft orders endpoint
type DraftOrderCountOptions struct {
	Fields  string `url:"fields,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	SinceID int64  `url:"since_id,omitempty"`
	IDs     string `url:"ids,omitempty"`
	Status  string `url:"status,omitempty"`
}

// Create draft order
func (s *DraftOrderServiceOp) Create(draftOrder DraftOrder) (*DraftOrder, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	wrappedData := DraftOrderResource{DraftOrder: &draftOrder}
	resource := new(DraftOrderResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.DraftOrder, err
}

// List draft orders
func (s *DraftOrderServiceOp) List(options interface{}) ([]DraftOrder, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	resource := new(DraftOrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.DraftOrders, err
}

func (s *DraftOrderServiceOp) ListWithPagination(options interface{}) ([]DraftOrder, *Pagination, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	resource := new(DraftOrdersResource)
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

	return resource.DraftOrders, pagination, nil
}

// Count draft orders
func (s *DraftOrderServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", draftOrdersBasePath)
	return s.client.Count(path, options)
}

// Delete draft orders
func (s *DraftOrderServiceOp) Delete(draftOrderID int64) error {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, draftOrderID)
	return s.client.Delete(path)
}

// Invoice a draft order
func (s *DraftOrderServiceOp) Invoice(draftOrderID int64, draftOrderInvoice DraftOrderInvoice) (*DraftOrderInvoice, error) {
	path := fmt.Sprintf("%s/%d/send_invoice.json", draftOrdersBasePath, draftOrderID)
	wrappedData := DraftOrderInvoiceResource{DraftOrderInvoice: &draftOrderInvoice}
	resource := new(DraftOrderInvoiceResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.DraftOrderInvoice, err
}

// Get individual draft order
func (s *DraftOrderServiceOp) Get(draftOrderID int64, options interface{}) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, draftOrderID)
	resource := new(DraftOrderResource)
	err := s.client.Get(path, resource, options)
	return resource.DraftOrder, err
}

// Update draft order
func (s *DraftOrderServiceOp) Update(draftOrder DraftOrder) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, draftOrder.ID)
	wrappedData := DraftOrderResource{DraftOrder: &draftOrder}
	resource := new(DraftOrderResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.DraftOrder, err
}

// Complete draft order
func (s *DraftOrderServiceOp) Complete(draftOrderID int64, paymentPending bool) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d/complete.json?payment_pending=%t", draftOrdersBasePath, draftOrderID, paymentPending)
	resource := new(DraftOrderResource)
	err := s.client.Put(path, nil, resource)
	return resource.DraftOrder, err
}

// List metafields for an order
func (s *DraftOrderServiceOp) ListMetafields(draftOrderID int64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metafieldService.List(options)
}

// Count metafields for an order
func (s *DraftOrderServiceOp) CountMetafields(draftOrderID int64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metafieldService.Count(options)
}

// Get individual metafield for an order
func (s *DraftOrderServiceOp) GetMetafield(draftOrderID int64, metafieldID int64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metafieldService.Get(metafieldID, options)
}

// Create a new metafield for an order
func (s *DraftOrderServiceOp) CreateMetafield(draftOrderID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metafieldService.Create(metafield)
}

// Update an existing metafield for an order
func (s *DraftOrderServiceOp) UpdateMetafield(draftOrderID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metafieldService.Update(metafield)
}

// Delete an existing metafield for an order
func (s *DraftOrderServiceOp) DeleteMetafield(draftOrderID int64, metafieldID int64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metafieldService.Delete(metafieldID)
}

func (s *DraftOrderServiceOp) GetOrderList() []string {
	str := new(DraftOrder)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
