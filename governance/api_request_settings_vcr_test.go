//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestSettingsAPI_GetOrgRequestSettingsV2(t *testing.T) {
	cassetteName := CassetteName("request_settings", "get_org_request_settings")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.RequestSettingsAPI.GetOrgRequestSettingsV2(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
