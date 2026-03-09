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

// AccessScopeSettingsCreatableAccessScopeSettings - Settings for the access request scope (such as groups, entitlement bundles, or default resources)
type AccessScopeSettingsCreatableAccessScopeSettings struct {
	AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings
	AccessScopeSettingsCreatableGroupAccessScopeSettings             *AccessScopeSettingsCreatableGroupAccessScopeSettings
	AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings   *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings
}

// AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettingsAsAccessScopeSettingsCreatableAccessScopeSettings is a convenience function that returns AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings wrapped in AccessScopeSettingsCreatableAccessScopeSettings
func AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettingsAsAccessScopeSettingsCreatableAccessScopeSettings(v *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) AccessScopeSettingsCreatableAccessScopeSettings {
	return AccessScopeSettingsCreatableAccessScopeSettings{
		AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings: v,
	}
}

// AccessScopeSettingsCreatableGroupAccessScopeSettingsAsAccessScopeSettingsCreatableAccessScopeSettings is a convenience function that returns AccessScopeSettingsCreatableGroupAccessScopeSettings wrapped in AccessScopeSettingsCreatableAccessScopeSettings
func AccessScopeSettingsCreatableGroupAccessScopeSettingsAsAccessScopeSettingsCreatableAccessScopeSettings(v *AccessScopeSettingsCreatableGroupAccessScopeSettings) AccessScopeSettingsCreatableAccessScopeSettings {
	return AccessScopeSettingsCreatableAccessScopeSettings{
		AccessScopeSettingsCreatableGroupAccessScopeSettings: v,
	}
}

// AccessScopeSettingsCreatableResourceDefaultAccessScopeSettingsAsAccessScopeSettingsCreatableAccessScopeSettings is a convenience function that returns AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings wrapped in AccessScopeSettingsCreatableAccessScopeSettings
func AccessScopeSettingsCreatableResourceDefaultAccessScopeSettingsAsAccessScopeSettingsCreatableAccessScopeSettings(v *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) AccessScopeSettingsCreatableAccessScopeSettings {
	return AccessScopeSettingsCreatableAccessScopeSettings{
		AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccessScopeSettingsCreatableAccessScopeSettings) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ENTITLEMENT_BUNDLES'
	if jsonDict["type"] == "ENTITLEMENT_BUNDLES" {
		// try to unmarshal JSON data into AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings
		err = json.Unmarshal(data, &dst.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings)
		if err == nil {
			return nil // data stored in dst.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings, return on the first match
		} else {
			dst.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings = nil
			return fmt.Errorf("failed to unmarshal AccessScopeSettingsCreatableAccessScopeSettings as AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUPS'
	if jsonDict["type"] == "GROUPS" {
		// try to unmarshal JSON data into AccessScopeSettingsCreatableGroupAccessScopeSettings
		err = json.Unmarshal(data, &dst.AccessScopeSettingsCreatableGroupAccessScopeSettings)
		if err == nil {
			return nil // data stored in dst.AccessScopeSettingsCreatableGroupAccessScopeSettings, return on the first match
		} else {
			dst.AccessScopeSettingsCreatableGroupAccessScopeSettings = nil
			return fmt.Errorf("failed to unmarshal AccessScopeSettingsCreatableAccessScopeSettings as AccessScopeSettingsCreatableGroupAccessScopeSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESOURCE_DEFAULT'
	if jsonDict["type"] == "RESOURCE_DEFAULT" {
		// try to unmarshal JSON data into AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings
		err = json.Unmarshal(data, &dst.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings)
		if err == nil {
			return nil // data stored in dst.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings, return on the first match
		} else {
			dst.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings = nil
			return fmt.Errorf("failed to unmarshal AccessScopeSettingsCreatableAccessScopeSettings as AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccessScopeSettingsCreatableAccessScopeSettings) MarshalJSON() ([]byte, error) {
	if src.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings != nil {
		return json.Marshal(&src.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings)
	}

	if src.AccessScopeSettingsCreatableGroupAccessScopeSettings != nil {
		return json.Marshal(&src.AccessScopeSettingsCreatableGroupAccessScopeSettings)
	}

	if src.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings != nil {
		return json.Marshal(&src.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccessScopeSettingsCreatableAccessScopeSettings) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings != nil {
		return obj.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings
	}

	if obj.AccessScopeSettingsCreatableGroupAccessScopeSettings != nil {
		return obj.AccessScopeSettingsCreatableGroupAccessScopeSettings
	}

	if obj.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings != nil {
		return obj.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AccessScopeSettingsCreatableAccessScopeSettings) GetActualInstanceValue() interface{} {
	if obj.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings != nil {
		return *obj.AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings
	}

	if obj.AccessScopeSettingsCreatableGroupAccessScopeSettings != nil {
		return *obj.AccessScopeSettingsCreatableGroupAccessScopeSettings
	}

	if obj.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings != nil {
		return *obj.AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings
	}

	// all schemas are nil
	return nil
}

type NullableAccessScopeSettingsCreatableAccessScopeSettings struct {
	value *AccessScopeSettingsCreatableAccessScopeSettings
	isSet bool
}

func (v NullableAccessScopeSettingsCreatableAccessScopeSettings) Get() *AccessScopeSettingsCreatableAccessScopeSettings {
	return v.value
}

func (v *NullableAccessScopeSettingsCreatableAccessScopeSettings) Set(val *AccessScopeSettingsCreatableAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeSettingsCreatableAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeSettingsCreatableAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeSettingsCreatableAccessScopeSettings(val *AccessScopeSettingsCreatableAccessScopeSettings) *NullableAccessScopeSettingsCreatableAccessScopeSettings {
	return &NullableAccessScopeSettingsCreatableAccessScopeSettings{value: val, isSet: true}
}

func (v NullableAccessScopeSettingsCreatableAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeSettingsCreatableAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
