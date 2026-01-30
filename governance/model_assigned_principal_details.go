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

// checks if the AssignedPrincipalDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignedPrincipalDetails{}

// AssignedPrincipalDetails struct for AssignedPrincipalDetails
type AssignedPrincipalDetails struct {
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
	PrincipalProfile     *PrincipalProfile        `json:"principalProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignedPrincipalDetails AssignedPrincipalDetails

// NewAssignedPrincipalDetails instantiates a new AssignedPrincipalDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedPrincipalDetails() *AssignedPrincipalDetails {
	this := AssignedPrincipalDetails{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	var assignmentType PrincipalAssignmentType = PRINCIPALASSIGNMENTTYPE_INDIVIDUAL
	this.AssignmentType = &assignmentType
	return &this
}

// NewAssignedPrincipalDetailsWithDefaults instantiates a new AssignedPrincipalDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedPrincipalDetailsWithDefaults() *AssignedPrincipalDetails {
	this := AssignedPrincipalDetails{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	var assignmentType PrincipalAssignmentType = PRINCIPALASSIGNMENTTYPE_INDIVIDUAL
	this.AssignmentType = &assignmentType
	return &this
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *AssignedPrincipalDetails) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AssignedPrincipalDetails) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetPrincipal() TargetPrincipalFull {
	if o == nil || IsNil(o.Principal) {
		var ret TargetPrincipalFull
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given TargetPrincipalFull and assigns it to the Principal field.
func (o *AssignedPrincipalDetails) SetPrincipal(v TargetPrincipalFull) {
	o.Principal = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *AssignedPrincipalDetails) SetActor(v GrantActor) {
	o.Actor = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetCollectionId() string {
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *AssignedPrincipalDetails) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssignedPrincipalDetails) SetId(v string) {
	o.Id = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetAssignmentType() PrincipalAssignmentType {
	if o == nil || IsNil(o.AssignmentType) {
		var ret PrincipalAssignmentType
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetAssignmentTypeOk() (*PrincipalAssignmentType, bool) {
	if o == nil || IsNil(o.AssignmentType) {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasAssignmentType() bool {
	if o != nil && !IsNil(o.AssignmentType) {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given PrincipalAssignmentType and assigns it to the AssignmentType field.
func (o *AssignedPrincipalDetails) SetAssignmentType(v PrincipalAssignmentType) {
	o.AssignmentType = &v
}

// GetPrincipalProfile returns the PrincipalProfile field value if set, zero value otherwise.
func (o *AssignedPrincipalDetails) GetPrincipalProfile() PrincipalProfile {
	if o == nil || IsNil(o.PrincipalProfile) {
		var ret PrincipalProfile
		return ret
	}
	return *o.PrincipalProfile
}

// GetPrincipalProfileOk returns a tuple with the PrincipalProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipalDetails) GetPrincipalProfileOk() (*PrincipalProfile, bool) {
	if o == nil || IsNil(o.PrincipalProfile) {
		return nil, false
	}
	return o.PrincipalProfile, true
}

// HasPrincipalProfile returns a boolean if a field has been set.
func (o *AssignedPrincipalDetails) HasPrincipalProfile() bool {
	if o != nil && !IsNil(o.PrincipalProfile) {
		return true
	}

	return false
}

// SetPrincipalProfile gets a reference to the given PrincipalProfile and assigns it to the PrincipalProfile field.
func (o *AssignedPrincipalDetails) SetPrincipalProfile(v PrincipalProfile) {
	o.PrincipalProfile = &v
}

func (o AssignedPrincipalDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignedPrincipalDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.CollectionId) {
		toSerialize["collectionId"] = o.CollectionId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AssignmentType) {
		toSerialize["assignmentType"] = o.AssignmentType
	}
	if !IsNil(o.PrincipalProfile) {
		toSerialize["principalProfile"] = o.PrincipalProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignedPrincipalDetails) UnmarshalJSON(data []byte) (err error) {
	varAssignedPrincipalDetails := _AssignedPrincipalDetails{}

	err = json.Unmarshal(data, &varAssignedPrincipalDetails)

	if err != nil {
		return err
	}

	*o = AssignedPrincipalDetails(varAssignedPrincipalDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "collectionId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "principalProfile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignedPrincipalDetails struct {
	value *AssignedPrincipalDetails
	isSet bool
}

func (v NullableAssignedPrincipalDetails) Get() *AssignedPrincipalDetails {
	return v.value
}

func (v *NullableAssignedPrincipalDetails) Set(val *AssignedPrincipalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedPrincipalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedPrincipalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedPrincipalDetails(val *AssignedPrincipalDetails) *NullableAssignedPrincipalDetails {
	return &NullableAssignedPrincipalDetails{value: val, isSet: true}
}

func (v NullableAssignedPrincipalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedPrincipalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
