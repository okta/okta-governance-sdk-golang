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

// checks if the CollectionPropertiesWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionPropertiesWritable{}

// CollectionPropertiesWritable struct for CollectionPropertiesWritable
type CollectionPropertiesWritable struct {
	// The name of a resource collection
	Name *string `json:"name,omitempty"`
	// The human-readable description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionPropertiesWritable CollectionPropertiesWritable

// NewCollectionPropertiesWritable instantiates a new CollectionPropertiesWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionPropertiesWritable() *CollectionPropertiesWritable {
	this := CollectionPropertiesWritable{}
	return &this
}

// NewCollectionPropertiesWritableWithDefaults instantiates a new CollectionPropertiesWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionPropertiesWritableWithDefaults() *CollectionPropertiesWritable {
	this := CollectionPropertiesWritable{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectionPropertiesWritable) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPropertiesWritable) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectionPropertiesWritable) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectionPropertiesWritable) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionPropertiesWritable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPropertiesWritable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionPropertiesWritable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionPropertiesWritable) SetDescription(v string) {
	o.Description = &v
}

func (o CollectionPropertiesWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionPropertiesWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionPropertiesWritable) UnmarshalJSON(data []byte) (err error) {
	varCollectionPropertiesWritable := _CollectionPropertiesWritable{}

	err = json.Unmarshal(data, &varCollectionPropertiesWritable)

	if err != nil {
		return err
	}

	*o = CollectionPropertiesWritable(varCollectionPropertiesWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionPropertiesWritable struct {
	value *CollectionPropertiesWritable
	isSet bool
}

func (v NullableCollectionPropertiesWritable) Get() *CollectionPropertiesWritable {
	return v.value
}

func (v *NullableCollectionPropertiesWritable) Set(val *CollectionPropertiesWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPropertiesWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPropertiesWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPropertiesWritable(val *CollectionPropertiesWritable) *NullableCollectionPropertiesWritable {
	return &NullableCollectionPropertiesWritable{value: val, isSet: true}
}

func (v NullableCollectionPropertiesWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPropertiesWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
