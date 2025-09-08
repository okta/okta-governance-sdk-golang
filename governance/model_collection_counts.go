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

// CollectionCounts Collection count metadata
type CollectionCounts struct {
	// Number of principals assigned
	PrincipalAssignmentCount int32                          `json:"principalAssignmentCount"`
	ResourceCounts           CollectionCountsResourceCounts `json:"resourceCounts"`
	AdditionalProperties     map[string]interface{}
}

type _CollectionCounts CollectionCounts

// NewCollectionCounts instantiates a new CollectionCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCounts(principalAssignmentCount int32, resourceCounts CollectionCountsResourceCounts) *CollectionCounts {
	this := CollectionCounts{}
	this.PrincipalAssignmentCount = principalAssignmentCount
	this.ResourceCounts = resourceCounts
	return &this
}

// NewCollectionCountsWithDefaults instantiates a new CollectionCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCountsWithDefaults() *CollectionCounts {
	this := CollectionCounts{}
	return &this
}

// GetPrincipalAssignmentCount returns the PrincipalAssignmentCount field value
func (o *CollectionCounts) GetPrincipalAssignmentCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrincipalAssignmentCount
}

// GetPrincipalAssignmentCountOk returns a tuple with the PrincipalAssignmentCount field value
// and a boolean to check if the value has been set.
func (o *CollectionCounts) GetPrincipalAssignmentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalAssignmentCount, true
}

// SetPrincipalAssignmentCount sets field value
func (o *CollectionCounts) SetPrincipalAssignmentCount(v int32) {
	o.PrincipalAssignmentCount = v
}

// GetResourceCounts returns the ResourceCounts field value
func (o *CollectionCounts) GetResourceCounts() CollectionCountsResourceCounts {
	if o == nil {
		var ret CollectionCountsResourceCounts
		return ret
	}

	return o.ResourceCounts
}

// GetResourceCountsOk returns a tuple with the ResourceCounts field value
// and a boolean to check if the value has been set.
func (o *CollectionCounts) GetResourceCountsOk() (*CollectionCountsResourceCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceCounts, true
}

// SetResourceCounts sets field value
func (o *CollectionCounts) SetResourceCounts(v CollectionCountsResourceCounts) {
	o.ResourceCounts = v
}

func (o CollectionCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["principalAssignmentCount"] = o.PrincipalAssignmentCount
	}
	if true {
		toSerialize["resourceCounts"] = o.ResourceCounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionCounts) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionCounts := _CollectionCounts{}

	err = json.Unmarshal(bytes, &varCollectionCounts)
	if err == nil {
		*o = CollectionCounts(varCollectionCounts)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "principalAssignmentCount")
		delete(additionalProperties, "resourceCounts")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCollectionCounts struct {
	value *CollectionCounts
	isSet bool
}

func (v NullableCollectionCounts) Get() *CollectionCounts {
	return v.value
}

func (v *NullableCollectionCounts) Set(val *CollectionCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCounts(val *CollectionCounts) *NullableCollectionCounts {
	return &NullableCollectionCounts{value: val, isSet: true}
}

func (v NullableCollectionCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
