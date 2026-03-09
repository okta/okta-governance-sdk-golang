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

// checks if the PrincipalSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalSettings{}

// PrincipalSettings struct for PrincipalSettings
type PrincipalSettings struct {
	Delegates            *PrincipalSettingsDelegates `json:"delegates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalSettings PrincipalSettings

// NewPrincipalSettings instantiates a new PrincipalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalSettings() *PrincipalSettings {
	this := PrincipalSettings{}
	return &this
}

// NewPrincipalSettingsWithDefaults instantiates a new PrincipalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalSettingsWithDefaults() *PrincipalSettings {
	this := PrincipalSettings{}
	return &this
}

// GetDelegates returns the Delegates field value if set, zero value otherwise.
func (o *PrincipalSettings) GetDelegates() PrincipalSettingsDelegates {
	if o == nil || IsNil(o.Delegates) {
		var ret PrincipalSettingsDelegates
		return ret
	}
	return *o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalSettings) GetDelegatesOk() (*PrincipalSettingsDelegates, bool) {
	if o == nil || IsNil(o.Delegates) {
		return nil, false
	}
	return o.Delegates, true
}

// HasDelegates returns a boolean if a field has been set.
func (o *PrincipalSettings) HasDelegates() bool {
	if o != nil && !IsNil(o.Delegates) {
		return true
	}

	return false
}

// SetDelegates gets a reference to the given PrincipalSettingsDelegates and assigns it to the Delegates field.
func (o *PrincipalSettings) SetDelegates(v PrincipalSettingsDelegates) {
	o.Delegates = &v
}

func (o PrincipalSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Delegates) {
		toSerialize["delegates"] = o.Delegates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalSettings) UnmarshalJSON(data []byte) (err error) {
	varPrincipalSettings := _PrincipalSettings{}

	err = json.Unmarshal(data, &varPrincipalSettings)

	if err != nil {
		return err
	}

	*o = PrincipalSettings(varPrincipalSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegates")
		o.AdditionalProperties = additionalProperties
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
