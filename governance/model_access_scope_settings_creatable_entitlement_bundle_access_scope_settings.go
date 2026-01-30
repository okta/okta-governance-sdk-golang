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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings{}

// AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings Access scope settings that resource permissions are entitlement bundles.
type AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings struct {
	// Request for access to entitlement bundles
	Type string `json:"type"`
	// Entitlement bundles that are made requestable
	EntitlementBundles []EntitlementBundlesArrayCreatableInner `json:"entitlementBundles"`
}

type _AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings

// NewAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings instantiates a new AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings(type_ string, entitlementBundles []EntitlementBundlesArrayCreatableInner) *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings {
	this := AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings{}
	this.Type = type_
	this.EntitlementBundles = entitlementBundles
	return &this
}

// NewAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettingsWithDefaults instantiates a new AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettingsWithDefaults() *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings {
	this := AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings{}
	return &this
}

// GetType returns the Type field value
func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) SetType(v string) {
	o.Type = v
}

// GetEntitlementBundles returns the EntitlementBundles field value
func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) GetEntitlementBundles() []EntitlementBundlesArrayCreatableInner {
	if o == nil {
		var ret []EntitlementBundlesArrayCreatableInner
		return ret
	}

	return o.EntitlementBundles
}

// GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field value
// and a boolean to check if the value has been set.
func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) GetEntitlementBundlesOk() ([]EntitlementBundlesArrayCreatableInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementBundles, true
}

// SetEntitlementBundles sets field value
func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) SetEntitlementBundles(v []EntitlementBundlesArrayCreatableInner) {
	o.EntitlementBundles = v
}

func (o AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["entitlementBundles"] = o.EntitlementBundles
	return toSerialize, nil
}

func (o *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"entitlementBundles",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings := _AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings)

	if err != nil {
		return err
	}

	*o = AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings(varAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings)

	return err
}

type NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings struct {
	value *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings
	isSet bool
}

func (v NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) Get() *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings {
	return v.value
}

func (v *NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) Set(val *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings(val *AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) *NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings {
	return &NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings{value: val, isSet: true}
}

func (v NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
