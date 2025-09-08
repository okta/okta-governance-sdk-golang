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

// AccessDurationType Identifies by whom the time settings were specified
type AccessDurationType string

// List of access-duration-type
const (
	ACCESSDURATIONTYPE_ADMIN_FIXED_DURATION         AccessDurationType = "ADMIN_FIXED_DURATION"
	ACCESSDURATIONTYPE_REQUESTER_SPECIFIED_DURATION AccessDurationType = "REQUESTER_SPECIFIED_DURATION"
)

// All allowed values of AccessDurationType enum
var AllowedAccessDurationTypeEnumValues = []AccessDurationType{
	"ADMIN_FIXED_DURATION",
	"REQUESTER_SPECIFIED_DURATION",
}

func (v *AccessDurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessDurationType(value)
	for _, existing := range AllowedAccessDurationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessDurationType", value)
}

// NewAccessDurationTypeFromValue returns a pointer to a valid AccessDurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessDurationTypeFromValue(v string) (*AccessDurationType, error) {
	ev := AccessDurationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessDurationType: valid values are %v", v, AllowedAccessDurationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessDurationType) IsValid() bool {
	for _, existing := range AllowedAccessDurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to access-duration-type value
func (v AccessDurationType) Ptr() *AccessDurationType {
	return &v
}

type NullableAccessDurationType struct {
	value *AccessDurationType
	isSet bool
}

func (v NullableAccessDurationType) Get() *AccessDurationType {
	return v.value
}

func (v *NullableAccessDurationType) Set(val *AccessDurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDurationType(val *AccessDurationType) *NullableAccessDurationType {
	return &NullableAccessDurationType{value: val, isSet: true}
}

func (v NullableAccessDurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
