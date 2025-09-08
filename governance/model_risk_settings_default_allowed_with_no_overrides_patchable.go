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

// RiskSettingsDefaultAllowedWithNoOverridesPatchable Risk settings where request submission is allowed with no overrides
type RiskSettingsDefaultAllowedWithNoOverridesPatchable struct {
	RequestSubmissionType string `json:"requestSubmissionType"`
	AdditionalProperties  map[string]interface{}
}

type _RiskSettingsDefaultAllowedWithNoOverridesPatchable RiskSettingsDefaultAllowedWithNoOverridesPatchable

// NewRiskSettingsDefaultAllowedWithNoOverridesPatchable instantiates a new RiskSettingsDefaultAllowedWithNoOverridesPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDefaultAllowedWithNoOverridesPatchable(requestSubmissionType string) *RiskSettingsDefaultAllowedWithNoOverridesPatchable {
	this := RiskSettingsDefaultAllowedWithNoOverridesPatchable{}
	this.RequestSubmissionType = requestSubmissionType
	return &this
}

// NewRiskSettingsDefaultAllowedWithNoOverridesPatchableWithDefaults instantiates a new RiskSettingsDefaultAllowedWithNoOverridesPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDefaultAllowedWithNoOverridesPatchableWithDefaults() *RiskSettingsDefaultAllowedWithNoOverridesPatchable {
	this := RiskSettingsDefaultAllowedWithNoOverridesPatchable{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskSettingsDefaultAllowedWithNoOverridesPatchable) GetRequestSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithNoOverridesPatchable) GetRequestSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskSettingsDefaultAllowedWithNoOverridesPatchable) SetRequestSubmissionType(v string) {
	o.RequestSubmissionType = v
}

func (o RiskSettingsDefaultAllowedWithNoOverridesPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSettingsDefaultAllowedWithNoOverridesPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDefaultAllowedWithNoOverridesPatchable := _RiskSettingsDefaultAllowedWithNoOverridesPatchable{}

	err = json.Unmarshal(bytes, &varRiskSettingsDefaultAllowedWithNoOverridesPatchable)
	if err == nil {
		*o = RiskSettingsDefaultAllowedWithNoOverridesPatchable(varRiskSettingsDefaultAllowedWithNoOverridesPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestSubmissionType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable struct {
	value *RiskSettingsDefaultAllowedWithNoOverridesPatchable
	isSet bool
}

func (v NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable) Get() *RiskSettingsDefaultAllowedWithNoOverridesPatchable {
	return v.value
}

func (v *NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable) Set(val *RiskSettingsDefaultAllowedWithNoOverridesPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultAllowedWithNoOverridesPatchable(val *RiskSettingsDefaultAllowedWithNoOverridesPatchable) *NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable {
	return &NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultAllowedWithNoOverridesPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
