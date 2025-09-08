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

// RequestFieldType Type of value for the requester field.
type RequestFieldType string

// List of request-field-type
const (
	REQUESTFIELDTYPE_TEXT         RequestFieldType = "TEXT"
	REQUESTFIELDTYPE_SELECT       RequestFieldType = "SELECT"
	REQUESTFIELDTYPE_MULTISELECT  RequestFieldType = "MULTISELECT"
	REQUESTFIELDTYPE_ISO_DATE     RequestFieldType = "ISO_DATE"
	REQUESTFIELDTYPE_DURATION     RequestFieldType = "DURATION"
	REQUESTFIELDTYPE_OKTA_USER_ID RequestFieldType = "OKTA_USER_ID"
)

// All allowed values of RequestFieldType enum
var AllowedRequestFieldTypeEnumValues = []RequestFieldType{
	"TEXT",
	"SELECT",
	"MULTISELECT",
	"ISO_DATE",
	"DURATION",
	"OKTA_USER_ID",
}

func (v *RequestFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestFieldType(value)
	for _, existing := range AllowedRequestFieldTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestFieldType", value)
}

// NewRequestFieldTypeFromValue returns a pointer to a valid RequestFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestFieldTypeFromValue(v string) (*RequestFieldType, error) {
	ev := RequestFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestFieldType: valid values are %v", v, AllowedRequestFieldTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestFieldType) IsValid() bool {
	for _, existing := range AllowedRequestFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-field-type value
func (v RequestFieldType) Ptr() *RequestFieldType {
	return &v
}

type NullableRequestFieldType struct {
	value *RequestFieldType
	isSet bool
}

func (v NullableRequestFieldType) Get() *RequestFieldType {
	return v.value
}

func (v *NullableRequestFieldType) Set(val *RequestFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFieldType(val *RequestFieldType) *NullableRequestFieldType {
	return &NullableRequestFieldType{value: val, isSet: true}
}

func (v NullableRequestFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
