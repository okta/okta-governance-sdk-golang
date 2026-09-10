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

// checks if the ResourceAssetTypesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetTypesList{}

// ResourceAssetTypesList struct for ResourceAssetTypesList
type ResourceAssetTypesList struct {
	// List of resource asset types for the resource
	Data                 []ResourceAssetType `json:"data"`
	Links                ListLinks           `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetTypesList ResourceAssetTypesList

// NewResourceAssetTypesList instantiates a new ResourceAssetTypesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetTypesList(data []ResourceAssetType, links ListLinks) *ResourceAssetTypesList {
	this := ResourceAssetTypesList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewResourceAssetTypesListWithDefaults instantiates a new ResourceAssetTypesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetTypesListWithDefaults() *ResourceAssetTypesList {
	this := ResourceAssetTypesList{}
	return &this
}

// GetData returns the Data field value
func (o *ResourceAssetTypesList) GetData() []ResourceAssetType {
	if o == nil {
		var ret []ResourceAssetType
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetTypesList) GetDataOk() ([]ResourceAssetType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ResourceAssetTypesList) SetData(v []ResourceAssetType) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ResourceAssetTypesList) GetLinks() ListLinks {
	if o == nil {
		var ret ListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetTypesList) GetLinksOk() (*ListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ResourceAssetTypesList) SetLinks(v ListLinks) {
	o.Links = v
}

func (o ResourceAssetTypesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetTypesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetTypesList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
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

	varResourceAssetTypesList := _ResourceAssetTypesList{}

	err = json.Unmarshal(data, &varResourceAssetTypesList)

	if err != nil {
		return err
	}

	*o = ResourceAssetTypesList(varResourceAssetTypesList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetTypesList struct {
	value *ResourceAssetTypesList
	isSet bool
}

func (v NullableResourceAssetTypesList) Get() *ResourceAssetTypesList {
	return v.value
}

func (v *NullableResourceAssetTypesList) Set(val *ResourceAssetTypesList) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetTypesList) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetTypesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetTypesList(val *ResourceAssetTypesList) *NullableResourceAssetTypesList {
	return &NullableResourceAssetTypesList{value: val, isSet: true}
}

func (v NullableResourceAssetTypesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetTypesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
