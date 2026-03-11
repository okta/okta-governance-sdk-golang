//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCampaignsAPI_ListCampaigns(t *testing.T) {
	cassetteName := CassetteName("campaigns", "list_campaigns")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.CampaignsAPI.ListCampaigns(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.NotNil(t, resp.Data)
}

func TestCampaignsAPI_GetCampaign(t *testing.T) {
	cassetteName := CassetteName("campaigns", "get_campaign")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list campaigns to get a valid campaign ID
	listResp, _, err := client.CampaignsAPI.ListCampaigns(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No campaigns available to test GetCampaign")
	}

	campaignId := listResp.Data[0].Id

	// Get the campaign
	gotCampaign, httpRes, err := client.CampaignsAPI.GetCampaign(ctx, campaignId).Execute()

	require.NoError(t, err)
	require.NotNil(t, gotCampaign)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, campaignId, gotCampaign.Id)
}
