//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrincipalEntitlementsAPI_GetPrincipalEntitlements(t *testing.T) {
	cassetteName := CassetteName("principal_entitlements", "get_principal_entitlements")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Get principal entitlements requires filter with principal and resource
	// Format: principal.id eq "{userId}" AND resource.externalId eq "{appId}" AND resource.type eq "APPLICATION"
	filter := `principal.id eq "` + TestUserId + `" AND resource.externalId eq "` + TestApplicationId + `" AND resource.type eq "APPLICATION"`
	resp, httpRes, err := client.PrincipalEntitlementsAPI.GetPrincipalEntitlements(ctx).Filter(filter).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestPrincipalEntitlementsAPI_GetPrincipalEntitlementsHistory(t *testing.T) {
	cassetteName := CassetteName("principal_entitlements", "get_principal_entitlements_history")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Get principal entitlements history requires filter with principal and resource
	filter := `principal.id eq "` + TestUserId + `" AND resource.externalId eq "` + TestApplicationId + `" AND resource.type eq "APPLICATION"`
	resp, httpRes, err := client.PrincipalEntitlementsAPI.GetPrincipalEntitlementsHistory(ctx).Filter(filter).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
