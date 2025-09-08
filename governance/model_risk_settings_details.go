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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["defaultSetting"] = o.DefaultSetting
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSettingsDetails) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDetails := _RiskSettingsDetails{}

	err = json.Unmarshal(bytes, &varRiskSettingsDetails)
	if err == nil {
		*o = RiskSettingsDetails(varRiskSettingsDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "defaultSetting")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
