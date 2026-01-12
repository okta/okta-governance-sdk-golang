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

// AssignmentMethod Assignment method
type AssignmentMethod string

// List of assignment-method
const (
	ASSIGNMENTMETHOD_POLICY         AssignmentMethod = "POLICY"
	ASSIGNMENTMETHOD_ACCESS_REQUEST AssignmentMethod = "ACCESS_REQUEST"
	ASSIGNMENTMETHOD_ADMIN          AssignmentMethod = "ADMIN"
	ASSIGNMENTMETHOD_API            AssignmentMethod = "API"
)

// All allowed values of AssignmentMethod enum
var AllowedAssignmentMethodEnumValues = []AssignmentMethod{
	"POLICY",
	"ACCESS_REQUEST",
	"ADMIN",
	"API",
}

func (v *AssignmentMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssignmentMethod(value)
	for _, existing := range AllowedAssignmentMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssignmentMethod", value)
}

// NewAssignmentMethodFromValue returns a pointer to a valid AssignmentMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssignmentMethodFromValue(v string) (*AssignmentMethod, error) {
	ev := AssignmentMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssignmentMethod: valid values are %v", v, AllowedAssignmentMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssignmentMethod) IsValid() bool {
	for _, existing := range AllowedAssignmentMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assignment-method value
func (v AssignmentMethod) Ptr() *AssignmentMethod {
	return &v
}

type NullableAssignmentMethod struct {
	value *AssignmentMethod
	isSet bool
}

func (v NullableAssignmentMethod) Get() *AssignmentMethod {
	return v.value
}

func (v *NullableAssignmentMethod) Set(val *AssignmentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentMethod(val *AssignmentMethod) *NullableAssignmentMethod {
	return &NullableAssignmentMethod{value: val, isSet: true}
}

func (v NullableAssignmentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
