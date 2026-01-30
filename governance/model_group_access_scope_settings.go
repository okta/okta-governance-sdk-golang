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

// checks if the GroupAccessScopeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupAccessScopeSettings{}

// GroupAccessScopeSettings Settings when the resource uses groups for access scoping.
type GroupAccessScopeSettings struct {
	Type string `json:"type"`
	// A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.
	Groups []GroupsArrayFullInner `json:"groups"`
}

type _GroupAccessScopeSettings GroupAccessScopeSettings

// NewGroupAccessScopeSettings instantiates a new GroupAccessScopeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupAccessScopeSettings(type_ string, groups []GroupsArrayFullInner) *GroupAccessScopeSettings {
	this := GroupAccessScopeSettings{}
	this.Type = type_
	this.Groups = groups
	return &this
}

// NewGroupAccessScopeSettingsWithDefaults instantiates a new GroupAccessScopeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupAccessScopeSettingsWithDefaults() *GroupAccessScopeSettings {
	this := GroupAccessScopeSettings{}
	return &this
}

// GetType returns the Type field value
func (o *GroupAccessScopeSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GroupAccessScopeSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GroupAccessScopeSettings) SetType(v string) {
	o.Type = v
}

// GetGroups returns the Groups field value
func (o *GroupAccessScopeSettings) GetGroups() []GroupsArrayFullInner {
	if o == nil {
		var ret []GroupsArrayFullInner
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *GroupAccessScopeSettings) GetGroupsOk() ([]GroupsArrayFullInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *GroupAccessScopeSettings) SetGroups(v []GroupsArrayFullInner) {
	o.Groups = v
}

func (o GroupAccessScopeSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupAccessScopeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["groups"] = o.Groups
	return toSerialize, nil
}

func (o *GroupAccessScopeSettings) UnmarshalJSON(data []byte) (err error) {
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

	varGroupAccessScopeSettings := _GroupAccessScopeSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupAccessScopeSettings)

	if err != nil {
		return err
	}

	*o = GroupAccessScopeSettings(varGroupAccessScopeSettings)

	return err
}

type NullableGroupAccessScopeSettings struct {
	value *GroupAccessScopeSettings
	isSet bool
}

func (v NullableGroupAccessScopeSettings) Get() *GroupAccessScopeSettings {
	return v.value
}

func (v *NullableGroupAccessScopeSettings) Set(val *GroupAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupAccessScopeSettings(val *GroupAccessScopeSettings) *NullableGroupAccessScopeSettings {
	return &NullableGroupAccessScopeSettings{value: val, isSet: true}
}

func (v NullableGroupAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
