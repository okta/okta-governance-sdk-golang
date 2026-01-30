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

// checks if the GrantMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantMetadata{}

// GrantMetadata Grant metadata properties
type GrantMetadata struct {
	Collection           *CollectionMetadata `json:"collection,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantMetadata GrantMetadata

// NewGrantMetadata instantiates a new GrantMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantMetadata() *GrantMetadata {
	this := GrantMetadata{}
	return &this
}

// NewGrantMetadataWithDefaults instantiates a new GrantMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantMetadataWithDefaults() *GrantMetadata {
	this := GrantMetadata{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *GrantMetadata) GetCollection() CollectionMetadata {
	if o == nil || IsNil(o.Collection) {
		var ret CollectionMetadata
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantMetadata) GetCollectionOk() (*CollectionMetadata, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *GrantMetadata) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given CollectionMetadata and assigns it to the Collection field.
func (o *GrantMetadata) SetCollection(v CollectionMetadata) {
	o.Collection = &v
}

func (o GrantMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantMetadata) UnmarshalJSON(data []byte) (err error) {
	varGrantMetadata := _GrantMetadata{}

	err = json.Unmarshal(data, &varGrantMetadata)

	if err != nil {
		return err
	}

	*o = GrantMetadata(varGrantMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "collection")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantMetadata struct {
	value *GrantMetadata
	isSet bool
}

func (v NullableGrantMetadata) Get() *GrantMetadata {
	return v.value
}

func (v *NullableGrantMetadata) Set(val *GrantMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantMetadata(val *GrantMetadata) *NullableGrantMetadata {
	return &NullableGrantMetadata{value: val, isSet: true}
}

func (v NullableGrantMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
