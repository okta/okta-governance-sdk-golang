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

// PrincipalProfileType The type of principal
type PrincipalProfileType string

// List of principal-profile-type
const (
	PRINCIPALPROFILETYPE_USER       PrincipalProfileType = "USER"
	PRINCIPALPROFILETYPE_CLIENT_APP PrincipalProfileType = "CLIENT_APP"
	PRINCIPALPROFILETYPE_UNKNOWN    PrincipalProfileType = "UNKNOWN"
)

// All allowed values of PrincipalProfileType enum
var AllowedPrincipalProfileTypeEnumValues = []PrincipalProfileType{
	"USER",
	"CLIENT_APP",
	"UNKNOWN",
}

func (v *PrincipalProfileType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrincipalProfileType(value)
	for _, existing := range AllowedPrincipalProfileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrincipalProfileType", value)
}

// NewPrincipalProfileTypeFromValue returns a pointer to a valid PrincipalProfileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrincipalProfileTypeFromValue(v string) (*PrincipalProfileType, error) {
	ev := PrincipalProfileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrincipalProfileType: valid values are %v", v, AllowedPrincipalProfileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrincipalProfileType) IsValid() bool {
	for _, existing := range AllowedPrincipalProfileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to principal-profile-type value
func (v PrincipalProfileType) Ptr() *PrincipalProfileType {
	return &v
}

type NullablePrincipalProfileType struct {
	value *PrincipalProfileType
	isSet bool
}

func (v NullablePrincipalProfileType) Get() *PrincipalProfileType {
	return v.value
}

func (v *NullablePrincipalProfileType) Set(val *PrincipalProfileType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalProfileType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalProfileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalProfileType(val *PrincipalProfileType) *NullablePrincipalProfileType {
	return &NullablePrincipalProfileType{value: val, isSet: true}
}

func (v NullablePrincipalProfileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalProfileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
