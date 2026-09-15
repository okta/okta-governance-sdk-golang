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

// checks if the CollectionResourceRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceRelationship{}

// CollectionResourceRelationship Resource relationship details. Only present when filtering by `resourceOrn`.
type CollectionResourceRelationship struct {
	// The ORN of the resource in this collection
	ResourceOrn *string `json:"resourceOrn,omitempty"`
	// When the resource was added to this collection
	AddedToCollectionDate *time.Time `json:"addedToCollectionDate,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _CollectionResourceRelationship CollectionResourceRelationship

// NewCollectionResourceRelationship instantiates a new CollectionResourceRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceRelationship() *CollectionResourceRelationship {
	this := CollectionResourceRelationship{}
	return &this
}

// NewCollectionResourceRelationshipWithDefaults instantiates a new CollectionResourceRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceRelationshipWithDefaults() *CollectionResourceRelationship {
	this := CollectionResourceRelationship{}
	return &this
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *CollectionResourceRelationship) GetResourceOrn() string {
	if o == nil || IsNil(o.ResourceOrn) {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceRelationship) GetResourceOrnOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceOrn) {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *CollectionResourceRelationship) HasResourceOrn() bool {
	if o != nil && !IsNil(o.ResourceOrn) {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *CollectionResourceRelationship) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

// GetAddedToCollectionDate returns the AddedToCollectionDate field value if set, zero value otherwise.
func (o *CollectionResourceRelationship) GetAddedToCollectionDate() time.Time {
	if o == nil || IsNil(o.AddedToCollectionDate) {
		var ret time.Time
		return ret
	}
	return *o.AddedToCollectionDate
}

// GetAddedToCollectionDateOk returns a tuple with the AddedToCollectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceRelationship) GetAddedToCollectionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddedToCollectionDate) {
		return nil, false
	}
	return o.AddedToCollectionDate, true
}

// HasAddedToCollectionDate returns a boolean if a field has been set.
func (o *CollectionResourceRelationship) HasAddedToCollectionDate() bool {
	if o != nil && !IsNil(o.AddedToCollectionDate) {
		return true
	}

	return false
}

// SetAddedToCollectionDate gets a reference to the given time.Time and assigns it to the AddedToCollectionDate field.
func (o *CollectionResourceRelationship) SetAddedToCollectionDate(v time.Time) {
	o.AddedToCollectionDate = &v
}

func (o CollectionResourceRelationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceOrn) {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if !IsNil(o.AddedToCollectionDate) {
		toSerialize["addedToCollectionDate"] = o.AddedToCollectionDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceRelationship) UnmarshalJSON(data []byte) (err error) {
	varCollectionResourceRelationship := _CollectionResourceRelationship{}

	err = json.Unmarshal(data, &varCollectionResourceRelationship)

	if err != nil {
		return err
	}

	*o = CollectionResourceRelationship(varCollectionResourceRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "addedToCollectionDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceRelationship struct {
	value *CollectionResourceRelationship
	isSet bool
}

func (v NullableCollectionResourceRelationship) Get() *CollectionResourceRelationship {
	return v.value
}

func (v *NullableCollectionResourceRelationship) Set(val *CollectionResourceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceRelationship(val *CollectionResourceRelationship) *NullableCollectionResourceRelationship {
	return &NullableCollectionResourceRelationship{value: val, isSet: true}
}

func (v NullableCollectionResourceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
