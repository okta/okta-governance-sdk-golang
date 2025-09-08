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
	"time"
)

// ScheduleSettingsMutable Scheduler specific settings.    A campaign can be a `ONE_OFF` or a `RECURRING` campaign.    You can't provide both in the campaign definition.
type ScheduleSettingsMutable struct {
	Type ScheduleType `json:"type"`
	// The date on which the campaign is supposed to start. Accepts date in ISO 8601 format.
	StartDate time.Time `json:"startDate"`
	// The duration (in days) that the campaign is active.
	DurationInDays float32 `json:"durationInDays"`
	// The time zone, in IANA format, for the start date of the campaign.
	TimeZone             string                       `json:"timeZone"`
	Recurrence           *RecurrenceDefinitionMutable `json:"recurrence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleSettingsMutable ScheduleSettingsMutable

// NewScheduleSettingsMutable instantiates a new ScheduleSettingsMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleSettingsMutable(type_ ScheduleType, startDate time.Time, durationInDays float32, timeZone string) *ScheduleSettingsMutable {
	this := ScheduleSettingsMutable{}
	this.Type = type_
	this.StartDate = startDate
	this.DurationInDays = durationInDays
	this.TimeZone = timeZone
	return &this
}

// NewScheduleSettingsMutableWithDefaults instantiates a new ScheduleSettingsMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleSettingsMutableWithDefaults() *ScheduleSettingsMutable {
	this := ScheduleSettingsMutable{}
	return &this
}

// GetType returns the Type field value
func (o *ScheduleSettingsMutable) GetType() ScheduleType {
	if o == nil {
		var ret ScheduleType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsMutable) GetTypeOk() (*ScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleSettingsMutable) SetType(v ScheduleType) {
	o.Type = v
}

// GetStartDate returns the StartDate field value
func (o *ScheduleSettingsMutable) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsMutable) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ScheduleSettingsMutable) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetDurationInDays returns the DurationInDays field value
func (o *ScheduleSettingsMutable) GetDurationInDays() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DurationInDays
}

// GetDurationInDaysOk returns a tuple with the DurationInDays field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsMutable) GetDurationInDaysOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInDays, true
}

// SetDurationInDays sets field value
func (o *ScheduleSettingsMutable) SetDurationInDays(v float32) {
	o.DurationInDays = v
}

// GetTimeZone returns the TimeZone field value
func (o *ScheduleSettingsMutable) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsMutable) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *ScheduleSettingsMutable) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *ScheduleSettingsMutable) GetRecurrence() RecurrenceDefinitionMutable {
	if o == nil || o.Recurrence == nil {
		var ret RecurrenceDefinitionMutable
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsMutable) GetRecurrenceOk() (*RecurrenceDefinitionMutable, bool) {
	if o == nil || o.Recurrence == nil {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *ScheduleSettingsMutable) HasRecurrence() bool {
	if o != nil && o.Recurrence != nil {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given RecurrenceDefinitionMutable and assigns it to the Recurrence field.
func (o *ScheduleSettingsMutable) SetRecurrence(v RecurrenceDefinitionMutable) {
	o.Recurrence = &v
}

func (o ScheduleSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["durationInDays"] = o.DurationInDays
	}
	if true {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Recurrence != nil {
		toSerialize["recurrence"] = o.Recurrence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScheduleSettingsMutable) UnmarshalJSON(bytes []byte) (err error) {
	varScheduleSettingsMutable := _ScheduleSettingsMutable{}

	err = json.Unmarshal(bytes, &varScheduleSettingsMutable)
	if err == nil {
		*o = ScheduleSettingsMutable(varScheduleSettingsMutable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "durationInDays")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "recurrence")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableScheduleSettingsMutable struct {
	value *ScheduleSettingsMutable
	isSet bool
}

func (v NullableScheduleSettingsMutable) Get() *ScheduleSettingsMutable {
	return v.value
}

func (v *NullableScheduleSettingsMutable) Set(val *ScheduleSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleSettingsMutable(val *ScheduleSettingsMutable) *NullableScheduleSettingsMutable {
	return &NullableScheduleSettingsMutable{value: val, isSet: true}
}

func (v NullableScheduleSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
