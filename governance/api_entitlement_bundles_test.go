package governance

import (
	"net/http"
	"testing"
)

func TestGetEntitlementBundleExecute(t *testing.T) {
	t.Skip("Skipping test")
	execute, a, err := apiClient.EntitlementBundlesAPI.GetentitlementBundle(apiClient.cfg.Context, "enbzcbqe3Ts4wdw1swD1d6").Execute()
	if err != nil {
		t.Errorf("Error getting entitlement bundle: %v", err)
		return
	}
	if execute.Entitlements == nil {
		t.Errorf("Expected entitlements to be present in the response")
		return
	}

	if a.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}
}

func TestCreateEntitlementBundleRequest_Execute(t *testing.T) {
	t.Skip("Skipping test")
	_, a, err := apiClient.EntitlementBundlesAPI.CreateEntitlementBundle(apiClient.cfg.Context).EntitlementBundleCreatable(buildEntitlementBundle("entitlement bundle test")).Execute()
	if err != nil {
		t.Errorf("Error creating entitlement bundle: %v", err)
		return
	}

	if a.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", a.StatusCode)
		return
	}
}

func TestDeleteEntitlementBundleRequest_Execute(t *testing.T) {
	t.Skip("Skipping test")
	entitlementBundle, _, _ := apiClient.EntitlementBundlesAPI.CreateEntitlementBundle(apiClient.cfg.Context).EntitlementBundleCreatable(buildEntitlementBundle("entitlement Bundle delete")).Execute()
	execute, err := apiClient.EntitlementBundlesAPI.DeleteEntitlementBundle(apiClient.cfg.Context, entitlementBundle.Id).Execute()
	if err != nil {
		t.Errorf("Error deleting entitlement bundle: %v", err)
		return
	}

	if execute.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %d", execute.StatusCode)
		return
	}
}

func buildEntitlementBundle(name string) EntitlementBundleCreatable {
	entitlementId := "espzcbqd7Suwp4Y7A1d6"
	entitlementValueId := "entzcbqd8lcD3BRWR1d6"

	return EntitlementBundleCreatable{
		Name: name,
		Entitlements: []EntitlementCreatable{
			{
				Id: &entitlementId,
				Values: []EntitlementValueCreatable{
					{
						Id: &entitlementValueId,
					},
				},
			},
		},
		Target: TargetResource{
			Type:       "APPLICATION",
			ExternalId: "0oao01ardu8r8qUP91d7",
		},
	}
}

func TestListEntitlementBundlesRequest_Execute(t *testing.T) {
	t.Skip("Skipping test")
	// filter := `target.externalId eq "0oao01ardu8r8qUP91d7" AND target.type eq "APPLICATION"`
	// include := []string{"full_entitlements"}
	_, a, err := apiClient.EntitlementBundlesAPI.GetentitlementBundle(apiClient.cfg.Context, "enbzng5wgTiu7d04L1d6").Execute()
	if err != nil {
		t.Errorf("Error listing entitlement bundles: %v", err)
		return
	}

	if a.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}
}
