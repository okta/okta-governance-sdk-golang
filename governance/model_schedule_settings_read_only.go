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

// ScheduleSettingsReadOnly Scheduler specific settings.  A campaign can be a `ONE_OFF` or a `RECURRING` campaign.  You can't provide both in the campaign definition.
type ScheduleSettingsReadOnly struct {
	Type ScheduleType `json:"type"`
	// The date on which the campaign is supposed to start. Accepts date in ISO 8601 format.
	StartDate time.Time `json:"startDate"`
	// The duration (in days) that the campaign is active.
	DurationInDays float32 `json:"durationInDays"`
	// The time zone, in IANA format, for the start date of the campaign.
	TimeZone string `json:"timeZone"`
	// The date on which the campaign is supposed to end. Date in ISO 8601 format.
	EndDate              *time.Time                   `json:"endDate,omitempty"`
	Recurrence           *RecurrenceDefinitionMutable `json:"recurrence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleSettingsReadOnly ScheduleSettingsReadOnly

// NewScheduleSettingsReadOnly instantiates a new ScheduleSettingsReadOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleSettingsReadOnly(type_ ScheduleType, startDate time.Time, durationInDays float32, timeZone string) *ScheduleSettingsReadOnly {
	this := ScheduleSettingsReadOnly{}
	this.Type = type_
	this.StartDate = startDate
	this.DurationInDays = durationInDays
	this.TimeZone = timeZone
	return &this
}

// NewScheduleSettingsReadOnlyWithDefaults instantiates a new ScheduleSettingsReadOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleSettingsReadOnlyWithDefaults() *ScheduleSettingsReadOnly {
	this := ScheduleSettingsReadOnly{}
	return &this
}

// GetType returns the Type field value
func (o *ScheduleSettingsReadOnly) GetType() ScheduleType {
	if o == nil {
		var ret ScheduleType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsReadOnly) GetTypeOk() (*ScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleSettingsReadOnly) SetType(v ScheduleType) {
	o.Type = v
}

// GetStartDate returns the StartDate field value
func (o *ScheduleSettingsReadOnly) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsReadOnly) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ScheduleSettingsReadOnly) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetDurationInDays returns the DurationInDays field value
func (o *ScheduleSettingsReadOnly) GetDurationInDays() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DurationInDays
}

// GetDurationInDaysOk returns a tuple with the DurationInDays field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsReadOnly) GetDurationInDaysOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInDays, true
}

// SetDurationInDays sets field value
func (o *ScheduleSettingsReadOnly) SetDurationInDays(v float32) {
	o.DurationInDays = v
}

// GetTimeZone returns the TimeZone field value
func (o *ScheduleSettingsReadOnly) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsReadOnly) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *ScheduleSettingsReadOnly) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ScheduleSettingsReadOnly) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsReadOnly) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ScheduleSettingsReadOnly) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ScheduleSettingsReadOnly) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *ScheduleSettingsReadOnly) GetRecurrence() RecurrenceDefinitionMutable {
	if o == nil || o.Recurrence == nil {
		var ret RecurrenceDefinitionMutable
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleSettingsReadOnly) GetRecurrenceOk() (*RecurrenceDefinitionMutable, bool) {
	if o == nil || o.Recurrence == nil {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *ScheduleSettingsReadOnly) HasRecurrence() bool {
	if o != nil && o.Recurrence != nil {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given RecurrenceDefinitionMutable and assigns it to the Recurrence field.
func (o *ScheduleSettingsReadOnly) SetRecurrence(v RecurrenceDefinitionMutable) {
	o.Recurrence = &v
}

func (o ScheduleSettingsReadOnly) MarshalJSON() ([]byte, error) {
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
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Recurrence != nil {
		toSerialize["recurrence"] = o.Recurrence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScheduleSettingsReadOnly) UnmarshalJSON(bytes []byte) (err error) {
	varScheduleSettingsReadOnly := _ScheduleSettingsReadOnly{}

	err = json.Unmarshal(bytes, &varScheduleSettingsReadOnly)
	if err == nil {
		*o = ScheduleSettingsReadOnly(varScheduleSettingsReadOnly)
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
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "recurrence")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableScheduleSettingsReadOnly struct {
	value *ScheduleSettingsReadOnly
	isSet bool
}

func (v NullableScheduleSettingsReadOnly) Get() *ScheduleSettingsReadOnly {
	return v.value
}

func (v *NullableScheduleSettingsReadOnly) Set(val *ScheduleSettingsReadOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleSettingsReadOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleSettingsReadOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleSettingsReadOnly(val *ScheduleSettingsReadOnly) *NullableScheduleSettingsReadOnly {
	return &NullableScheduleSettingsReadOnly{value: val, isSet: true}
}

func (v NullableScheduleSettingsReadOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleSettingsReadOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
