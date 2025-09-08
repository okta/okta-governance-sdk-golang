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

// CollectionInfoSparse struct for CollectionInfoSparse
type CollectionInfoSparse struct {
	// collection id
	Id string `json:"id"`
	// name of the collection
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CollectionInfoSparse CollectionInfoSparse

// NewCollectionInfoSparse instantiates a new CollectionInfoSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionInfoSparse(id string, name string) *CollectionInfoSparse {
	this := CollectionInfoSparse{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCollectionInfoSparseWithDefaults instantiates a new CollectionInfoSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionInfoSparseWithDefaults() *CollectionInfoSparse {
	this := CollectionInfoSparse{}
	return &this
}

// GetId returns the Id field value
func (o *CollectionInfoSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CollectionInfoSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CollectionInfoSparse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CollectionInfoSparse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CollectionInfoSparse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CollectionInfoSparse) SetName(v string) {
	o.Name = v
}

func (o CollectionInfoSparse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionInfoSparse) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionInfoSparse := _CollectionInfoSparse{}

	err = json.Unmarshal(bytes, &varCollectionInfoSparse)
	if err == nil {
		*o = CollectionInfoSparse(varCollectionInfoSparse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCollectionInfoSparse struct {
	value *CollectionInfoSparse
	isSet bool
}

func (v NullableCollectionInfoSparse) Get() *CollectionInfoSparse {
	return v.value
}

func (v *NullableCollectionInfoSparse) Set(val *CollectionInfoSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionInfoSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionInfoSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionInfoSparse(val *CollectionInfoSparse) *NullableCollectionInfoSparse {
	return &NullableCollectionInfoSparse{value: val, isSet: true}
}

func (v NullableCollectionInfoSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionInfoSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
