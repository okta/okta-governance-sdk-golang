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

// checks if the ResourceAssetTypeSparse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetTypeSparse{}

// ResourceAssetTypeSparse The type of the resource asset
type ResourceAssetTypeSparse struct {
	// The unique identifier for the resource asset type
	Id string `json:"id"`
	// The display name of the resource asset type
	DisplayName          string `json:"displayName"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetTypeSparse ResourceAssetTypeSparse

// NewResourceAssetTypeSparse instantiates a new ResourceAssetTypeSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetTypeSparse(id string, displayName string) *ResourceAssetTypeSparse {
	this := ResourceAssetTypeSparse{}
	this.Id = id
	this.DisplayName = displayName
	return &this
}

// NewResourceAssetTypeSparseWithDefaults instantiates a new ResourceAssetTypeSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetTypeSparseWithDefaults() *ResourceAssetTypeSparse {
	this := ResourceAssetTypeSparse{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceAssetTypeSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetTypeSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceAssetTypeSparse) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *ResourceAssetTypeSparse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetTypeSparse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ResourceAssetTypeSparse) SetDisplayName(v string) {
	o.DisplayName = v
}

func (o ResourceAssetTypeSparse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetTypeSparse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetTypeSparse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"displayName",
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

	varResourceAssetTypeSparse := _ResourceAssetTypeSparse{}

	err = json.Unmarshal(data, &varResourceAssetTypeSparse)

	if err != nil {
		return err
	}

	*o = ResourceAssetTypeSparse(varResourceAssetTypeSparse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetTypeSparse struct {
	value *ResourceAssetTypeSparse
	isSet bool
}

func (v NullableResourceAssetTypeSparse) Get() *ResourceAssetTypeSparse {
	return v.value
}

func (v *NullableResourceAssetTypeSparse) Set(val *ResourceAssetTypeSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetTypeSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetTypeSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetTypeSparse(val *ResourceAssetTypeSparse) *NullableResourceAssetTypeSparse {
	return &NullableResourceAssetTypeSparse{value: val, isSet: true}
}

func (v NullableResourceAssetTypeSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetTypeSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
