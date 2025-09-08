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

// AccessScopeType The access scope type
type AccessScopeType string

// List of access-scope-type
const (
	ACCESSSCOPETYPE_ENTITLEMENT_BUNDLE AccessScopeType = "ENTITLEMENT_BUNDLE"
	ACCESSSCOPETYPE_GROUP              AccessScopeType = "GROUP"
	ACCESSSCOPETYPE_APPLICATION        AccessScopeType = "APPLICATION"
)

// All allowed values of AccessScopeType enum
var AllowedAccessScopeTypeEnumValues = []AccessScopeType{
	"ENTITLEMENT_BUNDLE",
	"GROUP",
	"APPLICATION",
}

func (v *AccessScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessScopeType(value)
	for _, existing := range AllowedAccessScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessScopeType", value)
}

// NewAccessScopeTypeFromValue returns a pointer to a valid AccessScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessScopeTypeFromValue(v string) (*AccessScopeType, error) {
	ev := AccessScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessScopeType: valid values are %v", v, AllowedAccessScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessScopeType) IsValid() bool {
	for _, existing := range AllowedAccessScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to access-scope-type value
func (v AccessScopeType) Ptr() *AccessScopeType {
	return &v
}

type NullableAccessScopeType struct {
	value *AccessScopeType
	isSet bool
}

func (v NullableAccessScopeType) Get() *AccessScopeType {
	return v.value
}

func (v *NullableAccessScopeType) Set(val *AccessScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeType(val *AccessScopeType) *NullableAccessScopeType {
	return &NullableAccessScopeType{value: val, isSet: true}
}

func (v NullableAccessScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
