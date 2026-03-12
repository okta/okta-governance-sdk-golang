//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReviewsAPI_ListReviews(t *testing.T) {
	cassetteName := CassetteName("reviews", "list_reviews")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.ReviewsAPI.ListReviews(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.NotNil(t, resp.Data)
}

func TestReviewsAPI_GetReview(t *testing.T) {
	cassetteName := CassetteName("reviews", "get_review")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list reviews to get a valid review ID
	listResp, _, err := client.ReviewsAPI.ListReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No reviews available to test GetReview")
	}

	reviewId := listResp.Data[0].Id

	// Get the specific review
	resp, httpRes, err := client.ReviewsAPI.GetReview(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, reviewId, resp.Id)
}

func TestReviewsAPI_GetReviewWithResources(t *testing.T) {
	cassetteName := CassetteName("reviews", "get_review_with_resources")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list reviews to get a valid review ID
	listResp, _, err := client.ReviewsAPI.ListReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No reviews available to test GetReview with resources")
	}

	reviewId := listResp.Data[0].Id

	// Get the specific review (the API doesn't have a separate GetReviewAssignedResources)
	resp, httpRes, err := client.ReviewsAPI.GetReview(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestReviewsAPI_ListReviewsToReassign(t *testing.T) {
	cassetteName := CassetteName("reviews", "list_reviews_to_reassign")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list campaigns to get a valid campaign ID
	campaignsResp, _, err := client.CampaignsAPI.ListCampaigns(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(campaignsResp.Data) == 0 {
		t.Skip("No campaigns available to test ReassignReviews")
	}

	campaignId := campaignsResp.Data[0].Id

	// List reviews for the campaign
	resp, httpRes, err := client.ReviewsAPI.ListReviews(ctx).
		Filter(`campaignId eq "` + campaignId + `"`).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
