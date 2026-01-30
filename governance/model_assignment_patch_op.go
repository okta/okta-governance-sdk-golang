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

// AssignmentPatchOp The operation performed for the update
type AssignmentPatchOp string

// List of assignment-patch-op
const (
	ASSIGNMENTPATCHOP_ADD     AssignmentPatchOp = "ADD"
	ASSIGNMENTPATCHOP_REMOVE  AssignmentPatchOp = "REMOVE"
	ASSIGNMENTPATCHOP_REPLACE AssignmentPatchOp = "REPLACE"
)

// All allowed values of AssignmentPatchOp enum
var AllowedAssignmentPatchOpEnumValues = []AssignmentPatchOp{
	"ADD",
	"REMOVE",
	"REPLACE",
}

func (v *AssignmentPatchOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssignmentPatchOp(value)
	for _, existing := range AllowedAssignmentPatchOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssignmentPatchOp", value)
}

// NewAssignmentPatchOpFromValue returns a pointer to a valid AssignmentPatchOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssignmentPatchOpFromValue(v string) (*AssignmentPatchOp, error) {
	ev := AssignmentPatchOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssignmentPatchOp: valid values are %v", v, AllowedAssignmentPatchOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssignmentPatchOp) IsValid() bool {
	for _, existing := range AllowedAssignmentPatchOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assignment-patch-op value
func (v AssignmentPatchOp) Ptr() *AssignmentPatchOp {
	return &v
}

type NullableAssignmentPatchOp struct {
	value *AssignmentPatchOp
	isSet bool
}

func (v NullableAssignmentPatchOp) Get() *AssignmentPatchOp {
	return v.value
}

func (v *NullableAssignmentPatchOp) Set(val *AssignmentPatchOp) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentPatchOp) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentPatchOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentPatchOp(val *AssignmentPatchOp) *NullableAssignmentPatchOp {
	return &NullableAssignmentPatchOp{value: val, isSet: true}
}

func (v NullableAssignmentPatchOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentPatchOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
