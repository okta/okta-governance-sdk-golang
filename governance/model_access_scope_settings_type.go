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

// AccessScopeSettingsType the model 'AccessScopeSettingsType'
type AccessScopeSettingsType string

// List of access-scope-settings-type
const (
	ACCESSSCOPESETTINGSTYPE_RESOURCE_DEFAULT    AccessScopeSettingsType = "RESOURCE_DEFAULT"
	ACCESSSCOPESETTINGSTYPE_ENTITLEMENT_BUNDLES AccessScopeSettingsType = "ENTITLEMENT_BUNDLES"
	ACCESSSCOPESETTINGSTYPE_GROUPS              AccessScopeSettingsType = "GROUPS"
)

// All allowed values of AccessScopeSettingsType enum
var AllowedAccessScopeSettingsTypeEnumValues = []AccessScopeSettingsType{
	"RESOURCE_DEFAULT",
	"ENTITLEMENT_BUNDLES",
	"GROUPS",
}

func (v *AccessScopeSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessScopeSettingsType(value)
	for _, existing := range AllowedAccessScopeSettingsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessScopeSettingsType", value)
}

// NewAccessScopeSettingsTypeFromValue returns a pointer to a valid AccessScopeSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessScopeSettingsTypeFromValue(v string) (*AccessScopeSettingsType, error) {
	ev := AccessScopeSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessScopeSettingsType: valid values are %v", v, AllowedAccessScopeSettingsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessScopeSettingsType) IsValid() bool {
	for _, existing := range AllowedAccessScopeSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to access-scope-settings-type value
func (v AccessScopeSettingsType) Ptr() *AccessScopeSettingsType {
	return &v
}

type NullableAccessScopeSettingsType struct {
	value *AccessScopeSettingsType
	isSet bool
}

func (v NullableAccessScopeSettingsType) Get() *AccessScopeSettingsType {
	return v.value
}

func (v *NullableAccessScopeSettingsType) Set(val *AccessScopeSettingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeSettingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeSettingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeSettingsType(val *AccessScopeSettingsType) *NullableAccessScopeSettingsType {
	return &NullableAccessScopeSettingsType{value: val, isSet: true}
}

func (v NullableAccessScopeSettingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeSettingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
