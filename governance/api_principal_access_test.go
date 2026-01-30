package governance

import (
	"testing"
)

func TestApiGetPrincipalAccessRequest_Execute(t *testing.T) {
	t.Skip("Skipping test")
	filter := `parent.externalId eq "0oao01ardu8rdw8qUP91d7" AND parent.type eq "APPLICATION" AND targetPrincipal.externalId eq "00unkw1sfbTw08c0g1d7" AND targetPrincipal.type eq "OKTA_USER"` // Replace with a valid principal ID
	execute, a, err := apiClient.PrincipalAccessAPI.GetPrincipalAccess(apiClient.cfg.Context).Filter(filter).Execute()
	if err != nil {
		t.Errorf("Error getting principal access: %v", err)
		return
	}

	if a.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}

	if execute.TargetPrincipalOrn == "" {
		t.Errorf("Response should have contained a target principal ORN")
		return
	}
}
