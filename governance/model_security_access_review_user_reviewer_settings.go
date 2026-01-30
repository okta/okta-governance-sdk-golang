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
)

// checks if the SecurityAccessReviewUserReviewerSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewUserReviewerSettings{}

// SecurityAccessReviewUserReviewerSettings struct for SecurityAccessReviewUserReviewerSettings
type SecurityAccessReviewUserReviewerSettings struct {
	// The list of reviewer user IDs for the security access review
	IncludedUserIds      []string `json:"includedUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewUserReviewerSettings SecurityAccessReviewUserReviewerSettings

// NewSecurityAccessReviewUserReviewerSettings instantiates a new SecurityAccessReviewUserReviewerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewUserReviewerSettings() *SecurityAccessReviewUserReviewerSettings {
	this := SecurityAccessReviewUserReviewerSettings{}
	return &this
}

// NewSecurityAccessReviewUserReviewerSettingsWithDefaults instantiates a new SecurityAccessReviewUserReviewerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewUserReviewerSettingsWithDefaults() *SecurityAccessReviewUserReviewerSettings {
	this := SecurityAccessReviewUserReviewerSettings{}
	return &this
}

// GetIncludedUserIds returns the IncludedUserIds field value if set, zero value otherwise.
func (o *SecurityAccessReviewUserReviewerSettings) GetIncludedUserIds() []string {
	if o == nil || IsNil(o.IncludedUserIds) {
		var ret []string
		return ret
	}
	return o.IncludedUserIds
}

// GetIncludedUserIdsOk returns a tuple with the IncludedUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewUserReviewerSettings) GetIncludedUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedUserIds) {
		return nil, false
	}
	return o.IncludedUserIds, true
}

// HasIncludedUserIds returns a boolean if a field has been set.
func (o *SecurityAccessReviewUserReviewerSettings) HasIncludedUserIds() bool {
	if o != nil && !IsNil(o.IncludedUserIds) {
		return true
	}

	return false
}

// SetIncludedUserIds gets a reference to the given []string and assigns it to the IncludedUserIds field.
func (o *SecurityAccessReviewUserReviewerSettings) SetIncludedUserIds(v []string) {
	o.IncludedUserIds = v
}

func (o SecurityAccessReviewUserReviewerSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewUserReviewerSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludedUserIds) {
		toSerialize["includedUserIds"] = o.IncludedUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewUserReviewerSettings) UnmarshalJSON(data []byte) (err error) {
	varSecurityAccessReviewUserReviewerSettings := _SecurityAccessReviewUserReviewerSettings{}

	err = json.Unmarshal(data, &varSecurityAccessReviewUserReviewerSettings)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewUserReviewerSettings(varSecurityAccessReviewUserReviewerSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "includedUserIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewUserReviewerSettings struct {
	value *SecurityAccessReviewUserReviewerSettings
	isSet bool
}

func (v NullableSecurityAccessReviewUserReviewerSettings) Get() *SecurityAccessReviewUserReviewerSettings {
	return v.value
}

func (v *NullableSecurityAccessReviewUserReviewerSettings) Set(val *SecurityAccessReviewUserReviewerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewUserReviewerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewUserReviewerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewUserReviewerSettings(val *SecurityAccessReviewUserReviewerSettings) *NullableSecurityAccessReviewUserReviewerSettings {
	return &NullableSecurityAccessReviewUserReviewerSettings{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewUserReviewerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewUserReviewerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
