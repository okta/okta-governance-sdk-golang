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

// checks if the MySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MySettings{}

// MySettings struct for MySettings
type MySettings struct {
	Delegates            *MySettingsDelegates `json:"delegates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MySettings MySettings

// NewMySettings instantiates a new MySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMySettings() *MySettings {
	this := MySettings{}
	return &this
}

// NewMySettingsWithDefaults instantiates a new MySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMySettingsWithDefaults() *MySettings {
	this := MySettings{}
	return &this
}

// GetDelegates returns the Delegates field value if set, zero value otherwise.
func (o *MySettings) GetDelegates() MySettingsDelegates {
	if o == nil || IsNil(o.Delegates) {
		var ret MySettingsDelegates
		return ret
	}
	return *o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettings) GetDelegatesOk() (*MySettingsDelegates, bool) {
	if o == nil || IsNil(o.Delegates) {
		return nil, false
	}
	return o.Delegates, true
}

// HasDelegates returns a boolean if a field has been set.
func (o *MySettings) HasDelegates() bool {
	if o != nil && !IsNil(o.Delegates) {
		return true
	}

	return false
}

// SetDelegates gets a reference to the given MySettingsDelegates and assigns it to the Delegates field.
func (o *MySettings) SetDelegates(v MySettingsDelegates) {
	o.Delegates = &v
}

func (o MySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Delegates) {
		toSerialize["delegates"] = o.Delegates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MySettings) UnmarshalJSON(data []byte) (err error) {
	varMySettings := _MySettings{}

	err = json.Unmarshal(data, &varMySettings)

	if err != nil {
		return err
	}

	*o = MySettings(varMySettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMySettings struct {
	value *MySettings
	isSet bool
}

func (v NullableMySettings) Get() *MySettings {
	return v.value
}

func (v *NullableMySettings) Set(val *MySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMySettings(val *MySettings) *NullableMySettings {
	return &NullableMySettings{value: val, isSet: true}
}

func (v NullableMySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
