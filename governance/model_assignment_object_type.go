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

// AssignmentObjectType Identifies assignment type. assignmentObject will be present when assignmentObjectType is Bundle or Collection.
type AssignmentObjectType string

// List of assignment-object-type
const (
	ASSIGNMENTOBJECTTYPE_ENTITLEMENT_BUNDLE AssignmentObjectType = "ENTITLEMENT_BUNDLE"
	ASSIGNMENTOBJECTTYPE_COLLECTION         AssignmentObjectType = "COLLECTION"
	ASSIGNMENTOBJECTTYPE_ENTITLEMENT        AssignmentObjectType = "ENTITLEMENT"
)

// All allowed values of AssignmentObjectType enum
var AllowedAssignmentObjectTypeEnumValues = []AssignmentObjectType{
	"ENTITLEMENT_BUNDLE",
	"COLLECTION",
	"ENTITLEMENT",
}

func (v *AssignmentObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssignmentObjectType(value)
	for _, existing := range AllowedAssignmentObjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssignmentObjectType", value)
}

// NewAssignmentObjectTypeFromValue returns a pointer to a valid AssignmentObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssignmentObjectTypeFromValue(v string) (*AssignmentObjectType, error) {
	ev := AssignmentObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssignmentObjectType: valid values are %v", v, AllowedAssignmentObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssignmentObjectType) IsValid() bool {
	for _, existing := range AllowedAssignmentObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assignment-object-type value
func (v AssignmentObjectType) Ptr() *AssignmentObjectType {
	return &v
}

type NullableAssignmentObjectType struct {
	value *AssignmentObjectType
	isSet bool
}

func (v NullableAssignmentObjectType) Get() *AssignmentObjectType {
	return v.value
}

func (v *NullableAssignmentObjectType) Set(val *AssignmentObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentObjectType(val *AssignmentObjectType) *NullableAssignmentObjectType {
	return &NullableAssignmentObjectType{value: val, isSet: true}
}

func (v NullableAssignmentObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
