package goshopify

import "reflect"

type AccessScopesService interface {
	List(interface{}) ([]AccessScope, error)
	GetOrderList() []string
}

type AccessScope struct {
	Handle string `json:"handle,omitempty"`
}

// AccessScopesResource represents the result from the oauth/access_scopes.json endpoint
type AccessScopesResource struct {
	AccessScopes []AccessScope `json:"access_scopes,omitempty"`
}

// AccessScopesServiceOp handles communication with the Access Scopes
// related methods of the Shopify API
type AccessScopesServiceOp struct {
	client *Client
}

// List gets access scopes based on used oauth token
func (s *AccessScopesServiceOp) List(options interface{}) ([]AccessScope, error) {
	path := "oauth/access_scopes.json"
	resource := new(AccessScopesResource)
	err := s.client.Get(path, resource, options)
	return resource.AccessScopes, err
}

func (s *AccessScopesServiceOp) GetOrderList() []string {
	str := new(AccessScope)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
