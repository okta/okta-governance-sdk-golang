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

func Test_governance_EntitlementsAPIService(t *testing.T) {

	configuration, err := okta.NewConfiguration()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EntitlementsAPIService CreateEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntitlementsAPI.CreateEntitlement(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService DeleteEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var entitlementId string

		httpRes, err := apiClient.EntitlementsAPI.DeleteEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService GetEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementsAPI.GetEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService GetEntitlementValue", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var entitlementId string
		var valueId string

		resp, httpRes, err := apiClient.EntitlementsAPI.GetEntitlementValue(context.Background(), entitlementId, valueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService ListAllEntitlementValues", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntitlementsAPI.ListAllEntitlementValues(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService ListEntitlementValues", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementsAPI.ListEntitlementValues(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService ListEntitlements", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EntitlementsAPI.ListEntitlements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService ReplaceEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementsAPI.ReplaceEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService UpdateEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementsAPI.UpdateEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
