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

// BulkReviewDecisionStatusType The bulk-review submission request status
type BulkReviewDecisionStatusType string

// List of bulk-review-decision-status-type
const (
	BULKREVIEWDECISIONSTATUSTYPE_ACCEPTED    BulkReviewDecisionStatusType = "ACCEPTED"
	BULKREVIEWDECISIONSTATUSTYPE_COMPLETED   BulkReviewDecisionStatusType = "COMPLETED"
	BULKREVIEWDECISIONSTATUSTYPE_ERROR       BulkReviewDecisionStatusType = "ERROR"
	BULKREVIEWDECISIONSTATUSTYPE_IN_PROGRESS BulkReviewDecisionStatusType = "IN_PROGRESS"
)

// All allowed values of BulkReviewDecisionStatusType enum
var AllowedBulkReviewDecisionStatusTypeEnumValues = []BulkReviewDecisionStatusType{
	"ACCEPTED",
	"COMPLETED",
	"ERROR",
	"IN_PROGRESS",
}

func (v *BulkReviewDecisionStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkReviewDecisionStatusType(value)
	for _, existing := range AllowedBulkReviewDecisionStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkReviewDecisionStatusType", value)
}

// NewBulkReviewDecisionStatusTypeFromValue returns a pointer to a valid BulkReviewDecisionStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkReviewDecisionStatusTypeFromValue(v string) (*BulkReviewDecisionStatusType, error) {
	ev := BulkReviewDecisionStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkReviewDecisionStatusType: valid values are %v", v, AllowedBulkReviewDecisionStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkReviewDecisionStatusType) IsValid() bool {
	for _, existing := range AllowedBulkReviewDecisionStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bulk-review-decision-status-type value
func (v BulkReviewDecisionStatusType) Ptr() *BulkReviewDecisionStatusType {
	return &v
}

type NullableBulkReviewDecisionStatusType struct {
	value *BulkReviewDecisionStatusType
	isSet bool
}

func (v NullableBulkReviewDecisionStatusType) Get() *BulkReviewDecisionStatusType {
	return v.value
}

func (v *NullableBulkReviewDecisionStatusType) Set(val *BulkReviewDecisionStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewDecisionStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewDecisionStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewDecisionStatusType(val *BulkReviewDecisionStatusType) *NullableBulkReviewDecisionStatusType {
	return &NullableBulkReviewDecisionStatusType{value: val, isSet: true}
}

func (v NullableBulkReviewDecisionStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewDecisionStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
