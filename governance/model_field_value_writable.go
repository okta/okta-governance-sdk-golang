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
)

// FieldValueWritable An `id` and `value` corresponding to the `id` of a field in the request type's `requestSettings.requesterFields`.  If the `id` corresponds to a `DATE-TIME` field, the value must an `ISO-8601` date-time. If the `id` corresponds to a `SELECT` field, the value must be an array of strings, where each string is a value from the select field's options. If the `id` corresponds to a `TEXT` field, the value must be a string.  A field `id` must only be provided once, duplicates will result in a `409 Conflict` response.  A non-required field may be omitted from the list.
type FieldValueWritable struct {
	// The `id` of a `requesterField` in the related request type's `requestSettings.requesterFields`.
	Id                   string                          `json:"id"`
	Value                FieldValueWritableAllowedValues `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FieldValueWritable FieldValueWritable

// NewFieldValueWritable instantiates a new FieldValueWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldValueWritable(id string, value FieldValueWritableAllowedValues) *FieldValueWritable {
	this := FieldValueWritable{}
	this.Id = id
	this.Value = value
	return &this
}

// NewFieldValueWritableWithDefaults instantiates a new FieldValueWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldValueWritableWithDefaults() *FieldValueWritable {
	this := FieldValueWritable{}
	return &this
}

// GetId returns the Id field value
func (o *FieldValueWritable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FieldValueWritable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FieldValueWritable) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *FieldValueWritable) GetValue() FieldValueWritableAllowedValues {
	if o == nil {
		var ret FieldValueWritableAllowedValues
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FieldValueWritable) GetValueOk() (*FieldValueWritableAllowedValues, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FieldValueWritable) SetValue(v FieldValueWritableAllowedValues) {
	o.Value = v
}

func (o FieldValueWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FieldValueWritable) UnmarshalJSON(bytes []byte) (err error) {
	varFieldValueWritable := _FieldValueWritable{}

	err = json.Unmarshal(bytes, &varFieldValueWritable)
	if err == nil {
		*o = FieldValueWritable(varFieldValueWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFieldValueWritable struct {
	value *FieldValueWritable
	isSet bool
}

func (v NullableFieldValueWritable) Get() *FieldValueWritable {
	return v.value
}

func (v *NullableFieldValueWritable) Set(val *FieldValueWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldValueWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldValueWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldValueWritable(val *FieldValueWritable) *NullableFieldValueWritable {
	return &NullableFieldValueWritable{value: val, isSet: true}
}

func (v NullableFieldValueWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldValueWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
