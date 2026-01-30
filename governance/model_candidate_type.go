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

// CandidateType The type of assignment candidate
type CandidateType string

// List of candidate-type
const (
	CANDIDATETYPE_OKTA_USER  CandidateType = "OKTA_USER"
	CANDIDATETYPE_OKTA_GROUP CandidateType = "OKTA_GROUP"
)

// All allowed values of CandidateType enum
var AllowedCandidateTypeEnumValues = []CandidateType{
	"OKTA_USER",
	"OKTA_GROUP",
}

func (v *CandidateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CandidateType(value)
	for _, existing := range AllowedCandidateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CandidateType", value)
}

// NewCandidateTypeFromValue returns a pointer to a valid CandidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCandidateTypeFromValue(v string) (*CandidateType, error) {
	ev := CandidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CandidateType: valid values are %v", v, AllowedCandidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CandidateType) IsValid() bool {
	for _, existing := range AllowedCandidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to candidate-type value
func (v CandidateType) Ptr() *CandidateType {
	return &v
}

type NullableCandidateType struct {
	value *CandidateType
	isSet bool
}

func (v NullableCandidateType) Get() *CandidateType {
	return v.value
}

func (v *NullableCandidateType) Set(val *CandidateType) {
	v.value = val
	v.isSet = true
}

func (v NullableCandidateType) IsSet() bool {
	return v.isSet
}

func (v *NullableCandidateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCandidateType(val *CandidateType) *NullableCandidateType {
	return &NullableCandidateType{value: val, isSet: true}
}

func (v NullableCandidateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCandidateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
