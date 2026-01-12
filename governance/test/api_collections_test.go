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

func Test_governance_CollectionsAPIService(t *testing.T) {

	configuration, err := okta.NewConfiguration()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CollectionsAPIService AddResourcesToCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.AddResourcesToCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService AssignCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.AssignCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService CreateCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CollectionsAPI.CreateCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService DeleteCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		httpRes, err := apiClient.CollectionsAPI.DeleteCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService DeleteCollectionResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string
		var resourceId string

		httpRes, err := apiClient.CollectionsAPI.DeleteCollectionResource(context.Background(), collectionId, resourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService DeletePrincipalAssignment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string
		var assignmentId string

		httpRes, err := apiClient.CollectionsAPI.DeletePrincipalAssignment(context.Background(), collectionId, assignmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService GetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.GetCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService GetCollectionResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string
		var resourceId string

		resp, httpRes, err := apiClient.CollectionsAPI.GetCollectionResource(context.Background(), collectionId, resourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService GetUnassignedUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.GetUnassignedUsers(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService ListCollectionAssignments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.ListCollectionAssignments(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService ListCollectionResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.ListCollectionResources(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService ListCollections", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CollectionsAPI.ListCollections(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService ListCollectionsAssignments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CollectionsAPI.ListCollectionsAssignments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService ReplaceCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string

		resp, httpRes, err := apiClient.CollectionsAPI.ReplaceCollection(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService ReplaceCollectionResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string
		var resourceId string

		resp, httpRes, err := apiClient.CollectionsAPI.ReplaceCollectionResource(context.Background(), collectionId, resourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollectionsAPIService UpdatePrincipalAssignment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var collectionId string
		var assignmentId string

		httpRes, err := apiClient.CollectionsAPI.UpdatePrincipalAssignment(context.Background(), collectionId, assignmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
