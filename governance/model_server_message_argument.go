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

// checks if the ServerMessageArgument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMessageArgument{}

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerMessageArgument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerMessageArgument) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"type",
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

	varServerMessageArgument := _ServerMessageArgument{}

	err = json.Unmarshal(data, &varServerMessageArgument)

	if err != nil {
		return err
	}

	*o = ServerMessageArgument(varServerMessageArgument)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
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
