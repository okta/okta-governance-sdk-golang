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

// RiskSettingsDefaultAllowedWithOverridesDetails Risk settings where request submission is allowed with specified approval sequence and optional access duration settings
type RiskSettingsDefaultAllowedWithOverridesDetails struct {
	RequestSubmissionType string `json:"requestSubmissionType"`
	// The ID of the approval sequence
	ApprovalSequenceId     string                      `json:"approvalSequenceId"`
	AccessDurationSettings *AccessDurationSettingsFull `json:"accessDurationSettings,omitempty"`
	Error                  []RiskSettingsError         `json:"error,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _RiskSettingsDefaultAllowedWithOverridesDetails RiskSettingsDefaultAllowedWithOverridesDetails

// NewRiskSettingsDefaultAllowedWithOverridesDetails instantiates a new RiskSettingsDefaultAllowedWithOverridesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDefaultAllowedWithOverridesDetails(requestSubmissionType string, approvalSequenceId string) *RiskSettingsDefaultAllowedWithOverridesDetails {
	this := RiskSettingsDefaultAllowedWithOverridesDetails{}
	this.RequestSubmissionType = requestSubmissionType
	this.ApprovalSequenceId = approvalSequenceId
	return &this
}

// NewRiskSettingsDefaultAllowedWithOverridesDetailsWithDefaults instantiates a new RiskSettingsDefaultAllowedWithOverridesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDefaultAllowedWithOverridesDetailsWithDefaults() *RiskSettingsDefaultAllowedWithOverridesDetails {
	this := RiskSettingsDefaultAllowedWithOverridesDetails{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetRequestSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetRequestSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetRequestSubmissionType(v string) {
	o.RequestSubmissionType = v
}

// GetApprovalSequenceId returns the ApprovalSequenceId field value
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetApprovalSequenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApprovalSequenceId
}

// GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetApprovalSequenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalSequenceId, true
}

// SetApprovalSequenceId sets field value
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetApprovalSequenceId(v string) {
	o.ApprovalSequenceId = v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetAccessDurationSettings() AccessDurationSettingsFull {
	if o == nil || o.AccessDurationSettings == nil {
		var ret AccessDurationSettingsFull
		return ret
	}
	return *o.AccessDurationSettings
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool) {
	if o == nil || o.AccessDurationSettings == nil {
		return nil, false
	}
	return o.AccessDurationSettings, true
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) HasAccessDurationSettings() bool {
	if o != nil && o.AccessDurationSettings != nil {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given AccessDurationSettingsFull and assigns it to the AccessDurationSettings field.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetAccessDurationSettings(v AccessDurationSettingsFull) {
	o.AccessDurationSettings = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetError() []RiskSettingsError {
	if o == nil || o.Error == nil {
		var ret []RiskSettingsError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetErrorOk() ([]RiskSettingsError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given []RiskSettingsError and assigns it to the Error field.
func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetError(v []RiskSettingsError) {
	o.Error = v
}

func (o RiskSettingsDefaultAllowedWithOverridesDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	}
	if true {
		toSerialize["approvalSequenceId"] = o.ApprovalSequenceId
	}
	if o.AccessDurationSettings != nil {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSettingsDefaultAllowedWithOverridesDetails) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDefaultAllowedWithOverridesDetails := _RiskSettingsDefaultAllowedWithOverridesDetails{}

	err = json.Unmarshal(bytes, &varRiskSettingsDefaultAllowedWithOverridesDetails)
	if err == nil {
		*o = RiskSettingsDefaultAllowedWithOverridesDetails(varRiskSettingsDefaultAllowedWithOverridesDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestSubmissionType")
		delete(additionalProperties, "approvalSequenceId")
		delete(additionalProperties, "accessDurationSettings")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskSettingsDefaultAllowedWithOverridesDetails struct {
	value *RiskSettingsDefaultAllowedWithOverridesDetails
	isSet bool
}

func (v NullableRiskSettingsDefaultAllowedWithOverridesDetails) Get() *RiskSettingsDefaultAllowedWithOverridesDetails {
	return v.value
}

func (v *NullableRiskSettingsDefaultAllowedWithOverridesDetails) Set(val *RiskSettingsDefaultAllowedWithOverridesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultAllowedWithOverridesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultAllowedWithOverridesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultAllowedWithOverridesDetails(val *RiskSettingsDefaultAllowedWithOverridesDetails) *NullableRiskSettingsDefaultAllowedWithOverridesDetails {
	return &NullableRiskSettingsDefaultAllowedWithOverridesDetails{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultAllowedWithOverridesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultAllowedWithOverridesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
