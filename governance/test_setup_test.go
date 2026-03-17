package governance

import (
	"os"
	"testing"

	"github.com/okta/okta-sdk-golang/v6/okta"
)

// Package-level test client setup
// This is used by integration tests in this package
type testClient struct {
	*OktaGovernanceAPIClient
	cfg *okta.Configuration
}

var apiClient *testClient

func init() {
	// Initialize test client with default configuration
	// Users should set up their own configuration via environment variables or config file
	cfg, err := okta.NewConfiguration()
	if err != nil {
		// If configuration fails, apiClient will be nil and tests will skip
		return
	}

	client := NewAPIClient(cfg)
	apiClient = &testClient{
		OktaGovernanceAPIClient: client,
		cfg:                     cfg,
	}
}

// hasOktaCredentials checks if Okta credentials are configured via environment variables
func hasOktaCredentials() bool {
	orgUrl := os.Getenv("OKTA_CLIENT_ORGURL")
	token := os.Getenv("OKTA_CLIENT_TOKEN")
	return orgUrl != "" && token != ""
}

// skipIfNoCredentials skips the test if Okta credentials are not configured
func skipIfNoCredentials(t *testing.T) {
	if apiClient == nil || !hasOktaCredentials() {
		t.Skip("Skipping integration test: Okta credentials not configured (set OKTA_CLIENT_ORGURL and OKTA_CLIENT_TOKEN)")
	}
}
