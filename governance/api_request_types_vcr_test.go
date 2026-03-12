//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestTypesAPI_ListAllRequestTeams(t *testing.T) {
	cassetteName := CassetteName("request_types", "list_all_request_teams")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.RequestTypesAPI.ListAllRequestTeams(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)

	// Verify we got real team data with expected fields
	require.True(t, len(resp.Data) > 0, "Expected at least one team")
	team := resp.Data[0]
	assert.NotEmpty(t, team.Id, "Team should have an ID")
	assert.NotEmpty(t, team.Name, "Team should have a name")
	assert.NotNil(t, team.Created, "Team should have a created timestamp")
}
