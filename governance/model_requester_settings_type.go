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

// RequesterSettingsType the model 'RequesterSettingsType'
type RequesterSettingsType string

// List of requester-settings-type
const (
	REQUESTERSETTINGSTYPE_EVERYONE RequesterSettingsType = "EVERYONE"
	REQUESTERSETTINGSTYPE_GROUPS   RequesterSettingsType = "GROUPS"
	REQUESTERSETTINGSTYPE_TEAMS    RequesterSettingsType = "TEAMS"
)

// All allowed values of RequesterSettingsType enum
var AllowedRequesterSettingsTypeEnumValues = []RequesterSettingsType{
	"EVERYONE",
	"GROUPS",
	"TEAMS",
}

func (v *RequesterSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequesterSettingsType(value)
	for _, existing := range AllowedRequesterSettingsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequesterSettingsType", value)
}

// NewRequesterSettingsTypeFromValue returns a pointer to a valid RequesterSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequesterSettingsTypeFromValue(v string) (*RequesterSettingsType, error) {
	ev := RequesterSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequesterSettingsType: valid values are %v", v, AllowedRequesterSettingsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequesterSettingsType) IsValid() bool {
	for _, existing := range AllowedRequesterSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to requester-settings-type value
func (v RequesterSettingsType) Ptr() *RequesterSettingsType {
	return &v
}

type NullableRequesterSettingsType struct {
	value *RequesterSettingsType
	isSet bool
}

func (v NullableRequesterSettingsType) Get() *RequesterSettingsType {
	return v.value
}

func (v *NullableRequesterSettingsType) Set(val *RequesterSettingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequesterSettingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequesterSettingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequesterSettingsType(val *RequesterSettingsType) *NullableRequesterSettingsType {
	return &NullableRequesterSettingsType{value: val, isSet: true}
}

func (v NullableRequesterSettingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequesterSettingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
