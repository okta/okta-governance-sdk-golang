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

// RiskSettingsDefaultRestrictedDetails Risk settings where request submission is restricted. Access duration settings can only be present when there is an error.
type RiskSettingsDefaultRestrictedDetails struct {
	RequestSubmissionType  string                      `json:"requestSubmissionType"`
	AccessDurationSettings *AccessDurationSettingsFull `json:"accessDurationSettings,omitempty"`
	Error                  []RiskSettingsError         `json:"error,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _RiskSettingsDefaultRestrictedDetails RiskSettingsDefaultRestrictedDetails

// NewRiskSettingsDefaultRestrictedDetails instantiates a new RiskSettingsDefaultRestrictedDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDefaultRestrictedDetails(requestSubmissionType string) *RiskSettingsDefaultRestrictedDetails {
	this := RiskSettingsDefaultRestrictedDetails{}
	this.RequestSubmissionType = requestSubmissionType
	return &this
}

// NewRiskSettingsDefaultRestrictedDetailsWithDefaults instantiates a new RiskSettingsDefaultRestrictedDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDefaultRestrictedDetailsWithDefaults() *RiskSettingsDefaultRestrictedDetails {
	this := RiskSettingsDefaultRestrictedDetails{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskSettingsDefaultRestrictedDetails) GetRequestSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultRestrictedDetails) GetRequestSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskSettingsDefaultRestrictedDetails) SetRequestSubmissionType(v string) {
	o.RequestSubmissionType = v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise.
func (o *RiskSettingsDefaultRestrictedDetails) GetAccessDurationSettings() AccessDurationSettingsFull {
	if o == nil || o.AccessDurationSettings == nil {
		var ret AccessDurationSettingsFull
		return ret
	}
	return *o.AccessDurationSettings
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultRestrictedDetails) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool) {
	if o == nil || o.AccessDurationSettings == nil {
		return nil, false
	}
	return o.AccessDurationSettings, true
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RiskSettingsDefaultRestrictedDetails) HasAccessDurationSettings() bool {
	if o != nil && o.AccessDurationSettings != nil {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given AccessDurationSettingsFull and assigns it to the AccessDurationSettings field.
func (o *RiskSettingsDefaultRestrictedDetails) SetAccessDurationSettings(v AccessDurationSettingsFull) {
	o.AccessDurationSettings = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RiskSettingsDefaultRestrictedDetails) GetError() []RiskSettingsError {
	if o == nil || o.Error == nil {
		var ret []RiskSettingsError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultRestrictedDetails) GetErrorOk() ([]RiskSettingsError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RiskSettingsDefaultRestrictedDetails) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given []RiskSettingsError and assigns it to the Error field.
func (o *RiskSettingsDefaultRestrictedDetails) SetError(v []RiskSettingsError) {
	o.Error = v
}

func (o RiskSettingsDefaultRestrictedDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
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

func (o *RiskSettingsDefaultRestrictedDetails) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDefaultRestrictedDetails := _RiskSettingsDefaultRestrictedDetails{}

	err = json.Unmarshal(bytes, &varRiskSettingsDefaultRestrictedDetails)
	if err == nil {
		*o = RiskSettingsDefaultRestrictedDetails(varRiskSettingsDefaultRestrictedDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestSubmissionType")
		delete(additionalProperties, "accessDurationSettings")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskSettingsDefaultRestrictedDetails struct {
	value *RiskSettingsDefaultRestrictedDetails
	isSet bool
}

func (v NullableRiskSettingsDefaultRestrictedDetails) Get() *RiskSettingsDefaultRestrictedDetails {
	return v.value
}

func (v *NullableRiskSettingsDefaultRestrictedDetails) Set(val *RiskSettingsDefaultRestrictedDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultRestrictedDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultRestrictedDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultRestrictedDetails(val *RiskSettingsDefaultRestrictedDetails) *NullableRiskSettingsDefaultRestrictedDetails {
	return &NullableRiskSettingsDefaultRestrictedDetails{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultRestrictedDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultRestrictedDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
