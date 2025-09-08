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

// BaseFieldValue The field value provided by the user
type BaseFieldValue struct {
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id string `json:"id"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             bool `json:"required"`
	AdditionalProperties map[string]interface{}
}

type _BaseFieldValue BaseFieldValue

// NewBaseFieldValue instantiates a new BaseFieldValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseFieldValue(id string, prompt string, required bool) *BaseFieldValue {
	this := BaseFieldValue{}
	this.Id = id
	this.Prompt = prompt
	this.Required = required
	return &this
}

// NewBaseFieldValueWithDefaults instantiates a new BaseFieldValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFieldValueWithDefaults() *BaseFieldValue {
	this := BaseFieldValue{}
	var required bool = true
	this.Required = required
	return &this
}

// GetId returns the Id field value
func (o *BaseFieldValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BaseFieldValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BaseFieldValue) SetId(v string) {
	o.Id = v
}

// GetPrompt returns the Prompt field value
func (o *BaseFieldValue) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *BaseFieldValue) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *BaseFieldValue) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value
func (o *BaseFieldValue) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *BaseFieldValue) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *BaseFieldValue) SetRequired(v bool) {
	o.Required = v
}

func (o BaseFieldValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *BaseFieldValue) UnmarshalJSON(bytes []byte) (err error) {
	varBaseFieldValue := _BaseFieldValue{}

	err = json.Unmarshal(bytes, &varBaseFieldValue)
	if err == nil {
		*o = BaseFieldValue(varBaseFieldValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBaseFieldValue struct {
	value *BaseFieldValue
	isSet bool
}

func (v NullableBaseFieldValue) Get() *BaseFieldValue {
	return v.value
}

func (v *NullableBaseFieldValue) Set(val *BaseFieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFieldValue(val *BaseFieldValue) *NullableBaseFieldValue {
	return &NullableBaseFieldValue{value: val, isSet: true}
}

func (v NullableBaseFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
