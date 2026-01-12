package governance

import (
	"testing"

	apiClient "github.com/okta/okta-governance-sdk-golang"
)

func TestEntitlementsAPIService_GetPrincipalEntitlements(t *testing.T) {
	filter := `parentResourceOrn eq "orn:oktapreview:idp:00onkdqw1sbuAh3Q06I1d7:apps:oidc_client:0oao01ardcdu8r8qUP91d7" AND targetPrincipalOrn eq "orn:oktapreview:directory:00onkw1sbuAh3Q06I1d7:users:00unkw1sfbTcew08c0g1d7"`
	_, a, err := apiClient.PrincipalEntitlementsAPI.GetPrincipalEntitlements(apiClient.cfg.Context).Filter(filter).Execute()
	if err != nil {
		t.Errorf("Error getting entitlements: %v", err)
		return
	}

	if a.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}
}
