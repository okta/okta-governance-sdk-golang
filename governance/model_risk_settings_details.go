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

// checks if the RiskSettingsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskSettingsDetails{}

// RiskSettingsDetails Risk settings that are valid for an access request when a risk has been detected for the resource and requesting user
type RiskSettingsDetails struct {
	DefaultSetting       RiskSettingsDefaultDetails `json:"defaultSetting"`
	AdditionalProperties map[string]interface{}
}

type _RiskSettingsDetails RiskSettingsDetails

// NewRiskSettingsDetails instantiates a new RiskSettingsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDetails(defaultSetting RiskSettingsDefaultDetails) *RiskSettingsDetails {
	this := RiskSettingsDetails{}
	this.DefaultSetting = defaultSetting
	return &this
}

// NewRiskSettingsDetailsWithDefaults instantiates a new RiskSettingsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDetailsWithDefaults() *RiskSettingsDetails {
	this := RiskSettingsDetails{}
	return &this
}

// GetDefaultSetting returns the DefaultSetting field value
func (o *RiskSettingsDetails) GetDefaultSetting() RiskSettingsDefaultDetails {
	if o == nil {
		var ret RiskSettingsDefaultDetails
		return ret
	}

	return o.DefaultSetting
}

// GetDefaultSettingOk returns a tuple with the DefaultSetting field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDetails) GetDefaultSettingOk() (*RiskSettingsDefaultDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSetting, true
}

// SetDefaultSetting sets field value
func (o *RiskSettingsDetails) SetDefaultSetting(v RiskSettingsDefaultDetails) {
	o.DefaultSetting = v
}

func (o RiskSettingsDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskSettingsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultSetting"] = o.DefaultSetting

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskSettingsDetails) UnmarshalJSON(data []byte) (err error) {
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

	varRiskSettingsDetails := _RiskSettingsDetails{}

	err = json.Unmarshal(data, &varRiskSettingsDetails)

	if err != nil {
		return err
	}

	*o = RiskSettingsDetails(varRiskSettingsDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultSetting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskSettingsDetails struct {
	value *RiskSettingsDetails
	isSet bool
}

func (v NullableRiskSettingsDetails) Get() *RiskSettingsDetails {
	return v.value
}

func (v *NullableRiskSettingsDetails) Set(val *RiskSettingsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDetails(val *RiskSettingsDetails) *NullableRiskSettingsDetails {
	return &NullableRiskSettingsDetails{value: val, isSet: true}
}

func (v NullableRiskSettingsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
