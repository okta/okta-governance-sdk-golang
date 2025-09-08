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

// AssignmentType the model 'AssignmentType'
type AssignmentType string

// List of assignment-type
const (
	ASSIGNMENTTYPE_INDIVIDUAL     AssignmentType = "INDIVIDUAL"
	ASSIGNMENTTYPE_GROUP          AssignmentType = "GROUP"
	ASSIGNMENTTYPE_GROUP_RULE     AssignmentType = "GROUP_RULE"
	ASSIGNMENTTYPE_APP_GROUP      AssignmentType = "APP_GROUP"
	ASSIGNMENTTYPE_ACCESS_REQUEST AssignmentType = "ACCESS_REQUEST"
)

// All allowed values of AssignmentType enum
var AllowedAssignmentTypeEnumValues = []AssignmentType{
	"INDIVIDUAL",
	"GROUP",
	"GROUP_RULE",
	"APP_GROUP",
	"ACCESS_REQUEST",
}

func (v *AssignmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssignmentType(value)
	for _, existing := range AllowedAssignmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssignmentType", value)
}

// NewAssignmentTypeFromValue returns a pointer to a valid AssignmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssignmentTypeFromValue(v string) (*AssignmentType, error) {
	ev := AssignmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssignmentType: valid values are %v", v, AllowedAssignmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssignmentType) IsValid() bool {
	for _, existing := range AllowedAssignmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assignment-type value
func (v AssignmentType) Ptr() *AssignmentType {
	return &v
}

type NullableAssignmentType struct {
	value *AssignmentType
	isSet bool
}

func (v NullableAssignmentType) Get() *AssignmentType {
	return v.value
}

func (v *NullableAssignmentType) Set(val *AssignmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentType(val *AssignmentType) *NullableAssignmentType {
	return &NullableAssignmentType{value: val, isSet: true}
}

func (v NullableAssignmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
