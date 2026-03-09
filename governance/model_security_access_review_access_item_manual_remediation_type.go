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

// SecurityAccessReviewAccessItemManualRemediationType The reason manual remediation is required
type SecurityAccessReviewAccessItemManualRemediationType string

// List of security-access-review-access-item-manual-remediation-type
const (
	SECURITYACCESSREVIEWACCESSITEMMANUALREMEDIATIONTYPE_GROUP_RULE         SecurityAccessReviewAccessItemManualRemediationType = "GROUP_RULE"
	SECURITYACCESSREVIEWACCESSITEMMANUALREMEDIATIONTYPE_ENTITLEMENT_POLICY SecurityAccessReviewAccessItemManualRemediationType = "ENTITLEMENT_POLICY"
	SECURITYACCESSREVIEWACCESSITEMMANUALREMEDIATIONTYPE_COLLECTION         SecurityAccessReviewAccessItemManualRemediationType = "COLLECTION"
	SECURITYACCESSREVIEWACCESSITEMMANUALREMEDIATIONTYPE_APP_GROUP          SecurityAccessReviewAccessItemManualRemediationType = "APP_GROUP"
	SECURITYACCESSREVIEWACCESSITEMMANUALREMEDIATIONTYPE_UNSUPPORTED_GROUP  SecurityAccessReviewAccessItemManualRemediationType = "UNSUPPORTED_GROUP"
	SECURITYACCESSREVIEWACCESSITEMMANUALREMEDIATIONTYPE_FIRST_PARTY_APP    SecurityAccessReviewAccessItemManualRemediationType = "FIRST_PARTY_APP"
)

// All allowed values of SecurityAccessReviewAccessItemManualRemediationType enum
var AllowedSecurityAccessReviewAccessItemManualRemediationTypeEnumValues = []SecurityAccessReviewAccessItemManualRemediationType{
	"GROUP_RULE",
	"ENTITLEMENT_POLICY",
	"COLLECTION",
	"APP_GROUP",
	"UNSUPPORTED_GROUP",
	"FIRST_PARTY_APP",
}

func (v *SecurityAccessReviewAccessItemManualRemediationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityAccessReviewAccessItemManualRemediationType(value)
	for _, existing := range AllowedSecurityAccessReviewAccessItemManualRemediationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityAccessReviewAccessItemManualRemediationType", value)
}

// NewSecurityAccessReviewAccessItemManualRemediationTypeFromValue returns a pointer to a valid SecurityAccessReviewAccessItemManualRemediationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityAccessReviewAccessItemManualRemediationTypeFromValue(v string) (*SecurityAccessReviewAccessItemManualRemediationType, error) {
	ev := SecurityAccessReviewAccessItemManualRemediationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityAccessReviewAccessItemManualRemediationType: valid values are %v", v, AllowedSecurityAccessReviewAccessItemManualRemediationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityAccessReviewAccessItemManualRemediationType) IsValid() bool {
	for _, existing := range AllowedSecurityAccessReviewAccessItemManualRemediationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to security-access-review-access-item-manual-remediation-type value
func (v SecurityAccessReviewAccessItemManualRemediationType) Ptr() *SecurityAccessReviewAccessItemManualRemediationType {
	return &v
}

type NullableSecurityAccessReviewAccessItemManualRemediationType struct {
	value *SecurityAccessReviewAccessItemManualRemediationType
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItemManualRemediationType) Get() *SecurityAccessReviewAccessItemManualRemediationType {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItemManualRemediationType) Set(val *SecurityAccessReviewAccessItemManualRemediationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItemManualRemediationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItemManualRemediationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItemManualRemediationType(val *SecurityAccessReviewAccessItemManualRemediationType) *NullableSecurityAccessReviewAccessItemManualRemediationType {
	return &NullableSecurityAccessReviewAccessItemManualRemediationType{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItemManualRemediationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItemManualRemediationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
