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

// RiskSettingsDefaultRestrictedPatchable Risk settings where request submission is restricted
type RiskSettingsDefaultRestrictedPatchable struct {
	RequestSubmissionType string `json:"requestSubmissionType"`
	AdditionalProperties  map[string]interface{}
}

type _RiskSettingsDefaultRestrictedPatchable RiskSettingsDefaultRestrictedPatchable

// NewRiskSettingsDefaultRestrictedPatchable instantiates a new RiskSettingsDefaultRestrictedPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDefaultRestrictedPatchable(requestSubmissionType string) *RiskSettingsDefaultRestrictedPatchable {
	this := RiskSettingsDefaultRestrictedPatchable{}
	this.RequestSubmissionType = requestSubmissionType
	return &this
}

// NewRiskSettingsDefaultRestrictedPatchableWithDefaults instantiates a new RiskSettingsDefaultRestrictedPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDefaultRestrictedPatchableWithDefaults() *RiskSettingsDefaultRestrictedPatchable {
	this := RiskSettingsDefaultRestrictedPatchable{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskSettingsDefaultRestrictedPatchable) GetRequestSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultRestrictedPatchable) GetRequestSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskSettingsDefaultRestrictedPatchable) SetRequestSubmissionType(v string) {
	o.RequestSubmissionType = v
}

func (o RiskSettingsDefaultRestrictedPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSettingsDefaultRestrictedPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDefaultRestrictedPatchable := _RiskSettingsDefaultRestrictedPatchable{}

	err = json.Unmarshal(bytes, &varRiskSettingsDefaultRestrictedPatchable)
	if err == nil {
		*o = RiskSettingsDefaultRestrictedPatchable(varRiskSettingsDefaultRestrictedPatchable)
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

type NullableRiskSettingsDefaultRestrictedPatchable struct {
	value *RiskSettingsDefaultRestrictedPatchable
	isSet bool
}

func (v NullableRiskSettingsDefaultRestrictedPatchable) Get() *RiskSettingsDefaultRestrictedPatchable {
	return v.value
}

func (v *NullableRiskSettingsDefaultRestrictedPatchable) Set(val *RiskSettingsDefaultRestrictedPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultRestrictedPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultRestrictedPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultRestrictedPatchable(val *RiskSettingsDefaultRestrictedPatchable) *NullableRiskSettingsDefaultRestrictedPatchable {
	return &NullableRiskSettingsDefaultRestrictedPatchable{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultRestrictedPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultRestrictedPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
