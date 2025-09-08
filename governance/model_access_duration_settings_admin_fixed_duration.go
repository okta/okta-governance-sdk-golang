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

// AccessDurationSettingsAdminFixedDuration Settings when the access duration is specified by the admin.
type AccessDurationSettingsAdminFixedDuration struct {
	Type string `json:"type"`
	// The duration set by the admin for access durations. Use ISO8061 notation for duration values, see https://tc39.es/proposal-temporal/docs/duration.html. You can set up an access duration to a maximum of 72 hours (`PT72H`), 90 days (`90D`), or 12 weeks (`P12W`). For example:    - 24 hours (`PT24H`)   - 7 days (`P7D`)   - 2 weeks (`P2W`)
	Duration string `json:"duration"`
}

// NewAccessDurationSettingsAdminFixedDuration instantiates a new AccessDurationSettingsAdminFixedDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDurationSettingsAdminFixedDuration(type_ string, duration string) *AccessDurationSettingsAdminFixedDuration {
	this := AccessDurationSettingsAdminFixedDuration{}
	this.Type = type_
	this.Duration = duration
	return &this
}

// NewAccessDurationSettingsAdminFixedDurationWithDefaults instantiates a new AccessDurationSettingsAdminFixedDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDurationSettingsAdminFixedDurationWithDefaults() *AccessDurationSettingsAdminFixedDuration {
	this := AccessDurationSettingsAdminFixedDuration{}
	return &this
}

// GetType returns the Type field value
func (o *AccessDurationSettingsAdminFixedDuration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessDurationSettingsAdminFixedDuration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessDurationSettingsAdminFixedDuration) SetType(v string) {
	o.Type = v
}

// GetDuration returns the Duration field value
func (o *AccessDurationSettingsAdminFixedDuration) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *AccessDurationSettingsAdminFixedDuration) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *AccessDurationSettingsAdminFixedDuration) SetDuration(v string) {
	o.Duration = v
}

func (o AccessDurationSettingsAdminFixedDuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableAccessDurationSettingsAdminFixedDuration struct {
	value *AccessDurationSettingsAdminFixedDuration
	isSet bool
}

func (v NullableAccessDurationSettingsAdminFixedDuration) Get() *AccessDurationSettingsAdminFixedDuration {
	return v.value
}

func (v *NullableAccessDurationSettingsAdminFixedDuration) Set(val *AccessDurationSettingsAdminFixedDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDurationSettingsAdminFixedDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDurationSettingsAdminFixedDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDurationSettingsAdminFixedDuration(val *AccessDurationSettingsAdminFixedDuration) *NullableAccessDurationSettingsAdminFixedDuration {
	return &NullableAccessDurationSettingsAdminFixedDuration{value: val, isSet: true}
}

func (v NullableAccessDurationSettingsAdminFixedDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDurationSettingsAdminFixedDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
