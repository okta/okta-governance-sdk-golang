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

// ScheduleType The type of of campaign.
type ScheduleType string

// List of schedule-type
const (
	SCHEDULETYPE_ONE_OFF   ScheduleType = "ONE_OFF"
	SCHEDULETYPE_RECURRING ScheduleType = "RECURRING"
)

// All allowed values of ScheduleType enum
var AllowedScheduleTypeEnumValues = []ScheduleType{
	"ONE_OFF",
	"RECURRING",
}

func (v *ScheduleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduleType(value)
	for _, existing := range AllowedScheduleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleType", value)
}

// NewScheduleTypeFromValue returns a pointer to a valid ScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduleTypeFromValue(v string) (*ScheduleType, error) {
	ev := ScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduleType: valid values are %v", v, AllowedScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduleType) IsValid() bool {
	for _, existing := range AllowedScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to schedule-type value
func (v ScheduleType) Ptr() *ScheduleType {
	return &v
}

type NullableScheduleType struct {
	value *ScheduleType
	isSet bool
}

func (v NullableScheduleType) Get() *ScheduleType {
	return v.value
}

func (v *NullableScheduleType) Set(val *ScheduleType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleType(val *ScheduleType) *NullableScheduleType {
	return &NullableScheduleType{value: val, isSet: true}
}

func (v NullableScheduleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
