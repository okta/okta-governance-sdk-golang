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

// BulkReviewCriteriaType Type of criteria for a bulk-review decision
type BulkReviewCriteriaType string

// List of bulk-review-criteria-type
const (
	BULKREVIEWCRITERIATYPE_GOVERNANCE_ANALYZER BulkReviewCriteriaType = "GOVERNANCE_ANALYZER"
)

// All allowed values of BulkReviewCriteriaType enum
var AllowedBulkReviewCriteriaTypeEnumValues = []BulkReviewCriteriaType{
	"GOVERNANCE_ANALYZER",
}

func (v *BulkReviewCriteriaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkReviewCriteriaType(value)
	for _, existing := range AllowedBulkReviewCriteriaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkReviewCriteriaType", value)
}

// NewBulkReviewCriteriaTypeFromValue returns a pointer to a valid BulkReviewCriteriaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkReviewCriteriaTypeFromValue(v string) (*BulkReviewCriteriaType, error) {
	ev := BulkReviewCriteriaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkReviewCriteriaType: valid values are %v", v, AllowedBulkReviewCriteriaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkReviewCriteriaType) IsValid() bool {
	for _, existing := range AllowedBulkReviewCriteriaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bulk-review-criteria-type value
func (v BulkReviewCriteriaType) Ptr() *BulkReviewCriteriaType {
	return &v
}

type NullableBulkReviewCriteriaType struct {
	value *BulkReviewCriteriaType
	isSet bool
}

func (v NullableBulkReviewCriteriaType) Get() *BulkReviewCriteriaType {
	return v.value
}

func (v *NullableBulkReviewCriteriaType) Set(val *BulkReviewCriteriaType) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReviewCriteriaType) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReviewCriteriaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReviewCriteriaType(val *BulkReviewCriteriaType) *NullableBulkReviewCriteriaType {
	return &NullableBulkReviewCriteriaType{value: val, isSet: true}
}

func (v NullableBulkReviewCriteriaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReviewCriteriaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
