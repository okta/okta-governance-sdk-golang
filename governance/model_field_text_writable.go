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
)

// checks if the FieldTextWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldTextWritable{}

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
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldTextWritable) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FieldTextWritable) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FieldTextWritable) SetRequired(v bool) {
	o.Required = &v
}

func (o FieldTextWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldTextWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["prompt"] = o.Prompt
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FieldTextWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"prompt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFieldTextWritable := _FieldTextWritable{}

	err = json.Unmarshal(data, &varFieldTextWritable)

	if err != nil {
		return err
	}

	*o = FieldTextWritable(varFieldTextWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
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
