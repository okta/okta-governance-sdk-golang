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

// checks if the RiskSettingsDefaultRestrictedDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskSettingsDefaultRestrictedDetails{}

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
	if o == nil || IsNil(o.AccessDurationSettings) {
		var ret AccessDurationSettingsFull
		return ret
	}
	return *o.AccessDurationSettings
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultRestrictedDetails) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool) {
	if o == nil || IsNil(o.AccessDurationSettings) {
		return nil, false
	}
	return o.AccessDurationSettings, true
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RiskSettingsDefaultRestrictedDetails) HasAccessDurationSettings() bool {
	if o != nil && !IsNil(o.AccessDurationSettings) {
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
	if o == nil || IsNil(o.Error) {
		var ret []RiskSettingsError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultRestrictedDetails) GetErrorOk() ([]RiskSettingsError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RiskSettingsDefaultRestrictedDetails) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []RiskSettingsError and assigns it to the Error field.
func (o *RiskSettingsDefaultRestrictedDetails) SetError(v []RiskSettingsError) {
	o.Error = v
}

func (o RiskSettingsDefaultRestrictedDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskSettingsDefaultRestrictedDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	if !IsNil(o.AccessDurationSettings) {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskSettingsDefaultRestrictedDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestSubmissionType",
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

	varRiskSettingsDefaultRestrictedDetails := _RiskSettingsDefaultRestrictedDetails{}

	err = json.Unmarshal(data, &varRiskSettingsDefaultRestrictedDetails)

	if err != nil {
		return err
	}

	*o = RiskSettingsDefaultRestrictedDetails(varRiskSettingsDefaultRestrictedDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestSubmissionType")
		delete(additionalProperties, "accessDurationSettings")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
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
