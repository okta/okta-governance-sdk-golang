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

// checks if the RecurrenceDefinitionMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecurrenceDefinitionMutable{}

// RecurrenceDefinitionMutable Applicable for recurring campaigns only (`type: RECURRING`). You can specify the campaign recurring frequency using `interval` and `repeatOn`. Optionally specify `ends` with a date-time, when the campaign schedule can end.
type RecurrenceDefinitionMutable struct {
	// Recurrence interval specified according to ISO8061 notation for durations. Refer https://tc39.es/proposal-temporal/docs/duration.html for all supported values. Recurring campaigns support - daily, weekly, monthly and yearly interval. One can setup a \"interval\" to max of -   - 182 days (That is `P182D`)   - 26 weeks (equivalent of 182 days, that is `P26W`)   - 24 months (That is `P24M`)   - 2 years (That is `P2Y`).  Few examples to better understand, how `interval` is used. If the `startDate` is `June 28th 2023` and the intent is to repeat the campaign -   1) Every 25 days, specify \"P25D\".   2) Every 4 weeks specify \"P4W\". In this example the `startDate` happens to be a `Wednesday`, the campaign would run every `Wednesday` of 4th week of the month.   3) Every 3 months specify \"P3M\". In this example the `startDate` is on 28th day. Hence, campaigns repeats every 3 months on 28th day.   4) Every 2 years specify \"P2Y\". In this example, the `startDate` is `Jun 28`. Hence repeats every 2 yrs on `Jun 28th`.    Also note that, the `interval` should not conflict `durationInDays`. For eg: If `durationInDays` is specified as 21 days, but interval is `P2W` (Every 2 weeks), it becomes an invalid request, as duration of the campaign conflicts with the next recurrence of the campaign.
	Interval     string                  `json:"interval"`
	RepeatOnType *RecurrenceRepeatOnType `json:"repeatOnType,omitempty"`
	// An optional field. Specifies when the recurring schedule can have an end. If not specified, the recurring schedule will run forever.
	Ends                 *time.Time `json:"ends,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecurrenceDefinitionMutable RecurrenceDefinitionMutable

// NewRecurrenceDefinitionMutable instantiates a new RecurrenceDefinitionMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrenceDefinitionMutable(interval string) *RecurrenceDefinitionMutable {
	this := RecurrenceDefinitionMutable{}
	this.Interval = interval
	return &this
}

// NewRecurrenceDefinitionMutableWithDefaults instantiates a new RecurrenceDefinitionMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceDefinitionMutableWithDefaults() *RecurrenceDefinitionMutable {
	this := RecurrenceDefinitionMutable{}
	return &this
}

// GetInterval returns the Interval field value
func (o *RecurrenceDefinitionMutable) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *RecurrenceDefinitionMutable) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *RecurrenceDefinitionMutable) SetInterval(v string) {
	o.Interval = v
}

// GetRepeatOnType returns the RepeatOnType field value if set, zero value otherwise.
func (o *RecurrenceDefinitionMutable) GetRepeatOnType() RecurrenceRepeatOnType {
	if o == nil || IsNil(o.RepeatOnType) {
		var ret RecurrenceRepeatOnType
		return ret
	}
	return *o.RepeatOnType
}

// GetRepeatOnTypeOk returns a tuple with the RepeatOnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceDefinitionMutable) GetRepeatOnTypeOk() (*RecurrenceRepeatOnType, bool) {
	if o == nil || IsNil(o.RepeatOnType) {
		return nil, false
	}
	return o.RepeatOnType, true
}

// HasRepeatOnType returns a boolean if a field has been set.
func (o *RecurrenceDefinitionMutable) HasRepeatOnType() bool {
	if o != nil && !IsNil(o.RepeatOnType) {
		return true
	}

	return false
}

// SetRepeatOnType gets a reference to the given RecurrenceRepeatOnType and assigns it to the RepeatOnType field.
func (o *RecurrenceDefinitionMutable) SetRepeatOnType(v RecurrenceRepeatOnType) {
	o.RepeatOnType = &v
}

// GetEnds returns the Ends field value if set, zero value otherwise.
func (o *RecurrenceDefinitionMutable) GetEnds() time.Time {
	if o == nil || IsNil(o.Ends) {
		var ret time.Time
		return ret
	}
	return *o.Ends
}

// GetEndsOk returns a tuple with the Ends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceDefinitionMutable) GetEndsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ends) {
		return nil, false
	}
	return o.Ends, true
}

// HasEnds returns a boolean if a field has been set.
func (o *RecurrenceDefinitionMutable) HasEnds() bool {
	if o != nil && !IsNil(o.Ends) {
		return true
	}

	return false
}

// SetEnds gets a reference to the given time.Time and assigns it to the Ends field.
func (o *RecurrenceDefinitionMutable) SetEnds(v time.Time) {
	o.Ends = &v
}

func (o RecurrenceDefinitionMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurrenceDefinitionMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	if !IsNil(o.RepeatOnType) {
		toSerialize["repeatOnType"] = o.RepeatOnType
	}
	if !IsNil(o.Ends) {
		toSerialize["ends"] = o.Ends
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecurrenceDefinitionMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
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

	varRecurrenceDefinitionMutable := _RecurrenceDefinitionMutable{}

	err = json.Unmarshal(data, &varRecurrenceDefinitionMutable)

	if err != nil {
		return err
	}

	*o = RecurrenceDefinitionMutable(varRecurrenceDefinitionMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interval")
		delete(additionalProperties, "repeatOnType")
		delete(additionalProperties, "ends")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecurrenceDefinitionMutable struct {
	value *RecurrenceDefinitionMutable
	isSet bool
}

func (v NullableRecurrenceDefinitionMutable) Get() *RecurrenceDefinitionMutable {
	return v.value
}

func (v *NullableRecurrenceDefinitionMutable) Set(val *RecurrenceDefinitionMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceDefinitionMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceDefinitionMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceDefinitionMutable(val *RecurrenceDefinitionMutable) *NullableRecurrenceDefinitionMutable {
	return &NullableRecurrenceDefinitionMutable{value: val, isSet: true}
}

func (v NullableRecurrenceDefinitionMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurrenceDefinitionMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
