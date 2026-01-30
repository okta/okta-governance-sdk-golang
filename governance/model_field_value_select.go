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

// checks if the FieldValueSelect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldValueSelect{}

// FieldValueSelect struct for FieldValueSelect
type FieldValueSelect struct {
	Type FieldSelectType `json:"type"`
	// Values provided by a user, only 1 supported.
	Value []string `json:"value"`
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id string `json:"id"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             bool `json:"required"`
	AdditionalProperties map[string]interface{}
}

type _FieldValueSelect FieldValueSelect

// NewFieldValueSelect instantiates a new FieldValueSelect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldValueSelect(type_ FieldSelectType, value []string, id string, prompt string, required bool) *FieldValueSelect {
	this := FieldValueSelect{}
	this.Id = id
	this.Prompt = prompt
	this.Required = required
	return &this
}

// NewFieldValueSelectWithDefaults instantiates a new FieldValueSelect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldValueSelectWithDefaults() *FieldValueSelect {
	this := FieldValueSelect{}
	var required bool = true
	this.Required = required
	return &this
}

// GetType returns the Type field value
func (o *FieldValueSelect) GetType() FieldSelectType {
	if o == nil {
		var ret FieldSelectType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FieldValueSelect) GetTypeOk() (*FieldSelectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FieldValueSelect) SetType(v FieldSelectType) {
	o.Type = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *FieldValueSelect) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldValueSelect) GetValueOk() ([]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *FieldValueSelect) SetValue(v []string) {
	o.Value = v
}

// GetId returns the Id field value
func (o *FieldValueSelect) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FieldValueSelect) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FieldValueSelect) SetId(v string) {
	o.Id = v
}

// GetPrompt returns the Prompt field value
func (o *FieldValueSelect) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *FieldValueSelect) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *FieldValueSelect) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value
func (o *FieldValueSelect) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FieldValueSelect) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *FieldValueSelect) SetRequired(v bool) {
	o.Required = v
}

func (o FieldValueSelect) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldValueSelect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	toSerialize["id"] = o.Id
	toSerialize["prompt"] = o.Prompt
	toSerialize["required"] = o.Required

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FieldValueSelect) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varFieldValueSelect := _FieldValueSelect{}

	err = json.Unmarshal(data, &varFieldValueSelect)

	if err != nil {
		return err
	}

	*o = FieldValueSelect(varFieldValueSelect)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "id")
		delete(additionalProperties, "prompt")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFieldValueSelect struct {
	value *FieldValueSelect
	isSet bool
}

func (v NullableFieldValueSelect) Get() *FieldValueSelect {
	return v.value
}

func (v *NullableFieldValueSelect) Set(val *FieldValueSelect) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldValueSelect) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldValueSelect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldValueSelect(val *FieldValueSelect) *NullableFieldValueSelect {
	return &NullableFieldValueSelect{value: val, isSet: true}
}

func (v NullableFieldValueSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldValueSelect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
