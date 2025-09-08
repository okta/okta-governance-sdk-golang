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

// ReviewersReviewerType Identifies the kind of reviewer for Access Certification.
type ReviewersReviewerType string

// List of reviewers-reviewer-type
const (
	REVIEWERSREVIEWERTYPE_USER           ReviewersReviewerType = "USER"
	REVIEWERSREVIEWERTYPE_GROUP          ReviewersReviewerType = "GROUP"
	REVIEWERSREVIEWERTYPE_RESOURCE_OWNER ReviewersReviewerType = "RESOURCE_OWNER"
)

// All allowed values of ReviewersReviewerType enum
var AllowedReviewersReviewerTypeEnumValues = []ReviewersReviewerType{
	"USER",
	"GROUP",
	"RESOURCE_OWNER",
}

func (v *ReviewersReviewerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReviewersReviewerType(value)
	for _, existing := range AllowedReviewersReviewerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReviewersReviewerType", value)
}

// NewReviewersReviewerTypeFromValue returns a pointer to a valid ReviewersReviewerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReviewersReviewerTypeFromValue(v string) (*ReviewersReviewerType, error) {
	ev := ReviewersReviewerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReviewersReviewerType: valid values are %v", v, AllowedReviewersReviewerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReviewersReviewerType) IsValid() bool {
	for _, existing := range AllowedReviewersReviewerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reviewers-reviewer-type value
func (v ReviewersReviewerType) Ptr() *ReviewersReviewerType {
	return &v
}

type NullableReviewersReviewerType struct {
	value *ReviewersReviewerType
	isSet bool
}

func (v NullableReviewersReviewerType) Get() *ReviewersReviewerType {
	return v.value
}

func (v *NullableReviewersReviewerType) Set(val *ReviewersReviewerType) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewersReviewerType) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewersReviewerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewersReviewerType(val *ReviewersReviewerType) *NullableReviewersReviewerType {
	return &NullableReviewersReviewerType{value: val, isSet: true}
}

func (v NullableReviewersReviewerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewersReviewerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
