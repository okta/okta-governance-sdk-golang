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

// ApprovalDecisionEnum the model 'ApprovalDecisionEnum'
type ApprovalDecisionEnum string

// List of approval-decision-enum
const (
	APPROVALDECISIONENUM_APPROVE ApprovalDecisionEnum = "APPROVE"
	APPROVALDECISIONENUM_DENY    ApprovalDecisionEnum = "DENY"
	APPROVALDECISIONENUM_CANCEL  ApprovalDecisionEnum = "CANCEL"
)

// All allowed values of ApprovalDecisionEnum enum
var AllowedApprovalDecisionEnumEnumValues = []ApprovalDecisionEnum{
	"APPROVE",
	"DENY",
	"CANCEL",
}

func (v *ApprovalDecisionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovalDecisionEnum(value)
	for _, existing := range AllowedApprovalDecisionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApprovalDecisionEnum", value)
}

// NewApprovalDecisionEnumFromValue returns a pointer to a valid ApprovalDecisionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovalDecisionEnumFromValue(v string) (*ApprovalDecisionEnum, error) {
	ev := ApprovalDecisionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovalDecisionEnum: valid values are %v", v, AllowedApprovalDecisionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovalDecisionEnum) IsValid() bool {
	for _, existing := range AllowedApprovalDecisionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to approval-decision-enum value
func (v ApprovalDecisionEnum) Ptr() *ApprovalDecisionEnum {
	return &v
}

type NullableApprovalDecisionEnum struct {
	value *ApprovalDecisionEnum
	isSet bool
}

func (v NullableApprovalDecisionEnum) Get() *ApprovalDecisionEnum {
	return v.value
}

func (v *NullableApprovalDecisionEnum) Set(val *ApprovalDecisionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalDecisionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalDecisionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalDecisionEnum(val *ApprovalDecisionEnum) *NullableApprovalDecisionEnum {
	return &NullableApprovalDecisionEnum{value: val, isSet: true}
}

func (v NullableApprovalDecisionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalDecisionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
