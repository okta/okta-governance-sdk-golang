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

// RequestResolutionTerminalState The terminal state for the access request in the request provider
type RequestResolutionTerminalState string

// List of request-resolution-terminal-state
const (
	REQUESTRESOLUTIONTERMINALSTATE_APPROVED RequestResolutionTerminalState = "APPROVED"
	REQUESTRESOLUTIONTERMINALSTATE_DENIED   RequestResolutionTerminalState = "DENIED"
	REQUESTRESOLUTIONTERMINALSTATE_CANCELED RequestResolutionTerminalState = "CANCELED"
	REQUESTRESOLUTIONTERMINALSTATE_EXPIRED  RequestResolutionTerminalState = "EXPIRED"
)

// All allowed values of RequestResolutionTerminalState enum
var AllowedRequestResolutionTerminalStateEnumValues = []RequestResolutionTerminalState{
	"APPROVED",
	"DENIED",
	"CANCELED",
	"EXPIRED",
}

func (v *RequestResolutionTerminalState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestResolutionTerminalState(value)
	for _, existing := range AllowedRequestResolutionTerminalStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestResolutionTerminalState", value)
}

// NewRequestResolutionTerminalStateFromValue returns a pointer to a valid RequestResolutionTerminalState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestResolutionTerminalStateFromValue(v string) (*RequestResolutionTerminalState, error) {
	ev := RequestResolutionTerminalState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestResolutionTerminalState: valid values are %v", v, AllowedRequestResolutionTerminalStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestResolutionTerminalState) IsValid() bool {
	for _, existing := range AllowedRequestResolutionTerminalStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to request-resolution-terminal-state value
func (v RequestResolutionTerminalState) Ptr() *RequestResolutionTerminalState {
	return &v
}

type NullableRequestResolutionTerminalState struct {
	value *RequestResolutionTerminalState
	isSet bool
}

func (v NullableRequestResolutionTerminalState) Get() *RequestResolutionTerminalState {
	return v.value
}

func (v *NullableRequestResolutionTerminalState) Set(val *RequestResolutionTerminalState) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResolutionTerminalState) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResolutionTerminalState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResolutionTerminalState(val *RequestResolutionTerminalState) *NullableRequestResolutionTerminalState {
	return &NullableRequestResolutionTerminalState{value: val, isSet: true}
}

func (v NullableRequestResolutionTerminalState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResolutionTerminalState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
