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

// SecurityAccessReviewAccessItemSupportedAction Supported actions
type SecurityAccessReviewAccessItemSupportedAction string

// List of security-access-review-access-item-supported-action
const (
	SECURITYACCESSREVIEWACCESSITEMSUPPORTEDACTION_REVOKE_ACCESS               SecurityAccessReviewAccessItemSupportedAction = "REVOKE_ACCESS"
	SECURITYACCESSREVIEWACCESSITEMSUPPORTEDACTION_RESTORE_ACCESS              SecurityAccessReviewAccessItemSupportedAction = "RESTORE_ACCESS"
	SECURITYACCESSREVIEWACCESSITEMSUPPORTEDACTION_FLAG_FOR_MANUAL_REMEDIATION SecurityAccessReviewAccessItemSupportedAction = "FLAG_FOR_MANUAL_REMEDIATION"
	SECURITYACCESSREVIEWACCESSITEMSUPPORTEDACTION_FLAG_FOR_MANUAL_RESTORATION SecurityAccessReviewAccessItemSupportedAction = "FLAG_FOR_MANUAL_RESTORATION"
)

// All allowed values of SecurityAccessReviewAccessItemSupportedAction enum
var AllowedSecurityAccessReviewAccessItemSupportedActionEnumValues = []SecurityAccessReviewAccessItemSupportedAction{
	"REVOKE_ACCESS",
	"RESTORE_ACCESS",
	"FLAG_FOR_MANUAL_REMEDIATION",
	"FLAG_FOR_MANUAL_RESTORATION",
}

func (v *SecurityAccessReviewAccessItemSupportedAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewAccessItemSupportedAction(value)
	for _, existing := range AllowedSecurityAccessReviewAccessItemSupportedActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewAccessItemSupportedAction", value)
}

// NewSecurityAccessReviewAccessItemSupportedActionFromValue returns a pointer to a valid SecurityAccessReviewAccessItemSupportedAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewAccessItemSupportedActionFromValue(v string) (*SecurityAccessReviewAccessItemSupportedAction, error) {
	ev := SecurityAccessReviewAccessItemSupportedAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewAccessItemSupportedAction: valid values are %v", v, AllowedSecurityAccessReviewAccessItemSupportedActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewAccessItemSupportedAction) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewAccessItemSupportedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-access-item-supported-action value
func (v SecurityAccessReviewAccessItemSupportedAction) Ptr() *SecurityAccessReviewAccessItemSupportedAction {
	return &v
}

type NullableSecurityAccessReviewAccessItemSupportedAction struct {
	value *SecurityAccessReviewAccessItemSupportedAction
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItemSupportedAction) Get() *SecurityAccessReviewAccessItemSupportedAction {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItemSupportedAction) Set(val *SecurityAccessReviewAccessItemSupportedAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItemSupportedAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItemSupportedAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItemSupportedAction(val *SecurityAccessReviewAccessItemSupportedAction) *NullableSecurityAccessReviewAccessItemSupportedAction {
	return &NullableSecurityAccessReviewAccessItemSupportedAction{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItemSupportedAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItemSupportedAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
