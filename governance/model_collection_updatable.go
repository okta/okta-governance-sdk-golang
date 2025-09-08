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

// CollectionUpdatable Full representation of a collection to be updated
type CollectionUpdatable struct {
	// The name of a resource collection
	Name string `json:"name"`
	// The human-readable description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionUpdatable CollectionUpdatable

// NewCollectionUpdatable instantiates a new CollectionUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionUpdatable(name string) *CollectionUpdatable {
	this := CollectionUpdatable{}
	this.Name = name
	return &this
}

// NewCollectionUpdatableWithDefaults instantiates a new CollectionUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionUpdatableWithDefaults() *CollectionUpdatable {
	this := CollectionUpdatable{}
	return &this
}

// GetName returns the Name field value
func (o *CollectionUpdatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CollectionUpdatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CollectionUpdatable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionUpdatable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionUpdatable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionUpdatable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionUpdatable) SetDescription(v string) {
	o.Description = &v
}

func (o CollectionUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionUpdatable) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionUpdatable := _CollectionUpdatable{}

	err = json.Unmarshal(bytes, &varCollectionUpdatable)
	if err == nil {
		*o = CollectionUpdatable(varCollectionUpdatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCollectionUpdatable struct {
	value *CollectionUpdatable
	isSet bool
}

func (v NullableCollectionUpdatable) Get() *CollectionUpdatable {
	return v.value
}

func (v *NullableCollectionUpdatable) Set(val *CollectionUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionUpdatable(val *CollectionUpdatable) *NullableCollectionUpdatable {
	return &NullableCollectionUpdatable{value: val, isSet: true}
}

func (v NullableCollectionUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
