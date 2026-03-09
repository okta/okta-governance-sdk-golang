//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrantsAPI_ListGrants(t *testing.T) {
	cassetteName := CassetteName("grants", "list_grants")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.GrantsAPI.ListGrants(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestGrantsAPI_ListGrantsWithFilter(t *testing.T) {
	cassetteName := CassetteName("grants", "list_grants_filtered")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Filter grants by target resource
	filter := `target.externalId eq "` + TestApplicationId + `" AND target.type eq "APPLICATION"`
	resp, httpRes, err := client.GrantsAPI.ListGrants(ctx).Filter(filter).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestGrantsAPI_ListGrantsWithEntitlements(t *testing.T) {
	cassetteName := CassetteName("grants", "list_grants_with_entitlements")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// List grants with full entitlements included
	fields := []string{"full_entitlements"}
	resp, httpRes, err := client.GrantsAPI.ListGrants(ctx).Include(fields).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestGrantsAPI_GetGrant(t *testing.T) {
	cassetteName := CassetteName("grants", "get_grant")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list grants to get a valid grant ID
	listResp, _, err := client.GrantsAPI.ListGrants(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if listResp.GrantsList == nil || len(listResp.GrantsList.Data) == 0 {
		if listResp.GrantsListWithEntitlements == nil || len(listResp.GrantsListWithEntitlements.Data) == 0 {
			t.Skip("No grants available to test GetGrant")
		}
	}

	var grantId string
	if listResp.GrantsList != nil && len(listResp.GrantsList.Data) > 0 {
		grantId = listResp.GrantsList.Data[0].Id
	} else if listResp.GrantsListWithEntitlements != nil && len(listResp.GrantsListWithEntitlements.Data) > 0 {
		grantId = listResp.GrantsListWithEntitlements.Data[0].Id
	}

	if grantId == "" {
		t.Skip("No grant ID available")
	}

	// Get the specific grant
	resp, httpRes, err := client.GrantsAPI.GetGrant(ctx, grantId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestGrantsAPI_CreateGrant(t *testing.T) {
	cassetteName := CassetteName("grants", "create_grant")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Create grant
	grant := BuildTestGrant()
	createdGrant, httpRes, err := client.GrantsAPI.CreateGrant(ctx).GrantCreatable(grant).Execute()

	require.NoError(t, err)
	require.NotNil(t, createdGrant)
	assert.Equal(t, 201, httpRes.StatusCode)
	assert.NotEmpty(t, createdGrant.Id)
}
