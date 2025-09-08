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

// model_oneof.mustache
// Field - The field to use when prompting the user
type Field struct {
	FieldDate   *FieldDate
	FieldSelect *FieldSelect
	FieldText   *FieldText
}

// FieldDateAsField is a convenience function that returns FieldDate wrapped in Field
func FieldDateAsField(v *FieldDate) Field {
	return Field{
		FieldDate: v,
	}
}

// FieldSelectAsField is a convenience function that returns FieldSelect wrapped in Field
func FieldSelectAsField(v *FieldSelect) Field {
	return Field{
		FieldSelect: v,
	}
}

// FieldTextAsField is a convenience function that returns FieldText wrapped in Field
func FieldTextAsField(v *FieldText) Field {
	return Field{
		FieldText: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *Field) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'DATE-TIME'
	if jsonDict["type"] == "DATE-TIME" {
		// try to unmarshal JSON data into FieldDate
		err = json.Unmarshal(data, &dst.FieldDate)
		if err == nil {
			return nil // data stored in dst.FieldDate, return on the first match
		} else {
			dst.FieldDate = nil
			return fmt.Errorf("Failed to unmarshal Field as FieldDate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SELECT'
	if jsonDict["type"] == "SELECT" {
		// try to unmarshal JSON data into FieldSelect
		err = json.Unmarshal(data, &dst.FieldSelect)
		if err == nil {
			return nil // data stored in dst.FieldSelect, return on the first match
		} else {
			dst.FieldSelect = nil
			return fmt.Errorf("Failed to unmarshal Field as FieldSelect: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TEXT'
	if jsonDict["type"] == "TEXT" {
		// try to unmarshal JSON data into FieldText
		err = json.Unmarshal(data, &dst.FieldText)
		if err == nil {
			return nil // data stored in dst.FieldText, return on the first match
		} else {
			dst.FieldText = nil
			return fmt.Errorf("Failed to unmarshal Field as FieldText: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-date'
	if jsonDict["type"] == "field-date" {
		// try to unmarshal JSON data into FieldDate
		err = json.Unmarshal(data, &dst.FieldDate)
		if err == nil {
			return nil // data stored in dst.FieldDate, return on the first match
		} else {
			dst.FieldDate = nil
			return fmt.Errorf("Failed to unmarshal Field as FieldDate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-select'
	if jsonDict["type"] == "field-select" {
		// try to unmarshal JSON data into FieldSelect
		err = json.Unmarshal(data, &dst.FieldSelect)
		if err == nil {
			return nil // data stored in dst.FieldSelect, return on the first match
		} else {
			dst.FieldSelect = nil
			return fmt.Errorf("Failed to unmarshal Field as FieldSelect: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-text'
	if jsonDict["type"] == "field-text" {
		// try to unmarshal JSON data into FieldText
		err = json.Unmarshal(data, &dst.FieldText)
		if err == nil {
			return nil // data stored in dst.FieldText, return on the first match
		} else {
			dst.FieldText = nil
			return fmt.Errorf("Failed to unmarshal Field as FieldText: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Field) MarshalJSON() ([]byte, error) {
	if src.FieldDate != nil {
		return json.Marshal(&src.FieldDate)
	}

	if src.FieldSelect != nil {
		return json.Marshal(&src.FieldSelect)
	}

	if src.FieldText != nil {
		return json.Marshal(&src.FieldText)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Field) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FieldDate != nil {
		return obj.FieldDate
	}

	if obj.FieldSelect != nil {
		return obj.FieldSelect
	}

	if obj.FieldText != nil {
		return obj.FieldText
	}

	// all schemas are nil
	return nil
}

type NullableField struct {
	value *Field
	isSet bool
}

func (v NullableField) Get() *Field {
	return v.value
}

func (v *NullableField) Set(val *Field) {
	v.value = val
	v.isSet = true
}

func (v NullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableField(val *Field) *NullableField {
	return &NullableField{value: val, isSet: true}
}

func (v NullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
