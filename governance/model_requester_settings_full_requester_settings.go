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
	"fmt"
)

// RequesterSettingsFullRequesterSettings - Requester settings define who may submit an access request for the related resource and access scopes.
type RequesterSettingsFullRequesterSettings struct {
	EveryoneRequesterSettings *EveryoneRequesterSettings
	GroupsRequesterSettings   *GroupsRequesterSettings
	TeamsRequesterSettings    *TeamsRequesterSettings
}

// EveryoneRequesterSettingsAsRequesterSettingsFullRequesterSettings is a convenience function that returns EveryoneRequesterSettings wrapped in RequesterSettingsFullRequesterSettings
func EveryoneRequesterSettingsAsRequesterSettingsFullRequesterSettings(v *EveryoneRequesterSettings) RequesterSettingsFullRequesterSettings {
	return RequesterSettingsFullRequesterSettings{
		EveryoneRequesterSettings: v,
	}
}

// GroupsRequesterSettingsAsRequesterSettingsFullRequesterSettings is a convenience function that returns GroupsRequesterSettings wrapped in RequesterSettingsFullRequesterSettings
func GroupsRequesterSettingsAsRequesterSettingsFullRequesterSettings(v *GroupsRequesterSettings) RequesterSettingsFullRequesterSettings {
	return RequesterSettingsFullRequesterSettings{
		GroupsRequesterSettings: v,
	}
}

// TeamsRequesterSettingsAsRequesterSettingsFullRequesterSettings is a convenience function that returns TeamsRequesterSettings wrapped in RequesterSettingsFullRequesterSettings
func TeamsRequesterSettingsAsRequesterSettingsFullRequesterSettings(v *TeamsRequesterSettings) RequesterSettingsFullRequesterSettings {
	return RequesterSettingsFullRequesterSettings{
		TeamsRequesterSettings: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequesterSettingsFullRequesterSettings) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVERYONE'
	if jsonDict["type"] == "EVERYONE" {
		// try to unmarshal JSON data into EveryoneRequesterSettings
		err = json.Unmarshal(data, &dst.EveryoneRequesterSettings)
		if err == nil {
			return nil // data stored in dst.EveryoneRequesterSettings, return on the first match
		} else {
			dst.EveryoneRequesterSettings = nil
			return fmt.Errorf("failed to unmarshal RequesterSettingsFullRequesterSettings as EveryoneRequesterSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUPS'
	if jsonDict["type"] == "GROUPS" {
		// try to unmarshal JSON data into GroupsRequesterSettings
		err = json.Unmarshal(data, &dst.GroupsRequesterSettings)
		if err == nil {
			return nil // data stored in dst.GroupsRequesterSettings, return on the first match
		} else {
			dst.GroupsRequesterSettings = nil
			return fmt.Errorf("failed to unmarshal RequesterSettingsFullRequesterSettings as GroupsRequesterSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TEAMS'
	if jsonDict["type"] == "TEAMS" {
		// try to unmarshal JSON data into TeamsRequesterSettings
		err = json.Unmarshal(data, &dst.TeamsRequesterSettings)
		if err == nil {
			return nil // data stored in dst.TeamsRequesterSettings, return on the first match
		} else {
			dst.TeamsRequesterSettings = nil
			return fmt.Errorf("failed to unmarshal RequesterSettingsFullRequesterSettings as TeamsRequesterSettings: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequesterSettingsFullRequesterSettings) MarshalJSON() ([]byte, error) {
	if src.EveryoneRequesterSettings != nil {
		return json.Marshal(&src.EveryoneRequesterSettings)
	}

	if src.GroupsRequesterSettings != nil {
		return json.Marshal(&src.GroupsRequesterSettings)
	}

	if src.TeamsRequesterSettings != nil {
		return json.Marshal(&src.TeamsRequesterSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequesterSettingsFullRequesterSettings) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EveryoneRequesterSettings != nil {
		return obj.EveryoneRequesterSettings
	}

	if obj.GroupsRequesterSettings != nil {
		return obj.GroupsRequesterSettings
	}

	if obj.TeamsRequesterSettings != nil {
		return obj.TeamsRequesterSettings
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequesterSettingsFullRequesterSettings) GetActualInstanceValue() interface{} {
	if obj.EveryoneRequesterSettings != nil {
		return *obj.EveryoneRequesterSettings
	}

	if obj.GroupsRequesterSettings != nil {
		return *obj.GroupsRequesterSettings
	}

	if obj.TeamsRequesterSettings != nil {
		return *obj.TeamsRequesterSettings
	}

	// all schemas are nil
	return nil
}

type NullableRequesterSettingsFullRequesterSettings struct {
	value *RequesterSettingsFullRequesterSettings
	isSet bool
}

func (v NullableRequesterSettingsFullRequesterSettings) Get() *RequesterSettingsFullRequesterSettings {
	return v.value
}

func (v *NullableRequesterSettingsFullRequesterSettings) Set(val *RequesterSettingsFullRequesterSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRequesterSettingsFullRequesterSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRequesterSettingsFullRequesterSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequesterSettingsFullRequesterSettings(val *RequesterSettingsFullRequesterSettings) *NullableRequesterSettingsFullRequesterSettings {
	return &NullableRequesterSettingsFullRequesterSettings{value: val, isSet: true}
}

func (v NullableRequesterSettingsFullRequesterSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequesterSettingsFullRequesterSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
