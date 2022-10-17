package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

const marketingEventBasePath = "marketing_events"

// MarketingEventService is an interface for interacting with the
// MarketingEvent endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/marketingevent
type MarketingEventService interface {
	Get(int64, interface{}) (*MarketingEvent, error)
	List(interface{}) ([]MarketingEvent, error)
	ListWithPagination(interface{}) ([]MarketingEvent, *Pagination, error)
	GetOrderList() []string
}

type MarketingEventServiceOp struct {
	client *Client
}

type MarketingEvent struct {
	ID                  int64                `json:"id"`
	EventType           string               `json:"event_type"`
	RemoteID            string               `json:"remote_id"`
	StartedAt           *time.Time           `json:"started_at"`
	EndedAt             *time.Time           `json:"ended_at"`
	ScheduledToEndAt    *time.Time           `json:"scheduled_to_end_at"`
	Budget              *decimal.Decimal     `json:"budget"`
	Currency            string               `json:"currency"`
	ManageURL           string               `json:"manage_url"`
	PreviewURL          string               `json:"preview_url"`
	UtmCampaign         string               `json:"utm_campaign"`
	UtmSource           string               `json:"utm_source"`
	UtmMedium           string               `json:"utm_medium"`
	BudgetType          string               `json:"budget_type"`
	Description         string               `json:"description"`
	MarketingChannel    string               `json:"marketing_channel"`
	Paid                bool                 `json:"paid"`
	ReferringDomain     string               `json:"referring_domain"`
	BreadcrumbID        string               `json:"breadcrumb_id"`
	MarketingActivityID string               `json:"marketing_activity_id"`
	AdminGraphqlAPIID   string               `json:"admin_graphql_api_id"`
	MarketedResources   []*MarketedResources `json:"marketed_resources"`
}

type MarketedResources struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// MarketingEventResource represents the result from the
// admin/marketingEvents{/X{/activate.json}.json}.json endpoints.
type MarketingEventResource struct {
	MarketingEvent *MarketingEvent `json:"marketing_event"`
}

// MarketingEventsResource represents the result from the
// admin/marketingEvents.json endpoint.
type MarketingEventsResource struct {
	MarketingEvents []MarketingEvent `json:"marketing_events"`
}

// Get gets individual application marketingEvent.
func (s MarketingEventServiceOp) Get(marketingEventID int64, options interface{}) (*MarketingEvent, error) {
	path := fmt.Sprintf("%s/%d.json", marketingEventBasePath, marketingEventID)
	resource := &MarketingEventResource{}
	return resource.MarketingEvent, s.client.Get(path, resource, options)
}

// List gets all application marketingEvents.
func (s MarketingEventServiceOp) List(options interface{}) ([]MarketingEvent, error) {
	path := fmt.Sprintf("%s.json", marketingEventBasePath)
	resource := &MarketingEventsResource{}
	return resource.MarketingEvents, s.client.Get(path, resource, options)
}

func (s *MarketingEventServiceOp) ListWithPagination(options interface{}) ([]MarketingEvent, *Pagination, error) {
	path := fmt.Sprintf("%s.json", marketingEventBasePath)
	resource := new(MarketingEventsResource)
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

	return resource.MarketingEvents, pagination, nil
}

func (s *MarketingEventServiceOp) GetOrderList() []string {
	str := new(MarketingEvent)

	var orderList []string
	for i := 0; i < reflect.TypeOf(*str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(*str).Field(i).Tag.Get("json"))
	}

	return orderList
}
