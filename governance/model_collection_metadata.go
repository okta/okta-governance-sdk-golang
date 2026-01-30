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
)

// checks if the CollectionMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionMetadata{}

// CollectionMetadata Collection metadata properties
type CollectionMetadata struct {
	// The resource collection `id`
	Id *string `json:"id,omitempty" validate:"regexp=col[0-9a-zA-Z]+"`
	// The name of a resource collection
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionMetadata CollectionMetadata

// NewCollectionMetadata instantiates a new CollectionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionMetadata() *CollectionMetadata {
	this := CollectionMetadata{}
	return &this
}

// NewCollectionMetadataWithDefaults instantiates a new CollectionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionMetadataWithDefaults() *CollectionMetadata {
	this := CollectionMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CollectionMetadata) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadata) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CollectionMetadata) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CollectionMetadata) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectionMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectionMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectionMetadata) SetName(v string) {
	o.Name = &v
}

func (o CollectionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionMetadata) UnmarshalJSON(data []byte) (err error) {
	varCollectionMetadata := _CollectionMetadata{}

	err = json.Unmarshal(data, &varCollectionMetadata)

	if err != nil {
		return err
	}

	*o = CollectionMetadata(varCollectionMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionMetadata struct {
	value *CollectionMetadata
	isSet bool
}

func (v NullableCollectionMetadata) Get() *CollectionMetadata {
	return v.value
}

func (v *NullableCollectionMetadata) Set(val *CollectionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionMetadata(val *CollectionMetadata) *NullableCollectionMetadata {
	return &NullableCollectionMetadata{value: val, isSet: true}
}

func (v NullableCollectionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
