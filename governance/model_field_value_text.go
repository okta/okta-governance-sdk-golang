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

// FieldValueText struct for FieldValueText
type FieldValueText struct {
	Type FieldTextType `json:"type"`
	// Value provided by a user.
	Value NullableString `json:"value"`
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id string `json:"id"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             bool `json:"required"`
	AdditionalProperties map[string]interface{}
}

type _FieldValueText FieldValueText

// NewFieldValueText instantiates a new FieldValueText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldValueText(type_ FieldTextType, value NullableString, id string, prompt string, required bool) *FieldValueText {
	this := FieldValueText{}
	this.Id = id
	this.Prompt = prompt
	this.Required = required
	return &this
}

// NewFieldValueTextWithDefaults instantiates a new FieldValueText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldValueTextWithDefaults() *FieldValueText {
	this := FieldValueText{}
	var required bool = true
	this.Required = required
	return &this
}

// GetType returns the Type field value
func (o *FieldValueText) GetType() FieldTextType {
	if o == nil {
		var ret FieldTextType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FieldValueText) GetTypeOk() (*FieldTextType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FieldValueText) SetType(v FieldTextType) {
	o.Type = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FieldValueText) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldValueText) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *FieldValueText) SetValue(v string) {
	o.Value.Set(&v)
}

// GetId returns the Id field value
func (o *FieldValueText) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FieldValueText) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FieldValueText) SetId(v string) {
	o.Id = v
}

// GetPrompt returns the Prompt field value
func (o *FieldValueText) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *FieldValueText) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *FieldValueText) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value
func (o *FieldValueText) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FieldValueText) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *FieldValueText) SetRequired(v bool) {
	o.Required = v
}

func (o FieldValueText) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["prompt"] = o.Prompt
	}
	if true {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FieldValueText) UnmarshalJSON(bytes []byte) (err error) {
	varFieldValueText := _FieldValueText{}

	err = json.Unmarshal(bytes, &varFieldValueText)
	if err == nil {
		*o = FieldValueText(varFieldValueText)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "id")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFieldValueText struct {
	value *FieldValueText
	isSet bool
}

func (v NullableFieldValueText) Get() *FieldValueText {
	return v.value
}

func (v *NullableFieldValueText) Set(val *FieldValueText) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldValueText) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldValueText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldValueText(val *FieldValueText) *NullableFieldValueText {
	return &NullableFieldValueText{value: val, isSet: true}
}

func (v NullableFieldValueText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldValueText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
