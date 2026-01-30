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
)

// checks if the PrincipalSettingsDelegates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalSettingsDelegates{}

// PrincipalSettingsDelegates Delegates for the principal
type PrincipalSettingsDelegates struct {
	// Delegate appointments
	Appointments         []DelegateAppointment `json:"appointments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalSettingsDelegates PrincipalSettingsDelegates

// NewPrincipalSettingsDelegates instantiates a new PrincipalSettingsDelegates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalSettingsDelegates() *PrincipalSettingsDelegates {
	this := PrincipalSettingsDelegates{}
	return &this
}

// NewPrincipalSettingsDelegatesWithDefaults instantiates a new PrincipalSettingsDelegates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalSettingsDelegatesWithDefaults() *PrincipalSettingsDelegates {
	this := PrincipalSettingsDelegates{}
	return &this
}

// GetAppointments returns the Appointments field value if set, zero value otherwise.
func (o *PrincipalSettingsDelegates) GetAppointments() []DelegateAppointment {
	if o == nil || IsNil(o.Appointments) {
		var ret []DelegateAppointment
		return ret
	}
	return o.Appointments
}

// GetAppointmentsOk returns a tuple with the Appointments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalSettingsDelegates) GetAppointmentsOk() ([]DelegateAppointment, bool) {
	if o == nil || IsNil(o.Appointments) {
		return nil, false
	}
	return o.Appointments, true
}

// HasAppointments returns a boolean if a field has been set.
func (o *PrincipalSettingsDelegates) HasAppointments() bool {
	if o != nil && !IsNil(o.Appointments) {
		return true
	}

	return false
}

// SetAppointments gets a reference to the given []DelegateAppointment and assigns it to the Appointments field.
func (o *PrincipalSettingsDelegates) SetAppointments(v []DelegateAppointment) {
	o.Appointments = v
}

func (o PrincipalSettingsDelegates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalSettingsDelegates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Appointments) {
		toSerialize["appointments"] = o.Appointments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalSettingsDelegates) UnmarshalJSON(data []byte) (err error) {
	varPrincipalSettingsDelegates := _PrincipalSettingsDelegates{}

	err = json.Unmarshal(data, &varPrincipalSettingsDelegates)

	if err != nil {
		return err
	}

	*o = PrincipalSettingsDelegates(varPrincipalSettingsDelegates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appointments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalSettingsDelegates struct {
	value *PrincipalSettingsDelegates
	isSet bool
}

func (v NullablePrincipalSettingsDelegates) Get() *PrincipalSettingsDelegates {
	return v.value
}

func (v *NullablePrincipalSettingsDelegates) Set(val *PrincipalSettingsDelegates) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalSettingsDelegates) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalSettingsDelegates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalSettingsDelegates(val *PrincipalSettingsDelegates) *NullablePrincipalSettingsDelegates {
	return &NullablePrincipalSettingsDelegates{value: val, isSet: true}
}

func (v NullablePrincipalSettingsDelegates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalSettingsDelegates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
