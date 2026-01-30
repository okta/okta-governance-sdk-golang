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

// checks if the RiskSettingsPatchable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskSettingsPatchable{}

// RiskSettingsPatchable Risk settings that are valid for an access request when a risk has been detected for the resource and requesting user
type RiskSettingsPatchable struct {
	DefaultSetting       RiskSettingsDefaultPatchable `json:"defaultSetting"`
	AdditionalProperties map[string]interface{}
}

type _RiskSettingsPatchable RiskSettingsPatchable

// NewRiskSettingsPatchable instantiates a new RiskSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsPatchable(defaultSetting RiskSettingsDefaultPatchable) *RiskSettingsPatchable {
	this := RiskSettingsPatchable{}
	this.DefaultSetting = defaultSetting
	return &this
}

// NewRiskSettingsPatchableWithDefaults instantiates a new RiskSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsPatchableWithDefaults() *RiskSettingsPatchable {
	this := RiskSettingsPatchable{}
	return &this
}

// GetDefaultSetting returns the DefaultSetting field value
func (o *RiskSettingsPatchable) GetDefaultSetting() RiskSettingsDefaultPatchable {
	if o == nil {
		var ret RiskSettingsDefaultPatchable
		return ret
	}

	return o.DefaultSetting
}

// GetDefaultSettingOk returns a tuple with the DefaultSetting field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsPatchable) GetDefaultSettingOk() (*RiskSettingsDefaultPatchable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSetting, true
}

// SetDefaultSetting sets field value
func (o *RiskSettingsPatchable) SetDefaultSetting(v RiskSettingsDefaultPatchable) {
	o.DefaultSetting = v
}

func (o RiskSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskSettingsPatchable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultSetting"] = o.DefaultSetting

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskSettingsPatchable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultSetting",
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

	varRiskSettingsPatchable := _RiskSettingsPatchable{}

	err = json.Unmarshal(data, &varRiskSettingsPatchable)

	if err != nil {
		return err
	}

	*o = RiskSettingsPatchable(varRiskSettingsPatchable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultSetting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskSettingsPatchable struct {
	value *RiskSettingsPatchable
	isSet bool
}

func (v NullableRiskSettingsPatchable) Get() *RiskSettingsPatchable {
	return v.value
}

func (v *NullableRiskSettingsPatchable) Set(val *RiskSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsPatchable(val *RiskSettingsPatchable) *NullableRiskSettingsPatchable {
	return &NullableRiskSettingsPatchable{value: val, isSet: true}
}

func (v NullableRiskSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
