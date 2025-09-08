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

// AccessScopeSettingsCreatableGroupAccessScopeSettings Access scope settings that resource permissions are groups.
type AccessScopeSettingsCreatableGroupAccessScopeSettings struct {
	Type   string                      `json:"type"`
	Groups []GroupsArrayCreatableInner `json:"groups"`
}

// NewAccessScopeSettingsCreatableGroupAccessScopeSettings instantiates a new AccessScopeSettingsCreatableGroupAccessScopeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessScopeSettingsCreatableGroupAccessScopeSettings(type_ string, groups []GroupsArrayCreatableInner) *AccessScopeSettingsCreatableGroupAccessScopeSettings {
	this := AccessScopeSettingsCreatableGroupAccessScopeSettings{}
	this.Type = type_
	this.Groups = groups
	return &this
}

// NewAccessScopeSettingsCreatableGroupAccessScopeSettingsWithDefaults instantiates a new AccessScopeSettingsCreatableGroupAccessScopeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessScopeSettingsCreatableGroupAccessScopeSettingsWithDefaults() *AccessScopeSettingsCreatableGroupAccessScopeSettings {
	this := AccessScopeSettingsCreatableGroupAccessScopeSettings{}
	return &this
}

// GetType returns the Type field value
func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) SetType(v string) {
	o.Type = v
}

// GetGroups returns the Groups field value
func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetGroups() []GroupsArrayCreatableInner {
	if o == nil {
		var ret []GroupsArrayCreatableInner
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetGroupsOk() ([]GroupsArrayCreatableInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) SetGroups(v []GroupsArrayCreatableInner) {
	o.Groups = v
}

func (o AccessScopeSettingsCreatableGroupAccessScopeSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableAccessScopeSettingsCreatableGroupAccessScopeSettings struct {
	value *AccessScopeSettingsCreatableGroupAccessScopeSettings
	isSet bool
}

func (v NullableAccessScopeSettingsCreatableGroupAccessScopeSettings) Get() *AccessScopeSettingsCreatableGroupAccessScopeSettings {
	return v.value
}

func (v *NullableAccessScopeSettingsCreatableGroupAccessScopeSettings) Set(val *AccessScopeSettingsCreatableGroupAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessScopeSettingsCreatableGroupAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessScopeSettingsCreatableGroupAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessScopeSettingsCreatableGroupAccessScopeSettings(val *AccessScopeSettingsCreatableGroupAccessScopeSettings) *NullableAccessScopeSettingsCreatableGroupAccessScopeSettings {
	return &NullableAccessScopeSettingsCreatableGroupAccessScopeSettings{value: val, isSet: true}
}

func (v NullableAccessScopeSettingsCreatableGroupAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessScopeSettingsCreatableGroupAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
