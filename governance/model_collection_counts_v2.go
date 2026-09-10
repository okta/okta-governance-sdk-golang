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

// checks if the CollectionCountsV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionCountsV2{}

// CollectionCountsV2 Collection count metadata
type CollectionCountsV2 struct {
	// Number of principals assigned
	PrincipalAssignmentCount int32 `json:"principalAssignmentCount"`
	// The total number of resources in this collection, across all resource types (apps, groups, and push groups).
	ResourceCount        int32 `json:"resourceCount"`
	AdditionalProperties map[string]interface{}
}

type _CollectionCountsV2 CollectionCountsV2

// NewCollectionCountsV2 instantiates a new CollectionCountsV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCountsV2(principalAssignmentCount int32, resourceCount int32) *CollectionCountsV2 {
	this := CollectionCountsV2{}
	this.PrincipalAssignmentCount = principalAssignmentCount
	this.ResourceCount = resourceCount
	return &this
}

// NewCollectionCountsV2WithDefaults instantiates a new CollectionCountsV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCountsV2WithDefaults() *CollectionCountsV2 {
	this := CollectionCountsV2{}
	return &this
}

// GetPrincipalAssignmentCount returns the PrincipalAssignmentCount field value
func (o *CollectionCountsV2) GetPrincipalAssignmentCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrincipalAssignmentCount
}

// GetPrincipalAssignmentCountOk returns a tuple with the PrincipalAssignmentCount field value
// and a boolean to check if the value has been set.
func (o *CollectionCountsV2) GetPrincipalAssignmentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalAssignmentCount, true
}

// SetPrincipalAssignmentCount sets field value
func (o *CollectionCountsV2) SetPrincipalAssignmentCount(v int32) {
	o.PrincipalAssignmentCount = v
}

// GetResourceCount returns the ResourceCount field value
func (o *CollectionCountsV2) GetResourceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value
// and a boolean to check if the value has been set.
func (o *CollectionCountsV2) GetResourceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceCount, true
}

// SetResourceCount sets field value
func (o *CollectionCountsV2) SetResourceCount(v int32) {
	o.ResourceCount = v
}

func (o CollectionCountsV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionCountsV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalAssignmentCount"] = o.PrincipalAssignmentCount
	toSerialize["resourceCount"] = o.ResourceCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionCountsV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalAssignmentCount",
		"resourceCount",
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

	varCollectionCountsV2 := _CollectionCountsV2{}

	err = json.Unmarshal(data, &varCollectionCountsV2)

	if err != nil {
		return err
	}

	*o = CollectionCountsV2(varCollectionCountsV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalAssignmentCount")
		delete(additionalProperties, "resourceCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionCountsV2 struct {
	value *CollectionCountsV2
	isSet bool
}

func (v NullableCollectionCountsV2) Get() *CollectionCountsV2 {
	return v.value
}

func (v *NullableCollectionCountsV2) Set(val *CollectionCountsV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCountsV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCountsV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCountsV2(val *CollectionCountsV2) *NullableCollectionCountsV2 {
	return &NullableCollectionCountsV2{value: val, isSet: true}
}

func (v NullableCollectionCountsV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCountsV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
