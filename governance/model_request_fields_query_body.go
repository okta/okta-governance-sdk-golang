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
)

// checks if the RequestFieldsQueryBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestFieldsQueryBody{}

// RequestFieldsQueryBody Request body for dynamic request field queries, providing previously selected field values.
type RequestFieldsQueryBody struct {
	// Previously selected field values, keyed by field ID.
	FieldValues          map[string]RequestFieldValuesMapValue `json:"fieldValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestFieldsQueryBody RequestFieldsQueryBody

// NewRequestFieldsQueryBody instantiates a new RequestFieldsQueryBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFieldsQueryBody() *RequestFieldsQueryBody {
	this := RequestFieldsQueryBody{}
	return &this
}

// NewRequestFieldsQueryBodyWithDefaults instantiates a new RequestFieldsQueryBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFieldsQueryBodyWithDefaults() *RequestFieldsQueryBody {
	this := RequestFieldsQueryBody{}
	return &this
}

// GetFieldValues returns the FieldValues field value if set, zero value otherwise.
func (o *RequestFieldsQueryBody) GetFieldValues() map[string]RequestFieldValuesMapValue {
	if o == nil || IsNil(o.FieldValues) {
		var ret map[string]RequestFieldValuesMapValue
		return ret
	}
	return o.FieldValues
}

// GetFieldValuesOk returns a tuple with the FieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldsQueryBody) GetFieldValuesOk() (map[string]RequestFieldValuesMapValue, bool) {
	if o == nil || IsNil(o.FieldValues) {
		return map[string]RequestFieldValuesMapValue{}, false
	}
	return o.FieldValues, true
}

// HasFieldValues returns a boolean if a field has been set.
func (o *RequestFieldsQueryBody) HasFieldValues() bool {
	if o != nil && !IsNil(o.FieldValues) {
		return true
	}

	return false
}

// SetFieldValues gets a reference to the given map[string]RequestFieldValuesMapValue and assigns it to the FieldValues field.
func (o *RequestFieldsQueryBody) SetFieldValues(v map[string]RequestFieldValuesMapValue) {
	o.FieldValues = v
}

func (o RequestFieldsQueryBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestFieldsQueryBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldValues) {
		toSerialize["fieldValues"] = o.FieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestFieldsQueryBody) UnmarshalJSON(data []byte) (err error) {
	varRequestFieldsQueryBody := _RequestFieldsQueryBody{}

	err = json.Unmarshal(data, &varRequestFieldsQueryBody)

	if err != nil {
		return err
	}

	*o = RequestFieldsQueryBody(varRequestFieldsQueryBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fieldValues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestFieldsQueryBody struct {
	value *RequestFieldsQueryBody
	isSet bool
}

func (v NullableRequestFieldsQueryBody) Get() *RequestFieldsQueryBody {
	return v.value
}

func (v *NullableRequestFieldsQueryBody) Set(val *RequestFieldsQueryBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFieldsQueryBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFieldsQueryBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFieldsQueryBody(val *RequestFieldsQueryBody) *NullableRequestFieldsQueryBody {
	return &NullableRequestFieldsQueryBody{value: val, isSet: true}
}

func (v NullableRequestFieldsQueryBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFieldsQueryBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
