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

// checks if the MySettingsGetDelegateReadonly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MySettingsGetDelegateReadonly{}

// MySettingsGetDelegateReadonly Delegate appointment settings
type MySettingsGetDelegateReadonly struct {
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

type _MySettingsGetDelegateReadonly MySettingsGetDelegateReadonly

// NewMySettingsGetDelegateReadonly instantiates a new MySettingsGetDelegateReadonly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMySettingsGetDelegateReadonly(id string, delegate DelegateAppointmentDelegate, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *MySettingsGetDelegateReadonly {
	this := MySettingsGetDelegateReadonly{}
	this.Id = id
	this.Delegate = delegate
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	return &this
}

// NewMySettingsGetDelegateReadonlyWithDefaults instantiates a new MySettingsGetDelegateReadonly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMySettingsGetDelegateReadonlyWithDefaults() *MySettingsGetDelegateReadonly {
	this := MySettingsGetDelegateReadonly{}
	return &this
}

// GetId returns the Id field value
func (o *MySettingsGetDelegateReadonly) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MySettingsGetDelegateReadonly) SetId(v string) {
	o.Id = v
}

// GetDelegate returns the Delegate field value
func (o *MySettingsGetDelegateReadonly) GetDelegate() DelegateAppointmentDelegate {
	if o == nil {
		var ret DelegateAppointmentDelegate
		return ret
	}

	return o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetDelegateOk() (*DelegateAppointmentDelegate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegate, true
}

// SetDelegate sets field value
func (o *MySettingsGetDelegateReadonly) SetDelegate(v DelegateAppointmentDelegate) {
	o.Delegate = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MySettingsGetDelegateReadonly) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MySettingsGetDelegateReadonly) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *MySettingsGetDelegateReadonly) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *MySettingsGetDelegateReadonly) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *MySettingsGetDelegateReadonly) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *MySettingsGetDelegateReadonly) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *MySettingsGetDelegateReadonly) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *MySettingsGetDelegateReadonly) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *MySettingsGetDelegateReadonly) SetNote(v string) {
	o.Note = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *MySettingsGetDelegateReadonly) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *MySettingsGetDelegateReadonly) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *MySettingsGetDelegateReadonly) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *MySettingsGetDelegateReadonly) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *MySettingsGetDelegateReadonly) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *MySettingsGetDelegateReadonly) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *MySettingsGetDelegateReadonly) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *MySettingsGetDelegateReadonly) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *MySettingsGetDelegateReadonly) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o MySettingsGetDelegateReadonly) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MySettingsGetDelegateReadonly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["delegate"] = o.Delegate
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MySettingsGetDelegateReadonly) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"delegate",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
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

	varMySettingsGetDelegateReadonly := _MySettingsGetDelegateReadonly{}

	err = json.Unmarshal(data, &varMySettingsGetDelegateReadonly)

	if err != nil {
		return err
	}

	*o = MySettingsGetDelegateReadonly(varMySettingsGetDelegateReadonly)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableMySettingsGetDelegateReadonly struct {
	value *MySettingsGetDelegateReadonly
	isSet bool
}

func (v NullableMySettingsGetDelegateReadonly) Get() *MySettingsGetDelegateReadonly {
	return v.value
}

func (v *NullableMySettingsGetDelegateReadonly) Set(val *MySettingsGetDelegateReadonly) {
	v.value = val
	v.isSet = true
}

func (v NullableMySettingsGetDelegateReadonly) IsSet() bool {
	return v.isSet
}

func (v *NullableMySettingsGetDelegateReadonly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMySettingsGetDelegateReadonly(val *MySettingsGetDelegateReadonly) *NullableMySettingsGetDelegateReadonly {
	return &NullableMySettingsGetDelegateReadonly{value: val, isSet: true}
}

func (v NullableMySettingsGetDelegateReadonly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMySettingsGetDelegateReadonly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
