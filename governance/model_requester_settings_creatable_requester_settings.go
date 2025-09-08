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
	"fmt"
)

// model_oneof.mustache
// RequesterSettingsCreatableRequesterSettings - Requester settings define who may submit an access request for the related resource and access scopes.
type RequesterSettingsCreatableRequesterSettings struct {
	EveryoneRequesterSettings                         *EveryoneRequesterSettings
	RequesterSettingsCreatableGroupsRequesterSettings *RequesterSettingsCreatableGroupsRequesterSettings
}

// EveryoneRequesterSettingsAsRequesterSettingsCreatableRequesterSettings is a convenience function that returns EveryoneRequesterSettings wrapped in RequesterSettingsCreatableRequesterSettings
func EveryoneRequesterSettingsAsRequesterSettingsCreatableRequesterSettings(v *EveryoneRequesterSettings) RequesterSettingsCreatableRequesterSettings {
	return RequesterSettingsCreatableRequesterSettings{
		EveryoneRequesterSettings: v,
	}
}

// RequesterSettingsCreatableGroupsRequesterSettingsAsRequesterSettingsCreatableRequesterSettings is a convenience function that returns RequesterSettingsCreatableGroupsRequesterSettings wrapped in RequesterSettingsCreatableRequesterSettings
func RequesterSettingsCreatableGroupsRequesterSettingsAsRequesterSettingsCreatableRequesterSettings(v *RequesterSettingsCreatableGroupsRequesterSettings) RequesterSettingsCreatableRequesterSettings {
	return RequesterSettingsCreatableRequesterSettings{
		RequesterSettingsCreatableGroupsRequesterSettings: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequesterSettingsCreatableRequesterSettings) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'EVERYONE'
	if jsonDict["type"] == "EVERYONE" {
		// try to unmarshal JSON data into EveryoneRequesterSettings
		err = json.Unmarshal(data, &dst.EveryoneRequesterSettings)
		if err == nil {
			return nil // data stored in dst.EveryoneRequesterSettings, return on the first match
		} else {
			dst.EveryoneRequesterSettings = nil
			return fmt.Errorf("Failed to unmarshal RequesterSettingsCreatableRequesterSettings as EveryoneRequesterSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GROUPS'
	if jsonDict["type"] == "GROUPS" {
		// try to unmarshal JSON data into RequesterSettingsCreatableGroupsRequesterSettings
		err = json.Unmarshal(data, &dst.RequesterSettingsCreatableGroupsRequesterSettings)
		if err == nil {
			return nil // data stored in dst.RequesterSettingsCreatableGroupsRequesterSettings, return on the first match
		} else {
			dst.RequesterSettingsCreatableGroupsRequesterSettings = nil
			return fmt.Errorf("Failed to unmarshal RequesterSettingsCreatableRequesterSettings as RequesterSettingsCreatableGroupsRequesterSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EveryoneRequesterSettings'
	if jsonDict["type"] == "EveryoneRequesterSettings" {
		// try to unmarshal JSON data into EveryoneRequesterSettings
		err = json.Unmarshal(data, &dst.EveryoneRequesterSettings)
		if err == nil {
			return nil // data stored in dst.EveryoneRequesterSettings, return on the first match
		} else {
			dst.EveryoneRequesterSettings = nil
			return fmt.Errorf("Failed to unmarshal RequesterSettingsCreatableRequesterSettings as EveryoneRequesterSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'requester-settings-creatable_GroupsRequesterSettings'
	if jsonDict["type"] == "requester-settings-creatable_GroupsRequesterSettings" {
		// try to unmarshal JSON data into RequesterSettingsCreatableGroupsRequesterSettings
		err = json.Unmarshal(data, &dst.RequesterSettingsCreatableGroupsRequesterSettings)
		if err == nil {
			return nil // data stored in dst.RequesterSettingsCreatableGroupsRequesterSettings, return on the first match
		} else {
			dst.RequesterSettingsCreatableGroupsRequesterSettings = nil
			return fmt.Errorf("Failed to unmarshal RequesterSettingsCreatableRequesterSettings as RequesterSettingsCreatableGroupsRequesterSettings: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequesterSettingsCreatableRequesterSettings) MarshalJSON() ([]byte, error) {
	if src.EveryoneRequesterSettings != nil {
		return json.Marshal(&src.EveryoneRequesterSettings)
	}

	if src.RequesterSettingsCreatableGroupsRequesterSettings != nil {
		return json.Marshal(&src.RequesterSettingsCreatableGroupsRequesterSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequesterSettingsCreatableRequesterSettings) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EveryoneRequesterSettings != nil {
		return obj.EveryoneRequesterSettings
	}

	if obj.RequesterSettingsCreatableGroupsRequesterSettings != nil {
		return obj.RequesterSettingsCreatableGroupsRequesterSettings
	}

	// all schemas are nil
	return nil
}

type NullableRequesterSettingsCreatableRequesterSettings struct {
	value *RequesterSettingsCreatableRequesterSettings
	isSet bool
}

func (v NullableRequesterSettingsCreatableRequesterSettings) Get() *RequesterSettingsCreatableRequesterSettings {
	return v.value
}

func (v *NullableRequesterSettingsCreatableRequesterSettings) Set(val *RequesterSettingsCreatableRequesterSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRequesterSettingsCreatableRequesterSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRequesterSettingsCreatableRequesterSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequesterSettingsCreatableRequesterSettings(val *RequesterSettingsCreatableRequesterSettings) *NullableRequesterSettingsCreatableRequesterSettings {
	return &NullableRequesterSettingsCreatableRequesterSettings{value: val, isSet: true}
}

func (v NullableRequesterSettingsCreatableRequesterSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequesterSettingsCreatableRequesterSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
