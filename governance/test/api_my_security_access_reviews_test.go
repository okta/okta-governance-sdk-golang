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

func Test_governance_MySecurityAccessReviewsAPIService(t *testing.T) {

	configuration, err := okta.NewConfiguration()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MySecurityAccessReviewsAPIService AddMySecurityAccessReviewComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		httpRes, err := apiClient.MySecurityAccessReviewsAPI.AddMySecurityAccessReviewComment(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ExecuteMySecurityAccessReviewAccessesAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewTargetId string

		httpRes, err := apiClient.MySecurityAccessReviewsAPI.ExecuteMySecurityAccessReviewAccessesAction(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ExecuteMySecurityAccessReviewAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		httpRes, err := apiClient.MySecurityAccessReviewsAPI.ExecuteMySecurityAccessReviewAction(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService GenerateMySecurityAccessReviewAccessesSummary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewTargetId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewAccessesSummary(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService GenerateMySecurityAccessReviewSummary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewSummary(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService GetMySecurityAccessReview", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.GetMySecurityAccessReview(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService GetMySecurityAccessReviewPrincipalDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.GetMySecurityAccessReviewPrincipalDetails(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService GetMySecurityAccessReviewsStats", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.GetMySecurityAccessReviewsStats(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ListMySecurityAccessReviewAccesses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccesses(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ListMySecurityAccessReviewAccessesAnomalies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewTargetId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccessesAnomalies(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ListMySecurityAccessReviewActions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewActions(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ListMySecurityAccessReviewHistory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewHistory(context.Background(), securityAccessReviewId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ListMySecurityAccessReviewSubAccesses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var securityAccessReviewId string
		var securityAccessReviewAccessId string

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewSubAccesses(context.Background(), securityAccessReviewId, securityAccessReviewAccessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MySecurityAccessReviewsAPIService ListMySecurityAccessReviews", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviews(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
