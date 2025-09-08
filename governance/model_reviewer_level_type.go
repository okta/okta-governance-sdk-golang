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

// ReviewerLevelType Identifies the reviewer level of each reviews during access certification. Applicable for multi level campaigns only.
type ReviewerLevelType string

// List of reviewer-level-type
const (
	REVIEWERLEVELTYPE_FIRST  ReviewerLevelType = "FIRST"
	REVIEWERLEVELTYPE_SECOND ReviewerLevelType = "SECOND"
)

// All allowed values of ReviewerLevelType enum
var AllowedReviewerLevelTypeEnumValues = []ReviewerLevelType{
	"FIRST",
	"SECOND",
}

func (v *ReviewerLevelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReviewerLevelType(value)
	for _, existing := range AllowedReviewerLevelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReviewerLevelType", value)
}

// NewReviewerLevelTypeFromValue returns a pointer to a valid ReviewerLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReviewerLevelTypeFromValue(v string) (*ReviewerLevelType, error) {
	ev := ReviewerLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReviewerLevelType: valid values are %v", v, AllowedReviewerLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReviewerLevelType) IsValid() bool {
	for _, existing := range AllowedReviewerLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reviewer-level-type value
func (v ReviewerLevelType) Ptr() *ReviewerLevelType {
	return &v
}

type NullableReviewerLevelType struct {
	value *ReviewerLevelType
	isSet bool
}

func (v NullableReviewerLevelType) Get() *ReviewerLevelType {
	return v.value
}

func (v *NullableReviewerLevelType) Set(val *ReviewerLevelType) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerLevelType) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerLevelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerLevelType(val *ReviewerLevelType) *NullableReviewerLevelType {
	return &NullableReviewerLevelType{value: val, isSet: true}
}

func (v NullableReviewerLevelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerLevelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
