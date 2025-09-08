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

// CollectionCreatable The properties expected when creating a new empty collection
type CollectionCreatable struct {
	// The name of a resource collection
	Name string `json:"name"`
	// The human-readable description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionCreatable CollectionCreatable

// NewCollectionCreatable instantiates a new CollectionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCreatable(name string) *CollectionCreatable {
	this := CollectionCreatable{}
	this.Name = name
	return &this
}

// NewCollectionCreatableWithDefaults instantiates a new CollectionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCreatableWithDefaults() *CollectionCreatable {
	this := CollectionCreatable{}
	return &this
}

// GetName returns the Name field value
func (o *CollectionCreatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CollectionCreatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CollectionCreatable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionCreatable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCreatable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionCreatable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionCreatable) SetDescription(v string) {
	o.Description = &v
}

func (o CollectionCreatable) MarshalJSON() ([]byte, error) {
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

func (o *CollectionCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionCreatable := _CollectionCreatable{}

	err = json.Unmarshal(bytes, &varCollectionCreatable)
	if err == nil {
		*o = CollectionCreatable(varCollectionCreatable)
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

type NullableCollectionCreatable struct {
	value *CollectionCreatable
	isSet bool
}

func (v NullableCollectionCreatable) Get() *CollectionCreatable {
	return v.value
}

func (v *NullableCollectionCreatable) Set(val *CollectionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCreatable(val *CollectionCreatable) *NullableCollectionCreatable {
	return &NullableCollectionCreatable{value: val, isSet: true}
}

func (v NullableCollectionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
