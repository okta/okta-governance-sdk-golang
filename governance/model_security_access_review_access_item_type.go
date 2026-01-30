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

// SecurityAccessReviewAccessItemType Access item type
type SecurityAccessReviewAccessItemType string

// List of security-access-review-access-item-type
const (
	SECURITYACCESSREVIEWACCESSITEMTYPE_APPLICATION SecurityAccessReviewAccessItemType = "APPLICATION"
)

// All allowed values of SecurityAccessReviewAccessItemType enum
var AllowedSecurityAccessReviewAccessItemTypeEnumValues = []SecurityAccessReviewAccessItemType{
	"APPLICATION",
}

func (v *SecurityAccessReviewAccessItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewAccessItemType(value)
	for _, existing := range AllowedSecurityAccessReviewAccessItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewAccessItemType", value)
}

// NewSecurityAccessReviewAccessItemTypeFromValue returns a pointer to a valid SecurityAccessReviewAccessItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewAccessItemTypeFromValue(v string) (*SecurityAccessReviewAccessItemType, error) {
	ev := SecurityAccessReviewAccessItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewAccessItemType: valid values are %v", v, AllowedSecurityAccessReviewAccessItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewAccessItemType) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewAccessItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-access-item-type value
func (v SecurityAccessReviewAccessItemType) Ptr() *SecurityAccessReviewAccessItemType {
	return &v
}

type NullableSecurityAccessReviewAccessItemType struct {
	value *SecurityAccessReviewAccessItemType
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItemType) Get() *SecurityAccessReviewAccessItemType {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItemType) Set(val *SecurityAccessReviewAccessItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItemType(val *SecurityAccessReviewAccessItemType) *NullableSecurityAccessReviewAccessItemType {
	return &NullableSecurityAccessReviewAccessItemType{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
