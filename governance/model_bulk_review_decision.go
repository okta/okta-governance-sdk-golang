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

// BulkReviewDecision The decision applied to all reviews matching the bulk-review criteria
type BulkReviewDecision string

// List of bulk-review-decision
const (
	BULKREVIEWDECISION_APPROVE BulkReviewDecision = "APPROVE"
)

// All allowed values of BulkReviewDecision enum
var AllowedBulkReviewDecisionEnumValues = []BulkReviewDecision{
	"APPROVE",
}

func (v *BulkReviewDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkReviewDecision(value)
	for _, existing := range AllowedBulkReviewDecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkReviewDecision", value)
}

// NewBulkReviewDecisionFromValue returns a pointer to a valid BulkReviewDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkReviewDecisionFromValue(v string) (*BulkReviewDecision, error) {
	ev := BulkReviewDecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkReviewDecision: valid values are %v", v, AllowedBulkReviewDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkReviewDecision) IsValid() bool {
	for _, existing := range AllowedBulkReviewDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bulk-review-decision value
func (v BulkReviewDecision) Ptr() *BulkReviewDecision {
	return &v
}

type NullableBulkReviewDecision struct {
	value *BulkReviewDecision
	isSet bool
}

func (v NullableBulkReviewDecision) Get() *BulkReviewDecision {
	return v.value
}

func (v *NullableBulkReviewDecision) Set(val *BulkReviewDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewDecision(val *BulkReviewDecision) *NullableBulkReviewDecision {
	return &NullableBulkReviewDecision{value: val, isSet: true}
}

func (v NullableBulkReviewDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
