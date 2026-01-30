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

// checks if the EntitlementBundleAccessScopeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementBundleAccessScopeSettings{}

// EntitlementBundleAccessScopeSettings Settings when the resource uses entitlement bundles for access scoping.
type EntitlementBundleAccessScopeSettings struct {
	Type string `json:"type"`
	// A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.
	EntitlementBundles []EntitlementBundlesArrayFullInner `json:"entitlementBundles"`
}

type _EntitlementBundleAccessScopeSettings EntitlementBundleAccessScopeSettings

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementBundleAccessScopeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["entitlementBundles"] = o.EntitlementBundles
	return toSerialize, nil
}

func (o *EntitlementBundleAccessScopeSettings) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementBundleAccessScopeSettings := _EntitlementBundleAccessScopeSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntitlementBundleAccessScopeSettings)

	if err != nil {
		return err
	}

	*o = EntitlementBundleAccessScopeSettings(varEntitlementBundleAccessScopeSettings)

	return err
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
