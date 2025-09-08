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

// SecurityAccessReviewSubAccessItemType the model 'SecurityAccessReviewSubAccessItemType'
type SecurityAccessReviewSubAccessItemType string

// List of security-access-review-sub-access-item-type
const (
	SECURITYACCESSREVIEWSUBACCESSITEMTYPE_ENTITLEMENT SecurityAccessReviewSubAccessItemType = "ENTITLEMENT"
	SECURITYACCESSREVIEWSUBACCESSITEMTYPE_GROUP       SecurityAccessReviewSubAccessItemType = "GROUP"
)

// All allowed values of SecurityAccessReviewSubAccessItemType enum
var AllowedSecurityAccessReviewSubAccessItemTypeEnumValues = []SecurityAccessReviewSubAccessItemType{
	"ENTITLEMENT",
	"GROUP",
}

func (v *SecurityAccessReviewSubAccessItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewSubAccessItemType(value)
	for _, existing := range AllowedSecurityAccessReviewSubAccessItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewSubAccessItemType", value)
}

// NewSecurityAccessReviewSubAccessItemTypeFromValue returns a pointer to a valid SecurityAccessReviewSubAccessItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewSubAccessItemTypeFromValue(v string) (*SecurityAccessReviewSubAccessItemType, error) {
	ev := SecurityAccessReviewSubAccessItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewSubAccessItemType: valid values are %v", v, AllowedSecurityAccessReviewSubAccessItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewSubAccessItemType) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewSubAccessItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-sub-access-item-type value
func (v SecurityAccessReviewSubAccessItemType) Ptr() *SecurityAccessReviewSubAccessItemType {
	return &v
}

type NullableSecurityAccessReviewSubAccessItemType struct {
	value *SecurityAccessReviewSubAccessItemType
	isSet bool
}

func (v NullableSecurityAccessReviewSubAccessItemType) Get() *SecurityAccessReviewSubAccessItemType {
	return v.value
}

func (v *NullableSecurityAccessReviewSubAccessItemType) Set(val *SecurityAccessReviewSubAccessItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSubAccessItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSubAccessItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSubAccessItemType(val *SecurityAccessReviewSubAccessItemType) *NullableSecurityAccessReviewSubAccessItemType {
	return &NullableSecurityAccessReviewSubAccessItemType{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSubAccessItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSubAccessItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
