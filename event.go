package goshopify

import (
	"fmt"
	"net/http"
	"time"
)

const eventBasePath = "events"

// EventService is an interface for interacting with the
// Event endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/events
type EventService interface {
	Get(int64, interface{}) (*Event, error)
	List(interface{}) ([]Event, error)
	ListWithPagination(interface{}) ([]Event, *Pagination, error)
}

type EventServiceOp struct {
	client *Client
}

type Event struct {
	ID           int64      `json:"id"`
	SubjectID    int64      `json:"subject_id"`
	CreatedAt    *time.Time `json:"created_at"`
	Subject_type string     `json:"subject_type"`
	Verb         string     `json:"verb"`
	Arguments    []string   `json:"arguments"`
	Body         string     `json:"body"`
	Message      string     `json:"message"`
	Author       string     `json:"author"`
	Description  string     `json:"description"`
	Path         string     `json:"path"`
}

// EventResource represents the result from the
// admin/events{/X{/activate.json}.json}.json endpoints.
type EventResource struct {
	Event *Event `json:"event"`
}

// EventsResource represents the result from the
// admin/events.json endpoint.
type EventsResource struct {
	Events []Event `json:"events"`
}

// Get gets individual application event.
func (s EventServiceOp) Get(eventID int64, options interface{}) (*Event, error) {
	path := fmt.Sprintf("%s/%d.json", eventBasePath, eventID)
	resource := &EventResource{}
	return resource.Event, s.client.Get(path, resource, options)
}

// List gets all application events.
func (s EventServiceOp) List(options interface{}) ([]Event, error) {
	path := fmt.Sprintf("%s.json", eventBasePath)
	resource := &EventsResource{}
	return resource.Events, s.client.Get(path, resource, options)
}

func (s *EventServiceOp) ListWithPagination(options interface{}) ([]Event, *Pagination, error) {
	path := fmt.Sprintf("%s.json", eventBasePath)
	resource := new(EventsResource)
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

	return resource.Events, pagination, nil
}
