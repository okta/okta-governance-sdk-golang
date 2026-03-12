//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityAccessReviewsAPI_ListSecurityAccessReviews(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "list_security_access_reviews")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_GetSecurityAccessReviewsStats(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "get_stats")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.SecurityAccessReviewsAPI.GetSecurityAccessReviewsStats(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_GetSecurityAccessReview(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "get_security_access_review")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	listResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No security access reviews available to test GetSecurityAccessReview")
	}

	reviewId := listResp.Data[0].Id

	// Get the specific review
	resp, httpRes, err := client.SecurityAccessReviewsAPI.GetSecurityAccessReview(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, reviewId, resp.Id)
}

func TestSecurityAccessReviewsAPI_GetSecurityAccessReviewPrincipalDetails(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "get_principal_details")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	listResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No security access reviews available to test GetSecurityAccessReviewPrincipalDetails")
	}

	reviewId := listResp.Data[0].Id

	// Get the principal details
	resp, httpRes, err := client.SecurityAccessReviewsAPI.GetSecurityAccessReviewPrincipalDetails(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_ListSecurityAccessReviewAccesses(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "list_accesses")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	listResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No security access reviews available to test ListSecurityAccessReviewAccesses")
	}

	reviewId := listResp.Data[0].Id

	// List accesses for the review
	resp, httpRes, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_ListSecurityAccessReviewActions(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "list_actions")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	listResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No security access reviews available to test ListSecurityAccessReviewActions")
	}

	reviewId := listResp.Data[0].Id

	// List actions for the review
	resp, httpRes, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewActions(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_ListSecurityAccessReviewHistory(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "list_history")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	listResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No security access reviews available to test ListSecurityAccessReviewHistory")
	}

	reviewId := listResp.Data[0].Id

	// List history for the review
	resp, httpRes, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewHistory(ctx, reviewId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_ListSecurityAccessReviewSubAccesses(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "list_sub_accesses")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	reviewsResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(reviewsResp.Data) == 0 {
		t.Skip("No security access reviews available to test ListSecurityAccessReviewSubAccesses")
	}

	reviewId := reviewsResp.Data[0].Id

	// List accesses to get a valid access ID
	accessesResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses(ctx, reviewId).Execute()
	require.NoError(t, err)

	if len(accessesResp.Data) == 0 {
		t.Skip("No accesses available to test ListSecurityAccessReviewSubAccesses")
	}

	accessId := accessesResp.Data[0].Id

	// List sub-accesses for the access
	resp, httpRes, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewSubAccesses(ctx, reviewId, accessId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestSecurityAccessReviewsAPI_ListSecurityAccessReviewAccessesAnomalies(t *testing.T) {
	cassetteName := CassetteName("security_access_reviews", "list_accesses_anomalies")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list security access reviews to get a valid ID
	reviewsResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviews(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(reviewsResp.Data) == 0 {
		t.Skip("No security access reviews available to test ListSecurityAccessReviewAccessesAnomalies")
	}

	reviewId := reviewsResp.Data[0].Id

	// List accesses to get a valid access ID
	accessesResp, _, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses(ctx, reviewId).Execute()
	require.NoError(t, err)

	if len(accessesResp.Data) == 0 {
		t.Skip("No accesses available to test ListSecurityAccessReviewAccessesAnomalies")
	}

	accessId := accessesResp.Data[0].Id

	// List anomalies for the access
	resp, httpRes, err := client.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccessesAnomalies(ctx, reviewId, accessId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
