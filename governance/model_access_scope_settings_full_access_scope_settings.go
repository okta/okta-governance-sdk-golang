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

// AccessScopeSettingsFullAccessScopeSettings - Settings specifying if groups or entitlements may be used for requesting finer grained access
type AccessScopeSettingsFullAccessScopeSettings struct {
	EntitlementBundleAccessScopeSettings *EntitlementBundleAccessScopeSettings
	GroupAccessScopeSettings             *GroupAccessScopeSettings
	ResourceDefaultAccessScopeSettings   *ResourceDefaultAccessScopeSettings
}

// EntitlementBundleAccessScopeSettingsAsAccessScopeSettingsFullAccessScopeSettings is a convenience function that returns EntitlementBundleAccessScopeSettings wrapped in AccessScopeSettingsFullAccessScopeSettings
func EntitlementBundleAccessScopeSettingsAsAccessScopeSettingsFullAccessScopeSettings(v *EntitlementBundleAccessScopeSettings) AccessScopeSettingsFullAccessScopeSettings {
	return AccessScopeSettingsFullAccessScopeSettings{
		EntitlementBundleAccessScopeSettings: v,
	}
}

// GroupAccessScopeSettingsAsAccessScopeSettingsFullAccessScopeSettings is a convenience function that returns GroupAccessScopeSettings wrapped in AccessScopeSettingsFullAccessScopeSettings
func GroupAccessScopeSettingsAsAccessScopeSettingsFullAccessScopeSettings(v *GroupAccessScopeSettings) AccessScopeSettingsFullAccessScopeSettings {
	return AccessScopeSettingsFullAccessScopeSettings{
		GroupAccessScopeSettings: v,
	}
}

// ResourceDefaultAccessScopeSettingsAsAccessScopeSettingsFullAccessScopeSettings is a convenience function that returns ResourceDefaultAccessScopeSettings wrapped in AccessScopeSettingsFullAccessScopeSettings
func ResourceDefaultAccessScopeSettingsAsAccessScopeSettingsFullAccessScopeSettings(v *ResourceDefaultAccessScopeSettings) AccessScopeSettingsFullAccessScopeSettings {
	return AccessScopeSettingsFullAccessScopeSettings{
		ResourceDefaultAccessScopeSettings: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccessScopeSettingsFullAccessScopeSettings) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ENTITLEMENT_BUNDLES'
	if jsonDict["type"] == "ENTITLEMENT_BUNDLES" {
		// try to unmarshal JSON data into EntitlementBundleAccessScopeSettings
		err = json.Unmarshal(data, &dst.EntitlementBundleAccessScopeSettings)
		if err == nil {
			return nil // data stored in dst.EntitlementBundleAccessScopeSettings, return on the first match
		} else {
			dst.EntitlementBundleAccessScopeSettings = nil
			return fmt.Errorf("failed to unmarshal AccessScopeSettingsFullAccessScopeSettings as EntitlementBundleAccessScopeSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUPS'
	if jsonDict["type"] == "GROUPS" {
		// try to unmarshal JSON data into GroupAccessScopeSettings
		err = json.Unmarshal(data, &dst.GroupAccessScopeSettings)
		if err == nil {
			return nil // data stored in dst.GroupAccessScopeSettings, return on the first match
		} else {
			dst.GroupAccessScopeSettings = nil
			return fmt.Errorf("failed to unmarshal AccessScopeSettingsFullAccessScopeSettings as GroupAccessScopeSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESOURCE_DEFAULT'
	if jsonDict["type"] == "RESOURCE_DEFAULT" {
		// try to unmarshal JSON data into ResourceDefaultAccessScopeSettings
		err = json.Unmarshal(data, &dst.ResourceDefaultAccessScopeSettings)
		if err == nil {
			return nil // data stored in dst.ResourceDefaultAccessScopeSettings, return on the first match
		} else {
			dst.ResourceDefaultAccessScopeSettings = nil
			return fmt.Errorf("failed to unmarshal AccessScopeSettingsFullAccessScopeSettings as ResourceDefaultAccessScopeSettings: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccessScopeSettingsFullAccessScopeSettings) MarshalJSON() ([]byte, error) {
	if src.EntitlementBundleAccessScopeSettings != nil {
		return json.Marshal(&src.EntitlementBundleAccessScopeSettings)
	}

	if src.GroupAccessScopeSettings != nil {
		return json.Marshal(&src.GroupAccessScopeSettings)
	}

	if src.ResourceDefaultAccessScopeSettings != nil {
		return json.Marshal(&src.ResourceDefaultAccessScopeSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccessScopeSettingsFullAccessScopeSettings) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementBundleAccessScopeSettings != nil {
		return obj.EntitlementBundleAccessScopeSettings
	}

	if obj.GroupAccessScopeSettings != nil {
		return obj.GroupAccessScopeSettings
	}

	if obj.ResourceDefaultAccessScopeSettings != nil {
		return obj.ResourceDefaultAccessScopeSettings
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AccessScopeSettingsFullAccessScopeSettings) GetActualInstanceValue() interface{} {
	if obj.EntitlementBundleAccessScopeSettings != nil {
		return *obj.EntitlementBundleAccessScopeSettings
	}

	if obj.GroupAccessScopeSettings != nil {
		return *obj.GroupAccessScopeSettings
	}

	if obj.ResourceDefaultAccessScopeSettings != nil {
		return *obj.ResourceDefaultAccessScopeSettings
	}

	// all schemas are nil
	return nil
}

type NullableAccessScopeSettingsFullAccessScopeSettings struct {
	value *AccessScopeSettingsFullAccessScopeSettings
	isSet bool
}

func (v NullableAccessScopeSettingsFullAccessScopeSettings) Get() *AccessScopeSettingsFullAccessScopeSettings {
	return v.value
}

func (v *NullableAccessScopeSettingsFullAccessScopeSettings) Set(val *AccessScopeSettingsFullAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeSettingsFullAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeSettingsFullAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeSettingsFullAccessScopeSettings(val *AccessScopeSettingsFullAccessScopeSettings) *NullableAccessScopeSettingsFullAccessScopeSettings {
	return &NullableAccessScopeSettingsFullAccessScopeSettings{value: val, isSet: true}
}

func (v NullableAccessScopeSettingsFullAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeSettingsFullAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
