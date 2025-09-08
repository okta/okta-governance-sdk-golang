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

// ReviewerLowerLevelCondition The condition for which, the lower level reviews will move to that level for further review.  This property should not be set for `FIRST` reviewer level. Only `SECOND` reviewer level will have the appropriate condition.  `LOWER_LEVEL_APPROVES` means only approved reviews from lower level moves to that level.  `LOWER_LEVEL_APPROVES_OR_REJECTS` means both approved and revoked reviews from lower level moves to that level.
type ReviewerLowerLevelCondition string

// List of reviewer-lower-level-condition
const (
	REVIEWERLOWERLEVELCONDITION_APPROVES            ReviewerLowerLevelCondition = "LOWER_LEVEL_APPROVES"
	REVIEWERLOWERLEVELCONDITION_APPROVES_OR_REJECTS ReviewerLowerLevelCondition = "LOWER_LEVEL_APPROVES_OR_REJECTS"
)

// All allowed values of ReviewerLowerLevelCondition enum
var AllowedReviewerLowerLevelConditionEnumValues = []ReviewerLowerLevelCondition{
	"LOWER_LEVEL_APPROVES",
	"LOWER_LEVEL_APPROVES_OR_REJECTS",
}

func (v *ReviewerLowerLevelCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReviewerLowerLevelCondition(value)
	for _, existing := range AllowedReviewerLowerLevelConditionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReviewerLowerLevelCondition", value)
}

// NewReviewerLowerLevelConditionFromValue returns a pointer to a valid ReviewerLowerLevelCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReviewerLowerLevelConditionFromValue(v string) (*ReviewerLowerLevelCondition, error) {
	ev := ReviewerLowerLevelCondition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReviewerLowerLevelCondition: valid values are %v", v, AllowedReviewerLowerLevelConditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReviewerLowerLevelCondition) IsValid() bool {
	for _, existing := range AllowedReviewerLowerLevelConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reviewer-lower-level-condition value
func (v ReviewerLowerLevelCondition) Ptr() *ReviewerLowerLevelCondition {
	return &v
}

type NullableReviewerLowerLevelCondition struct {
	value *ReviewerLowerLevelCondition
	isSet bool
}

func (v NullableReviewerLowerLevelCondition) Get() *ReviewerLowerLevelCondition {
	return v.value
}

func (v *NullableReviewerLowerLevelCondition) Set(val *ReviewerLowerLevelCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerLowerLevelCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerLowerLevelCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerLowerLevelCondition(val *ReviewerLowerLevelCondition) *NullableReviewerLowerLevelCondition {
	return &NullableReviewerLowerLevelCondition{value: val, isSet: true}
}

func (v NullableReviewerLowerLevelCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerLowerLevelCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
