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
	"time"
)

// checks if the AssignmentProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignmentProperties{}

// AssignmentProperties struct for AssignmentProperties
type AssignmentProperties struct {
	// The date on which the principal's access expires. This property is specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	// The time zone, in IANA format, for the end date of the user access.
	TimeZone             *string `json:"timeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignmentProperties AssignmentProperties

// NewAssignmentProperties instantiates a new AssignmentProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentProperties() *AssignmentProperties {
	this := AssignmentProperties{}
	return &this
}

// NewAssignmentPropertiesWithDefaults instantiates a new AssignmentProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentPropertiesWithDefaults() *AssignmentProperties {
	this := AssignmentProperties{}
	return &this
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AssignmentProperties) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentProperties) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AssignmentProperties) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *AssignmentProperties) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AssignmentProperties) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentProperties) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AssignmentProperties) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AssignmentProperties) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o AssignmentProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignmentProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignmentProperties) UnmarshalJSON(data []byte) (err error) {
	varAssignmentProperties := _AssignmentProperties{}

	err = json.Unmarshal(data, &varAssignmentProperties)

	if err != nil {
		return err
	}

	*o = AssignmentProperties(varAssignmentProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignmentProperties struct {
	value *AssignmentProperties
	isSet bool
}

func (v NullableAssignmentProperties) Get() *AssignmentProperties {
	return v.value
}

func (v *NullableAssignmentProperties) Set(val *AssignmentProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentProperties(val *AssignmentProperties) *NullableAssignmentProperties {
	return &NullableAssignmentProperties{value: val, isSet: true}
}

func (v NullableAssignmentProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
