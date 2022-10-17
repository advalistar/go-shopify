package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const giftCardBasePath = "gift_cards"

// GiftCardService is an interface for interacting with the
// GiftCard endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/giftCards
type GiftCardService interface {
	Get(int64, interface{}) (*GiftCard, error)
	List(interface{}) ([]GiftCard, error)
	ListWithPagination(interface{}) ([]GiftCard, *Pagination, error)
	GetOrderList() []string
}

type GiftCardServiceOp struct {
	client *Client
}

type GiftCard struct {
	ID             int64            `json:"id"`
	Balance        *decimal.Decimal `json:"balance"`
	CreatedAt      *time.Time       `json:"created_at"`
	UpdatedAt      *time.Time       `json:"updated_at"`
	Currency       string           `json:"currency"`
	InitialValue   *decimal.Decimal `json:"initial_value"`
	DisabledAt     *time.Time       `json:"disabled_at"`
	LineItemID     int64            `json:"line_item_id"`
	APIClientID    int64            `json:"api_client_id"`
	UserID         int64            `json:"user_id"`
	CustomerID     interface{}      `json:"customer_id"`
	Note           string           `json:"note"`
	ExpiresOn      *time.Time       `json:"expires_on"`
	TemplateSuffix string           `json:"template_suffix"`
	LastCharacters string           `json:"last_characters"`
	OrderID        int64            `json:"order_id"`
}

// GiftCardResource represents the result from the
// admin/giftCards{/X{/activate.json}.json}.json endpoints.
type GiftCardResource struct {
	GiftCard *GiftCard `json:"gift_card"`
}

// GiftCardsResource represents the result from the
// admin/giftCards.json endpoint.
type GiftCardsResource struct {
	GiftCards []GiftCard `json:"gift_cards"`
}

// Get gets individual application giftCard.
func (s GiftCardServiceOp) Get(giftCardID int64, options interface{}) (*GiftCard, error) {
	path := fmt.Sprintf("%s/%d.json", giftCardBasePath, giftCardID)
	resource := &GiftCardResource{}
	return resource.GiftCard, s.client.Get(path, resource, options)
}

// List gets all application giftCards.
func (s GiftCardServiceOp) List(options interface{}) ([]GiftCard, error) {
	path := fmt.Sprintf("%s.json", giftCardBasePath)
	resource := &GiftCardsResource{}
	return resource.GiftCards, s.client.Get(path, resource, options)
}

func (s *GiftCardServiceOp) ListWithPagination(options interface{}) ([]GiftCard, *Pagination, error) {
	path := fmt.Sprintf("%s.json", giftCardBasePath)
	resource := new(GiftCardsResource)
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

	return resource.GiftCards, pagination, nil
}

func (s *GiftCardServiceOp) GetOrderList() []string {
	str := new(GiftCard)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
