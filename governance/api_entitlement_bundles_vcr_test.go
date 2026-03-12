//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntitlementBundlesAPI_ListEntitlementBundles(t *testing.T) {
	cassetteName := CassetteName("entitlement_bundles", "list_entitlement_bundles")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.EntitlementBundlesAPI.ListEntitlementBundles(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	// Verify we got real bundle data with expected structure
	require.True(t, len(resp.Data) > 0, "Expected at least one entitlement bundle")
	bundle := resp.Data[0]
	assert.NotEmpty(t, bundle.Id, "Bundle should have an ID")
	assert.NotEmpty(t, bundle.Name, "Bundle should have a name")
	assert.NotEmpty(t, bundle.Orn, "Bundle should have an ORN")
	assert.NotEmpty(t, bundle.Status, "Bundle should have a status")
}

func TestEntitlementBundlesAPI_GetEntitlementBundle(t *testing.T) {
	cassetteName := CassetteName("entitlement_bundles", "get_entitlement_bundle")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list entitlement bundles to get a valid ID
	listResp, _, err := client.EntitlementBundlesAPI.ListEntitlementBundles(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No entitlement bundles available to test GetEntitlementBundle")
	}

	bundleFromList := listResp.Data[0]
	bundleId := bundleFromList.Id

	// Get the specific entitlement bundle
	resp, httpRes, err := client.EntitlementBundlesAPI.GetEntitlementBundle(ctx, bundleId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	// Verify the retrieved bundle matches what we listed
	assert.Equal(t, bundleId, resp.Id)
	assert.Equal(t, bundleFromList.Name, resp.Name)
	assert.Equal(t, bundleFromList.Orn, resp.Orn)
	assert.NotEmpty(t, resp.Description, "Bundle should have a description")
	assert.NotNil(t, resp.Created, "Bundle should have a created timestamp")
}
