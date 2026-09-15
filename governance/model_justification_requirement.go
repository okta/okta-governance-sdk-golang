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

// JustificationRequirement Controls when justification is required for review decisions.  For resource campaigns that have the Okta Admin Console as one of the resources, this property must be one of `REQUIRED_FOR_ALL`, `REQUIRED_FOR_REVOKE`, or `REQUIRED_FOR_REVOKE_OPTIONAL_FOR_APPROVE`.
type JustificationRequirement string

// List of justification-requirement
const (
	JUSTIFICATIONREQUIREMENT_OPTIONAL                                 JustificationRequirement = "OPTIONAL"
	JUSTIFICATIONREQUIREMENT_REQUIRED_FOR_REVOKE                      JustificationRequirement = "REQUIRED_FOR_REVOKE"
	JUSTIFICATIONREQUIREMENT_REQUIRED_FOR_REVOKE_OPTIONAL_FOR_APPROVE JustificationRequirement = "REQUIRED_FOR_REVOKE_OPTIONAL_FOR_APPROVE"
	JUSTIFICATIONREQUIREMENT_REQUIRED_FOR_ALL                         JustificationRequirement = "REQUIRED_FOR_ALL"
	JUSTIFICATIONREQUIREMENT_DISABLED                                 JustificationRequirement = "DISABLED"
)

// All allowed values of JustificationRequirement enum
var AllowedJustificationRequirementEnumValues = []JustificationRequirement{
	"OPTIONAL",
	"REQUIRED_FOR_REVOKE",
	"REQUIRED_FOR_REVOKE_OPTIONAL_FOR_APPROVE",
	"REQUIRED_FOR_ALL",
	"DISABLED",
}

func (v *JustificationRequirement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JustificationRequirement(value)
	for _, existing := range AllowedJustificationRequirementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JustificationRequirement", value)
}

// NewJustificationRequirementFromValue returns a pointer to a valid JustificationRequirement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJustificationRequirementFromValue(v string) (*JustificationRequirement, error) {
	ev := JustificationRequirement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JustificationRequirement: valid values are %v", v, AllowedJustificationRequirementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JustificationRequirement) IsValid() bool {
	for _, existing := range AllowedJustificationRequirementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to justification-requirement value
func (v JustificationRequirement) Ptr() *JustificationRequirement {
	return &v
}

type NullableJustificationRequirement struct {
	value *JustificationRequirement
	isSet bool
}

func (v NullableJustificationRequirement) Get() *JustificationRequirement {
	return v.value
}

func (v *NullableJustificationRequirement) Set(val *JustificationRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableJustificationRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableJustificationRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJustificationRequirement(val *JustificationRequirement) *NullableJustificationRequirement {
	return &NullableJustificationRequirement{value: val, isSet: true}
}

func (v NullableJustificationRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJustificationRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
