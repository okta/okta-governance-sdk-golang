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

// RequestResolutionCreatable The properties expected in a request resolution create request
type RequestResolutionCreatable struct {
	TerminalState        *RequestResolutionTerminalState `json:"terminalState,omitempty"`
	Decisions            []RequestDecisionCreatable      `json:"decisions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestResolutionCreatable RequestResolutionCreatable

// NewRequestResolutionCreatable instantiates a new RequestResolutionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResolutionCreatable() *RequestResolutionCreatable {
	this := RequestResolutionCreatable{}
	return &this
}

// NewRequestResolutionCreatableWithDefaults instantiates a new RequestResolutionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResolutionCreatableWithDefaults() *RequestResolutionCreatable {
	this := RequestResolutionCreatable{}
	return &this
}

// GetTerminalState returns the TerminalState field value if set, zero value otherwise.
func (o *RequestResolutionCreatable) GetTerminalState() RequestResolutionTerminalState {
	if o == nil || o.TerminalState == nil {
		var ret RequestResolutionTerminalState
		return ret
	}
	return *o.TerminalState
}

// GetTerminalStateOk returns a tuple with the TerminalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResolutionCreatable) GetTerminalStateOk() (*RequestResolutionTerminalState, bool) {
	if o == nil || o.TerminalState == nil {
		return nil, false
	}
	return o.TerminalState, true
}

// HasTerminalState returns a boolean if a field has been set.
func (o *RequestResolutionCreatable) HasTerminalState() bool {
	if o != nil && o.TerminalState != nil {
		return true
	}

	return false
}

// SetTerminalState gets a reference to the given RequestResolutionTerminalState and assigns it to the TerminalState field.
func (o *RequestResolutionCreatable) SetTerminalState(v RequestResolutionTerminalState) {
	o.TerminalState = &v
}

// GetDecisions returns the Decisions field value if set, zero value otherwise.
func (o *RequestResolutionCreatable) GetDecisions() []RequestDecisionCreatable {
	if o == nil || o.Decisions == nil {
		var ret []RequestDecisionCreatable
		return ret
	}
	return o.Decisions
}

// GetDecisionsOk returns a tuple with the Decisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResolutionCreatable) GetDecisionsOk() ([]RequestDecisionCreatable, bool) {
	if o == nil || o.Decisions == nil {
		return nil, false
	}
	return o.Decisions, true
}

// HasDecisions returns a boolean if a field has been set.
func (o *RequestResolutionCreatable) HasDecisions() bool {
	if o != nil && o.Decisions != nil {
		return true
	}

	return false
}

// SetDecisions gets a reference to the given []RequestDecisionCreatable and assigns it to the Decisions field.
func (o *RequestResolutionCreatable) SetDecisions(v []RequestDecisionCreatable) {
	o.Decisions = v
}

func (o RequestResolutionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TerminalState != nil {
		toSerialize["terminalState"] = o.TerminalState
	}
	if o.Decisions != nil {
		toSerialize["decisions"] = o.Decisions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestResolutionCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestResolutionCreatable := _RequestResolutionCreatable{}

	err = json.Unmarshal(bytes, &varRequestResolutionCreatable)
	if err == nil {
		*o = RequestResolutionCreatable(varRequestResolutionCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "terminalState")
		delete(additionalProperties, "decisions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestResolutionCreatable struct {
	value *RequestResolutionCreatable
	isSet bool
}

func (v NullableRequestResolutionCreatable) Get() *RequestResolutionCreatable {
	return v.value
}

func (v *NullableRequestResolutionCreatable) Set(val *RequestResolutionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResolutionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResolutionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResolutionCreatable(val *RequestResolutionCreatable) *NullableRequestResolutionCreatable {
	return &NullableRequestResolutionCreatable{value: val, isSet: true}
}

func (v NullableRequestResolutionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResolutionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
