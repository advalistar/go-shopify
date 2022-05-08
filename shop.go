package goshopify

import (
	"time"
)

// ShopService is an interface for interfacing with the shop endpoint of the
// Shopify API.
// See: https://help.shopify.com/api/reference/shop
type ShopService interface {
	Get(options interface{}) (*Shop, error)
}

// ShopServiceOp handles communication with the shop related methods of the
// Shopify API.
type ShopServiceOp struct {
	client *Client
}

// Shop represents a Shopify shop
type Shop struct {
	ID                              int64      `json:"id"`
	Name                            string     `json:"name"`
	Email                           string     `json:"email"`
	Domain                          string     `json:"domain"`
	Province                        string     `json:"province"`
	Country                         string     `json:"country"`
	Address1                        string     `json:"address1"`
	Zip                             string     `json:"zip"`
	City                            string     `json:"city"`
	Source                          string     `json:"source"`
	Phone                           string     `json:"phone"`
	Latitude                        float64    `json:"latitude"`
	Longitude                       float64    `json:"longitude"`
	PrimaryLocale                   string     `json:"primary_locale"`
	Address2                        string     `json:"address2"`
	CreatedAt                       *time.Time `json:"created_at"`
	UpdatedAt                       *time.Time `json:"updated_at"`
	CountryCode                     string     `json:"country_code"`
	CountryName                     string     `json:"country_name"`
	Currency                        string     `json:"currency"`
	CustomerEmail                   string     `json:"customer_email"`
	Timezone                        string     `json:"timezone"`
	IanaTimezone                    string     `json:"iana_timezone"`
	ShopOwner                       string     `json:"shop_owner"`
	MoneyFormat                     string     `json:"money_format"`
	MoneyWithCurrencyFormat         string     `json:"money_with_currency_format"`
	WeightUnit                      string     `json:"weight_unit"`
	ProvinceCode                    string     `json:"province_code"`
	TaxesIncluded                   bool       `json:"taxes_included"`
	TaxShipping                     bool       `json:"tax_shipping"`
	CountyTaxes                     bool       `json:"county_taxes"`
	PlanDisplayName                 string     `json:"plan_display_name"`
	PlanName                        string     `json:"plan_name"`
	HasDiscounts                    bool       `json:"has_discounts"`
	HasGiftcards                    bool       `json:"has_gift_cards"`
	MyshopifyDomain                 string     `json:"myshopify_domain"`
	GoogleAppsDomain                string     `json:"google_apps_domain"`
	GoogleAppsLoginEnabled          bool       `json:"google_apps_login_enabled"`
	MoneyInEmailsFormat             string     `json:"money_in_emails_format"`
	MoneyWithCurrencyInEmailsFormat string     `json:"money_with_currency_in_emails_format"`
	EligibleForPayments             bool       `json:"eligible_for_payments"`
	RequiresExtraPaymentsAgreement  bool       `json:"requires_extra_payments_agreement"`
	PasswordEnabled                 bool       `json:"password_enabled"`
	HasStorefront                   bool       `json:"has_storefront"`
	EligibleForCardReaderGiveaway   bool       `json:"eligible_for_card_reader_giveaway"`
	Finances                        bool       `json:"finances"`
	PrimaryLocationID               int64      `json:"primary_location_id"`
	CookieConsentLevel              string     `json:"cookie_consent_level"`
	CheckoutAPISupported            bool       `json:"checkout_api_supported"`
	MultiLocationEnabled            bool       `json:"multi_location_enabled"`
	SetupRequire                    bool       `json:"setup_required"`
	PreLaunchEnabled                bool       `json:"pre_launch_enabled"`
	EnabledPresentmentCurrencies    []string   `json:"enabled_presentment_currencies"`
	TransactionalSmsDisabled        bool       `json:"transactional_sms_disabled"`
	ForceSSL                        bool       `json:"force_ssl,omitempty"`
}

// Represents the result from the admin/shop.json endpoint
type ShopResource struct {
	Shop *Shop `json:"shop"`
}

// Get shop
func (s *ShopServiceOp) Get(options interface{}) (*Shop, error) {
	resource := new(ShopResource)
	err := s.client.Get("shop.json", resource, options)
	return resource.Shop, err
}
