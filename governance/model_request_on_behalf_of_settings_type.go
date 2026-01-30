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

// RequestOnBehalfOfSettingsType Resource that can request access on behalf of others
type RequestOnBehalfOfSettingsType string

// List of request-on-behalf-of-settings-type
const (
	REQUESTONBEHALFOFSETTINGSTYPE_EVERYONE      RequestOnBehalfOfSettingsType = "EVERYONE"
	REQUESTONBEHALFOFSETTINGSTYPE_DIRECT_REPORT RequestOnBehalfOfSettingsType = "DIRECT_REPORT"
)

// All allowed values of RequestOnBehalfOfSettingsType enum
var AllowedRequestOnBehalfOfSettingsTypeEnumValues = []RequestOnBehalfOfSettingsType{
	"EVERYONE",
	"DIRECT_REPORT",
}

func (v *RequestOnBehalfOfSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestOnBehalfOfSettingsType(value)
	for _, existing := range AllowedRequestOnBehalfOfSettingsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestOnBehalfOfSettingsType", value)
}

// NewRequestOnBehalfOfSettingsTypeFromValue returns a pointer to a valid RequestOnBehalfOfSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestOnBehalfOfSettingsTypeFromValue(v string) (*RequestOnBehalfOfSettingsType, error) {
	ev := RequestOnBehalfOfSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestOnBehalfOfSettingsType: valid values are %v", v, AllowedRequestOnBehalfOfSettingsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestOnBehalfOfSettingsType) IsValid() bool {
	for _, existing := range AllowedRequestOnBehalfOfSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-on-behalf-of-settings-type value
func (v RequestOnBehalfOfSettingsType) Ptr() *RequestOnBehalfOfSettingsType {
	return &v
}

type NullableRequestOnBehalfOfSettingsType struct {
	value *RequestOnBehalfOfSettingsType
	isSet bool
}

func (v NullableRequestOnBehalfOfSettingsType) Get() *RequestOnBehalfOfSettingsType {
	return v.value
}

func (v *NullableRequestOnBehalfOfSettingsType) Set(val *RequestOnBehalfOfSettingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestOnBehalfOfSettingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestOnBehalfOfSettingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestOnBehalfOfSettingsType(val *RequestOnBehalfOfSettingsType) *NullableRequestOnBehalfOfSettingsType {
	return &NullableRequestOnBehalfOfSettingsType{value: val, isSet: true}
}

func (v NullableRequestOnBehalfOfSettingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestOnBehalfOfSettingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
