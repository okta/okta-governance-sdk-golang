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

// checks if the FieldText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldText{}

// FieldText A text field
type FieldText struct {
	Type FieldTextType `json:"type"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required *bool `json:"required,omitempty"`
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _FieldText FieldText

// NewFieldText instantiates a new FieldText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldText(type_ FieldTextType, prompt string, id string) *FieldText {
	this := FieldText{}
	this.Prompt = prompt
	var required bool = true
	this.Required = &required
	this.Id = id
	return &this
}

// NewFieldTextWithDefaults instantiates a new FieldText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldTextWithDefaults() *FieldText {
	this := FieldText{}
	var required bool = true
	this.Required = &required
	return &this
}

// GetType returns the Type field value
func (o *FieldText) GetType() FieldTextType {
	if o == nil {
		var ret FieldTextType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FieldText) GetTypeOk() (*FieldTextType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FieldText) SetType(v FieldTextType) {
	o.Type = v
}

// GetPrompt returns the Prompt field value
func (o *FieldText) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *FieldText) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *FieldText) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FieldText) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldText) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FieldText) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FieldText) SetRequired(v bool) {
	o.Required = &v
}

// GetId returns the Id field value
func (o *FieldText) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FieldText) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FieldText) SetId(v string) {
	o.Id = v
}

func (o FieldText) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["prompt"] = o.Prompt
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FieldText) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"prompt",
		"id",
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

	varFieldText := _FieldText{}

	err = json.Unmarshal(data, &varFieldText)

	if err != nil {
		return err
	}

	*o = FieldText(varFieldText)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFieldText struct {
	value *FieldText
	isSet bool
}

func (v NullableFieldText) Get() *FieldText {
	return v.value
}

func (v *NullableFieldText) Set(val *FieldText) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldText) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldText(val *FieldText) *NullableFieldText {
	return &NullableFieldText{value: val, isSet: true}
}

func (v NullableFieldText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
