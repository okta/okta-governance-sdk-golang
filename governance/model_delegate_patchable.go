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

// DelegatePatchable struct for DelegatePatchable
type DelegatePatchable struct {
	Delegate DelegateAppointmentDelegate `json:"delegate"`
	// A note that describes the delegate appointment
	Note *string `json:"note,omitempty"`
	// The start time of the delegate appointment, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format
	StartTime *time.Time `json:"startTime,omitempty"`
	// The time when the delegate appointment expires, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format
	EndTime              *time.Time `json:"endTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DelegatePatchable DelegatePatchable

// NewDelegatePatchable instantiates a new DelegatePatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatePatchable(delegate DelegateAppointmentDelegate) *DelegatePatchable {
	this := DelegatePatchable{}
	this.Delegate = delegate
	return &this
}

// NewDelegatePatchableWithDefaults instantiates a new DelegatePatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatePatchableWithDefaults() *DelegatePatchable {
	this := DelegatePatchable{}
	return &this
}

// GetDelegate returns the Delegate field value
func (o *DelegatePatchable) GetDelegate() DelegateAppointmentDelegate {
	if o == nil {
		var ret DelegateAppointmentDelegate
		return ret
	}

	return o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value
// and a boolean to check if the value has been set.
func (o *DelegatePatchable) GetDelegateOk() (*DelegateAppointmentDelegate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegate, true
}

// SetDelegate sets field value
func (o *DelegatePatchable) SetDelegate(v DelegateAppointmentDelegate) {
	o.Delegate = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *DelegatePatchable) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatePatchable) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *DelegatePatchable) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *DelegatePatchable) SetNote(v string) {
	o.Note = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DelegatePatchable) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatePatchable) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DelegatePatchable) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *DelegatePatchable) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DelegatePatchable) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatePatchable) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DelegatePatchable) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *DelegatePatchable) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o DelegatePatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["delegate"] = o.Delegate
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DelegatePatchable) UnmarshalJSON(bytes []byte) (err error) {
	varDelegatePatchable := _DelegatePatchable{}

	err = json.Unmarshal(bytes, &varDelegatePatchable)
	if err == nil {
		*o = DelegatePatchable(varDelegatePatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "delegate")
		delete(additionalProperties, "note")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDelegatePatchable struct {
	value *DelegatePatchable
	isSet bool
}

func (v NullableDelegatePatchable) Get() *DelegatePatchable {
	return v.value
}

func (v *NullableDelegatePatchable) Set(val *DelegatePatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatePatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatePatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatePatchable(val *DelegatePatchable) *NullableDelegatePatchable {
	return &NullableDelegatePatchable{value: val, isSet: true}
}

func (v NullableDelegatePatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatePatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
