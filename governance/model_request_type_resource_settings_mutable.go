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

// RequestTypeResourceSettingsMutable - Which resource(s) are requestable
type RequestTypeResourceSettingsMutable struct {
	RequestTypeResourceSettingsApps               *RequestTypeResourceSettingsApps
	RequestTypeResourceSettingsEntitlementBundles *RequestTypeResourceSettingsEntitlementBundles
	RequestTypeResourceSettingsGroups             *RequestTypeResourceSettingsGroups
}

// RequestTypeResourceSettingsAppsAsRequestTypeResourceSettingsMutable is a convenience function that returns RequestTypeResourceSettingsApps wrapped in RequestTypeResourceSettingsMutable
func RequestTypeResourceSettingsAppsAsRequestTypeResourceSettingsMutable(v *RequestTypeResourceSettingsApps) RequestTypeResourceSettingsMutable {
	return RequestTypeResourceSettingsMutable{
		RequestTypeResourceSettingsApps: v,
	}
}

// RequestTypeResourceSettingsEntitlementBundlesAsRequestTypeResourceSettingsMutable is a convenience function that returns RequestTypeResourceSettingsEntitlementBundles wrapped in RequestTypeResourceSettingsMutable
func RequestTypeResourceSettingsEntitlementBundlesAsRequestTypeResourceSettingsMutable(v *RequestTypeResourceSettingsEntitlementBundles) RequestTypeResourceSettingsMutable {
	return RequestTypeResourceSettingsMutable{
		RequestTypeResourceSettingsEntitlementBundles: v,
	}
}

// RequestTypeResourceSettingsGroupsAsRequestTypeResourceSettingsMutable is a convenience function that returns RequestTypeResourceSettingsGroups wrapped in RequestTypeResourceSettingsMutable
func RequestTypeResourceSettingsGroupsAsRequestTypeResourceSettingsMutable(v *RequestTypeResourceSettingsGroups) RequestTypeResourceSettingsMutable {
	return RequestTypeResourceSettingsMutable{
		RequestTypeResourceSettingsGroups: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestTypeResourceSettingsMutable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'APPS'
	if jsonDict["type"] == "APPS" {
		// try to unmarshal JSON data into RequestTypeResourceSettingsApps
		err = json.Unmarshal(data, &dst.RequestTypeResourceSettingsApps)
		if err == nil {
			return nil // data stored in dst.RequestTypeResourceSettingsApps, return on the first match
		} else {
			dst.RequestTypeResourceSettingsApps = nil
			return fmt.Errorf("failed to unmarshal RequestTypeResourceSettingsMutable as RequestTypeResourceSettingsApps: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITLEMENT_BUNDLES'
	if jsonDict["type"] == "ENTITLEMENT_BUNDLES" {
		// try to unmarshal JSON data into RequestTypeResourceSettingsEntitlementBundles
		err = json.Unmarshal(data, &dst.RequestTypeResourceSettingsEntitlementBundles)
		if err == nil {
			return nil // data stored in dst.RequestTypeResourceSettingsEntitlementBundles, return on the first match
		} else {
			dst.RequestTypeResourceSettingsEntitlementBundles = nil
			return fmt.Errorf("failed to unmarshal RequestTypeResourceSettingsMutable as RequestTypeResourceSettingsEntitlementBundles: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUPS'
	if jsonDict["type"] == "GROUPS" {
		// try to unmarshal JSON data into RequestTypeResourceSettingsGroups
		err = json.Unmarshal(data, &dst.RequestTypeResourceSettingsGroups)
		if err == nil {
			return nil // data stored in dst.RequestTypeResourceSettingsGroups, return on the first match
		} else {
			dst.RequestTypeResourceSettingsGroups = nil
			return fmt.Errorf("failed to unmarshal RequestTypeResourceSettingsMutable as RequestTypeResourceSettingsGroups: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeResourceSettingsMutable) MarshalJSON() ([]byte, error) {
	if src.RequestTypeResourceSettingsApps != nil {
		return json.Marshal(&src.RequestTypeResourceSettingsApps)
	}

	if src.RequestTypeResourceSettingsEntitlementBundles != nil {
		return json.Marshal(&src.RequestTypeResourceSettingsEntitlementBundles)
	}

	if src.RequestTypeResourceSettingsGroups != nil {
		return json.Marshal(&src.RequestTypeResourceSettingsGroups)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeResourceSettingsMutable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeResourceSettingsApps != nil {
		return obj.RequestTypeResourceSettingsApps
	}

	if obj.RequestTypeResourceSettingsEntitlementBundles != nil {
		return obj.RequestTypeResourceSettingsEntitlementBundles
	}

	if obj.RequestTypeResourceSettingsGroups != nil {
		return obj.RequestTypeResourceSettingsGroups
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequestTypeResourceSettingsMutable) GetActualInstanceValue() interface{} {
	if obj.RequestTypeResourceSettingsApps != nil {
		return *obj.RequestTypeResourceSettingsApps
	}

	if obj.RequestTypeResourceSettingsEntitlementBundles != nil {
		return *obj.RequestTypeResourceSettingsEntitlementBundles
	}

	if obj.RequestTypeResourceSettingsGroups != nil {
		return *obj.RequestTypeResourceSettingsGroups
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeResourceSettingsMutable struct {
	value *RequestTypeResourceSettingsMutable
	isSet bool
}

func (v NullableRequestTypeResourceSettingsMutable) Get() *RequestTypeResourceSettingsMutable {
	return v.value
}

func (v *NullableRequestTypeResourceSettingsMutable) Set(val *RequestTypeResourceSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeResourceSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeResourceSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeResourceSettingsMutable(val *RequestTypeResourceSettingsMutable) *NullableRequestTypeResourceSettingsMutable {
	return &NullableRequestTypeResourceSettingsMutable{value: val, isSet: true}
}

func (v NullableRequestTypeResourceSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeResourceSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
