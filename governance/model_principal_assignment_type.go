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

// PrincipalAssignmentType Collection assignment type or source
type PrincipalAssignmentType string

// List of principal-assignment-type
const (
	PRINCIPALASSIGNMENTTYPE_INDIVIDUAL PrincipalAssignmentType = "INDIVIDUAL"
)

// All allowed values of PrincipalAssignmentType enum
var AllowedPrincipalAssignmentTypeEnumValues = []PrincipalAssignmentType{
	"INDIVIDUAL",
}

func (v *PrincipalAssignmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrincipalAssignmentType(value)
	for _, existing := range AllowedPrincipalAssignmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrincipalAssignmentType", value)
}

// NewPrincipalAssignmentTypeFromValue returns a pointer to a valid PrincipalAssignmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrincipalAssignmentTypeFromValue(v string) (*PrincipalAssignmentType, error) {
	ev := PrincipalAssignmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrincipalAssignmentType: valid values are %v", v, AllowedPrincipalAssignmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrincipalAssignmentType) IsValid() bool {
	for _, existing := range AllowedPrincipalAssignmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to principal-assignment-type value
func (v PrincipalAssignmentType) Ptr() *PrincipalAssignmentType {
	return &v
}

type NullablePrincipalAssignmentType struct {
	value *PrincipalAssignmentType
	isSet bool
}

func (v NullablePrincipalAssignmentType) Get() *PrincipalAssignmentType {
	return v.value
}

func (v *NullablePrincipalAssignmentType) Set(val *PrincipalAssignmentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalAssignmentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalAssignmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalAssignmentType(val *PrincipalAssignmentType) *NullablePrincipalAssignmentType {
	return &NullablePrincipalAssignmentType{value: val, isSet: true}
}

func (v NullablePrincipalAssignmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalAssignmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
