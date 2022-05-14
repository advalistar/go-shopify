package goshopify

import (
	"fmt"
	"time"
)

const mobilePlatformApplicationBasePath = "mobile_platform_applications"

type MobilePlatformApplicationService interface {
	List(interface{}) ([]MobilePlatformApplication, error)
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
