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

// DelegatesPatchable Delegates for the principal
type DelegatesPatchable struct {
	// Delegate appointments
	Appointments         []DelegatePatchable `json:"appointments"`
	AdditionalProperties map[string]interface{}
}

type _DelegatesPatchable DelegatesPatchable

// NewDelegatesPatchable instantiates a new DelegatesPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatesPatchable(appointments []DelegatePatchable) *DelegatesPatchable {
	this := DelegatesPatchable{}
	this.Appointments = appointments
	return &this
}

// NewDelegatesPatchableWithDefaults instantiates a new DelegatesPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatesPatchableWithDefaults() *DelegatesPatchable {
	this := DelegatesPatchable{}
	return &this
}

// GetAppointments returns the Appointments field value
// If the value is explicit nil, the zero value for []DelegatePatchable will be returned
func (o *DelegatesPatchable) GetAppointments() []DelegatePatchable {
	if o == nil {
		var ret []DelegatePatchable
		return ret
	}

	return o.Appointments
}

// GetAppointmentsOk returns a tuple with the Appointments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DelegatesPatchable) GetAppointmentsOk() ([]DelegatePatchable, bool) {
	if o == nil || o.Appointments == nil {
		return nil, false
	}
	return o.Appointments, true
}

// SetAppointments sets field value
func (o *DelegatesPatchable) SetAppointments(v []DelegatePatchable) {
	o.Appointments = v
}

func (o DelegatesPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appointments != nil {
		toSerialize["appointments"] = o.Appointments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DelegatesPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varDelegatesPatchable := _DelegatesPatchable{}

	err = json.Unmarshal(bytes, &varDelegatesPatchable)
	if err == nil {
		*o = DelegatesPatchable(varDelegatesPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appointments")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDelegatesPatchable struct {
	value *DelegatesPatchable
	isSet bool
}

func (v NullableDelegatesPatchable) Get() *DelegatesPatchable {
	return v.value
}

func (v *NullableDelegatesPatchable) Set(val *DelegatesPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatesPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatesPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatesPatchable(val *DelegatesPatchable) *NullableDelegatesPatchable {
	return &NullableDelegatesPatchable{value: val, isSet: true}
}

func (v NullableDelegatesPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatesPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
