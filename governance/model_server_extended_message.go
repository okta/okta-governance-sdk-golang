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

// checks if the ServerExtendedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerExtendedMessage{}

// ServerExtendedMessage An extended version of server message which includes few additional properties.
type ServerExtendedMessage struct {
	// Server message with detailed content.
	Message string `json:"message"`
	// Dynamic arguments, used to construct the whole message, are supplied as an array of values.
	Args []ServerMessageArgument `json:"args,omitempty"`
	// A unique message code. You can use the message code for additional processes, such as localization.
	MessageCode          string `json:"messageCode"`
	AdditionalProperties map[string]interface{}
}

type _ServerExtendedMessage ServerExtendedMessage

// NewServerExtendedMessage instantiates a new ServerExtendedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerExtendedMessage(message string, messageCode string) *ServerExtendedMessage {
	this := ServerExtendedMessage{}
	this.Message = message
	this.MessageCode = messageCode
	return &this
}

// NewServerExtendedMessageWithDefaults instantiates a new ServerExtendedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerExtendedMessageWithDefaults() *ServerExtendedMessage {
	this := ServerExtendedMessage{}
	return &this
}

// GetMessage returns the Message field value
func (o *ServerExtendedMessage) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ServerExtendedMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ServerExtendedMessage) SetMessage(v string) {
	o.Message = v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *ServerExtendedMessage) GetArgs() []ServerMessageArgument {
	if o == nil || IsNil(o.Args) {
		var ret []ServerMessageArgument
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerExtendedMessage) GetArgsOk() ([]ServerMessageArgument, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ServerExtendedMessage) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []ServerMessageArgument and assigns it to the Args field.
func (o *ServerExtendedMessage) SetArgs(v []ServerMessageArgument) {
	o.Args = v
}

// GetMessageCode returns the MessageCode field value
func (o *ServerExtendedMessage) GetMessageCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value
// and a boolean to check if the value has been set.
func (o *ServerExtendedMessage) GetMessageCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageCode, true
}

// SetMessageCode sets field value
func (o *ServerExtendedMessage) SetMessageCode(v string) {
	o.MessageCode = v
}

func (o ServerExtendedMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerExtendedMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	toSerialize["messageCode"] = o.MessageCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerExtendedMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"messageCode",
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

	varServerExtendedMessage := _ServerExtendedMessage{}

	err = json.Unmarshal(data, &varServerExtendedMessage)

	if err != nil {
		return err
	}

	*o = ServerExtendedMessage(varServerExtendedMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "args")
		delete(additionalProperties, "messageCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerExtendedMessage struct {
	value *ServerExtendedMessage
	isSet bool
}

func (v NullableServerExtendedMessage) Get() *ServerExtendedMessage {
	return v.value
}

func (v *NullableServerExtendedMessage) Set(val *ServerExtendedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableServerExtendedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableServerExtendedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerExtendedMessage(val *ServerExtendedMessage) *NullableServerExtendedMessage {
	return &NullableServerExtendedMessage{value: val, isSet: true}
}

func (v NullableServerExtendedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerExtendedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
