package governance

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/okta/okta-sdk-golang/v6/okta"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

const (
	// CassettePath is the directory where cassettes are stored
	CassettePath = "testdata/cassettes"
)

// RecorderMode returns the recorder mode based on the VCR_RECORD environment variable
// Set VCR_RECORD=true to record new cassettes, otherwise playback mode is used
func RecorderMode() recorder.Mode {
	if os.Getenv("VCR_RECORD") == "true" {
		return recorder.ModeRecordOnly
	}
	return recorder.ModeReplayOnly
}

// IsRecording returns true if currently in recording mode
func IsRecording() bool {
	return os.Getenv("VCR_RECORD") == "true"
}

// testdataDir returns the absolute path to the testdata directory
func testdataDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return CassettePath
	}
	return filepath.Join(filepath.Dir(filename), CassettePath)
}

// CassetteName returns the full path to a cassette file
func CassetteName(apiGroup, testName string) string {
	return filepath.Join(testdataDir(), apiGroup, testName)
}

// RecorderTestClient wraps the API client with a recorder for unit testing
type RecorderTestClient struct {
	*OktaGovernanceAPIClient
	Recorder *recorder.Recorder
	cfg      *okta.Configuration
}

// NewRecorderClient creates a new API client configured with HTTP recording/playback.
// The cassette name should be in the format "api_group/test_name" (e.g., "campaigns/list_campaigns").
// When VCR_RECORD=true, real HTTP calls are made and recorded.
// When VCR_RECORD is not set, previously recorded responses are played back.
//
// Recording requires valid Okta credentials configured via:
//   - Environment variables: OKTA_CLIENT_ORGURL and OKTA_CLIENT_TOKEN
//   - Or config file: ~/.okta/okta.yaml
func NewRecorderClient(t *testing.T, cassetteName string) *RecorderTestClient {
	t.Helper()

	mode := RecorderMode()

	opts := []recorder.Option{
		recorder.WithMode(mode),
	}

	// Add hook to remove sensitive headers from recorded cassettes
	if mode == recorder.ModeRecordOnly {
		opts = append(opts, recorder.WithHook(sanitizeHook, recorder.BeforeSaveHook))
	}

	rec, err := recorder.New(cassetteName, opts...)
	if err != nil {
		t.Fatalf("Failed to create recorder: %v", err)
	}

	// Check for environment variable overrides first
	orgUrl := os.Getenv("OKTA_CLIENT_ORGURL")
	token := os.Getenv("OKTA_CLIENT_TOKEN")

	var cfg *okta.Configuration
	if orgUrl != "" && token != "" {
		// Use environment variables
		cfg, err = okta.NewConfiguration(
			okta.WithOrgUrl(orgUrl),
			okta.WithToken(token),
		)
		if err != nil {
			t.Fatalf("Failed to create Okta configuration from env vars: %v", err)
		}
	} else {
		// Fall back to config file (~/.okta/okta.yaml)
		cfg, err = okta.NewConfiguration()
		if err != nil {
			if mode == recorder.ModeRecordOnly {
				t.Fatalf("Failed to get Okta configuration for recording.\n"+
					"Ensure valid credentials are configured via:\n"+
					"  - Environment: OKTA_CLIENT_ORGURL and OKTA_CLIENT_TOKEN\n"+
					"  - Or file: ~/.okta/okta.yaml\n"+
					"Error: %v", err)
			}
			// For playback, create a minimal configuration with empty values
			cfg, _ = okta.NewConfiguration(
				okta.WithOrgUrl("https://test.okta.com"),
				okta.WithToken("test-token"),
			)
		}
	}

	// Log the org URL being used when recording
	if mode == recorder.ModeRecordOnly {
		t.Logf("Recording mode: using Okta org %s", cfg.Okta.Client.OrgUrl)
	}

	// Create HTTP client with recorder transport
	httpClient := &http.Client{
		Transport: rec,
	}
	cfg.HTTPClient = httpClient

	client := NewAPIClient(cfg)

	return &RecorderTestClient{
		OktaGovernanceAPIClient: client,
		Recorder:                rec,
		cfg:                     cfg,
	}
}

// Stop stops the recorder. Must be called at the end of the test.
func (c *RecorderTestClient) Stop() error {
	return c.Recorder.Stop()
}

// sanitizeHook removes sensitive data from recorded cassettes
func sanitizeHook(i *cassette.Interaction) error {
	// Sanitize request headers
	delete(i.Request.Headers, "Authorization")
	delete(i.Request.Headers, "Cookie")
	delete(i.Request.Headers, "X-Okta-User-Agent-Extended")

	// Sanitize response headers
	delete(i.Response.Headers, "Set-Cookie")
	delete(i.Response.Headers, "X-Okta-Request-Id")

	return nil
}

// skipIfNotRecording skips the test if not in recording mode and cassette doesn't exist
func skipIfNotRecording(t *testing.T, cassetteName string) {
	t.Helper()
	if IsRecording() {
		return
	}
	// Check if cassette file exists
	cassettePath := cassetteName + ".yaml"
	if _, err := os.Stat(cassettePath); os.IsNotExist(err) {
		t.Skipf("Cassette %s does not exist. Run with VCR_RECORD=true to record.", cassetteName)
	}
}
