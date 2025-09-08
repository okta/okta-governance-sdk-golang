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
// FieldWritable - The field to use when prompting the user
type FieldWritable struct {
	FieldDateWritable   *FieldDateWritable
	FieldSelectWritable *FieldSelectWritable
	FieldTextWritable   *FieldTextWritable
}

// FieldDateWritableAsFieldWritable is a convenience function that returns FieldDateWritable wrapped in FieldWritable
func FieldDateWritableAsFieldWritable(v *FieldDateWritable) FieldWritable {
	return FieldWritable{
		FieldDateWritable: v,
	}
}

// FieldSelectWritableAsFieldWritable is a convenience function that returns FieldSelectWritable wrapped in FieldWritable
func FieldSelectWritableAsFieldWritable(v *FieldSelectWritable) FieldWritable {
	return FieldWritable{
		FieldSelectWritable: v,
	}
}

// FieldTextWritableAsFieldWritable is a convenience function that returns FieldTextWritable wrapped in FieldWritable
func FieldTextWritableAsFieldWritable(v *FieldTextWritable) FieldWritable {
	return FieldWritable{
		FieldTextWritable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *FieldWritable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'DATE-TIME'
	if jsonDict["type"] == "DATE-TIME" {
		// try to unmarshal JSON data into FieldDateWritable
		err = json.Unmarshal(data, &dst.FieldDateWritable)
		if err == nil {
			return nil // data stored in dst.FieldDateWritable, return on the first match
		} else {
			dst.FieldDateWritable = nil
			return fmt.Errorf("Failed to unmarshal FieldWritable as FieldDateWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SELECT'
	if jsonDict["type"] == "SELECT" {
		// try to unmarshal JSON data into FieldSelectWritable
		err = json.Unmarshal(data, &dst.FieldSelectWritable)
		if err == nil {
			return nil // data stored in dst.FieldSelectWritable, return on the first match
		} else {
			dst.FieldSelectWritable = nil
			return fmt.Errorf("Failed to unmarshal FieldWritable as FieldSelectWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TEXT'
	if jsonDict["type"] == "TEXT" {
		// try to unmarshal JSON data into FieldTextWritable
		err = json.Unmarshal(data, &dst.FieldTextWritable)
		if err == nil {
			return nil // data stored in dst.FieldTextWritable, return on the first match
		} else {
			dst.FieldTextWritable = nil
			return fmt.Errorf("Failed to unmarshal FieldWritable as FieldTextWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-date-writable'
	if jsonDict["type"] == "field-date-writable" {
		// try to unmarshal JSON data into FieldDateWritable
		err = json.Unmarshal(data, &dst.FieldDateWritable)
		if err == nil {
			return nil // data stored in dst.FieldDateWritable, return on the first match
		} else {
			dst.FieldDateWritable = nil
			return fmt.Errorf("Failed to unmarshal FieldWritable as FieldDateWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-select-writable'
	if jsonDict["type"] == "field-select-writable" {
		// try to unmarshal JSON data into FieldSelectWritable
		err = json.Unmarshal(data, &dst.FieldSelectWritable)
		if err == nil {
			return nil // data stored in dst.FieldSelectWritable, return on the first match
		} else {
			dst.FieldSelectWritable = nil
			return fmt.Errorf("Failed to unmarshal FieldWritable as FieldSelectWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'field-text-writable'
	if jsonDict["type"] == "field-text-writable" {
		// try to unmarshal JSON data into FieldTextWritable
		err = json.Unmarshal(data, &dst.FieldTextWritable)
		if err == nil {
			return nil // data stored in dst.FieldTextWritable, return on the first match
		} else {
			dst.FieldTextWritable = nil
			return fmt.Errorf("Failed to unmarshal FieldWritable as FieldTextWritable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FieldWritable) MarshalJSON() ([]byte, error) {
	if src.FieldDateWritable != nil {
		return json.Marshal(&src.FieldDateWritable)
	}

	if src.FieldSelectWritable != nil {
		return json.Marshal(&src.FieldSelectWritable)
	}

	if src.FieldTextWritable != nil {
		return json.Marshal(&src.FieldTextWritable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FieldWritable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FieldDateWritable != nil {
		return obj.FieldDateWritable
	}

	if obj.FieldSelectWritable != nil {
		return obj.FieldSelectWritable
	}

	if obj.FieldTextWritable != nil {
		return obj.FieldTextWritable
	}

	// all schemas are nil
	return nil
}

type NullableFieldWritable struct {
	value *FieldWritable
	isSet bool
}

func (v NullableFieldWritable) Get() *FieldWritable {
	return v.value
}

func (v *NullableFieldWritable) Set(val *FieldWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldWritable(val *FieldWritable) *NullableFieldWritable {
	return &NullableFieldWritable{value: val, isSet: true}
}

func (v NullableFieldWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
