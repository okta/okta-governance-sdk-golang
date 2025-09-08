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
	"fmt"
)

// ServerMessageArgumentType Type of the server message argument
type ServerMessageArgumentType string

// List of server-message-argument-type
const (
	SERVERMESSAGEARGUMENTTYPE_STRING               ServerMessageArgumentType = "STRING"
	SERVERMESSAGEARGUMENTTYPE_RELATIVE_DATE_TO_NOW ServerMessageArgumentType = "RELATIVE_DATE_TO_NOW"
	SERVERMESSAGEARGUMENTTYPE_NUMBER               ServerMessageArgumentType = "NUMBER"
)

// All allowed values of ServerMessageArgumentType enum
var AllowedServerMessageArgumentTypeEnumValues = []ServerMessageArgumentType{
	"STRING",
	"RELATIVE_DATE_TO_NOW",
	"NUMBER",
}

func (v *ServerMessageArgumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServerMessageArgumentType(value)
	for _, existing := range AllowedServerMessageArgumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServerMessageArgumentType", value)
}

// NewServerMessageArgumentTypeFromValue returns a pointer to a valid ServerMessageArgumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServerMessageArgumentTypeFromValue(v string) (*ServerMessageArgumentType, error) {
	ev := ServerMessageArgumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServerMessageArgumentType: valid values are %v", v, AllowedServerMessageArgumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServerMessageArgumentType) IsValid() bool {
	for _, existing := range AllowedServerMessageArgumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to server-message-argument-type value
func (v ServerMessageArgumentType) Ptr() *ServerMessageArgumentType {
	return &v
}

type NullableServerMessageArgumentType struct {
	value *ServerMessageArgumentType
	isSet bool
}

func (v NullableServerMessageArgumentType) Get() *ServerMessageArgumentType {
	return v.value
}

func (v *NullableServerMessageArgumentType) Set(val *ServerMessageArgumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMessageArgumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMessageArgumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMessageArgumentType(val *ServerMessageArgumentType) *NullableServerMessageArgumentType {
	return &NullableServerMessageArgumentType{value: val, isSet: true}
}

func (v NullableServerMessageArgumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMessageArgumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
