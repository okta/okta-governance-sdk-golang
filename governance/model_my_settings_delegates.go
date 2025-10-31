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

// checks if the MySettingsDelegates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MySettingsDelegates{}

// MySettingsDelegates My delegates
type MySettingsDelegates struct {
	// My delegate appointments
	Appointments         []DelegateReadonly `json:"appointments"`
	AdditionalProperties map[string]interface{}
}

type _MySettingsDelegates MySettingsDelegates

// NewMySettingsDelegates instantiates a new MySettingsDelegates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMySettingsDelegates(appointments []DelegateReadonly) *MySettingsDelegates {
	this := MySettingsDelegates{}
	this.Appointments = appointments
	return &this
}

// NewMySettingsDelegatesWithDefaults instantiates a new MySettingsDelegates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMySettingsDelegatesWithDefaults() *MySettingsDelegates {
	this := MySettingsDelegates{}
	return &this
}

// GetAppointments returns the Appointments field value
func (o *MySettingsDelegates) GetAppointments() []DelegateReadonly {
	if o == nil {
		var ret []DelegateReadonly
		return ret
	}

	return o.Appointments
}

// GetAppointmentsOk returns a tuple with the Appointments field value
// and a boolean to check if the value has been set.
func (o *MySettingsDelegates) GetAppointmentsOk() ([]DelegateReadonly, bool) {
	if o == nil {
		return nil, false
	}
	return o.Appointments, true
}

// SetAppointments sets field value
func (o *MySettingsDelegates) SetAppointments(v []DelegateReadonly) {
	o.Appointments = v
}

func (o MySettingsDelegates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MySettingsDelegates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appointments"] = o.Appointments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MySettingsDelegates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appointments",
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

	varMySettingsDelegates := _MySettingsDelegates{}

	err = json.Unmarshal(data, &varMySettingsDelegates)

	if err != nil {
		return err
	}

	*o = MySettingsDelegates(varMySettingsDelegates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appointments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMySettingsDelegates struct {
	value *MySettingsDelegates
	isSet bool
}

func (v NullableMySettingsDelegates) Get() *MySettingsDelegates {
	return v.value
}

func (v *NullableMySettingsDelegates) Set(val *MySettingsDelegates) {
	v.value = val
	v.isSet = true
}

func (v NullableMySettingsDelegates) IsSet() bool {
	return v.isSet
}

func (v *NullableMySettingsDelegates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMySettingsDelegates(val *MySettingsDelegates) *NullableMySettingsDelegates {
	return &NullableMySettingsDelegates{value: val, isSet: true}
}

func (v NullableMySettingsDelegates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMySettingsDelegates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
