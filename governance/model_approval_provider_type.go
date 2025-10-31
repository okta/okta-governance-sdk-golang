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

// ApprovalProviderType The approval system type
type ApprovalProviderType string

// List of approval-provider-type
const (
	APPROVALPROVIDERTYPE_OKTA     ApprovalProviderType = "OKTA"
	APPROVALPROVIDERTYPE_EXTERNAL ApprovalProviderType = "EXTERNAL"
)

// All allowed values of ApprovalProviderType enum
var AllowedApprovalProviderTypeEnumValues = []ApprovalProviderType{
	"OKTA",
	"EXTERNAL",
}

func (v *ApprovalProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovalProviderType(value)
	for _, existing := range AllowedApprovalProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApprovalProviderType", value)
}

// NewApprovalProviderTypeFromValue returns a pointer to a valid ApprovalProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovalProviderTypeFromValue(v string) (*ApprovalProviderType, error) {
	ev := ApprovalProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovalProviderType: valid values are %v", v, AllowedApprovalProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovalProviderType) IsValid() bool {
	for _, existing := range AllowedApprovalProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to approval-provider-type value
func (v ApprovalProviderType) Ptr() *ApprovalProviderType {
	return &v
}

type NullableApprovalProviderType struct {
	value *ApprovalProviderType
	isSet bool
}

func (v NullableApprovalProviderType) Get() *ApprovalProviderType {
	return v.value
}

func (v *NullableApprovalProviderType) Set(val *ApprovalProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalProviderType(val *ApprovalProviderType) *NullableApprovalProviderType {
	return &NullableApprovalProviderType{value: val, isSet: true}
}

func (v NullableApprovalProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
