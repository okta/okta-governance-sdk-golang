package governance

import (
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
