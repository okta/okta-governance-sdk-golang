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

// checks if the ScheduleSettingsWriteable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleSettingsWriteable{}

// ScheduleSettingsWriteable Scheduler specific settings applicable to a grant.
type ScheduleSettingsWriteable struct {
	// An optional expiration date. If provided, the entitlement bundle grant will be revoked at the specified date and time.
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	// The time zone, in IANA format, in which expirationDate was configured
	TimeZone             *string `json:"timeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleSettingsWriteable ScheduleSettingsWriteable

// NewScheduleSettingsWriteable instantiates a new ScheduleSettingsWriteable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleSettingsWriteable() *ScheduleSettingsWriteable {
	this := ScheduleSettingsWriteable{}
	return &this
}

// NewScheduleSettingsWriteableWithDefaults instantiates a new ScheduleSettingsWriteable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleSettingsWriteableWithDefaults() *ScheduleSettingsWriteable {
	this := ScheduleSettingsWriteable{}
	return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ScheduleSettingsWriteable) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsWriteable) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ScheduleSettingsWriteable) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *ScheduleSettingsWriteable) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ScheduleSettingsWriteable) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsWriteable) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ScheduleSettingsWriteable) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ScheduleSettingsWriteable) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o ScheduleSettingsWriteable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleSettingsWriteable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduleSettingsWriteable) UnmarshalJSON(data []byte) (err error) {
	varScheduleSettingsWriteable := _ScheduleSettingsWriteable{}

	err = json.Unmarshal(data, &varScheduleSettingsWriteable)

	if err != nil {
		return err
	}

	*o = ScheduleSettingsWriteable(varScheduleSettingsWriteable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduleSettingsWriteable struct {
	value *ScheduleSettingsWriteable
	isSet bool
}

func (v NullableScheduleSettingsWriteable) Get() *ScheduleSettingsWriteable {
	return v.value
}

func (v *NullableScheduleSettingsWriteable) Set(val *ScheduleSettingsWriteable) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleSettingsWriteable) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleSettingsWriteable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleSettingsWriteable(val *ScheduleSettingsWriteable) *NullableScheduleSettingsWriteable {
	return &NullableScheduleSettingsWriteable{value: val, isSet: true}
}

func (v NullableScheduleSettingsWriteable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleSettingsWriteable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
