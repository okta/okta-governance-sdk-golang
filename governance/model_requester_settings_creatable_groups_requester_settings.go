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

// RequesterSettingsCreatableGroupsRequesterSettings A requester settings indicating that access request can be submitted by specific groups.
type RequesterSettingsCreatableGroupsRequesterSettings struct {
	Type   string                      `json:"type"`
	Groups []GroupsArrayCreatableInner `json:"groups"`
}

// NewRequesterSettingsCreatableGroupsRequesterSettings instantiates a new RequesterSettingsCreatableGroupsRequesterSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequesterSettingsCreatableGroupsRequesterSettings(type_ string, groups []GroupsArrayCreatableInner) *RequesterSettingsCreatableGroupsRequesterSettings {
	this := RequesterSettingsCreatableGroupsRequesterSettings{}
	this.Type = type_
	this.Groups = groups
	return &this
}

// NewRequesterSettingsCreatableGroupsRequesterSettingsWithDefaults instantiates a new RequesterSettingsCreatableGroupsRequesterSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequesterSettingsCreatableGroupsRequesterSettingsWithDefaults() *RequesterSettingsCreatableGroupsRequesterSettings {
	this := RequesterSettingsCreatableGroupsRequesterSettings{}
	return &this
}

// GetType returns the Type field value
func (o *RequesterSettingsCreatableGroupsRequesterSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequesterSettingsCreatableGroupsRequesterSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequesterSettingsCreatableGroupsRequesterSettings) SetType(v string) {
	o.Type = v
}

// GetGroups returns the Groups field value
func (o *RequesterSettingsCreatableGroupsRequesterSettings) GetGroups() []GroupsArrayCreatableInner {
	if o == nil {
		var ret []GroupsArrayCreatableInner
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *RequesterSettingsCreatableGroupsRequesterSettings) GetGroupsOk() ([]GroupsArrayCreatableInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *RequesterSettingsCreatableGroupsRequesterSettings) SetGroups(v []GroupsArrayCreatableInner) {
	o.Groups = v
}

func (o RequesterSettingsCreatableGroupsRequesterSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableRequesterSettingsCreatableGroupsRequesterSettings struct {
	value *RequesterSettingsCreatableGroupsRequesterSettings
	isSet bool
}

func (v NullableRequesterSettingsCreatableGroupsRequesterSettings) Get() *RequesterSettingsCreatableGroupsRequesterSettings {
	return v.value
}

func (v *NullableRequesterSettingsCreatableGroupsRequesterSettings) Set(val *RequesterSettingsCreatableGroupsRequesterSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRequesterSettingsCreatableGroupsRequesterSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRequesterSettingsCreatableGroupsRequesterSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequesterSettingsCreatableGroupsRequesterSettings(val *RequesterSettingsCreatableGroupsRequesterSettings) *NullableRequesterSettingsCreatableGroupsRequesterSettings {
	return &NullableRequesterSettingsCreatableGroupsRequesterSettings{value: val, isSet: true}
}

func (v NullableRequesterSettingsCreatableGroupsRequesterSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequesterSettingsCreatableGroupsRequesterSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
