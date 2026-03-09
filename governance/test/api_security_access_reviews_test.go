/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.2.0
Contact: devex-public@okta.com
*/

package governance_test

import (
	"context"
	"testing"

	openapiclient "github.com/okta/okta-governance-sdk-golang/governance"
	"github.com/okta/okta-sdk-golang/v6/okta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_governance_SecurityAccessReviewsAPIService(t *testing.T) {

	configuration, err := okta.NewConfiguration()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecurityAccessReviewsAPIService AddSecurityAccessReviewComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		httpRes, err := apiClient.SecurityAccessReviewsAPI.AddSecurityAccessReviewComment(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService CreateSecurityAccessReview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.CreateSecurityAccessReview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ExecuteSecurityAccessReviewAccessesAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewTargetId string

		httpRes, err := apiClient.SecurityAccessReviewsAPI.ExecuteSecurityAccessReviewAccessesAction(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ExecuteSecurityAccessReviewAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		httpRes, err := apiClient.SecurityAccessReviewsAPI.ExecuteSecurityAccessReviewAction(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService GenerateSecurityAccessReviewAccessesSummary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewTargetId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.GenerateSecurityAccessReviewAccessesSummary(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService GenerateSecurityAccessReviewSummary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.GenerateSecurityAccessReviewSummary(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService GetSecurityAccessReview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.GetSecurityAccessReview(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService GetSecurityAccessReviewPrincipalDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.GetSecurityAccessReviewPrincipalDetails(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService GetSecurityAccessReviewsStats", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.GetSecurityAccessReviewsStats(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ListSecurityAccessReviewAccesses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ListSecurityAccessReviewAccessesAnomalies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewTargetId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccessesAnomalies(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ListSecurityAccessReviewActions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewActions(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ListSecurityAccessReviewHistory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewHistory(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ListSecurityAccessReviewSubAccesses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewAccessId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewSubAccesses(context.Background(), securityAccessReviewId, securityAccessReviewAccessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService ListSecurityAccessReviews", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviews(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAccessReviewsAPIService UpdateSecurityAccessReview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.SecurityAccessReviewsAPI.UpdateSecurityAccessReview(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
