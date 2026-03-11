//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntitlementsAPI_ListEntitlements(t *testing.T) {
	cassetteName := CassetteName("entitlements", "list_entitlements")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	// Filter is required - must have both parent.externalId and parent.type
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	resp, httpRes, err := client.EntitlementsAPI.ListEntitlements(ctx).Filter(filter).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.NotNil(t, resp.Data)
}

func TestEntitlementsAPI_ListEntitlementsWithFilter(t *testing.T) {
	cassetteName := CassetteName("entitlements", "list_entitlements_filtered")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Filter by application
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	resp, httpRes, err := client.EntitlementsAPI.ListEntitlements(ctx).Filter(filter).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestEntitlementsAPI_GetEntitlement(t *testing.T) {
	cassetteName := CassetteName("entitlements", "get_entitlement")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list entitlements to get a valid entitlement ID (filter is required)
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	listResp, _, err := client.EntitlementsAPI.ListEntitlements(ctx).Filter(filter).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No entitlements available to test GetEntitlement")
	}

	entitlementId := listResp.Data[0].Id

	// Get the specific entitlement
	resp, httpRes, err := client.EntitlementsAPI.GetEntitlement(ctx, entitlementId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, entitlementId, resp.Id)
}

func TestEntitlementsAPI_ListEntitlementValues(t *testing.T) {
	cassetteName := CassetteName("entitlements", "list_entitlement_values")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list entitlements to get a valid entitlement ID (filter is required)
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	listResp, _, err := client.EntitlementsAPI.ListEntitlements(ctx).Filter(filter).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No entitlements available to test ListEntitlementValues")
	}

	entitlementId := listResp.Data[0].Id

	// List entitlement values
	resp, httpRes, err := client.EntitlementsAPI.ListEntitlementValues(ctx, entitlementId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestEntitlementsAPI_ListAllEntitlementValues(t *testing.T) {
	cassetteName := CassetteName("entitlements", "list_all_entitlement_values")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// List all entitlement values with filter
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	resp, httpRes, err := client.EntitlementsAPI.ListAllEntitlementValues(ctx).Filter(filter).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestEntitlementsAPI_GetEntitlementValue(t *testing.T) {
	cassetteName := CassetteName("entitlements", "get_entitlement_value")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list entitlements to get a valid entitlement ID (filter is required)
	filter := `parent.externalId eq "` + TestApplicationId + `" AND parent.type eq "APPLICATION"`
	listResp, _, err := client.EntitlementsAPI.ListEntitlements(ctx).Filter(filter).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No entitlements available to test GetEntitlementValue")
	}

	entitlementId := listResp.Data[0].Id

	// List entitlement values to get a valid value ID
	valuesResp, _, err := client.EntitlementsAPI.ListEntitlementValues(ctx, entitlementId).Limit(1).Execute()
	require.NoError(t, err)

	if len(valuesResp.Data) == 0 {
		t.Skip("No entitlement values available to test GetEntitlementValue")
	}

	valueId := valuesResp.Data[0].Id

	// Get the specific entitlement value
	resp, httpRes, err := client.EntitlementsAPI.GetEntitlementValue(ctx, entitlementId, valueId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, valueId, resp.Id)
}
