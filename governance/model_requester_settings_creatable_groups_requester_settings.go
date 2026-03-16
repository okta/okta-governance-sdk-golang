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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RequesterSettingsCreatableGroupsRequesterSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequesterSettingsCreatableGroupsRequesterSettings{}

// RequesterSettingsCreatableGroupsRequesterSettings A requester settings indicating that access request can be submitted by specific groups.
type RequesterSettingsCreatableGroupsRequesterSettings struct {
	Type string `json:"type"`
	// List of requestable groups  > **Note:** Both standard Okta groups and AD-sourced groups are supported in Access Requests. > Standard Okta groups have the `okta:user_group` value, whereas AD-sourced groups have the `okta:windows_security_principal` value in their `objectClass` property.
	Groups []GroupsArrayCreatableInner `json:"groups"`
}

type _RequesterSettingsCreatableGroupsRequesterSettings RequesterSettingsCreatableGroupsRequesterSettings

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequesterSettingsCreatableGroupsRequesterSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["groups"] = o.Groups
	return toSerialize, nil
}

func (o *RequesterSettingsCreatableGroupsRequesterSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"groups",
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

	varRequesterSettingsCreatableGroupsRequesterSettings := _RequesterSettingsCreatableGroupsRequesterSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequesterSettingsCreatableGroupsRequesterSettings)

	if err != nil {
		return err
	}

	*o = RequesterSettingsCreatableGroupsRequesterSettings(varRequesterSettingsCreatableGroupsRequesterSettings)

	return err
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
