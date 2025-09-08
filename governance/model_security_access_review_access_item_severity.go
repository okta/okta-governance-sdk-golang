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

// SecurityAccessReviewAccessItemSeverity The access severity
type SecurityAccessReviewAccessItemSeverity string

// List of security-access-review-access-item-severity
const (
	SECURITYACCESSREVIEWACCESSITEMSEVERITY_LOW    SecurityAccessReviewAccessItemSeverity = "LOW"
	SECURITYACCESSREVIEWACCESSITEMSEVERITY_MEDIUM SecurityAccessReviewAccessItemSeverity = "MEDIUM"
	SECURITYACCESSREVIEWACCESSITEMSEVERITY_HIGH   SecurityAccessReviewAccessItemSeverity = "HIGH"
)

// All allowed values of SecurityAccessReviewAccessItemSeverity enum
var AllowedSecurityAccessReviewAccessItemSeverityEnumValues = []SecurityAccessReviewAccessItemSeverity{
	"LOW",
	"MEDIUM",
	"HIGH",
}

func (v *SecurityAccessReviewAccessItemSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewAccessItemSeverity(value)
	for _, existing := range AllowedSecurityAccessReviewAccessItemSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewAccessItemSeverity", value)
}

// NewSecurityAccessReviewAccessItemSeverityFromValue returns a pointer to a valid SecurityAccessReviewAccessItemSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewAccessItemSeverityFromValue(v string) (*SecurityAccessReviewAccessItemSeverity, error) {
	ev := SecurityAccessReviewAccessItemSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewAccessItemSeverity: valid values are %v", v, AllowedSecurityAccessReviewAccessItemSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewAccessItemSeverity) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewAccessItemSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-access-item-severity value
func (v SecurityAccessReviewAccessItemSeverity) Ptr() *SecurityAccessReviewAccessItemSeverity {
	return &v
}

type NullableSecurityAccessReviewAccessItemSeverity struct {
	value *SecurityAccessReviewAccessItemSeverity
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItemSeverity) Get() *SecurityAccessReviewAccessItemSeverity {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItemSeverity) Set(val *SecurityAccessReviewAccessItemSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItemSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItemSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItemSeverity(val *SecurityAccessReviewAccessItemSeverity) *NullableSecurityAccessReviewAccessItemSeverity {
	return &NullableSecurityAccessReviewAccessItemSeverity{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItemSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItemSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
