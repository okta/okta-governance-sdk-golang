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

// SecurityAccessReviewAnomalyType The access item target type
type SecurityAccessReviewAnomalyType string

// List of security-access-review-anomaly-type
const (
	SECURITYACCESSREVIEWANOMALYTYPE_SOD_CONFLICT              SecurityAccessReviewAnomalyType = "SOD_CONFLICT"
	SECURITYACCESSREVIEWANOMALYTYPE_ASSIGNMENT_METHOD         SecurityAccessReviewAnomalyType = "ASSIGNMENT_METHOD"
	SECURITYACCESSREVIEWANOMALYTYPE_PAST_GOVERNANCE_DECISIONS SecurityAccessReviewAnomalyType = "PAST_GOVERNANCE_DECISIONS"
	SECURITYACCESSREVIEWANOMALYTYPE_USAGE_HISTORY             SecurityAccessReviewAnomalyType = "USAGE_HISTORY"
)

// All allowed values of SecurityAccessReviewAnomalyType enum
var AllowedSecurityAccessReviewAnomalyTypeEnumValues = []SecurityAccessReviewAnomalyType{
	"SOD_CONFLICT",
	"ASSIGNMENT_METHOD",
	"PAST_GOVERNANCE_DECISIONS",
	"USAGE_HISTORY",
}

func (v *SecurityAccessReviewAnomalyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewAnomalyType(value)
	for _, existing := range AllowedSecurityAccessReviewAnomalyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewAnomalyType", value)
}

// NewSecurityAccessReviewAnomalyTypeFromValue returns a pointer to a valid SecurityAccessReviewAnomalyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewAnomalyTypeFromValue(v string) (*SecurityAccessReviewAnomalyType, error) {
	ev := SecurityAccessReviewAnomalyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewAnomalyType: valid values are %v", v, AllowedSecurityAccessReviewAnomalyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewAnomalyType) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewAnomalyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-anomaly-type value
func (v SecurityAccessReviewAnomalyType) Ptr() *SecurityAccessReviewAnomalyType {
	return &v
}

type NullableSecurityAccessReviewAnomalyType struct {
	value *SecurityAccessReviewAnomalyType
	isSet bool
}

func (v NullableSecurityAccessReviewAnomalyType) Get() *SecurityAccessReviewAnomalyType {
	return v.value
}

func (v *NullableSecurityAccessReviewAnomalyType) Set(val *SecurityAccessReviewAnomalyType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAnomalyType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAnomalyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAnomalyType(val *SecurityAccessReviewAnomalyType) *NullableSecurityAccessReviewAnomalyType {
	return &NullableSecurityAccessReviewAnomalyType{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAnomalyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAnomalyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
