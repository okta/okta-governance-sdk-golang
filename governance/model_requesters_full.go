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
)

// RequestersFull struct for RequestersFull
type RequestersFull struct {
	Id                   string             `json:"id"`
	Members              []RequestersMember `json:"members"`
	AdditionalProperties map[string]interface{}
}

type _RequestersFull RequestersFull

// NewRequestersFull instantiates a new RequestersFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestersFull(id string, members []RequestersMember) *RequestersFull {
	this := RequestersFull{}
	this.Id = id
	this.Members = members
	return &this
}

// NewRequestersFullWithDefaults instantiates a new RequestersFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestersFullWithDefaults() *RequestersFull {
	this := RequestersFull{}
	return &this
}

// GetId returns the Id field value
func (o *RequestersFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestersFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestersFull) SetId(v string) {
	o.Id = v
}

// GetMembers returns the Members field value
func (o *RequestersFull) GetMembers() []RequestersMember {
	if o == nil {
		var ret []RequestersMember
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *RequestersFull) GetMembersOk() ([]RequestersMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *RequestersFull) SetMembers(v []RequestersMember) {
	o.Members = v
}

func (o RequestersFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestersFull) UnmarshalJSON(bytes []byte) (err error) {
	varRequestersFull := _RequestersFull{}

	err = json.Unmarshal(bytes, &varRequestersFull)
	if err == nil {
		*o = RequestersFull(varRequestersFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestersFull struct {
	value *RequestersFull
	isSet bool
}

func (v NullableRequestersFull) Get() *RequestersFull {
	return v.value
}

func (v *NullableRequestersFull) Set(val *RequestersFull) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestersFull) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestersFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestersFull(val *RequestersFull) *NullableRequestersFull {
	return &NullableRequestersFull{value: val, isSet: true}
}

func (v NullableRequestersFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestersFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
