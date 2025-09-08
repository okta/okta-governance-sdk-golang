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

// FieldSelectType the model 'FieldSelectType'
type FieldSelectType string

// List of field-select-type
const (
	FIELDSELECTTYPE_SELECT FieldSelectType = "SELECT"
)

// All allowed values of FieldSelectType enum
var AllowedFieldSelectTypeEnumValues = []FieldSelectType{
	"SELECT",
}

func (v *FieldSelectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldSelectType(value)
	for _, existing := range AllowedFieldSelectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldSelectType", value)
}

// NewFieldSelectTypeFromValue returns a pointer to a valid FieldSelectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldSelectTypeFromValue(v string) (*FieldSelectType, error) {
	ev := FieldSelectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldSelectType: valid values are %v", v, AllowedFieldSelectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldSelectType) IsValid() bool {
	for _, existing := range AllowedFieldSelectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to field-select-type value
func (v FieldSelectType) Ptr() *FieldSelectType {
	return &v
}

type NullableFieldSelectType struct {
	value *FieldSelectType
	isSet bool
}

func (v NullableFieldSelectType) Get() *FieldSelectType {
	return v.value
}

func (v *NullableFieldSelectType) Set(val *FieldSelectType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldSelectType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldSelectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldSelectType(val *FieldSelectType) *NullableFieldSelectType {
	return &NullableFieldSelectType{value: val, isSet: true}
}

func (v NullableFieldSelectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldSelectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
