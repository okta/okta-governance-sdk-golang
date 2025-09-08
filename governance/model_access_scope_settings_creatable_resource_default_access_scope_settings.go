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

// AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings Default Access scope settings associated with the requesting resource.
type AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings struct {
	Type string `json:"type"`
}

// NewAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings instantiates a new AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings(type_ string) *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings {
	this := AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings{}
	this.Type = type_
	return &this
}

// NewAccessScopeSettingsCreatableResourceDefaultAccessScopeSettingsWithDefaults instantiates a new AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessScopeSettingsCreatableResourceDefaultAccessScopeSettingsWithDefaults() *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings {
	this := AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings{}
	return &this
}

// GetType returns the Type field value
func (o *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) SetType(v string) {
	o.Type = v
}

func (o AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings struct {
	value *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings
	isSet bool
}

func (v NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) Get() *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings {
	return v.value
}

func (v *NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) Set(val *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings(val *AccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) *NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings {
	return &NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings{value: val, isSet: true}
}

func (v NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeSettingsCreatableResourceDefaultAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
