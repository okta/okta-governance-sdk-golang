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

// PrincipalSettings struct for PrincipalSettings
type PrincipalSettings struct {
	Delegates            PrincipalSettingsDelegates `json:"delegates"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalSettings PrincipalSettings

// NewPrincipalSettings instantiates a new PrincipalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalSettings(delegates PrincipalSettingsDelegates) *PrincipalSettings {
	this := PrincipalSettings{}
	this.Delegates = delegates
	return &this
}

// NewPrincipalSettingsWithDefaults instantiates a new PrincipalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalSettingsWithDefaults() *PrincipalSettings {
	this := PrincipalSettings{}
	return &this
}

// GetDelegates returns the Delegates field value
func (o *PrincipalSettings) GetDelegates() PrincipalSettingsDelegates {
	if o == nil {
		var ret PrincipalSettingsDelegates
		return ret
	}

	return o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value
// and a boolean to check if the value has been set.
func (o *PrincipalSettings) GetDelegatesOk() (*PrincipalSettingsDelegates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegates, true
}

// SetDelegates sets field value
func (o *PrincipalSettings) SetDelegates(v PrincipalSettingsDelegates) {
	o.Delegates = v
}

func (o PrincipalSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["delegates"] = o.Delegates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalSettings) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalSettings := _PrincipalSettings{}

	err = json.Unmarshal(bytes, &varPrincipalSettings)
	if err == nil {
		*o = PrincipalSettings(varPrincipalSettings)
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

type NullablePrincipalSettings struct {
	value *PrincipalSettings
	isSet bool
}

func (v NullablePrincipalSettings) Get() *PrincipalSettings {
	return v.value
}

func (v *NullablePrincipalSettings) Set(val *PrincipalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalSettings(val *PrincipalSettings) *NullablePrincipalSettings {
	return &NullablePrincipalSettings{value: val, isSet: true}
}

func (v NullablePrincipalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
