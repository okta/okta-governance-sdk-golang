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

// checks if the AccessDurationSettingsAdminFixedDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDurationSettingsAdminFixedDuration{}

// AccessDurationSettingsAdminFixedDuration Settings when the access duration is specified by the admin.
type AccessDurationSettingsAdminFixedDuration struct {
	Type string `json:"type"`
	// The duration set by the admin for access durations. Use ISO8061 notation for duration values, see https://tc39.es/proposal-temporal/docs/duration.html. You can set up an access duration to a maximum of 72 hours (`PT72H`), 90 days (`90D`), or 12 weeks (`P12W`). For example:    - 24 hours (`PT24H`)   - 7 days (`P7D`)   - 2 weeks (`P2W`)
	Duration string `json:"duration"`
}

type _AccessDurationSettingsAdminFixedDuration AccessDurationSettingsAdminFixedDuration

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessDurationSettingsAdminFixedDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["duration"] = o.Duration
	return toSerialize, nil
}

func (o *AccessDurationSettingsAdminFixedDuration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"duration",
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

	varAccessDurationSettingsAdminFixedDuration := _AccessDurationSettingsAdminFixedDuration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessDurationSettingsAdminFixedDuration)

	if err != nil {
		return err
	}

	*o = AccessDurationSettingsAdminFixedDuration(varAccessDurationSettingsAdminFixedDuration)

	return err
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
