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

// RiskSettingsDefaultAllowedWithOverridesPatchable Risk settings where request submission is allowed with specified approval sequence and optional access duration settings
type RiskSettingsDefaultAllowedWithOverridesPatchable struct {
	RequestSubmissionType string `json:"requestSubmissionType"`
	// The ID of the approval sequence
	ApprovalSequenceId     string                                  `json:"approvalSequenceId"`
	AccessDurationSettings NullableAccessDurationSettingsPatchable `json:"accessDurationSettings,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _RiskSettingsDefaultAllowedWithOverridesPatchable RiskSettingsDefaultAllowedWithOverridesPatchable

// NewRiskSettingsDefaultAllowedWithOverridesPatchable instantiates a new RiskSettingsDefaultAllowedWithOverridesPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDefaultAllowedWithOverridesPatchable(requestSubmissionType string, approvalSequenceId string) *RiskSettingsDefaultAllowedWithOverridesPatchable {
	this := RiskSettingsDefaultAllowedWithOverridesPatchable{}
	this.RequestSubmissionType = requestSubmissionType
	this.ApprovalSequenceId = approvalSequenceId
	return &this
}

// NewRiskSettingsDefaultAllowedWithOverridesPatchableWithDefaults instantiates a new RiskSettingsDefaultAllowedWithOverridesPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDefaultAllowedWithOverridesPatchableWithDefaults() *RiskSettingsDefaultAllowedWithOverridesPatchable {
	this := RiskSettingsDefaultAllowedWithOverridesPatchable{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetRequestSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetRequestSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetRequestSubmissionType(v string) {
	o.RequestSubmissionType = v
}

// GetApprovalSequenceId returns the ApprovalSequenceId field value
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetApprovalSequenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApprovalSequenceId
}

// GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetApprovalSequenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalSequenceId, true
}

// SetApprovalSequenceId sets field value
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetApprovalSequenceId(v string) {
	o.ApprovalSequenceId = v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetAccessDurationSettings() AccessDurationSettingsPatchable {
	if o == nil || o.AccessDurationSettings.Get() == nil {
		var ret AccessDurationSettingsPatchable
		return ret
	}
	return *o.AccessDurationSettings.Get()
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetAccessDurationSettingsOk() (*AccessDurationSettingsPatchable, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessDurationSettings.Get(), o.AccessDurationSettings.IsSet()
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) HasAccessDurationSettings() bool {
	if o != nil && o.AccessDurationSettings.IsSet() {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given NullableAccessDurationSettingsPatchable and assigns it to the AccessDurationSettings field.
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetAccessDurationSettings(v AccessDurationSettingsPatchable) {
	o.AccessDurationSettings.Set(&v)
}

// SetAccessDurationSettingsNil sets the value for AccessDurationSettings to be an explicit nil
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetAccessDurationSettingsNil() {
	o.AccessDurationSettings.Set(nil)
}

// UnsetAccessDurationSettings ensures that no value is present for AccessDurationSettings, not even an explicit nil
func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) UnsetAccessDurationSettings() {
	o.AccessDurationSettings.Unset()
}

func (o RiskSettingsDefaultAllowedWithOverridesPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	}
	if true {
		toSerialize["approvalSequenceId"] = o.ApprovalSequenceId
	}
	if o.AccessDurationSettings.IsSet() {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDefaultAllowedWithOverridesPatchable := _RiskSettingsDefaultAllowedWithOverridesPatchable{}

	err = json.Unmarshal(bytes, &varRiskSettingsDefaultAllowedWithOverridesPatchable)
	if err == nil {
		*o = RiskSettingsDefaultAllowedWithOverridesPatchable(varRiskSettingsDefaultAllowedWithOverridesPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestSubmissionType")
		delete(additionalProperties, "approvalSequenceId")
		delete(additionalProperties, "accessDurationSettings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskSettingsDefaultAllowedWithOverridesPatchable struct {
	value *RiskSettingsDefaultAllowedWithOverridesPatchable
	isSet bool
}

func (v NullableRiskSettingsDefaultAllowedWithOverridesPatchable) Get() *RiskSettingsDefaultAllowedWithOverridesPatchable {
	return v.value
}

func (v *NullableRiskSettingsDefaultAllowedWithOverridesPatchable) Set(val *RiskSettingsDefaultAllowedWithOverridesPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultAllowedWithOverridesPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultAllowedWithOverridesPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultAllowedWithOverridesPatchable(val *RiskSettingsDefaultAllowedWithOverridesPatchable) *NullableRiskSettingsDefaultAllowedWithOverridesPatchable {
	return &NullableRiskSettingsDefaultAllowedWithOverridesPatchable{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultAllowedWithOverridesPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultAllowedWithOverridesPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
