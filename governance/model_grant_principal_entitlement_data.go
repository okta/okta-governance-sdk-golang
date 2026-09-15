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

package governance

import (
	"encoding/json"
	"fmt"
)

// GrantPrincipalEntitlementData - Specifies the entitlements to grant to the principal. Discriminated by `type`: - `ENTITLEMENT`: Direct entitlement assignment with grouped entitlement values - `ENTITLEMENT_BUNDLE`: Entitlement bundle assignment
type GrantPrincipalEntitlementData struct {
	GrantPrincipalEntitlementDataBundle      *GrantPrincipalEntitlementDataBundle
	GrantPrincipalEntitlementDataEntitlement *GrantPrincipalEntitlementDataEntitlement
}

// GrantPrincipalEntitlementDataBundleAsGrantPrincipalEntitlementData is a convenience function that returns GrantPrincipalEntitlementDataBundle wrapped in GrantPrincipalEntitlementData
func GrantPrincipalEntitlementDataBundleAsGrantPrincipalEntitlementData(v *GrantPrincipalEntitlementDataBundle) GrantPrincipalEntitlementData {
	return GrantPrincipalEntitlementData{
		GrantPrincipalEntitlementDataBundle: v,
	}
}

// GrantPrincipalEntitlementDataEntitlementAsGrantPrincipalEntitlementData is a convenience function that returns GrantPrincipalEntitlementDataEntitlement wrapped in GrantPrincipalEntitlementData
func GrantPrincipalEntitlementDataEntitlementAsGrantPrincipalEntitlementData(v *GrantPrincipalEntitlementDataEntitlement) GrantPrincipalEntitlementData {
	return GrantPrincipalEntitlementData{
		GrantPrincipalEntitlementDataEntitlement: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GrantPrincipalEntitlementData) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ENTITLEMENT'
	if jsonDict["type"] == "ENTITLEMENT" {
		// try to unmarshal JSON data into GrantPrincipalEntitlementDataEntitlement
		err = json.Unmarshal(data, &dst.GrantPrincipalEntitlementDataEntitlement)
		if err == nil {
			return nil // data stored in dst.GrantPrincipalEntitlementDataEntitlement, return on the first match
		} else {
			dst.GrantPrincipalEntitlementDataEntitlement = nil
			return fmt.Errorf("failed to unmarshal GrantPrincipalEntitlementData as GrantPrincipalEntitlementDataEntitlement: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITLEMENT_BUNDLE'
	if jsonDict["type"] == "ENTITLEMENT_BUNDLE" {
		// try to unmarshal JSON data into GrantPrincipalEntitlementDataBundle
		err = json.Unmarshal(data, &dst.GrantPrincipalEntitlementDataBundle)
		if err == nil {
			return nil // data stored in dst.GrantPrincipalEntitlementDataBundle, return on the first match
		} else {
			dst.GrantPrincipalEntitlementDataBundle = nil
			return fmt.Errorf("failed to unmarshal GrantPrincipalEntitlementData as GrantPrincipalEntitlementDataBundle: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GrantPrincipalEntitlementData) MarshalJSON() ([]byte, error) {
	if src.GrantPrincipalEntitlementDataBundle != nil {
		return json.Marshal(&src.GrantPrincipalEntitlementDataBundle)
	}

	if src.GrantPrincipalEntitlementDataEntitlement != nil {
		return json.Marshal(&src.GrantPrincipalEntitlementDataEntitlement)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GrantPrincipalEntitlementData) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GrantPrincipalEntitlementDataBundle != nil {
		return obj.GrantPrincipalEntitlementDataBundle
	}

	if obj.GrantPrincipalEntitlementDataEntitlement != nil {
		return obj.GrantPrincipalEntitlementDataEntitlement
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GrantPrincipalEntitlementData) GetActualInstanceValue() interface{} {
	if obj.GrantPrincipalEntitlementDataBundle != nil {
		return *obj.GrantPrincipalEntitlementDataBundle
	}

	if obj.GrantPrincipalEntitlementDataEntitlement != nil {
		return *obj.GrantPrincipalEntitlementDataEntitlement
	}

	// all schemas are nil
	return nil
}

type NullableGrantPrincipalEntitlementData struct {
	value *GrantPrincipalEntitlementData
	isSet bool
}

func (v NullableGrantPrincipalEntitlementData) Get() *GrantPrincipalEntitlementData {
	return v.value
}

func (v *NullableGrantPrincipalEntitlementData) Set(val *GrantPrincipalEntitlementData) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalEntitlementData) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalEntitlementData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalEntitlementData(val *GrantPrincipalEntitlementData) *NullableGrantPrincipalEntitlementData {
	return &NullableGrantPrincipalEntitlementData{value: val, isSet: true}
}

func (v NullableGrantPrincipalEntitlementData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalEntitlementData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
