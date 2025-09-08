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

// ReviewerType Identifies the kind of reviewer for Access Certification. For example, a reviewer can be a user or an expression.
type ReviewerType string

// List of reviewer-type
const (
	REVIEWERTYPE_USER                ReviewerType = "USER"
	REVIEWERTYPE_REVIEWER_EXPRESSION ReviewerType = "REVIEWER_EXPRESSION"
	REVIEWERTYPE_GROUP               ReviewerType = "GROUP"
	REVIEWERTYPE_RESOURCE_OWNER      ReviewerType = "RESOURCE_OWNER"
)

// All allowed values of ReviewerType enum
var AllowedReviewerTypeEnumValues = []ReviewerType{
	"USER",
	"REVIEWER_EXPRESSION",
	"GROUP",
	"RESOURCE_OWNER",
}

func (v *ReviewerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReviewerType(value)
	for _, existing := range AllowedReviewerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReviewerType", value)
}

// NewReviewerTypeFromValue returns a pointer to a valid ReviewerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReviewerTypeFromValue(v string) (*ReviewerType, error) {
	ev := ReviewerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReviewerType: valid values are %v", v, AllowedReviewerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReviewerType) IsValid() bool {
	for _, existing := range AllowedReviewerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reviewer-type value
func (v ReviewerType) Ptr() *ReviewerType {
	return &v
}

type NullableReviewerType struct {
	value *ReviewerType
	isSet bool
}

func (v NullableReviewerType) Get() *ReviewerType {
	return v.value
}

func (v *NullableReviewerType) Set(val *ReviewerType) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerType) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerType(val *ReviewerType) *NullableReviewerType {
	return &NullableReviewerType{value: val, isSet: true}
}

func (v NullableReviewerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
