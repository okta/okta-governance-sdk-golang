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

// FieldDateTimeType the model 'FieldDateTimeType'
type FieldDateTimeType string

// List of field-date-time-type
const (
	FIELDDATETIMETYPE_DATE_TIME FieldDateTimeType = "DATE-TIME"
)

// All allowed values of FieldDateTimeType enum
var AllowedFieldDateTimeTypeEnumValues = []FieldDateTimeType{
	"DATE-TIME",
}

func (v *FieldDateTimeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldDateTimeType(value)
	for _, existing := range AllowedFieldDateTimeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldDateTimeType", value)
}

// NewFieldDateTimeTypeFromValue returns a pointer to a valid FieldDateTimeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldDateTimeTypeFromValue(v string) (*FieldDateTimeType, error) {
	ev := FieldDateTimeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldDateTimeType: valid values are %v", v, AllowedFieldDateTimeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldDateTimeType) IsValid() bool {
	for _, existing := range AllowedFieldDateTimeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to field-date-time-type value
func (v FieldDateTimeType) Ptr() *FieldDateTimeType {
	return &v
}

type NullableFieldDateTimeType struct {
	value *FieldDateTimeType
	isSet bool
}

func (v NullableFieldDateTimeType) Get() *FieldDateTimeType {
	return v.value
}

func (v *NullableFieldDateTimeType) Set(val *FieldDateTimeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldDateTimeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldDateTimeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldDateTimeType(val *FieldDateTimeType) *NullableFieldDateTimeType {
	return &NullableFieldDateTimeType{value: val, isSet: true}
}

func (v NullableFieldDateTimeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldDateTimeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
