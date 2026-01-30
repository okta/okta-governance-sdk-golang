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

// checks if the RequestMessageMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestMessageMutable{}

// RequestMessageMutable struct for RequestMessageMutable
type RequestMessageMutable struct {
	// Message that will be created for the request. Newline can be specified by characters \"\\n\". Message will be visible for all users who can view the request.
	Message              string `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _RequestMessageMutable RequestMessageMutable

// NewRequestMessageMutable instantiates a new RequestMessageMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMessageMutable(message string) *RequestMessageMutable {
	this := RequestMessageMutable{}
	this.Message = message
	return &this
}

// NewRequestMessageMutableWithDefaults instantiates a new RequestMessageMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMessageMutableWithDefaults() *RequestMessageMutable {
	this := RequestMessageMutable{}
	return &this
}

// GetMessage returns the Message field value
func (o *RequestMessageMutable) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RequestMessageMutable) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RequestMessageMutable) SetMessage(v string) {
	o.Message = v
}

func (o RequestMessageMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestMessageMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestMessageMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varRequestMessageMutable := _RequestMessageMutable{}

	err = json.Unmarshal(data, &varRequestMessageMutable)

	if err != nil {
		return err
	}

	*o = RequestMessageMutable(varRequestMessageMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestMessageMutable struct {
	value *RequestMessageMutable
	isSet bool
}

func (v NullableRequestMessageMutable) Get() *RequestMessageMutable {
	return v.value
}

func (v *NullableRequestMessageMutable) Set(val *RequestMessageMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMessageMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMessageMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMessageMutable(val *RequestMessageMutable) *NullableRequestMessageMutable {
	return &NullableRequestMessageMutable{value: val, isSet: true}
}

func (v NullableRequestMessageMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMessageMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
