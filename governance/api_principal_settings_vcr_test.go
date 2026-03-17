//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrincipalSettingsAPI_UpdatePrincipalSettings(t *testing.T) {
	cassetteName := CassetteName("principal_settings", "update_principal_settings")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list delegate appointments to find a valid principal ID
	delegatesResp, _, err := client.DelegatesAPI.ListDelegateAppointments(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(delegatesResp.Data) == 0 {
		t.Skip("No delegate appointments available to get a valid principal ID")
	}

	// Use the delegator's external ID (Okta user ID)
	principalId := delegatesResp.Data[0].Delegator.ExternalId

	// Create update request with delegates settings
	updateRequest := PrincipalSettingsPatchable{
		Delegates: NewDelegatesPatchable(),
	}

	resp, httpRes, err := client.PrincipalSettingsAPI.UpdatePrincipalSettings(ctx, principalId).
		PrincipalSettingsPatchable(updateRequest).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
