package goshopify

import (
	"fmt"
	"net/http"
	"time"
)

const mobilePlatformApplicationBasePath = "mobile_platform_applications"

type MobilePlatformApplicationService interface {
	List(interface{}) ([]MobilePlatformApplication, error)
	ListWithPagination(interface{}) ([]MobilePlatformApplication, *Pagination, error)
}

type MobilePlatformApplicationServiceOp struct {
	client *Client
}

type MobilePlatformApplication struct {
	ID                          int64      `json:"id"`
	Application_id              string     `json:"application_id"`
	Platform                    string     `json:"platform"`
	CreatedAt                   *time.Time `json:"created_at"`
	UpdatedAt                   *time.Time `json:"updated_at"`
	Sha256CertFingerprints      []string   `json:"sha256_cert_fingerprints"`
	EnabledUniversalOrAppLinks  bool       `json:"enabled_universal_or_app_links"`
	EnabledSharedWebcredentials bool       `json:"enabled_shared_webcredentials"`
}

type MobilePlatformApplicationsResource struct {
	MobilePlatformApplications []MobilePlatformApplication `json:"mobile_platform_applications"`
}

// List of discount codes
func (s *MobilePlatformApplicationServiceOp) List(options interface{}) ([]MobilePlatformApplication, error) {
	path := fmt.Sprintf("%s.json", mobilePlatformApplicationBasePath)
	resource := new(MobilePlatformApplicationsResource)
	err := s.client.Get(path, resource, options)
	return resource.MobilePlatformApplications, err
}

func (s *MobilePlatformApplicationServiceOp) ListWithPagination(options interface{}) ([]MobilePlatformApplication, *Pagination, error) {
	path := fmt.Sprintf("%s.json", mobilePlatformApplicationBasePath)
	resource := new(MobilePlatformApplicationsResource)
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

	return resource.MobilePlatformApplications, pagination, nil
}
