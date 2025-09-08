/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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

package governance

import (
	"encoding/json"
	"fmt"
)

// model_oneof.mustache
// GetentitlementBundle200Response - struct for GetentitlementBundle200Response
type GetentitlementBundle200Response struct {
	EntitlementBundleFull                 *EntitlementBundleFull
	EntitlementBundleFullWithEntitlements *EntitlementBundleFullWithEntitlements
}

// EntitlementBundleFullAsGetentitlementBundle200Response is a convenience function that returns EntitlementBundleFull wrapped in GetentitlementBundle200Response
func EntitlementBundleFullAsGetentitlementBundle200Response(v *EntitlementBundleFull) GetentitlementBundle200Response {
	return GetentitlementBundle200Response{
		EntitlementBundleFull: v,
	}
}

// EntitlementBundleFullWithEntitlementsAsGetentitlementBundle200Response is a convenience function that returns EntitlementBundleFullWithEntitlements wrapped in GetentitlementBundle200Response
func EntitlementBundleFullWithEntitlementsAsGetentitlementBundle200Response(v *EntitlementBundleFullWithEntitlements) GetentitlementBundle200Response {
	return GetentitlementBundle200Response{
		EntitlementBundleFullWithEntitlements: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *GetentitlementBundle200Response) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'entitlement-bundle-full'
	if jsonDict["entitlementsObjectType"] == "entitlement-bundle-full" {
		// try to unmarshal JSON data into EntitlementBundleFull
		err = json.Unmarshal(data, &dst.EntitlementBundleFull)
		if err == nil {
			return nil // data stored in dst.EntitlementBundleFull, return on the first match
		} else {
			dst.EntitlementBundleFull = nil
			return fmt.Errorf("Failed to unmarshal GetentitlementBundle200Response as EntitlementBundleFull: %s", err.Error())
		}
	}

	// check if the discriminator value is 'entitlement-bundle-full-with-entitlements'
	if jsonDict["entitlementsObjectType"] == "entitlement-bundle-full-with-entitlements" {
		// try to unmarshal JSON data into EntitlementBundleFullWithEntitlements
		err = json.Unmarshal(data, &dst.EntitlementBundleFullWithEntitlements)
		if err == nil {
			return nil // data stored in dst.EntitlementBundleFullWithEntitlements, return on the first match
		} else {
			dst.EntitlementBundleFullWithEntitlements = nil
			return fmt.Errorf("Failed to unmarshal GetentitlementBundle200Response as EntitlementBundleFullWithEntitlements: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetentitlementBundle200Response) MarshalJSON() ([]byte, error) {
	if src.EntitlementBundleFull != nil {
		return json.Marshal(&src.EntitlementBundleFull)
	}

	if src.EntitlementBundleFullWithEntitlements != nil {
		return json.Marshal(&src.EntitlementBundleFullWithEntitlements)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetentitlementBundle200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementBundleFull != nil {
		return obj.EntitlementBundleFull
	}

	if obj.EntitlementBundleFullWithEntitlements != nil {
		return obj.EntitlementBundleFullWithEntitlements
	}

	// all schemas are nil
	return nil
}

type NullableGetentitlementBundle200Response struct {
	value *GetentitlementBundle200Response
	isSet bool
}

func (v NullableGetentitlementBundle200Response) Get() *GetentitlementBundle200Response {
	return v.value
}

func (v *NullableGetentitlementBundle200Response) Set(val *GetentitlementBundle200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetentitlementBundle200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetentitlementBundle200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetentitlementBundle200Response(val *GetentitlementBundle200Response) *NullableGetentitlementBundle200Response {
	return &NullableGetentitlementBundle200Response{value: val, isSet: true}
}

func (v NullableGetentitlementBundle200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetentitlementBundle200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
