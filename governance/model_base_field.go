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

// BaseField The field to use when prompting the user
type BaseField struct {
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required *bool `json:"required,omitempty"`
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _BaseField BaseField

// NewBaseField instantiates a new BaseField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseField(prompt string, id string) *BaseField {
	this := BaseField{}
	this.Prompt = prompt
	var required bool = true
	this.Required = &required
	this.Id = id
	return &this
}

// NewBaseFieldWithDefaults instantiates a new BaseField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFieldWithDefaults() *BaseField {
	this := BaseField{}
	var required bool = true
	this.Required = &required
	return &this
}

// GetPrompt returns the Prompt field value
func (o *BaseField) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *BaseField) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *BaseField) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BaseField) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseField) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BaseField) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BaseField) SetRequired(v bool) {
	o.Required = &v
}

// GetId returns the Id field value
func (o *BaseField) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseField) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseField) SetId(v string) {
	o.Id = v
}

func (o BaseField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prompt"] = o.Prompt
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseField) UnmarshalJSON(bytes []byte) (err error) {
	varBaseField := _BaseField{}

	err = json.Unmarshal(bytes, &varBaseField)
	if err == nil {
		*o = BaseField(varBaseField)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBaseField struct {
	value *BaseField
	isSet bool
}

func (v NullableBaseField) Get() *BaseField {
	return v.value
}

func (v *NullableBaseField) Set(val *BaseField) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseField) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseField(val *BaseField) *NullableBaseField {
	return &NullableBaseField{value: val, isSet: true}
}

func (v NullableBaseField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
