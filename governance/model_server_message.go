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

// checks if the ServerMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMessage{}

// ServerMessage struct for ServerMessage
type ServerMessage struct {
	// Server message with detailed content.
	Message string `json:"message"`
	// Dynamic arguments, used to construct the whole message, are supplied as an array of values.
	Args                 []ServerMessageArgument `json:"args,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerMessage ServerMessage

// NewServerMessage instantiates a new ServerMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMessage(message string) *ServerMessage {
	this := ServerMessage{}
	this.Message = message
	return &this
}

// NewServerMessageWithDefaults instantiates a new ServerMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMessageWithDefaults() *ServerMessage {
	this := ServerMessage{}
	return &this
}

// GetMessage returns the Message field value
func (o *ServerMessage) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ServerMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ServerMessage) SetMessage(v string) {
	o.Message = v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *ServerMessage) GetArgs() []ServerMessageArgument {
	if o == nil || IsNil(o.Args) {
		var ret []ServerMessageArgument
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMessage) GetArgsOk() ([]ServerMessageArgument, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ServerMessage) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []ServerMessageArgument and assigns it to the Args field.
func (o *ServerMessage) SetArgs(v []ServerMessageArgument) {
	o.Args = v
}

func (o ServerMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerMessage) UnmarshalJSON(data []byte) (err error) {
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

	varServerMessage := _ServerMessage{}

	err = json.Unmarshal(data, &varServerMessage)

	if err != nil {
		return err
	}

	*o = ServerMessage(varServerMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "args")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerMessage struct {
	value *ServerMessage
	isSet bool
}

func (v NullableServerMessage) Get() *ServerMessage {
	return v.value
}

func (v *NullableServerMessage) Set(val *ServerMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMessage(val *ServerMessage) *NullableServerMessage {
	return &NullableServerMessage{value: val, isSet: true}
}

func (v NullableServerMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
