package goshopify

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	collectionListingBasePath = "collection_listings"
)

type CollectionListingService interface {
	List(interface{}) ([]CollectionListing, error)
	ListWithPagination(interface{}) ([]CollectionListing, *Pagination, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*CollectionListing, error)
	GetOrderList() []string
}

// CollectionListingServiceOp handles communication with the collection related methods of
// the Shopify API.
type CollectionListingServiceOp struct {
	client *Client
}

// CollectionListing represents a Shopify collection published to your sales channel app
type CollectionListing struct {
	CollectionID        int64      `json:"collection_id"`
	UpdatedAt           *time.Time `json:"updated_at"`
	BodyHTML            string     `json:"body_html"`
	DefaultProductImage []*Image   `json:"default_product_image"`
	Handle              string     `json:"handle"`
	Images              []*Image   `json:"images"`
	Title               string     `json:"title"`
	SortOrder           string     `json:"sort_order"`
	PublishedAt         *time.Time `json:"published_at"`
}

// Represents the result from the collection_listings/X.json endpoint
type CollectionListingResource struct {
	CollectionListing *CollectionListing `json:"collection_listing"`
}

// Represents the result from the collection_listings.json endpoint
type CollectionsListingsResource struct {
	CollectionListings []CollectionListing `json:"collection_listings"`
}

// List collections
func (s *CollectionListingServiceOp) List(options interface{}) ([]CollectionListing, error) {
	collections, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// ListWithPagination lists collections and return pagination to retrieve next/previous results.
func (s *CollectionListingServiceOp) ListWithPagination(options interface{}) ([]CollectionListing, *Pagination, error) {
	path := fmt.Sprintf("%s.json", collectionListingBasePath)
	resource := new(CollectionsListingsResource)
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

	return resource.CollectionListings, pagination, nil
}

// Count collections listings published to your sales channel app
func (s *CollectionListingServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", collectionListingBasePath)
	return s.client.Count(path, options)
}

// Get individual collection_listing by collection ID
func (s *CollectionListingServiceOp) Get(collectionID int64, options interface{}) (*CollectionListing, error) {
	path := fmt.Sprintf("%s/%d.json", collectionListingBasePath, collectionID)
	resource := new(CollectionListingResource)
	err := s.client.Get(path, resource, options)
	return resource.CollectionListing, err
}

func (s *CollectionListingServiceOp) GetOrderList() []string {
	str := new(CollectionListing)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
