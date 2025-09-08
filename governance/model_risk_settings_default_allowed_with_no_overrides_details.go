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

// RiskSettingsDefaultAllowedWithNoOverridesDetails Risk settings where request submission is allowed with no overrides.
type RiskSettingsDefaultAllowedWithNoOverridesDetails struct {
	RequestSubmissionType string              `json:"requestSubmissionType"`
	Error                 []RiskSettingsError `json:"error,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _RiskSettingsDefaultAllowedWithNoOverridesDetails RiskSettingsDefaultAllowedWithNoOverridesDetails

// NewRiskSettingsDefaultAllowedWithNoOverridesDetails instantiates a new RiskSettingsDefaultAllowedWithNoOverridesDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskSettingsDefaultAllowedWithNoOverridesDetails(requestSubmissionType string) *RiskSettingsDefaultAllowedWithNoOverridesDetails {
	this := RiskSettingsDefaultAllowedWithNoOverridesDetails{}
	this.RequestSubmissionType = requestSubmissionType
	return &this
}

// NewRiskSettingsDefaultAllowedWithNoOverridesDetailsWithDefaults instantiates a new RiskSettingsDefaultAllowedWithNoOverridesDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskSettingsDefaultAllowedWithNoOverridesDetailsWithDefaults() *RiskSettingsDefaultAllowedWithNoOverridesDetails {
	this := RiskSettingsDefaultAllowedWithNoOverridesDetails{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) GetRequestSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) GetRequestSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) SetRequestSubmissionType(v string) {
	o.RequestSubmissionType = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) GetError() []RiskSettingsError {
	if o == nil || o.Error == nil {
		var ret []RiskSettingsError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) GetErrorOk() ([]RiskSettingsError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given []RiskSettingsError and assigns it to the Error field.
func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) SetError(v []RiskSettingsError) {
	o.Error = v
}

func (o RiskSettingsDefaultAllowedWithNoOverridesDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskSettingsDefaultAllowedWithNoOverridesDetails) UnmarshalJSON(bytes []byte) (err error) {
	varRiskSettingsDefaultAllowedWithNoOverridesDetails := _RiskSettingsDefaultAllowedWithNoOverridesDetails{}

	err = json.Unmarshal(bytes, &varRiskSettingsDefaultAllowedWithNoOverridesDetails)
	if err == nil {
		*o = RiskSettingsDefaultAllowedWithNoOverridesDetails(varRiskSettingsDefaultAllowedWithNoOverridesDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestSubmissionType")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskSettingsDefaultAllowedWithNoOverridesDetails struct {
	value *RiskSettingsDefaultAllowedWithNoOverridesDetails
	isSet bool
}

func (v NullableRiskSettingsDefaultAllowedWithNoOverridesDetails) Get() *RiskSettingsDefaultAllowedWithNoOverridesDetails {
	return v.value
}

func (v *NullableRiskSettingsDefaultAllowedWithNoOverridesDetails) Set(val *RiskSettingsDefaultAllowedWithNoOverridesDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsDefaultAllowedWithNoOverridesDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsDefaultAllowedWithNoOverridesDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsDefaultAllowedWithNoOverridesDetails(val *RiskSettingsDefaultAllowedWithNoOverridesDetails) *NullableRiskSettingsDefaultAllowedWithNoOverridesDetails {
	return &NullableRiskSettingsDefaultAllowedWithNoOverridesDetails{value: val, isSet: true}
}

func (v NullableRiskSettingsDefaultAllowedWithNoOverridesDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsDefaultAllowedWithNoOverridesDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
