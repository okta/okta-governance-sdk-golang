package governance

import (
	apiClient "github.com/okta/okta-governance-sdk-golang"
	"testing"
)

func TestCreateEntitlementExecute(t *testing.T) {
	applicationList, _, _ := idaasAPIClient.ApplicationAPI.ListApplications(apiClient.cfg.Context).Execute()

	for _, application := range applicationList {
		if application.OpenIdConnectApplication == nil {
			continue // Skip applications that are not SAML
		}
	}

	createdEntitlement, resp, err := apiClient.EntitlementsAPI.CreateEntitlement(apiClient.cfg.Context).EntitlementCreate(buildEntitlement()).Execute()
	if err != nil {
		t.Errorf("Error creating entitlement: %v", err)
		return
	}
	if resp.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
		return
	}

	t.Cleanup(func() {
		// Best effort cleanup: attempt to delete the entitlement if it exists and is deletable
		if createdEntitlement.Id != "" {
			// Only SCHEDULED or ERROR entitlements can be deleted.
			// If the test failed midway, it might be in ACTIVE.
			// For acceptance tests, it's often good practice to ensure cleanup,
			// but direct deletion might not always be possible based on status.
			execute, err := apiClient.EntitlementsAPI.DeleteEntitlement(apiClient.cfg.Context, createdEntitlement.Id).Execute()
			if execute.StatusCode != 201 || err != nil {
				return
			} // nolint
		}
	})
}

func TestListEntitlements(t *testing.T) {
	//createdEntitlement, resp, err := apiClient.EntitlementsAPI.CreateEntitlement(apiClient.cfg.Context).EntitlementCreate(buildEntitlement()).Execute()
	//if err != nil {
	//	t.Errorf("Error creating entitlement: %v", err)
	//	return
	//}
	//if resp.StatusCode != 201 {
	//	t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	//	return
	//}

	filter := `parent.externalId eq "0oankw1sc6z965dIC1d7" AND parent.type eq "APPLICATION"`
	// List entitlements
	entitlements, resp, err := apiClient.EntitlementsAPI.ListEntitlements(apiClient.cfg.Context).Filter(filter).Execute()
	if err != nil {
		t.Errorf("Error listing entitlements: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	if len(entitlements.Data) != 14 {
		t.Errorf("Expected 14 Entitlements, got %d", len(entitlements.Data))
		return
	}

	t.Cleanup(func() {
		execute, err := apiClient.EntitlementsAPI.DeleteEntitlement(apiClient.cfg.Context, "0oankw1sc6z965dIC1d7").Execute()
		if execute.StatusCode != 201 || err != nil {
			return
		} // nolint
	})
}

func TestGetEntitlement(t *testing.T) {
	entitlementId := "espyyzpbwwzhBGRAf1d6" // Replace with a valid entitlement ID

	entitlement, resp, err := apiClient.EntitlementsAPI.GetEntitlement(apiClient.cfg.Context, entitlementId).Execute()
	if err != nil {
		t.Errorf("Error getting entitlement: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	if entitlement.Id != entitlementId {
		t.Errorf("Expected entitlement ID %s, got %s", entitlementId, entitlement.Id)
		return
	}
}

func TestReplaceEntitlement(t *testing.T) {
	entitlementId := "espyyzpbwwzhBGRAf1d6" // Replace with a valid entitlement ID

	// Build the new entitlement data

	// Replace the entitlement
	replacedEntitlement, resp, err := apiClient.EntitlementsAPI.ReplaceEntitlement(apiClient.cfg.Context, entitlementId).EntitlementsFullWithParent(buildEntitlementFullWithParent(entitlementId)).Execute()
	if err != nil {
		t.Errorf("Error replacing entitlement: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	if replacedEntitlement.Id != entitlementId {
		t.Errorf("Expected replaced entitlement ID %s, got %s", entitlementId, replacedEntitlement.Id)
		return
	}
}

func TestUpdateEntitlement(t *testing.T) {
	//entitlementId := "espyyzpbwwzhBGRAf1d6" // Replace with a valid entitlement ID
	//
	//// Build the updated entitlement data
	//updatedEntitlement := buildEntitlementFullWithParent(entitlementId)
	//
	//// Update the entitlement
	//updatedResponse, resp, err := apiClient.EntitlementsAPI.UpdateEntitlement(apiClient.cfg.Context, entitlementId).EntitlementPatchInner().Execute()
	//if err != nil {
	//	t.Errorf("Error updating entitlement: %v", err)
	//	return
	//}
	//if resp.StatusCode != 200 {
	//	t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	//	return
	//}
	//
	//if updatedResponse.Id != entitlementId {
	//	t.Errorf("Expected updated entitlement ID %s, got %s", entitlementId, updatedResponse.Id)
	//	return
	//}
}

func TestApiListAllEntitlementValuesRequest_Execute(t *testing.T) {
	entitlementId := "espyyzpbwwzhBGRAf1d6"

	entitlementValues, resp, err := apiClient.EntitlementsAPI.ListEntitlementValues(apiClient.cfg.Context, entitlementId).Execute()
	if err != nil {
		t.Errorf("Error replacing entitlement: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	if len(entitlementValues.Data) == 0 {
		t.Errorf("Expected at least one entitlement value, got %d", len(entitlementValues.Data))
		return
	}
}

func TestRetrieveEntitlementValue(t *testing.T) {
	entitlementId := "espyyzpbwwzhBGRAf1d6" // Replace with a valid entitlement ID
	valueId := "entyyzpbxQsQP2ej61d6"       // Replace with a valid value ID

	entitlementValue, resp, err := apiClient.EntitlementsAPI.GetEntitlementValue(apiClient.cfg.Context, entitlementId, valueId).Execute()
	if err != nil {
		t.Errorf("Error retrieving entitlement value: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	if entitlementValue.Id != valueId {
		t.Errorf("Expected entitlement value ID %s, got %s", valueId, entitlementValue.Id)
		return
	}
}

func TestRetrieveAllEntitlementValuesRequest_Execute(t *testing.T) {
	filter := `parent.externalId eq "0oao1otjwdwdc8OEd6z61d7" AND parent.type eq "APPLICATION"`
	entitlementValues, resp, err := apiClient.EntitlementsAPI.ListAllEntitlementValues(apiClient.cfg.Context).Filter(filter).Execute()
	if err != nil {
		t.Errorf("Error retrieving all entitlement values: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	if len(entitlementValues.Data) == 0 {
		t.Errorf("Expected at least one entitlement value, got %d", len(entitlementValues.Data))
		return
	}
}

func buildEntitlementFullWithParent(entitlementId string) EntitlementsFullWithParent {
	str := "Test Value Full"
	return EntitlementsFullWithParent{
		Name:              "Test Entitlement Full",
		ExternalValue:     "test-external-value-full",
		Id:                entitlementId,
		Description:       nil,      // Optional field, can be set to a string if needed
		DataType:          "string", // Example data type, can be adjusted as needed
		Parent:            TargetResource{Type: "APPLICATION", ExternalId: "0oankw1sc6z965dIC1d7"},
		ParentResourceOrn: "orn:oktapreview:idp:00onkw1sbuAh3Q06I1d7:apps:saasure:0oankw1sc6z965dIC1d7",
		Values: []EntitlementValueFull{
			{
				Name:          &str,
				ExternalValue: &str,
			},
		},
	}
}

func buildEntitlement() EntitlementCreate {
	return EntitlementCreate{
		Name:          "Test Entitlement",
		ExternalValue: "test-external-value",
		Description:   nil,      // Optional field, can be set to a string if needed
		DataType:      "string", // Example data type, can be adjusted as needed
		Parent:        &TargetResource{Type: "APPLICATION", ExternalId: "0oao1otjwc8OEd6z61d7"},
		Values: []EntitlementValueWritableProperties{
			{
				Name:          "Test Value",
				ExternalValue: "test-value",
				Description:   nil, // Optional field, can be set to a string if needed
			},
		},
		MultiValue: true,
	}
}
