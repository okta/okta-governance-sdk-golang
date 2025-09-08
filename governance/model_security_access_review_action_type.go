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

// SecurityAccessReviewActionType The type of action for the security access review
type SecurityAccessReviewActionType string

// List of security-access-review-action-type
const (
	SECURITYACCESSREVIEWACTIONTYPE_UNIVERSAL_LOGOUT     SecurityAccessReviewActionType = "UNIVERSAL_LOGOUT"
	SECURITYACCESSREVIEWACTIONTYPE_SUSPEND_ACCOUNT      SecurityAccessReviewActionType = "SUSPEND_ACCOUNT"
	SECURITYACCESSREVIEWACTIONTYPE_REACTIVATE_ACCOUNT   SecurityAccessReviewActionType = "REACTIVATE_ACCOUNT"
	SECURITYACCESSREVIEWACTIONTYPE_CLOSE_REVIEW         SecurityAccessReviewActionType = "CLOSE_REVIEW"
	SECURITYACCESSREVIEWACTIONTYPE_REVOKE_OKTA_SESSIONS SecurityAccessReviewActionType = "REVOKE_OKTA_SESSIONS"
	SECURITYACCESSREVIEWACTIONTYPE_RESET_MFA            SecurityAccessReviewActionType = "RESET_MFA"
	SECURITYACCESSREVIEWACTIONTYPE_RESET_PASSWORD       SecurityAccessReviewActionType = "RESET_PASSWORD"
	SECURITYACCESSREVIEWACTIONTYPE_RESTORE_ALL_ACCESS   SecurityAccessReviewActionType = "RESTORE_ALL_ACCESS"
)

// All allowed values of SecurityAccessReviewActionType enum
var AllowedSecurityAccessReviewActionTypeEnumValues = []SecurityAccessReviewActionType{
	"UNIVERSAL_LOGOUT",
	"SUSPEND_ACCOUNT",
	"REACTIVATE_ACCOUNT",
	"CLOSE_REVIEW",
	"REVOKE_OKTA_SESSIONS",
	"RESET_MFA",
	"RESET_PASSWORD",
	"RESTORE_ALL_ACCESS",
}

func (v *SecurityAccessReviewActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewActionType(value)
	for _, existing := range AllowedSecurityAccessReviewActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewActionType", value)
}

// NewSecurityAccessReviewActionTypeFromValue returns a pointer to a valid SecurityAccessReviewActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewActionTypeFromValue(v string) (*SecurityAccessReviewActionType, error) {
	ev := SecurityAccessReviewActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewActionType: valid values are %v", v, AllowedSecurityAccessReviewActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewActionType) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-action-type value
func (v SecurityAccessReviewActionType) Ptr() *SecurityAccessReviewActionType {
	return &v
}

type NullableSecurityAccessReviewActionType struct {
	value *SecurityAccessReviewActionType
	isSet bool
}

func (v NullableSecurityAccessReviewActionType) Get() *SecurityAccessReviewActionType {
	return v.value
}

func (v *NullableSecurityAccessReviewActionType) Set(val *SecurityAccessReviewActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewActionType(val *SecurityAccessReviewActionType) *NullableSecurityAccessReviewActionType {
	return &NullableSecurityAccessReviewActionType{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
