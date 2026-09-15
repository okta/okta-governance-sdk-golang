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
)

// checks if the CollectionAssignmentOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionAssignmentOrigin{}

// CollectionAssignmentOrigin Collection origin for a single (principal, resource) assignment
type CollectionAssignmentOrigin struct {
	PrincipalOrn string `json:"principalOrn"`
	ResourceOrn  string `json:"resourceOrn"`
	// True if the principal receives this resource through a collection.
	CollectionManaged    bool                                  `json:"collectionManaged"`
	Collection           *CollectionAssignmentOriginCollection `json:"collection,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionAssignmentOrigin CollectionAssignmentOrigin

// NewCollectionAssignmentOrigin instantiates a new CollectionAssignmentOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionAssignmentOrigin(principalOrn string, resourceOrn string, collectionManaged bool) *CollectionAssignmentOrigin {
	this := CollectionAssignmentOrigin{}
	this.PrincipalOrn = principalOrn
	this.ResourceOrn = resourceOrn
	this.CollectionManaged = collectionManaged
	return &this
}

// NewCollectionAssignmentOriginWithDefaults instantiates a new CollectionAssignmentOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionAssignmentOriginWithDefaults() *CollectionAssignmentOrigin {
	this := CollectionAssignmentOrigin{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *CollectionAssignmentOrigin) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOrigin) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *CollectionAssignmentOrigin) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionAssignmentOrigin) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOrigin) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionAssignmentOrigin) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetCollectionManaged returns the CollectionManaged field value
func (o *CollectionAssignmentOrigin) GetCollectionManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CollectionManaged
}

// GetCollectionManagedOk returns a tuple with the CollectionManaged field value
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOrigin) GetCollectionManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionManaged, true
}

// SetCollectionManaged sets field value
func (o *CollectionAssignmentOrigin) SetCollectionManaged(v bool) {
	o.CollectionManaged = v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *CollectionAssignmentOrigin) GetCollection() CollectionAssignmentOriginCollection {
	if o == nil || IsNil(o.Collection) {
		var ret CollectionAssignmentOriginCollection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOrigin) GetCollectionOk() (*CollectionAssignmentOriginCollection, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *CollectionAssignmentOrigin) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given CollectionAssignmentOriginCollection and assigns it to the Collection field.
func (o *CollectionAssignmentOrigin) SetCollection(v CollectionAssignmentOriginCollection) {
	o.Collection = &v
}

func (o CollectionAssignmentOrigin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionAssignmentOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["resourceOrn"] = o.ResourceOrn
	toSerialize["collectionManaged"] = o.CollectionManaged
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionAssignmentOrigin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"resourceOrn",
		"collectionManaged",
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

	varCollectionAssignmentOrigin := _CollectionAssignmentOrigin{}

	err = json.Unmarshal(data, &varCollectionAssignmentOrigin)

	if err != nil {
		return err
	}

	*o = CollectionAssignmentOrigin(varCollectionAssignmentOrigin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "collectionManaged")
		delete(additionalProperties, "collection")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionAssignmentOrigin struct {
	value *CollectionAssignmentOrigin
	isSet bool
}

func (v NullableCollectionAssignmentOrigin) Get() *CollectionAssignmentOrigin {
	return v.value
}

func (v *NullableCollectionAssignmentOrigin) Set(val *CollectionAssignmentOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionAssignmentOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionAssignmentOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionAssignmentOrigin(val *CollectionAssignmentOrigin) *NullableCollectionAssignmentOrigin {
	return &NullableCollectionAssignmentOrigin{value: val, isSet: true}
}

func (v NullableCollectionAssignmentOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionAssignmentOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
