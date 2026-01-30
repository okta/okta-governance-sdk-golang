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

// checks if the AssignedPrincipal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignedPrincipal{}

// AssignedPrincipal struct for AssignedPrincipal
type AssignedPrincipal struct {
	// The date on which the principal's access expires. This property is specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	// The time zone, in IANA format, for the end date of the user access.
	TimeZone  *string              `json:"timeZone,omitempty"`
	Principal *TargetPrincipalFull `json:"principal,omitempty"`
	Actor     *GrantActor          `json:"actor,omitempty"`
	// The resource collection `id`
	CollectionId         *string `json:"collectionId,omitempty" validate:"regexp=col[0-9a-zA-Z]+"`
	AdditionalProperties map[string]interface{}
}

type _AssignedPrincipal AssignedPrincipal

// NewAssignedPrincipal instantiates a new AssignedPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedPrincipal() *AssignedPrincipal {
	this := AssignedPrincipal{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// NewAssignedPrincipalWithDefaults instantiates a new AssignedPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedPrincipalWithDefaults() *AssignedPrincipal {
	this := AssignedPrincipal{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AssignedPrincipal) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipal) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AssignedPrincipal) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *AssignedPrincipal) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AssignedPrincipal) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipal) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AssignedPrincipal) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AssignedPrincipal) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *AssignedPrincipal) GetPrincipal() TargetPrincipalFull {
	if o == nil || IsNil(o.Principal) {
		var ret TargetPrincipalFull
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipal) GetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *AssignedPrincipal) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given TargetPrincipalFull and assigns it to the Principal field.
func (o *AssignedPrincipal) SetPrincipal(v TargetPrincipalFull) {
	o.Principal = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *AssignedPrincipal) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipal) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *AssignedPrincipal) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *AssignedPrincipal) SetActor(v GrantActor) {
	o.Actor = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *AssignedPrincipal) GetCollectionId() string {
	if o == nil || IsNil(o.CollectionId) {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedPrincipal) GetCollectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionId) {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *AssignedPrincipal) HasCollectionId() bool {
	if o != nil && !IsNil(o.CollectionId) {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *AssignedPrincipal) SetCollectionId(v string) {
	o.CollectionId = &v
}

func (o AssignedPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignedPrincipal) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignedPrincipal) UnmarshalJSON(data []byte) (err error) {
	varAssignedPrincipal := _AssignedPrincipal{}

	err = json.Unmarshal(data, &varAssignedPrincipal)

	if err != nil {
		return err
	}

	*o = AssignedPrincipal(varAssignedPrincipal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "collectionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignedPrincipal struct {
	value *AssignedPrincipal
	isSet bool
}

func (v NullableAssignedPrincipal) Get() *AssignedPrincipal {
	return v.value
}

func (v *NullableAssignedPrincipal) Set(val *AssignedPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedPrincipal(val *AssignedPrincipal) *NullableAssignedPrincipal {
	return &NullableAssignedPrincipal{value: val, isSet: true}
}

func (v NullableAssignedPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
