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
// FieldValue - The field value provided by the user
type FieldValue struct {
	FieldValueDate   *FieldValueDate
	FieldValueSelect *FieldValueSelect
	FieldValueText   *FieldValueText
}

// FieldValueDateAsFieldValue is a convenience function that returns FieldValueDate wrapped in FieldValue
func FieldValueDateAsFieldValue(v *FieldValueDate) FieldValue {
	return FieldValue{
		FieldValueDate: v,
	}
}

// FieldValueSelectAsFieldValue is a convenience function that returns FieldValueSelect wrapped in FieldValue
func FieldValueSelectAsFieldValue(v *FieldValueSelect) FieldValue {
	return FieldValue{
		FieldValueSelect: v,
	}
}

// FieldValueTextAsFieldValue is a convenience function that returns FieldValueText wrapped in FieldValue
func FieldValueTextAsFieldValue(v *FieldValueText) FieldValue {
	return FieldValue{
		FieldValueText: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *FieldValue) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'DATE-TIME'
	if jsonDict["type"] == "DATE-TIME" {
		// try to unmarshal JSON data into FieldValueDate
		err = json.Unmarshal(data, &dst.FieldValueDate)
		if err == nil {
			return nil // data stored in dst.FieldValueDate, return on the first match
		} else {
			dst.FieldValueDate = nil
			return fmt.Errorf("Failed to unmarshal FieldValue as FieldValueDate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SELECT'
	if jsonDict["type"] == "SELECT" {
		// try to unmarshal JSON data into FieldValueSelect
		err = json.Unmarshal(data, &dst.FieldValueSelect)
		if err == nil {
			return nil // data stored in dst.FieldValueSelect, return on the first match
		} else {
			dst.FieldValueSelect = nil
			return fmt.Errorf("Failed to unmarshal FieldValue as FieldValueSelect: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TEXT'
	if jsonDict["type"] == "TEXT" {
		// try to unmarshal JSON data into FieldValueText
		err = json.Unmarshal(data, &dst.FieldValueText)
		if err == nil {
			return nil // data stored in dst.FieldValueText, return on the first match
		} else {
			dst.FieldValueText = nil
			return fmt.Errorf("Failed to unmarshal FieldValue as FieldValueText: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-value-date'
	if jsonDict["type"] == "field-value-date" {
		// try to unmarshal JSON data into FieldValueDate
		err = json.Unmarshal(data, &dst.FieldValueDate)
		if err == nil {
			return nil // data stored in dst.FieldValueDate, return on the first match
		} else {
			dst.FieldValueDate = nil
			return fmt.Errorf("Failed to unmarshal FieldValue as FieldValueDate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-value-select'
	if jsonDict["type"] == "field-value-select" {
		// try to unmarshal JSON data into FieldValueSelect
		err = json.Unmarshal(data, &dst.FieldValueSelect)
		if err == nil {
			return nil // data stored in dst.FieldValueSelect, return on the first match
		} else {
			dst.FieldValueSelect = nil
			return fmt.Errorf("Failed to unmarshal FieldValue as FieldValueSelect: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-value-text'
	if jsonDict["type"] == "field-value-text" {
		// try to unmarshal JSON data into FieldValueText
		err = json.Unmarshal(data, &dst.FieldValueText)
		if err == nil {
			return nil // data stored in dst.FieldValueText, return on the first match
		} else {
			dst.FieldValueText = nil
			return fmt.Errorf("Failed to unmarshal FieldValue as FieldValueText: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FieldValue) MarshalJSON() ([]byte, error) {
	if src.FieldValueDate != nil {
		return json.Marshal(&src.FieldValueDate)
	}

	if src.FieldValueSelect != nil {
		return json.Marshal(&src.FieldValueSelect)
	}

	if src.FieldValueText != nil {
		return json.Marshal(&src.FieldValueText)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FieldValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FieldValueDate != nil {
		return obj.FieldValueDate
	}

	if obj.FieldValueSelect != nil {
		return obj.FieldValueSelect
	}

	if obj.FieldValueText != nil {
		return obj.FieldValueText
	}

	// all schemas are nil
	return nil
}

type NullableFieldValue struct {
	value *FieldValue
	isSet bool
}

func (v NullableFieldValue) Get() *FieldValue {
	return v.value
}

func (v *NullableFieldValue) Set(val *FieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldValue(val *FieldValue) *NullableFieldValue {
	return &NullableFieldValue{value: val, isSet: true}
}

func (v NullableFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
