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

// checks if the CollectionCountsResourceCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionCountsResourceCounts{}

// CollectionCountsResourceCounts Collection resource counts
type CollectionCountsResourceCounts struct {
	// Number of app resources in a collection
	Applications         int32 `json:"applications"`
	AdditionalProperties map[string]interface{}
}

type _CollectionCountsResourceCounts CollectionCountsResourceCounts

// NewCollectionCountsResourceCounts instantiates a new CollectionCountsResourceCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCountsResourceCounts(applications int32) *CollectionCountsResourceCounts {
	this := CollectionCountsResourceCounts{}
	this.Applications = applications
	return &this
}

// NewCollectionCountsResourceCountsWithDefaults instantiates a new CollectionCountsResourceCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCountsResourceCountsWithDefaults() *CollectionCountsResourceCounts {
	this := CollectionCountsResourceCounts{}
	return &this
}

// GetApplications returns the Applications field value
func (o *CollectionCountsResourceCounts) GetApplications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value
// and a boolean to check if the value has been set.
func (o *CollectionCountsResourceCounts) GetApplicationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Applications, true
}

// SetApplications sets field value
func (o *CollectionCountsResourceCounts) SetApplications(v int32) {
	o.Applications = v
}

func (o CollectionCountsResourceCounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionCountsResourceCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applications"] = o.Applications

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionCountsResourceCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applications",
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

	varCollectionCountsResourceCounts := _CollectionCountsResourceCounts{}

	err = json.Unmarshal(data, &varCollectionCountsResourceCounts)

	if err != nil {
		return err
	}

	*o = CollectionCountsResourceCounts(varCollectionCountsResourceCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionCountsResourceCounts struct {
	value *CollectionCountsResourceCounts
	isSet bool
}

func (v NullableCollectionCountsResourceCounts) Get() *CollectionCountsResourceCounts {
	return v.value
}

func (v *NullableCollectionCountsResourceCounts) Set(val *CollectionCountsResourceCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCountsResourceCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCountsResourceCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCountsResourceCounts(val *CollectionCountsResourceCounts) *NullableCollectionCountsResourceCounts {
	return &NullableCollectionCountsResourceCounts{value: val, isSet: true}
}

func (v NullableCollectionCountsResourceCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCountsResourceCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
