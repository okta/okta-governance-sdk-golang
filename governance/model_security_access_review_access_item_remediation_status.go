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

// SecurityAccessReviewAccessItemRemediationStatus Remediation status
type SecurityAccessReviewAccessItemRemediationStatus string

// List of security-access-review-access-item-remediation-status
const (
	SECURITYACCESSREVIEWACCESSITEMREMEDIATIONSTATUS_PENDING_MANUAL_REMEDIATION SecurityAccessReviewAccessItemRemediationStatus = "PENDING_MANUAL_REMEDIATION"
	SECURITYACCESSREVIEWACCESSITEMREMEDIATIONSTATUS_PENDING_MANUAL_RESTORATION SecurityAccessReviewAccessItemRemediationStatus = "PENDING_MANUAL_RESTORATION"
	SECURITYACCESSREVIEWACCESSITEMREMEDIATIONSTATUS_REMEDIATED                 SecurityAccessReviewAccessItemRemediationStatus = "REMEDIATED"
	SECURITYACCESSREVIEWACCESSITEMREMEDIATIONSTATUS_IN_PROGRESS                SecurityAccessReviewAccessItemRemediationStatus = "IN_PROGRESS"
	SECURITYACCESSREVIEWACCESSITEMREMEDIATIONSTATUS_NONE                       SecurityAccessReviewAccessItemRemediationStatus = "NONE"
)

// All allowed values of SecurityAccessReviewAccessItemRemediationStatus enum
var AllowedSecurityAccessReviewAccessItemRemediationStatusEnumValues = []SecurityAccessReviewAccessItemRemediationStatus{
	"PENDING_MANUAL_REMEDIATION",
	"PENDING_MANUAL_RESTORATION",
	"REMEDIATED",
	"IN_PROGRESS",
	"NONE",
}

func (v *SecurityAccessReviewAccessItemRemediationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewAccessItemRemediationStatus(value)
	for _, existing := range AllowedSecurityAccessReviewAccessItemRemediationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewAccessItemRemediationStatus", value)
}

// NewSecurityAccessReviewAccessItemRemediationStatusFromValue returns a pointer to a valid SecurityAccessReviewAccessItemRemediationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewAccessItemRemediationStatusFromValue(v string) (*SecurityAccessReviewAccessItemRemediationStatus, error) {
	ev := SecurityAccessReviewAccessItemRemediationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewAccessItemRemediationStatus: valid values are %v", v, AllowedSecurityAccessReviewAccessItemRemediationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewAccessItemRemediationStatus) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewAccessItemRemediationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-access-item-remediation-status value
func (v SecurityAccessReviewAccessItemRemediationStatus) Ptr() *SecurityAccessReviewAccessItemRemediationStatus {
	return &v
}

type NullableSecurityAccessReviewAccessItemRemediationStatus struct {
	value *SecurityAccessReviewAccessItemRemediationStatus
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItemRemediationStatus) Get() *SecurityAccessReviewAccessItemRemediationStatus {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItemRemediationStatus) Set(val *SecurityAccessReviewAccessItemRemediationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItemRemediationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItemRemediationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItemRemediationStatus(val *SecurityAccessReviewAccessItemRemediationStatus) *NullableSecurityAccessReviewAccessItemRemediationStatus {
	return &NullableSecurityAccessReviewAccessItemRemediationStatus{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItemRemediationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItemRemediationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
