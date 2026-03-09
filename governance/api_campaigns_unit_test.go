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

func TestCampaignsAPI_CreateAndDeleteCampaign(t *testing.T) {
	cassetteName := CassetteName("campaigns", "create_delete_campaign")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Create campaign
	campaign := BuildTestCampaign()
	createdCampaign, httpRes, err := client.CampaignsAPI.CreateCampaign(ctx).CampaignMutable(campaign).Execute()

	require.NoError(t, err)
	require.NotNil(t, createdCampaign)
	assert.Equal(t, 201, httpRes.StatusCode)
	assert.NotEmpty(t, createdCampaign.Id)

	// Clean up - delete the campaign
	t.Cleanup(func() {
		if createdCampaign.Id != "" {
			_, _ = client.CampaignsAPI.DeleteCampaign(ctx, createdCampaign.Id).Execute()
		}
	})
}

func TestCampaignsAPI_GetCampaign(t *testing.T) {
	cassetteName := CassetteName("campaigns", "get_campaign")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First create a campaign
	campaign := BuildTestCampaign()
	createdCampaign, _, err := client.CampaignsAPI.CreateCampaign(ctx).CampaignMutable(campaign).Execute()
	require.NoError(t, err)
	require.NotNil(t, createdCampaign)

	t.Cleanup(func() {
		if createdCampaign.Id != "" {
			_, _ = client.CampaignsAPI.DeleteCampaign(ctx, createdCampaign.Id).Execute()
		}
	})

	// Get the campaign
	gotCampaign, httpRes, err := client.CampaignsAPI.GetCampaign(ctx, createdCampaign.Id).Execute()

	require.NoError(t, err)
	require.NotNil(t, gotCampaign)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, createdCampaign.Id, gotCampaign.Id)
}

func TestCampaignsAPI_LaunchAndEndCampaign(t *testing.T) {
	cassetteName := CassetteName("campaigns", "launch_end_campaign")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Create a campaign
	campaign := BuildTestCampaign()
	createdCampaign, _, err := client.CampaignsAPI.CreateCampaign(ctx).CampaignMutable(campaign).Execute()
	require.NoError(t, err)
	require.NotNil(t, createdCampaign)

	// Launch the campaign
	httpRes, err := client.CampaignsAPI.LaunchCampaign(ctx, createdCampaign.Id).Execute()
	require.NoError(t, err)
	assert.Equal(t, 202, httpRes.StatusCode)

	// End the campaign
	httpRes, err = client.CampaignsAPI.EndCampaign(ctx, createdCampaign.Id).Execute()
	require.NoError(t, err)
	assert.Equal(t, 202, httpRes.StatusCode)
}
