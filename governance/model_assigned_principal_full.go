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

// AssignedPrincipalFull struct for AssignedPrincipalFull
type AssignedPrincipalFull struct {
	// The date on which the principal's access expires. This property is specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	// The time zone, in IANA format, for the end date of the user access.
	TimeZone  *string              `json:"timeZone,omitempty"`
	Principal *TargetPrincipalFull `json:"principal,omitempty"`
	Actor     *GrantActor          `json:"actor,omitempty"`
	// The resource collection `id`
	CollectionId *string `json:"collectionId,omitempty" validate:"regexp=col[0-9a-zA-Z]+"`
	// The assignment `id`
	Id                   *string                  `json:"id,omitempty"`
	AssignmentType       *PrincipalAssignmentType `json:"assignmentType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignedPrincipalFull AssignedPrincipalFull

// NewAssignedPrincipalFull instantiates a new AssignedPrincipalFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedPrincipalFull() *AssignedPrincipalFull {
	this := AssignedPrincipalFull{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	var assignmentType PrincipalAssignmentType = PRINCIPALASSIGNMENTTYPE_INDIVIDUAL
	this.AssignmentType = &assignmentType
	return &this
}

// NewAssignedPrincipalFullWithDefaults instantiates a new AssignedPrincipalFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedPrincipalFullWithDefaults() *AssignedPrincipalFull {
	this := AssignedPrincipalFull{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	var assignmentType PrincipalAssignmentType = PRINCIPALASSIGNMENTTYPE_INDIVIDUAL
	this.AssignmentType = &assignmentType
	return &this
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *AssignedPrincipalFull) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AssignedPrincipalFull) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetPrincipal() TargetPrincipalFull {
	if o == nil || o.Principal == nil {
		var ret TargetPrincipalFull
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given TargetPrincipalFull and assigns it to the Principal field.
func (o *AssignedPrincipalFull) SetPrincipal(v TargetPrincipalFull) {
	o.Principal = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetActor() GrantActor {
	if o == nil || o.Actor == nil {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetActorOk() (*GrantActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *AssignedPrincipalFull) SetActor(v GrantActor) {
	o.Actor = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *AssignedPrincipalFull) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssignedPrincipalFull) SetId(v string) {
	o.Id = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *AssignedPrincipalFull) GetAssignmentType() PrincipalAssignmentType {
	if o == nil || o.AssignmentType == nil {
		var ret PrincipalAssignmentType
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalFull) GetAssignmentTypeOk() (*PrincipalAssignmentType, bool) {
	if o == nil || o.AssignmentType == nil {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *AssignedPrincipalFull) HasAssignmentType() bool {
	if o != nil && o.AssignmentType != nil {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given PrincipalAssignmentType and assigns it to the AssignmentType field.
func (o *AssignedPrincipalFull) SetAssignmentType(v PrincipalAssignmentType) {
	o.AssignmentType = &v
}

func (o AssignedPrincipalFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationTime != nil {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AssignmentType != nil {
		toSerialize["assignmentType"] = o.AssignmentType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssignedPrincipalFull) UnmarshalJSON(bytes []byte) (err error) {
	varAssignedPrincipalFull := _AssignedPrincipalFull{}

	err = json.Unmarshal(bytes, &varAssignedPrincipalFull)
	if err == nil {
		*o = AssignedPrincipalFull(varAssignedPrincipalFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "collectionId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "assignmentType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAssignedPrincipalFull struct {
	value *AssignedPrincipalFull
	isSet bool
}

func (v NullableAssignedPrincipalFull) Get() *AssignedPrincipalFull {
	return v.value
}

func (v *NullableAssignedPrincipalFull) Set(val *AssignedPrincipalFull) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedPrincipalFull) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedPrincipalFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedPrincipalFull(val *AssignedPrincipalFull) *NullableAssignedPrincipalFull {
	return &NullableAssignedPrincipalFull{value: val, isSet: true}
}

func (v NullableAssignedPrincipalFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedPrincipalFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
