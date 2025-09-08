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

// DelegateReadonly Delegate settings specify what delegates are associated with the user
type DelegateReadonly struct {
	// Unique identifier for the delegate appointment
	Id       string                      `json:"id"`
	Delegate DelegateAppointmentDelegate `json:"delegate"`
	// The start time of the delegate appointment, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format
	StartTime *time.Time `json:"startTime,omitempty"`
	// The time when the delegate appointment expires, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format
	EndTime *time.Time `json:"endTime,omitempty"`
	// A note that describes the delegate appointment
	Note *string `json:"note,omitempty"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy        string `json:"lastUpdatedBy"`
	AdditionalProperties map[string]interface{}
}

type _DelegateReadonly DelegateReadonly

// NewDelegateReadonly instantiates a new DelegateReadonly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateReadonly(id string, delegate DelegateAppointmentDelegate, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *DelegateReadonly {
	this := DelegateReadonly{}
	this.Id = id
	this.Delegate = delegate
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	return &this
}

// NewDelegateReadonlyWithDefaults instantiates a new DelegateReadonly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateReadonlyWithDefaults() *DelegateReadonly {
	this := DelegateReadonly{}
	return &this
}

// GetId returns the Id field value
func (o *DelegateReadonly) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelegateReadonly) SetId(v string) {
	o.Id = v
}

// GetDelegate returns the Delegate field value
func (o *DelegateReadonly) GetDelegate() DelegateAppointmentDelegate {
	if o == nil {
		var ret DelegateAppointmentDelegate
		return ret
	}

	return o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetDelegateOk() (*DelegateAppointmentDelegate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegate, true
}

// SetDelegate sets field value
func (o *DelegateReadonly) SetDelegate(v DelegateAppointmentDelegate) {
	o.Delegate = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DelegateReadonly) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DelegateReadonly) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *DelegateReadonly) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DelegateReadonly) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DelegateReadonly) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *DelegateReadonly) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *DelegateReadonly) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *DelegateReadonly) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *DelegateReadonly) SetNote(v string) {
	o.Note = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *DelegateReadonly) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *DelegateReadonly) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *DelegateReadonly) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DelegateReadonly) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *DelegateReadonly) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *DelegateReadonly) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *DelegateReadonly) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *DelegateReadonly) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *DelegateReadonly) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o DelegateReadonly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["delegate"] = o.Delegate
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DelegateReadonly) UnmarshalJSON(bytes []byte) (err error) {
	varDelegateReadonly := _DelegateReadonly{}

	err = json.Unmarshal(bytes, &varDelegateReadonly)
	if err == nil {
		*o = DelegateReadonly(varDelegateReadonly)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "delegate")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "note")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDelegateReadonly struct {
	value *DelegateReadonly
	isSet bool
}

func (v NullableDelegateReadonly) Get() *DelegateReadonly {
	return v.value
}

func (v *NullableDelegateReadonly) Set(val *DelegateReadonly) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateReadonly) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateReadonly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateReadonly(val *DelegateReadonly) *NullableDelegateReadonly {
	return &NullableDelegateReadonly{value: val, isSet: true}
}

func (v NullableDelegateReadonly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateReadonly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
