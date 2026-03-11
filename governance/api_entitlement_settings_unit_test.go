//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntitlementSettingsAPI_GetEntitlementSettings(t *testing.T) {
	cassetteName := CassetteName("entitlement_settings", "get_entitlement_settings")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list entitlements to get a valid resource ORN (filter is required)
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	listResp, _, err := client.EntitlementsAPI.ListEntitlements(ctx).Filter(filter).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No entitlements available to test GetEntitlementSettings")
	}

	// Get the parent resource ORN from the first entitlement
	entitlement := listResp.Data[0]
	resourceOrn := entitlement.ParentResourceOrn

	// Get entitlement settings for the resource
	resp, httpRes, err := client.EntitlementSettingsAPI.GetEntitlementSettings(ctx, resourceOrn).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
