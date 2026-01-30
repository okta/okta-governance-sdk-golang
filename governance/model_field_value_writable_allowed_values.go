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
	"time"
)

// FieldValueWritableAllowedValues struct for FieldValueWritableAllowedValues
type FieldValueWritableAllowedValues struct {
	ArrayOfString *[]string
	String        *string
	TimeTime      *time.Time
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FieldValueWritableAllowedValues) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ArrayOfString
	err = json.Unmarshal(data, &dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	// try to unmarshal JSON data into TimeTime
	err = json.Unmarshal(data, &dst.TimeTime)
	if err == nil {
		jsonTimeTime, _ := json.Marshal(dst.TimeTime)
		if string(jsonTimeTime) == "{}" { // empty struct
			dst.TimeTime = nil
		} else {
			return nil // data stored in dst.TimeTime, return on the first match
		}
	} else {
		dst.TimeTime = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(FieldValueWritableAllowedValues)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FieldValueWritableAllowedValues) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	if src.TimeTime != nil {
		return json.Marshal(&src.TimeTime)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFieldValueWritableAllowedValues struct {
	value *FieldValueWritableAllowedValues
	isSet bool
}

func (v NullableFieldValueWritableAllowedValues) Get() *FieldValueWritableAllowedValues {
	return v.value
}

func (v *NullableFieldValueWritableAllowedValues) Set(val *FieldValueWritableAllowedValues) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldValueWritableAllowedValues) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldValueWritableAllowedValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldValueWritableAllowedValues(val *FieldValueWritableAllowedValues) *NullableFieldValueWritableAllowedValues {
	return &NullableFieldValueWritableAllowedValues{value: val, isSet: true}
}

func (v NullableFieldValueWritableAllowedValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldValueWritableAllowedValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
