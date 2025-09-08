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

// RequestMessageCreatable The properties expected in a request message create request
type RequestMessageCreatable struct {
	// Message that will be created for the request. Newline can be specified by characters \"\\n\". Message will be visible for all users who can view the request.
	Message              string `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _RequestMessageCreatable RequestMessageCreatable

// NewRequestMessageCreatable instantiates a new RequestMessageCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMessageCreatable(message string) *RequestMessageCreatable {
	this := RequestMessageCreatable{}
	this.Message = message
	return &this
}

// NewRequestMessageCreatableWithDefaults instantiates a new RequestMessageCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMessageCreatableWithDefaults() *RequestMessageCreatable {
	this := RequestMessageCreatable{}
	return &this
}

// GetMessage returns the Message field value
func (o *RequestMessageCreatable) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RequestMessageCreatable) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RequestMessageCreatable) SetMessage(v string) {
	o.Message = v
}

func (o RequestMessageCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestMessageCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestMessageCreatable := _RequestMessageCreatable{}

	err = json.Unmarshal(bytes, &varRequestMessageCreatable)
	if err == nil {
		*o = RequestMessageCreatable(varRequestMessageCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestMessageCreatable struct {
	value *RequestMessageCreatable
	isSet bool
}

func (v NullableRequestMessageCreatable) Get() *RequestMessageCreatable {
	return v.value
}

func (v *NullableRequestMessageCreatable) Set(val *RequestMessageCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMessageCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMessageCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMessageCreatable(val *RequestMessageCreatable) *NullableRequestMessageCreatable {
	return &NullableRequestMessageCreatable{value: val, isSet: true}
}

func (v NullableRequestMessageCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMessageCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
