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

// SecurityAccessReviewConflictedEntitlementListScope Shows the scope used to calculate the conflict between two entitlement lists
type SecurityAccessReviewConflictedEntitlementListScope string

// List of security-access-review-conflicted-entitlement-list-scope
const (
	SECURITYACCESSREVIEWCONFLICTEDENTITLEMENTLISTSCOPE_ANY_ONE_OF SecurityAccessReviewConflictedEntitlementListScope = "ANY_ONE_OF"
	SECURITYACCESSREVIEWCONFLICTEDENTITLEMENTLISTSCOPE_ALL_OF     SecurityAccessReviewConflictedEntitlementListScope = "ALL_OF"
)

// All allowed values of SecurityAccessReviewConflictedEntitlementListScope enum
var AllowedSecurityAccessReviewConflictedEntitlementListScopeEnumValues = []SecurityAccessReviewConflictedEntitlementListScope{
	"ANY_ONE_OF",
	"ALL_OF",
}

func (v *SecurityAccessReviewConflictedEntitlementListScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewConflictedEntitlementListScope(value)
	for _, existing := range AllowedSecurityAccessReviewConflictedEntitlementListScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewConflictedEntitlementListScope", value)
}

// NewSecurityAccessReviewConflictedEntitlementListScopeFromValue returns a pointer to a valid SecurityAccessReviewConflictedEntitlementListScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewConflictedEntitlementListScopeFromValue(v string) (*SecurityAccessReviewConflictedEntitlementListScope, error) {
	ev := SecurityAccessReviewConflictedEntitlementListScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewConflictedEntitlementListScope: valid values are %v", v, AllowedSecurityAccessReviewConflictedEntitlementListScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewConflictedEntitlementListScope) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewConflictedEntitlementListScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-conflicted-entitlement-list-scope value
func (v SecurityAccessReviewConflictedEntitlementListScope) Ptr() *SecurityAccessReviewConflictedEntitlementListScope {
	return &v
}

type NullableSecurityAccessReviewConflictedEntitlementListScope struct {
	value *SecurityAccessReviewConflictedEntitlementListScope
	isSet bool
}

func (v NullableSecurityAccessReviewConflictedEntitlementListScope) Get() *SecurityAccessReviewConflictedEntitlementListScope {
	return v.value
}

func (v *NullableSecurityAccessReviewConflictedEntitlementListScope) Set(val *SecurityAccessReviewConflictedEntitlementListScope) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewConflictedEntitlementListScope) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewConflictedEntitlementListScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewConflictedEntitlementListScope(val *SecurityAccessReviewConflictedEntitlementListScope) *NullableSecurityAccessReviewConflictedEntitlementListScope {
	return &NullableSecurityAccessReviewConflictedEntitlementListScope{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewConflictedEntitlementListScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewConflictedEntitlementListScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
