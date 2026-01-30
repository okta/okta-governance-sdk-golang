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

// checks if the BaseFieldValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseFieldValue{}

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseFieldValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["prompt"] = o.Prompt
	toSerialize["required"] = o.Required

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseFieldValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"prompt",
		"required",
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

	varBaseFieldValue := _BaseFieldValue{}

	err = json.Unmarshal(data, &varBaseFieldValue)

	if err != nil {
		return err
	}

	*o = BaseFieldValue(varBaseFieldValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
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
