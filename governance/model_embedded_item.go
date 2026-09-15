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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the EmbeddedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedItem{}

// EmbeddedItem struct for EmbeddedItem
type EmbeddedItem struct {
	// Unique identifier for the embedded resource
	Id string `json:"id"`
	// Display name of the embedded resource
	Name string `json:"name"`
	// Additional metadata specific to this embedded resource type
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _EmbeddedItem EmbeddedItem

// NewEmbeddedItem instantiates a new EmbeddedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedItem(id string, name string) *EmbeddedItem {
	this := EmbeddedItem{}
	this.Id = id
	this.Name = name
	return &this
}

// NewEmbeddedItemWithDefaults instantiates a new EmbeddedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedItemWithDefaults() *EmbeddedItem {
	this := EmbeddedItem{}
	return &this
}

// GetId returns the Id field value
func (o *EmbeddedItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmbeddedItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmbeddedItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EmbeddedItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmbeddedItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmbeddedItem) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EmbeddedItem) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedItem) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EmbeddedItem) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *EmbeddedItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o EmbeddedItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *EmbeddedItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varEmbeddedItem := _EmbeddedItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbeddedItem)

	if err != nil {
		return err
	}

	*o = EmbeddedItem(varEmbeddedItem)

	return err
}

type NullableEmbeddedItem struct {
	value *EmbeddedItem
	isSet bool
}

func (v NullableEmbeddedItem) Get() *EmbeddedItem {
	return v.value
}

func (v *NullableEmbeddedItem) Set(val *EmbeddedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedItem(val *EmbeddedItem) *NullableEmbeddedItem {
	return &NullableEmbeddedItem{value: val, isSet: true}
}

func (v NullableEmbeddedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
