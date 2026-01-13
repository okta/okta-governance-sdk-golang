package governance

import (
	"testing"
)

func TestGrantsAPIService_CreateGrant(t *testing.T) {
	t.Skip("Skipping test")
	_, a, err := apiClient.GrantsAPI.CreateGrant(apiClient.cfg.Context).GrantCreatable(buildGrant()).Execute()
	if err != nil {
		t.Errorf("Error getting grants: %v", err)
		return
	}

	if a.StatusCode != 201 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}
}

func buildGrant() GrantCreatable {
	return GrantCreatable{
		GrantTypeBundleWriteable: &GrantTypeBundleWriteable{
			TargetPrincipal: TargetPrincipal{
				Type:       "OKTA_USER",
				ExternalId: "00unkw1sfbTw2de08c0g1d7", // Replace with a valid principal ID
			},
			EntitlementBundleId: "enbzcco02Hw2i4qsDDE1d6",
			GrantType:           "ENTITLEMENT-BUNDLE",
		},
	}
}

func TestGrantsAPIService_ListGrantsExecute(t *testing.T) {
	t.Skip("Skipping test")
	fields := []string{"full_entitlements"}
	filter := `target.externalId eq "0oao01ardu8r8qdwUP91d7" AND target.type eq "APPLICATION" AND targetPrincipal.externalId eq "00unkw1sfbTw08c0g1d7" AND targetPrincipal.type eq "OKTA_USER"`
	_, a, err := apiClient.GrantsAPI.ListGrants(apiClient.cfg.Context).Filter(filter).Include(fields).Execute()
	if err != nil {
		t.Errorf("Error listing grants: %v", err)
		return
	}

	if a.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}
}
