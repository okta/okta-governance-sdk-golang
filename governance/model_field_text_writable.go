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

// FieldTextWritable A text field
type FieldTextWritable struct {
	Type FieldTextType `json:"type"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FieldTextWritable FieldTextWritable

// NewFieldTextWritable instantiates a new FieldTextWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldTextWritable(type_ FieldTextType, prompt string) *FieldTextWritable {
	this := FieldTextWritable{}
	this.Prompt = prompt
	var required bool = true
	this.Required = &required
	return &this
}

// NewFieldTextWritableWithDefaults instantiates a new FieldTextWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldTextWritableWithDefaults() *FieldTextWritable {
	this := FieldTextWritable{}
	var required bool = true
	this.Required = &required
	return &this
}

// GetType returns the Type field value
func (o *FieldTextWritable) GetType() FieldTextType {
	if o == nil {
		var ret FieldTextType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FieldTextWritable) GetTypeOk() (*FieldTextType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FieldTextWritable) SetType(v FieldTextType) {
	o.Type = v
}

// GetPrompt returns the Prompt field value
func (o *FieldTextWritable) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *FieldTextWritable) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *FieldTextWritable) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FieldTextWritable) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldTextWritable) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FieldTextWritable) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FieldTextWritable) SetRequired(v bool) {
	o.Required = &v
}

func (o FieldTextWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["prompt"] = o.Prompt
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FieldTextWritable) UnmarshalJSON(bytes []byte) (err error) {
	varFieldTextWritable := _FieldTextWritable{}

	err = json.Unmarshal(bytes, &varFieldTextWritable)
	if err == nil {
		*o = FieldTextWritable(varFieldTextWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFieldTextWritable struct {
	value *FieldTextWritable
	isSet bool
}

func (v NullableFieldTextWritable) Get() *FieldTextWritable {
	return v.value
}

func (v *NullableFieldTextWritable) Set(val *FieldTextWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldTextWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldTextWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldTextWritable(val *FieldTextWritable) *NullableFieldTextWritable {
	return &NullableFieldTextWritable{value: val, isSet: true}
}

func (v NullableFieldTextWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldTextWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
