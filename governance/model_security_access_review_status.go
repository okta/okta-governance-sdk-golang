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

// SecurityAccessReviewStatus Security access review status
type SecurityAccessReviewStatus string

// List of security-access-review-status
const (
	SECURITYACCESSREVIEWSTATUS_ACTIVE  SecurityAccessReviewStatus = "ACTIVE"
	SECURITYACCESSREVIEWSTATUS_CLOSED  SecurityAccessReviewStatus = "CLOSED"
	SECURITYACCESSREVIEWSTATUS_PENDING SecurityAccessReviewStatus = "PENDING"
	SECURITYACCESSREVIEWSTATUS_ERROR   SecurityAccessReviewStatus = "ERROR"
)

// All allowed values of SecurityAccessReviewStatus enum
var AllowedSecurityAccessReviewStatusEnumValues = []SecurityAccessReviewStatus{
	"ACTIVE",
	"CLOSED",
	"PENDING",
	"ERROR",
}

func (v *SecurityAccessReviewStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewStatus(value)
	for _, existing := range AllowedSecurityAccessReviewStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewStatus", value)
}

// NewSecurityAccessReviewStatusFromValue returns a pointer to a valid SecurityAccessReviewStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewStatusFromValue(v string) (*SecurityAccessReviewStatus, error) {
	ev := SecurityAccessReviewStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewStatus: valid values are %v", v, AllowedSecurityAccessReviewStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewStatus) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-status value
func (v SecurityAccessReviewStatus) Ptr() *SecurityAccessReviewStatus {
	return &v
}

type NullableSecurityAccessReviewStatus struct {
	value *SecurityAccessReviewStatus
	isSet bool
}

func (v NullableSecurityAccessReviewStatus) Get() *SecurityAccessReviewStatus {
	return v.value
}

func (v *NullableSecurityAccessReviewStatus) Set(val *SecurityAccessReviewStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewStatus(val *SecurityAccessReviewStatus) *NullableSecurityAccessReviewStatus {
	return &NullableSecurityAccessReviewStatus{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
