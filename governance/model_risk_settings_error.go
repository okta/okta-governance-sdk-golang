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
	"fmt"
)

// RiskSettingsError Error that has occurred with the risk settings. Currently, there is only error when an associated approval sequence was deleted.
type RiskSettingsError string

// List of risk-settings-error
const (
	RISKSETTINGSERROR_INVALID_SEQUENCE RiskSettingsError = "INVALID_SEQUENCE"
)

// All allowed values of RiskSettingsError enum
var AllowedRiskSettingsErrorEnumValues = []RiskSettingsError{
	"INVALID_SEQUENCE",
}

func (v *RiskSettingsError) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RiskSettingsError(value)
	for _, existing := range AllowedRiskSettingsErrorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RiskSettingsError", value)
}

// NewRiskSettingsErrorFromValue returns a pointer to a valid RiskSettingsError
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskSettingsErrorFromValue(v string) (*RiskSettingsError, error) {
	ev := RiskSettingsError(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RiskSettingsError: valid values are %v", v, AllowedRiskSettingsErrorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskSettingsError) IsValid() bool {
	for _, existing := range AllowedRiskSettingsErrorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to risk-settings-error value
func (v RiskSettingsError) Ptr() *RiskSettingsError {
	return &v
}

type NullableRiskSettingsError struct {
	value *RiskSettingsError
	isSet bool
}

func (v NullableRiskSettingsError) Get() *RiskSettingsError {
	return v.value
}

func (v *NullableRiskSettingsError) Set(val *RiskSettingsError) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSettingsError) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSettingsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSettingsError(val *RiskSettingsError) *NullableRiskSettingsError {
	return &NullableRiskSettingsError{value: val, isSet: true}
}

func (v NullableRiskSettingsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSettingsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
