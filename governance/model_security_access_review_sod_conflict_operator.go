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

// SecurityAccessReviewSodConflictOperator The boolean operator used to calculate the conflict between the given entitlement lists in the SOD rule
type SecurityAccessReviewSodConflictOperator string

// List of security-access-review-sod-conflict-operator
const (
	SECURITYACCESSREVIEWSODCONFLICTOPERATOR_AND SecurityAccessReviewSodConflictOperator = "AND"
)

// All allowed values of SecurityAccessReviewSodConflictOperator enum
var AllowedSecurityAccessReviewSodConflictOperatorEnumValues = []SecurityAccessReviewSodConflictOperator{
	"AND",
}

func (v *SecurityAccessReviewSodConflictOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewSodConflictOperator(value)
	for _, existing := range AllowedSecurityAccessReviewSodConflictOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewSodConflictOperator", value)
}

// NewSecurityAccessReviewSodConflictOperatorFromValue returns a pointer to a valid SecurityAccessReviewSodConflictOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewSodConflictOperatorFromValue(v string) (*SecurityAccessReviewSodConflictOperator, error) {
	ev := SecurityAccessReviewSodConflictOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewSodConflictOperator: valid values are %v", v, AllowedSecurityAccessReviewSodConflictOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewSodConflictOperator) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewSodConflictOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-sod-conflict-operator value
func (v SecurityAccessReviewSodConflictOperator) Ptr() *SecurityAccessReviewSodConflictOperator {
	return &v
}

type NullableSecurityAccessReviewSodConflictOperator struct {
	value *SecurityAccessReviewSodConflictOperator
	isSet bool
}

func (v NullableSecurityAccessReviewSodConflictOperator) Get() *SecurityAccessReviewSodConflictOperator {
	return v.value
}

func (v *NullableSecurityAccessReviewSodConflictOperator) Set(val *SecurityAccessReviewSodConflictOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSodConflictOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSodConflictOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSodConflictOperator(val *SecurityAccessReviewSodConflictOperator) *NullableSecurityAccessReviewSodConflictOperator {
	return &NullableSecurityAccessReviewSodConflictOperator{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSodConflictOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSodConflictOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
