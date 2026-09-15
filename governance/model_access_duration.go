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
	"time"
)

// checks if the AccessDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDuration{}

// AccessDuration Specifies the duration of access with an expiration time and timezone
type AccessDuration struct {
	// The date and time when access expires, in ISO 8601 format
	ExpirationTime time.Time `json:"expirationTime"`
	// The timezone for the expiration time, in IANA format
	Timezone             string `json:"timezone"`
	AdditionalProperties map[string]interface{}
}

type _AccessDuration AccessDuration

// NewAccessDuration instantiates a new AccessDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDuration(expirationTime time.Time, timezone string) *AccessDuration {
	this := AccessDuration{}
	this.ExpirationTime = expirationTime
	this.Timezone = timezone
	return &this
}

// NewAccessDurationWithDefaults instantiates a new AccessDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDurationWithDefaults() *AccessDuration {
	this := AccessDuration{}
	return &this
}

// GetExpirationTime returns the ExpirationTime field value
func (o *AccessDuration) GetExpirationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value
// and a boolean to check if the value has been set.
func (o *AccessDuration) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTime, true
}

// SetExpirationTime sets field value
func (o *AccessDuration) SetExpirationTime(v time.Time) {
	o.ExpirationTime = v
}

// GetTimezone returns the Timezone field value
func (o *AccessDuration) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *AccessDuration) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *AccessDuration) SetTimezone(v string) {
	o.Timezone = v
}

func (o AccessDuration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expirationTime"] = o.ExpirationTime
	toSerialize["timezone"] = o.Timezone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessDuration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expirationTime",
		"timezone",
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

	varAccessDuration := _AccessDuration{}

	err = json.Unmarshal(data, &varAccessDuration)

	if err != nil {
		return err
	}

	*o = AccessDuration(varAccessDuration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessDuration struct {
	value *AccessDuration
	isSet bool
}

func (v NullableAccessDuration) Get() *AccessDuration {
	return v.value
}

func (v *NullableAccessDuration) Set(val *AccessDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDuration(val *AccessDuration) *NullableAccessDuration {
	return &NullableAccessDuration{value: val, isSet: true}
}

func (v NullableAccessDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
