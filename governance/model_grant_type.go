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

// GrantType Type of grant
type GrantType string

// List of grant-type
const (
	GRANTTYPE_CUSTOM             GrantType = "CUSTOM"
	GRANTTYPE_ENTITLEMENT_BUNDLE GrantType = "ENTITLEMENT-BUNDLE"
	GRANTTYPE_POLICY             GrantType = "POLICY"
	GRANTTYPE_ENTITLEMENT        GrantType = "ENTITLEMENT"
)

// All allowed values of GrantType enum
var AllowedGrantTypeEnumValues = []GrantType{
	"CUSTOM",
	"ENTITLEMENT-BUNDLE",
	"POLICY",
	"ENTITLEMENT",
}

func (v *GrantType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantType(value)
	for _, existing := range AllowedGrantTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantType", value)
}

// NewGrantTypeFromValue returns a pointer to a valid GrantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantTypeFromValue(v string) (*GrantType, error) {
	ev := GrantType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantType: valid values are %v", v, AllowedGrantTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantType) IsValid() bool {
	for _, existing := range AllowedGrantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to grant-type value
func (v GrantType) Ptr() *GrantType {
	return &v
}

type NullableGrantType struct {
	value *GrantType
	isSet bool
}

func (v NullableGrantType) Get() *GrantType {
	return v.value
}

func (v *NullableGrantType) Set(val *GrantType) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantType) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantType(val *GrantType) *NullableGrantType {
	return &NullableGrantType{value: val, isSet: true}
}

func (v NullableGrantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
