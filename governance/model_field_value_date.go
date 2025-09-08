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
	"time"
)

// FieldValueDate struct for FieldValueDate
type FieldValueDate struct {
	Type FieldDateTimeType `json:"type"`
	// Value provided by a user in ISO 8601 string format
	Value NullableTime `json:"value"`
	// A `read-only` field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).
	Id string `json:"id"`
	// Text to prompt the user with
	Prompt string `json:"prompt"`
	// Whether a value to this field is required to advance the request
	Required             bool `json:"required"`
	AdditionalProperties map[string]interface{}
}

type _FieldValueDate FieldValueDate

// NewFieldValueDate instantiates a new FieldValueDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldValueDate(type_ FieldDateTimeType, value NullableTime, id string, prompt string, required bool) *FieldValueDate {
	this := FieldValueDate{}
	this.Id = id
	this.Prompt = prompt
	this.Required = required
	return &this
}

// NewFieldValueDateWithDefaults instantiates a new FieldValueDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldValueDateWithDefaults() *FieldValueDate {
	this := FieldValueDate{}
	var required bool = true
	this.Required = required
	return &this
}

// GetType returns the Type field value
func (o *FieldValueDate) GetType() FieldDateTimeType {
	if o == nil {
		var ret FieldDateTimeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FieldValueDate) GetTypeOk() (*FieldDateTimeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FieldValueDate) SetType(v FieldDateTimeType) {
	o.Type = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FieldValueDate) GetValue() time.Time {
	if o == nil || o.Value.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldValueDate) GetValueOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *FieldValueDate) SetValue(v time.Time) {
	o.Value.Set(&v)
}

// GetId returns the Id field value
func (o *FieldValueDate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FieldValueDate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FieldValueDate) SetId(v string) {
	o.Id = v
}

// GetPrompt returns the Prompt field value
func (o *FieldValueDate) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *FieldValueDate) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *FieldValueDate) SetPrompt(v string) {
	o.Prompt = v
}

// GetRequired returns the Required field value
func (o *FieldValueDate) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FieldValueDate) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *FieldValueDate) SetRequired(v bool) {
	o.Required = v
}

func (o FieldValueDate) MarshalJSON() ([]byte, error) {
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

func (o *FieldValueDate) UnmarshalJSON(bytes []byte) (err error) {
	varFieldValueDate := _FieldValueDate{}

	err = json.Unmarshal(bytes, &varFieldValueDate)
	if err == nil {
		*o = FieldValueDate(varFieldValueDate)
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

type NullableFieldValueDate struct {
	value *FieldValueDate
	isSet bool
}

func (v NullableFieldValueDate) Get() *FieldValueDate {
	return v.value
}

func (v *NullableFieldValueDate) Set(val *FieldValueDate) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldValueDate) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldValueDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldValueDate(val *FieldValueDate) *NullableFieldValueDate {
	return &NullableFieldValueDate{value: val, isSet: true}
}

func (v NullableFieldValueDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldValueDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
