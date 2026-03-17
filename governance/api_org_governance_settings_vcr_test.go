//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrgGovernanceSettingsAPI_GetOrgSettings(t *testing.T) {
	cassetteName := CassetteName("org_governance_settings", "get_org_settings")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.OrgGovernanceSettingsAPI.GetOrgSettings(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
