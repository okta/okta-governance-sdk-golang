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

// FieldOption An option to for the user to select for a field prompt.
type FieldOption struct {
	// The value of a select option
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FieldOption FieldOption

// NewFieldOption instantiates a new FieldOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldOption(value string) *FieldOption {
	this := FieldOption{}
	this.Value = value
	return &this
}

// NewFieldOptionWithDefaults instantiates a new FieldOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldOptionWithDefaults() *FieldOption {
	this := FieldOption{}
	return &this
}

// GetValue returns the Value field value
func (o *FieldOption) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FieldOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FieldOption) SetValue(v string) {
	o.Value = v
}

func (o FieldOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FieldOption) UnmarshalJSON(bytes []byte) (err error) {
	varFieldOption := _FieldOption{}

	err = json.Unmarshal(bytes, &varFieldOption)
	if err == nil {
		*o = FieldOption(varFieldOption)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFieldOption struct {
	value *FieldOption
	isSet bool
}

func (v NullableFieldOption) Get() *FieldOption {
	return v.value
}

func (v *NullableFieldOption) Set(val *FieldOption) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldOption) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldOption(val *FieldOption) *NullableFieldOption {
	return &NullableFieldOption{value: val, isSet: true}
}

func (v NullableFieldOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
