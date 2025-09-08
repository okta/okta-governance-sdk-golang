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

// RequestOnBehalfOfType the model 'RequestOnBehalfOfType'
type RequestOnBehalfOfType string

// List of request-on-behalf-of-type
const (
	REQUESTONBEHALFOFTYPE_DIRECT_REPORT RequestOnBehalfOfType = "DIRECT_REPORT"
)

// All allowed values of RequestOnBehalfOfType enum
var AllowedRequestOnBehalfOfTypeEnumValues = []RequestOnBehalfOfType{
	"DIRECT_REPORT",
}

func (v *RequestOnBehalfOfType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestOnBehalfOfType(value)
	for _, existing := range AllowedRequestOnBehalfOfTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestOnBehalfOfType", value)
}

// NewRequestOnBehalfOfTypeFromValue returns a pointer to a valid RequestOnBehalfOfType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestOnBehalfOfTypeFromValue(v string) (*RequestOnBehalfOfType, error) {
	ev := RequestOnBehalfOfType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestOnBehalfOfType: valid values are %v", v, AllowedRequestOnBehalfOfTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestOnBehalfOfType) IsValid() bool {
	for _, existing := range AllowedRequestOnBehalfOfTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-on-behalf-of-type value
func (v RequestOnBehalfOfType) Ptr() *RequestOnBehalfOfType {
	return &v
}

type NullableRequestOnBehalfOfType struct {
	value *RequestOnBehalfOfType
	isSet bool
}

func (v NullableRequestOnBehalfOfType) Get() *RequestOnBehalfOfType {
	return v.value
}

func (v *NullableRequestOnBehalfOfType) Set(val *RequestOnBehalfOfType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestOnBehalfOfType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestOnBehalfOfType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestOnBehalfOfType(val *RequestOnBehalfOfType) *NullableRequestOnBehalfOfType {
	return &NullableRequestOnBehalfOfType{value: val, isSet: true}
}

func (v NullableRequestOnBehalfOfType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestOnBehalfOfType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
