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

// checks if the BaseFieldWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseFieldWritable{}

// BaseFieldWritable The field to use when prompting the user
type BaseFieldWritable struct {
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseFieldWritable BaseFieldWritable

// NewBaseFieldWritable instantiates a new BaseFieldWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseFieldWritable(prompt string) *BaseFieldWritable {
	this := BaseFieldWritable{}
	this.Prompt = prompt
	var required bool = true
	this.Required = &required
	return &this
}

// NewBaseFieldWritableWithDefaults instantiates a new BaseFieldWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFieldWritableWithDefaults() *BaseFieldWritable {
	this := BaseFieldWritable{}
	var required bool = true
	this.Required = &required
	return &this
}

// GetPrompt returns the Prompt field value
func (o *BaseFieldWritable) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *BaseFieldWritable) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *BaseFieldWritable) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BaseFieldWritable) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFieldWritable) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BaseFieldWritable) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BaseFieldWritable) SetRequired(v bool) {
	o.Required = &v
}

func (o BaseFieldWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseFieldWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prompt"] = o.Prompt
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseFieldWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varBaseFieldWritable := _BaseFieldWritable{}

	err = json.Unmarshal(data, &varBaseFieldWritable)

	if err != nil {
		return err
	}

	*o = BaseFieldWritable(varBaseFieldWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseFieldWritable struct {
	value *BaseFieldWritable
	isSet bool
}

func (v NullableBaseFieldWritable) Get() *BaseFieldWritable {
	return v.value
}

func (v *NullableBaseFieldWritable) Set(val *BaseFieldWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFieldWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFieldWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFieldWritable(val *BaseFieldWritable) *NullableBaseFieldWritable {
	return &NullableBaseFieldWritable{value: val, isSet: true}
}

func (v NullableBaseFieldWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFieldWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
