package goshopify

import (
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
)

const applicationCreditsBasePath = "application_credits"

// ApplicationCreditService is an interface for interacting with the
// ApplicationCredit endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/billing/applicationcredit
type ApplicationCreditService interface {
	Create(ApplicationCredit) (*ApplicationCredit, error)
	Get(int64, interface{}) (*ApplicationCredit, error)
	List(interface{}) ([]ApplicationCredit, error)
	GetOrderList() []string
}

type ApplicationCreditServiceOp struct {
	client *Client
}

type ApplicationCredit struct {
	ID          int64            `json:"id"`
	Amount      *decimal.Decimal `json:"amount"`
	Description string           `json:"description"`
	Test        bool             `json:"test"`
}

// ApplicationCreditResource represents the result from the
// admin/application_credits{/X{/activate.json}.json}.json endpoints.
type ApplicationCreditResource struct {
	Credit *ApplicationCredit `json:"application_credit"`
}

// ApplicationCreditsResource represents the result from the
// admin/application_credits.json endpoint.
type ApplicationCreditsResource struct {
	Credits []ApplicationCredit `json:"application_credits"`
}

// Create creates new application credit.
func (s ApplicationCreditServiceOp) Create(credit ApplicationCredit) (*ApplicationCredit, error) {
	path := fmt.Sprintf("%s.json", applicationCreditsBasePath)
	resource := &ApplicationCreditResource{}
	return resource.Credit, s.client.Post(path, ApplicationCreditResource{Credit: &credit}, resource)
}

// Get gets individual application credit.
func (s ApplicationCreditServiceOp) Get(creditID int64, options interface{}) (*ApplicationCredit, error) {
	path := fmt.Sprintf("%s/%d.json", applicationCreditsBasePath, creditID)
	resource := &ApplicationCreditResource{}
	return resource.Credit, s.client.Get(path, resource, options)
}

// List gets all application credits.
func (s ApplicationCreditServiceOp) List(options interface{}) ([]ApplicationCredit, error) {
	path := fmt.Sprintf("%s.json", applicationCreditsBasePath)
	resource := &ApplicationCreditsResource{}
	return resource.Credits, s.client.Get(path, resource, options)
}

func (s *ApplicationCreditServiceOp) GetOrderList() []string {
	str := new(ApplicationCredit)

	var orderList []string
	for i := 0; i < reflect.TypeOf(str).NumField(); i++ {
		orderList = append(orderList, reflect.TypeOf(str).Field(i).Tag.Get("json"))
	}

	return orderList
}
