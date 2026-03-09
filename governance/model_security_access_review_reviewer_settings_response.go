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

// checks if the SecurityAccessReviewReviewerSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewReviewerSettingsResponse{}

// SecurityAccessReviewReviewerSettingsResponse The reviewer settings for a security access review. These include the type of reviewers and a list of reviewer IDs.
type SecurityAccessReviewReviewerSettingsResponse struct {
	Type                 SecurityAccessReviewReviewerType                      `json:"type"`
	UserSettings         *SecurityAccessReviewUserReviewerSettingsAddnlDetails `json:"userSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewReviewerSettingsResponse SecurityAccessReviewReviewerSettingsResponse

// NewSecurityAccessReviewReviewerSettingsResponse instantiates a new SecurityAccessReviewReviewerSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewReviewerSettingsResponse(type_ SecurityAccessReviewReviewerType) *SecurityAccessReviewReviewerSettingsResponse {
	this := SecurityAccessReviewReviewerSettingsResponse{}
	this.Type = type_
	return &this
}

// NewSecurityAccessReviewReviewerSettingsResponseWithDefaults instantiates a new SecurityAccessReviewReviewerSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewReviewerSettingsResponseWithDefaults() *SecurityAccessReviewReviewerSettingsResponse {
	this := SecurityAccessReviewReviewerSettingsResponse{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityAccessReviewReviewerSettingsResponse) GetType() SecurityAccessReviewReviewerType {
	if o == nil {
		var ret SecurityAccessReviewReviewerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewReviewerSettingsResponse) GetTypeOk() (*SecurityAccessReviewReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewReviewerSettingsResponse) SetType(v SecurityAccessReviewReviewerType) {
	o.Type = v
}

// GetUserSettings returns the UserSettings field value if set, zero value otherwise.
func (o *SecurityAccessReviewReviewerSettingsResponse) GetUserSettings() SecurityAccessReviewUserReviewerSettingsAddnlDetails {
	if o == nil || IsNil(o.UserSettings) {
		var ret SecurityAccessReviewUserReviewerSettingsAddnlDetails
		return ret
	}
	return *o.UserSettings
}

// GetUserSettingsOk returns a tuple with the UserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewReviewerSettingsResponse) GetUserSettingsOk() (*SecurityAccessReviewUserReviewerSettingsAddnlDetails, bool) {
	if o == nil || IsNil(o.UserSettings) {
		return nil, false
	}
	return o.UserSettings, true
}

// HasUserSettings returns a boolean if a field has been set.
func (o *SecurityAccessReviewReviewerSettingsResponse) HasUserSettings() bool {
	if o != nil && !IsNil(o.UserSettings) {
		return true
	}

	return false
}

// SetUserSettings gets a reference to the given SecurityAccessReviewUserReviewerSettingsAddnlDetails and assigns it to the UserSettings field.
func (o *SecurityAccessReviewReviewerSettingsResponse) SetUserSettings(v SecurityAccessReviewUserReviewerSettingsAddnlDetails) {
	o.UserSettings = &v
}

func (o SecurityAccessReviewReviewerSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewReviewerSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.UserSettings) {
		toSerialize["userSettings"] = o.UserSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewReviewerSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSecurityAccessReviewReviewerSettingsResponse := _SecurityAccessReviewReviewerSettingsResponse{}

	err = json.Unmarshal(data, &varSecurityAccessReviewReviewerSettingsResponse)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewReviewerSettingsResponse(varSecurityAccessReviewReviewerSettingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "userSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewReviewerSettingsResponse struct {
	value *SecurityAccessReviewReviewerSettingsResponse
	isSet bool
}

func (v NullableSecurityAccessReviewReviewerSettingsResponse) Get() *SecurityAccessReviewReviewerSettingsResponse {
	return v.value
}

func (v *NullableSecurityAccessReviewReviewerSettingsResponse) Set(val *SecurityAccessReviewReviewerSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewReviewerSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewReviewerSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewReviewerSettingsResponse(val *SecurityAccessReviewReviewerSettingsResponse) *NullableSecurityAccessReviewReviewerSettingsResponse {
	return &NullableSecurityAccessReviewReviewerSettingsResponse{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewReviewerSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewReviewerSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
