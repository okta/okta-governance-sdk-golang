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

// checks if the MySettingsGetDelegates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MySettingsGetDelegates{}

// MySettingsGetDelegates My delegates
type MySettingsGetDelegates struct {
	// My delegate appointments
	Appointments []MySettingsGetDelegateReadonly `json:"appointments"`
	// My delegate permission settings  | Permission | Description | |------------|-------------| | `READ` | I can view my delegates | | `WRITE` | I can view and set my own delegates |
	Permissions          []string `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MySettingsGetDelegates MySettingsGetDelegates

// NewMySettingsGetDelegates instantiates a new MySettingsGetDelegates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMySettingsGetDelegates(appointments []MySettingsGetDelegateReadonly) *MySettingsGetDelegates {
	this := MySettingsGetDelegates{}
	this.Appointments = appointments
	return &this
}

// NewMySettingsGetDelegatesWithDefaults instantiates a new MySettingsGetDelegates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMySettingsGetDelegatesWithDefaults() *MySettingsGetDelegates {
	this := MySettingsGetDelegates{}
	return &this
}

// GetAppointments returns the Appointments field value
func (o *MySettingsGetDelegates) GetAppointments() []MySettingsGetDelegateReadonly {
	if o == nil {
		var ret []MySettingsGetDelegateReadonly
		return ret
	}

	return o.Appointments
}

// GetAppointmentsOk returns a tuple with the Appointments field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegates) GetAppointmentsOk() ([]MySettingsGetDelegateReadonly, bool) {
	if o == nil {
		return nil, false
	}
	return o.Appointments, true
}

// SetAppointments sets field value
func (o *MySettingsGetDelegates) SetAppointments(v []MySettingsGetDelegateReadonly) {
	o.Appointments = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *MySettingsGetDelegates) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegates) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *MySettingsGetDelegates) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *MySettingsGetDelegates) SetPermissions(v []string) {
	o.Permissions = v
}

func (o MySettingsGetDelegates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MySettingsGetDelegates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appointments"] = o.Appointments
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MySettingsGetDelegates) UnmarshalJSON(data []byte) (err error) {
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

	varMySettingsGetDelegates := _MySettingsGetDelegates{}

	err = json.Unmarshal(data, &varMySettingsGetDelegates)

	if err != nil {
		return err
	}

	*o = MySettingsGetDelegates(varMySettingsGetDelegates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appointments")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMySettingsGetDelegates struct {
	value *MySettingsGetDelegates
	isSet bool
}

func (v NullableMySettingsGetDelegates) Get() *MySettingsGetDelegates {
	return v.value
}

func (v *NullableMySettingsGetDelegates) Set(val *MySettingsGetDelegates) {
	v.value = val
	v.isSet = true
}

func (v NullableMySettingsGetDelegates) IsSet() bool {
	return v.isSet
}

func (v *NullableMySettingsGetDelegates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMySettingsGetDelegates(val *MySettingsGetDelegates) *NullableMySettingsGetDelegates {
	return &NullableMySettingsGetDelegates{value: val, isSet: true}
}

func (v NullableMySettingsGetDelegates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMySettingsGetDelegates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
