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

// SecurityAccessReviewSubAccessItemEntitlementType The entitlement type
type SecurityAccessReviewSubAccessItemEntitlementType string

// List of security-access-review-sub-access-item-entitlement-type
const (
	SECURITYACCESSREVIEWSUBACCESSITEMENTITLEMENTTYPE_ENTITLEMENT_VALUE  SecurityAccessReviewSubAccessItemEntitlementType = "ENTITLEMENT_VALUE"
	SECURITYACCESSREVIEWSUBACCESSITEMENTITLEMENTTYPE_ENTITLEMENT_BUNDLE SecurityAccessReviewSubAccessItemEntitlementType = "ENTITLEMENT_BUNDLE"
)

// All allowed values of SecurityAccessReviewSubAccessItemEntitlementType enum
var AllowedSecurityAccessReviewSubAccessItemEntitlementTypeEnumValues = []SecurityAccessReviewSubAccessItemEntitlementType{
	"ENTITLEMENT_VALUE",
	"ENTITLEMENT_BUNDLE",
}

func (v *SecurityAccessReviewSubAccessItemEntitlementType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewSubAccessItemEntitlementType(value)
	for _, existing := range AllowedSecurityAccessReviewSubAccessItemEntitlementTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewSubAccessItemEntitlementType", value)
}

// NewSecurityAccessReviewSubAccessItemEntitlementTypeFromValue returns a pointer to a valid SecurityAccessReviewSubAccessItemEntitlementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewSubAccessItemEntitlementTypeFromValue(v string) (*SecurityAccessReviewSubAccessItemEntitlementType, error) {
	ev := SecurityAccessReviewSubAccessItemEntitlementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewSubAccessItemEntitlementType: valid values are %v", v, AllowedSecurityAccessReviewSubAccessItemEntitlementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewSubAccessItemEntitlementType) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewSubAccessItemEntitlementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-sub-access-item-entitlement-type value
func (v SecurityAccessReviewSubAccessItemEntitlementType) Ptr() *SecurityAccessReviewSubAccessItemEntitlementType {
	return &v
}

type NullableSecurityAccessReviewSubAccessItemEntitlementType struct {
	value *SecurityAccessReviewSubAccessItemEntitlementType
	isSet bool
}

func (v NullableSecurityAccessReviewSubAccessItemEntitlementType) Get() *SecurityAccessReviewSubAccessItemEntitlementType {
	return v.value
}

func (v *NullableSecurityAccessReviewSubAccessItemEntitlementType) Set(val *SecurityAccessReviewSubAccessItemEntitlementType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSubAccessItemEntitlementType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSubAccessItemEntitlementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSubAccessItemEntitlementType(val *SecurityAccessReviewSubAccessItemEntitlementType) *NullableSecurityAccessReviewSubAccessItemEntitlementType {
	return &NullableSecurityAccessReviewSubAccessItemEntitlementType{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSubAccessItemEntitlementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSubAccessItemEntitlementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
