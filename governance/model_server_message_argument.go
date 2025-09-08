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

// ServerMessageArgument struct for ServerMessageArgument
type ServerMessageArgument struct {
	Value                string                    `json:"value"`
	Type                 ServerMessageArgumentType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ServerMessageArgument ServerMessageArgument

// NewServerMessageArgument instantiates a new ServerMessageArgument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMessageArgument(value string, type_ ServerMessageArgumentType) *ServerMessageArgument {
	this := ServerMessageArgument{}
	this.Value = value
	this.Type = type_
	return &this
}

// NewServerMessageArgumentWithDefaults instantiates a new ServerMessageArgument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMessageArgumentWithDefaults() *ServerMessageArgument {
	this := ServerMessageArgument{}
	return &this
}

// GetValue returns the Value field value
func (o *ServerMessageArgument) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ServerMessageArgument) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ServerMessageArgument) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value
func (o *ServerMessageArgument) GetType() ServerMessageArgumentType {
	if o == nil {
		var ret ServerMessageArgumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServerMessageArgument) GetTypeOk() (*ServerMessageArgumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServerMessageArgument) SetType(v ServerMessageArgumentType) {
	o.Type = v
}

func (o ServerMessageArgument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerMessageArgument) UnmarshalJSON(bytes []byte) (err error) {
	varServerMessageArgument := _ServerMessageArgument{}

	err = json.Unmarshal(bytes, &varServerMessageArgument)
	if err == nil {
		*o = ServerMessageArgument(varServerMessageArgument)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableServerMessageArgument struct {
	value *ServerMessageArgument
	isSet bool
}

func (v NullableServerMessageArgument) Get() *ServerMessageArgument {
	return v.value
}

func (v *NullableServerMessageArgument) Set(val *ServerMessageArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMessageArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMessageArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMessageArgument(val *ServerMessageArgument) *NullableServerMessageArgument {
	return &NullableServerMessageArgument{value: val, isSet: true}
}

func (v NullableServerMessageArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMessageArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
