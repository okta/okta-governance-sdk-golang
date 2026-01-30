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

func Test_governance_GrantsAPIService(t *testing.T) {

	configuration, err := okta.NewConfiguration()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GrantsAPIService CreateGrant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GrantsAPI.CreateGrant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GrantsAPIService GetGrant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var grantId string

		resp, httpRes, err := apiClient.GrantsAPI.GetGrant(context.Background(), grantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GrantsAPIService ListGrants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GrantsAPI.ListGrants(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GrantsAPIService ReplaceGrant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var grantId string

		resp, httpRes, err := apiClient.GrantsAPI.ReplaceGrant(context.Background(), grantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GrantsAPIService UpdateGrant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var grantId string

		resp, httpRes, err := apiClient.GrantsAPI.UpdateGrant(context.Background(), grantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
