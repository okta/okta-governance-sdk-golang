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

// GrantMethod Type of grant assignment method
type GrantMethod string

// List of grantMethod
const (
	GRANTMETHOD_POLICY         GrantMethod = "POLICY"
	GRANTMETHOD_ACCESS_REQUEST GrantMethod = "ACCESS_REQUEST"
	GRANTMETHOD_ADMIN          GrantMethod = "ADMIN"
	GRANTMETHOD_API            GrantMethod = "API"
	GRANTMETHOD_NONE           GrantMethod = "NONE"
)

// All allowed values of GrantMethod enum
var AllowedGrantMethodEnumValues = []GrantMethod{
	"POLICY",
	"ACCESS_REQUEST",
	"ADMIN",
	"API",
	"NONE",
}

func (v *GrantMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantMethod(value)
	for _, existing := range AllowedGrantMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantMethod", value)
}

// NewGrantMethodFromValue returns a pointer to a valid GrantMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantMethodFromValue(v string) (*GrantMethod, error) {
	ev := GrantMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantMethod: valid values are %v", v, AllowedGrantMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantMethod) IsValid() bool {
	for _, existing := range AllowedGrantMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to grantMethod value
func (v GrantMethod) Ptr() *GrantMethod {
	return &v
}

type NullableGrantMethod struct {
	value *GrantMethod
	isSet bool
}

func (v NullableGrantMethod) Get() *GrantMethod {
	return v.value
}

func (v *NullableGrantMethod) Set(val *GrantMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantMethod(val *GrantMethod) *NullableGrantMethod {
	return &NullableGrantMethod{value: val, isSet: true}
}

func (v NullableGrantMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
