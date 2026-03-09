//go:build unit

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
	assert.NotNil(t, resp.Data)
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

	bundleId := listResp.Data[0].Id

	// Get the specific entitlement bundle
	resp, httpRes, err := client.EntitlementBundlesAPI.GetEntitlementBundle(ctx, bundleId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, bundleId, resp.Id)
}
