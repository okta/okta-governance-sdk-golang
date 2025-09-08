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

// FieldSelectWritable A select field
type FieldSelectWritable struct {
	Type FieldSelectType `json:"type"`
	// The options available for the select input.
	Options []FieldOption `json:"options"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FieldSelectWritable FieldSelectWritable

// NewFieldSelectWritable instantiates a new FieldSelectWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldSelectWritable(type_ FieldSelectType, options []FieldOption, prompt string) *FieldSelectWritable {
	this := FieldSelectWritable{}
	this.Prompt = prompt
	var required bool = true
	this.Required = &required
	return &this
}

// NewFieldSelectWritableWithDefaults instantiates a new FieldSelectWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldSelectWritableWithDefaults() *FieldSelectWritable {
	this := FieldSelectWritable{}
	var required bool = true
	this.Required = &required
	return &this
}

// GetType returns the Type field value
func (o *FieldSelectWritable) GetType() FieldSelectType {
	if o == nil {
		var ret FieldSelectType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FieldSelectWritable) GetTypeOk() (*FieldSelectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FieldSelectWritable) SetType(v FieldSelectType) {
	o.Type = v
}

// GetOptions returns the Options field value
func (o *FieldSelectWritable) GetOptions() []FieldOption {
	if o == nil {
		var ret []FieldOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *FieldSelectWritable) GetOptionsOk() ([]FieldOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *FieldSelectWritable) SetOptions(v []FieldOption) {
	o.Options = v
}

// GetPrompt returns the Prompt field value
func (o *FieldSelectWritable) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *FieldSelectWritable) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *FieldSelectWritable) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FieldSelectWritable) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldSelectWritable) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FieldSelectWritable) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FieldSelectWritable) SetRequired(v bool) {
	o.Required = &v
}

func (o FieldSelectWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["options"] = o.Options
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

func (o *FieldSelectWritable) UnmarshalJSON(bytes []byte) (err error) {
	varFieldSelectWritable := _FieldSelectWritable{}

	err = json.Unmarshal(bytes, &varFieldSelectWritable)
	if err == nil {
		*o = FieldSelectWritable(varFieldSelectWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "options")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFieldSelectWritable struct {
	value *FieldSelectWritable
	isSet bool
}

func (v NullableFieldSelectWritable) Get() *FieldSelectWritable {
	return v.value
}

func (v *NullableFieldSelectWritable) Set(val *FieldSelectWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldSelectWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldSelectWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldSelectWritable(val *FieldSelectWritable) *NullableFieldSelectWritable {
	return &NullableFieldSelectWritable{value: val, isSet: true}
}

func (v NullableFieldSelectWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldSelectWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
