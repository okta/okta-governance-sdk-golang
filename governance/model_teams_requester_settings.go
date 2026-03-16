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

// checks if the TeamsRequesterSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamsRequesterSettings{}

// TeamsRequesterSettings A requester settings indicating that access request can be submitted by specific teams.
type TeamsRequesterSettings struct {
	Type string `json:"type"`
	// A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.
	Teams []TeamsArrayFullInner `json:"teams"`
}

type _TeamsRequesterSettings TeamsRequesterSettings

// NewTeamsRequesterSettings instantiates a new TeamsRequesterSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsRequesterSettings(type_ string, teams []TeamsArrayFullInner) *TeamsRequesterSettings {
	this := TeamsRequesterSettings{}
	this.Type = type_
	this.Teams = teams
	return &this
}

// NewTeamsRequesterSettingsWithDefaults instantiates a new TeamsRequesterSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsRequesterSettingsWithDefaults() *TeamsRequesterSettings {
	this := TeamsRequesterSettings{}
	return &this
}

// GetType returns the Type field value
func (o *TeamsRequesterSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamsRequesterSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TeamsRequesterSettings) SetType(v string) {
	o.Type = v
}

// GetTeams returns the Teams field value
func (o *TeamsRequesterSettings) GetTeams() []TeamsArrayFullInner {
	if o == nil {
		var ret []TeamsArrayFullInner
		return ret
	}

	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *TeamsRequesterSettings) GetTeamsOk() ([]TeamsArrayFullInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teams, true
}

// SetTeams sets field value
func (o *TeamsRequesterSettings) SetTeams(v []TeamsArrayFullInner) {
	o.Teams = v
}

func (o TeamsRequesterSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamsRequesterSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["teams"] = o.Teams
	return toSerialize, nil
}

func (o *TeamsRequesterSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"teams",
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

	varTeamsRequesterSettings := _TeamsRequesterSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamsRequesterSettings)

	if err != nil {
		return err
	}

	*o = TeamsRequesterSettings(varTeamsRequesterSettings)

	return err
}

type NullableTeamsRequesterSettings struct {
	value *TeamsRequesterSettings
	isSet bool
}

func (v NullableTeamsRequesterSettings) Get() *TeamsRequesterSettings {
	return v.value
}

func (v *NullableTeamsRequesterSettings) Set(val *TeamsRequesterSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsRequesterSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsRequesterSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsRequesterSettings(val *TeamsRequesterSettings) *NullableTeamsRequesterSettings {
	return &NullableTeamsRequesterSettings{value: val, isSet: true}
}

func (v NullableTeamsRequesterSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsRequesterSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
