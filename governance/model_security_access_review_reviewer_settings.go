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

// SecurityAccessReviewReviewerSettings The settings for the reviewers a security access review. This includes the type of reviewer and specific user settings.
type SecurityAccessReviewReviewerSettings struct {
	Type                 SecurityAccessReviewReviewerType          `json:"type"`
	UserSettings         *SecurityAccessReviewUserReviewerSettings `json:"userSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewReviewerSettings SecurityAccessReviewReviewerSettings

// NewSecurityAccessReviewReviewerSettings instantiates a new SecurityAccessReviewReviewerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewReviewerSettings(type_ SecurityAccessReviewReviewerType) *SecurityAccessReviewReviewerSettings {
	this := SecurityAccessReviewReviewerSettings{}
	this.Type = type_
	return &this
}

// NewSecurityAccessReviewReviewerSettingsWithDefaults instantiates a new SecurityAccessReviewReviewerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewReviewerSettingsWithDefaults() *SecurityAccessReviewReviewerSettings {
	this := SecurityAccessReviewReviewerSettings{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityAccessReviewReviewerSettings) GetType() SecurityAccessReviewReviewerType {
	if o == nil {
		var ret SecurityAccessReviewReviewerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewReviewerSettings) GetTypeOk() (*SecurityAccessReviewReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewReviewerSettings) SetType(v SecurityAccessReviewReviewerType) {
	o.Type = v
}

// GetUserSettings returns the UserSettings field value if set, zero value otherwise.
func (o *SecurityAccessReviewReviewerSettings) GetUserSettings() SecurityAccessReviewUserReviewerSettings {
	if o == nil || o.UserSettings == nil {
		var ret SecurityAccessReviewUserReviewerSettings
		return ret
	}
	return *o.UserSettings
}

// GetUserSettingsOk returns a tuple with the UserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewReviewerSettings) GetUserSettingsOk() (*SecurityAccessReviewUserReviewerSettings, bool) {
	if o == nil || o.UserSettings == nil {
		return nil, false
	}
	return o.UserSettings, true
}

// HasUserSettings returns a boolean if a field has been set.
func (o *SecurityAccessReviewReviewerSettings) HasUserSettings() bool {
	if o != nil && o.UserSettings != nil {
		return true
	}

	return false
}

// SetUserSettings gets a reference to the given SecurityAccessReviewUserReviewerSettings and assigns it to the UserSettings field.
func (o *SecurityAccessReviewReviewerSettings) SetUserSettings(v SecurityAccessReviewUserReviewerSettings) {
	o.UserSettings = &v
}

func (o SecurityAccessReviewReviewerSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UserSettings != nil {
		toSerialize["userSettings"] = o.UserSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewReviewerSettings) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewReviewerSettings := _SecurityAccessReviewReviewerSettings{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewReviewerSettings)
	if err == nil {
		*o = SecurityAccessReviewReviewerSettings(varSecurityAccessReviewReviewerSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "userSettings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewReviewerSettings struct {
	value *SecurityAccessReviewReviewerSettings
	isSet bool
}

func (v NullableSecurityAccessReviewReviewerSettings) Get() *SecurityAccessReviewReviewerSettings {
	return v.value
}

func (v *NullableSecurityAccessReviewReviewerSettings) Set(val *SecurityAccessReviewReviewerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewReviewerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewReviewerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewReviewerSettings(val *SecurityAccessReviewReviewerSettings) *NullableSecurityAccessReviewReviewerSettings {
	return &NullableSecurityAccessReviewReviewerSettings{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewReviewerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewReviewerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
