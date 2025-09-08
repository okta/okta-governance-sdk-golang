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
)

// EntitlementBundleAccessScopeSettings Settings when the resource uses entitlement bundles for access scoping.
type EntitlementBundleAccessScopeSettings struct {
	Type string `json:"type"`
	// A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.
	EntitlementBundles []EntitlementBundlesArrayFullInner `json:"entitlementBundles"`
}

// NewEntitlementBundleAccessScopeSettings instantiates a new EntitlementBundleAccessScopeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBundleAccessScopeSettings(type_ string, entitlementBundles []EntitlementBundlesArrayFullInner) *EntitlementBundleAccessScopeSettings {
	this := EntitlementBundleAccessScopeSettings{}
	this.Type = type_
	this.EntitlementBundles = entitlementBundles
	return &this
}

// NewEntitlementBundleAccessScopeSettingsWithDefaults instantiates a new EntitlementBundleAccessScopeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBundleAccessScopeSettingsWithDefaults() *EntitlementBundleAccessScopeSettings {
	this := EntitlementBundleAccessScopeSettings{}
	return &this
}

// GetType returns the Type field value
func (o *EntitlementBundleAccessScopeSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleAccessScopeSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntitlementBundleAccessScopeSettings) SetType(v string) {
	o.Type = v
}

// GetEntitlementBundles returns the EntitlementBundles field value
func (o *EntitlementBundleAccessScopeSettings) GetEntitlementBundles() []EntitlementBundlesArrayFullInner {
	if o == nil {
		var ret []EntitlementBundlesArrayFullInner
		return ret
	}

	return o.EntitlementBundles
}

// GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleAccessScopeSettings) GetEntitlementBundlesOk() ([]EntitlementBundlesArrayFullInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementBundles, true
}

// SetEntitlementBundles sets field value
func (o *EntitlementBundleAccessScopeSettings) SetEntitlementBundles(v []EntitlementBundlesArrayFullInner) {
	o.EntitlementBundles = v
}

func (o EntitlementBundleAccessScopeSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["entitlementBundles"] = o.EntitlementBundles
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementBundleAccessScopeSettings struct {
	value *EntitlementBundleAccessScopeSettings
	isSet bool
}

func (v NullableEntitlementBundleAccessScopeSettings) Get() *EntitlementBundleAccessScopeSettings {
	return v.value
}

func (v *NullableEntitlementBundleAccessScopeSettings) Set(val *EntitlementBundleAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundleAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundleAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundleAccessScopeSettings(val *EntitlementBundleAccessScopeSettings) *NullableEntitlementBundleAccessScopeSettings {
	return &NullableEntitlementBundleAccessScopeSettings{value: val, isSet: true}
}

func (v NullableEntitlementBundleAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundleAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
