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

// checks if the CollectionPrincipalAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionPrincipalAssignment{}

// CollectionPrincipalAssignment Principal assignment details. Only present when filtering by `principalOrn`.
type CollectionPrincipalAssignment struct {
	// The ORN of the assigned principal
	PrincipalOrn *string `json:"principalOrn,omitempty"`
	// How the assignment was made
	Actor *string `json:"actor,omitempty"`
	// Type of assignment
	AssignmentType *string `json:"assignmentType,omitempty"`
	// When the assignment expires (null for indefinite)
	ExpirationTime       NullableTime `json:"expirationTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionPrincipalAssignment CollectionPrincipalAssignment

// NewCollectionPrincipalAssignment instantiates a new CollectionPrincipalAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionPrincipalAssignment() *CollectionPrincipalAssignment {
	this := CollectionPrincipalAssignment{}
	return &this
}

// NewCollectionPrincipalAssignmentWithDefaults instantiates a new CollectionPrincipalAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionPrincipalAssignmentWithDefaults() *CollectionPrincipalAssignment {
	this := CollectionPrincipalAssignment{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value if set, zero value otherwise.
func (o *CollectionPrincipalAssignment) GetPrincipalOrn() string {
	if o == nil || IsNil(o.PrincipalOrn) {
		var ret string
		return ret
	}
	return *o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPrincipalAssignment) GetPrincipalOrnOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalOrn) {
		return nil, false
	}
	return o.PrincipalOrn, true
}

// HasPrincipalOrn returns a boolean if a field has been set.
func (o *CollectionPrincipalAssignment) HasPrincipalOrn() bool {
	if o != nil && !IsNil(o.PrincipalOrn) {
		return true
	}

	return false
}

// SetPrincipalOrn gets a reference to the given string and assigns it to the PrincipalOrn field.
func (o *CollectionPrincipalAssignment) SetPrincipalOrn(v string) {
	o.PrincipalOrn = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *CollectionPrincipalAssignment) GetActor() string {
	if o == nil || IsNil(o.Actor) {
		var ret string
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPrincipalAssignment) GetActorOk() (*string, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *CollectionPrincipalAssignment) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given string and assigns it to the Actor field.
func (o *CollectionPrincipalAssignment) SetActor(v string) {
	o.Actor = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *CollectionPrincipalAssignment) GetAssignmentType() string {
	if o == nil || IsNil(o.AssignmentType) {
		var ret string
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPrincipalAssignment) GetAssignmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentType) {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *CollectionPrincipalAssignment) HasAssignmentType() bool {
	if o != nil && !IsNil(o.AssignmentType) {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given string and assigns it to the AssignmentType field.
func (o *CollectionPrincipalAssignment) SetAssignmentType(v string) {
	o.AssignmentType = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionPrincipalAssignment) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime.Get()
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionPrincipalAssignment) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationTime.Get(), o.ExpirationTime.IsSet()
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *CollectionPrincipalAssignment) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given NullableTime and assigns it to the ExpirationTime field.
func (o *CollectionPrincipalAssignment) SetExpirationTime(v time.Time) {
	o.ExpirationTime.Set(&v)
}

// SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil
func (o *CollectionPrincipalAssignment) SetExpirationTimeNil() {
	o.ExpirationTime.Set(nil)
}

// UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
func (o *CollectionPrincipalAssignment) UnsetExpirationTime() {
	o.ExpirationTime.Unset()
}

func (o CollectionPrincipalAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionPrincipalAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrincipalOrn) {
		toSerialize["principalOrn"] = o.PrincipalOrn
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.AssignmentType) {
		toSerialize["assignmentType"] = o.AssignmentType
	}
	if o.ExpirationTime.IsSet() {
		toSerialize["expirationTime"] = o.ExpirationTime.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionPrincipalAssignment) UnmarshalJSON(data []byte) (err error) {
	varCollectionPrincipalAssignment := _CollectionPrincipalAssignment{}

	err = json.Unmarshal(data, &varCollectionPrincipalAssignment)

	if err != nil {
		return err
	}

	*o = CollectionPrincipalAssignment(varCollectionPrincipalAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "expirationTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionPrincipalAssignment struct {
	value *CollectionPrincipalAssignment
	isSet bool
}

func (v NullableCollectionPrincipalAssignment) Get() *CollectionPrincipalAssignment {
	return v.value
}

func (v *NullableCollectionPrincipalAssignment) Set(val *CollectionPrincipalAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPrincipalAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPrincipalAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPrincipalAssignment(val *CollectionPrincipalAssignment) *NullableCollectionPrincipalAssignment {
	return &NullableCollectionPrincipalAssignment{value: val, isSet: true}
}

func (v NullableCollectionPrincipalAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPrincipalAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
