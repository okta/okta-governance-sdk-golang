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

// AuthorizationServerScopeCondition Scope condition defined for the authorization server
type AuthorizationServerScopeCondition string

// List of authorization-server-scope-condition
const (
	AUTHORIZATIONSERVERSCOPECONDITION_ALL_SCOPES   AuthorizationServerScopeCondition = "ALL_SCOPES"
	AUTHORIZATIONSERVERSCOPECONDITION_INCLUDE_ONLY AuthorizationServerScopeCondition = "INCLUDE_ONLY"
	AUTHORIZATIONSERVERSCOPECONDITION_EXCLUDE      AuthorizationServerScopeCondition = "EXCLUDE"
)

// All allowed values of AuthorizationServerScopeCondition enum
var AllowedAuthorizationServerScopeConditionEnumValues = []AuthorizationServerScopeCondition{
	"ALL_SCOPES",
	"INCLUDE_ONLY",
	"EXCLUDE",
}

func (v *AuthorizationServerScopeCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthorizationServerScopeCondition(value)
	for _, existing := range AllowedAuthorizationServerScopeConditionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthorizationServerScopeCondition", value)
}

// NewAuthorizationServerScopeConditionFromValue returns a pointer to a valid AuthorizationServerScopeCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthorizationServerScopeConditionFromValue(v string) (*AuthorizationServerScopeCondition, error) {
	ev := AuthorizationServerScopeCondition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthorizationServerScopeCondition: valid values are %v", v, AllowedAuthorizationServerScopeConditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthorizationServerScopeCondition) IsValid() bool {
	for _, existing := range AllowedAuthorizationServerScopeConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to authorization-server-scope-condition value
func (v AuthorizationServerScopeCondition) Ptr() *AuthorizationServerScopeCondition {
	return &v
}

type NullableAuthorizationServerScopeCondition struct {
	value *AuthorizationServerScopeCondition
	isSet bool
}

func (v NullableAuthorizationServerScopeCondition) Get() *AuthorizationServerScopeCondition {
	return v.value
}

func (v *NullableAuthorizationServerScopeCondition) Set(val *AuthorizationServerScopeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerScopeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerScopeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerScopeCondition(val *AuthorizationServerScopeCondition) *NullableAuthorizationServerScopeCondition {
	return &NullableAuthorizationServerScopeCondition{value: val, isSet: true}
}

func (v NullableAuthorizationServerScopeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerScopeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
