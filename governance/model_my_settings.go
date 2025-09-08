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

// MySettings struct for MySettings
type MySettings struct {
	Delegates            MySettingsDelegates `json:"delegates"`
	AdditionalProperties map[string]interface{}
}

type _MySettings MySettings

// NewMySettings instantiates a new MySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMySettings(delegates MySettingsDelegates) *MySettings {
	this := MySettings{}
	this.Delegates = delegates
	return &this
}

// NewMySettingsWithDefaults instantiates a new MySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMySettingsWithDefaults() *MySettings {
	this := MySettings{}
	return &this
}

// GetDelegates returns the Delegates field value
func (o *MySettings) GetDelegates() MySettingsDelegates {
	if o == nil {
		var ret MySettingsDelegates
		return ret
	}

	return o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value
// and a boolean to check if the value has been set.
func (o *MySettings) GetDelegatesOk() (*MySettingsDelegates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegates, true
}

// SetDelegates sets field value
func (o *MySettings) SetDelegates(v MySettingsDelegates) {
	o.Delegates = v
}

func (o MySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["delegates"] = o.Delegates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MySettings) UnmarshalJSON(bytes []byte) (err error) {
	varMySettings := _MySettings{}

	err = json.Unmarshal(bytes, &varMySettings)
	if err == nil {
		*o = MySettings(varMySettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "delegates")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
