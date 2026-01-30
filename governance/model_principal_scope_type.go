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

// PrincipalScopeType the model 'PrincipalScopeType'
type PrincipalScopeType string

// List of principal-scope-type
const (
	PRINCIPALSCOPETYPE_USERS PrincipalScopeType = "USERS"
)

// All allowed values of PrincipalScopeType enum
var AllowedPrincipalScopeTypeEnumValues = []PrincipalScopeType{
	"USERS",
}

func (v *PrincipalScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrincipalScopeType(value)
	for _, existing := range AllowedPrincipalScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrincipalScopeType", value)
}

// NewPrincipalScopeTypeFromValue returns a pointer to a valid PrincipalScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrincipalScopeTypeFromValue(v string) (*PrincipalScopeType, error) {
	ev := PrincipalScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrincipalScopeType: valid values are %v", v, AllowedPrincipalScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrincipalScopeType) IsValid() bool {
	for _, existing := range AllowedPrincipalScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to principal-scope-type value
func (v PrincipalScopeType) Ptr() *PrincipalScopeType {
	return &v
}

type NullablePrincipalScopeType struct {
	value *PrincipalScopeType
	isSet bool
}

func (v NullablePrincipalScopeType) Get() *PrincipalScopeType {
	return v.value
}

func (v *NullablePrincipalScopeType) Set(val *PrincipalScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalScopeType(val *PrincipalScopeType) *NullablePrincipalScopeType {
	return &NullablePrincipalScopeType{value: val, isSet: true}
}

func (v NullablePrincipalScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
