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

// checks if the ResourceInventoryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceInventoryList{}

// ResourceInventoryList struct for ResourceInventoryList
type ResourceInventoryList struct {
	// List of resources on the current page
	Data                 []Resource   `json:"data"`
	Links                ListLinks    `json:"_links"`
	Metadata             ListMetadata `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _ResourceInventoryList ResourceInventoryList

// NewResourceInventoryList instantiates a new ResourceInventoryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceInventoryList(data []Resource, links ListLinks, metadata ListMetadata) *ResourceInventoryList {
	this := ResourceInventoryList{}
	this.Data = data
	this.Links = links
	this.Metadata = metadata
	return &this
}

// NewResourceInventoryListWithDefaults instantiates a new ResourceInventoryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceInventoryListWithDefaults() *ResourceInventoryList {
	this := ResourceInventoryList{}
	return &this
}

// GetData returns the Data field value
func (o *ResourceInventoryList) GetData() []Resource {
	if o == nil {
		var ret []Resource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryList) GetDataOk() ([]Resource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ResourceInventoryList) SetData(v []Resource) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ResourceInventoryList) GetLinks() ListLinks {
	if o == nil {
		var ret ListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryList) GetLinksOk() (*ListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ResourceInventoryList) SetLinks(v ListLinks) {
	o.Links = v
}

// GetMetadata returns the Metadata field value
func (o *ResourceInventoryList) GetMetadata() ListMetadata {
	if o == nil {
		var ret ListMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryList) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ResourceInventoryList) SetMetadata(v ListMetadata) {
	o.Metadata = v
}

func (o ResourceInventoryList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceInventoryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceInventoryList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
		"metadata",
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

	varResourceInventoryList := _ResourceInventoryList{}

	err = json.Unmarshal(data, &varResourceInventoryList)

	if err != nil {
		return err
	}

	*o = ResourceInventoryList(varResourceInventoryList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceInventoryList struct {
	value *ResourceInventoryList
	isSet bool
}

func (v NullableResourceInventoryList) Get() *ResourceInventoryList {
	return v.value
}

func (v *NullableResourceInventoryList) Set(val *ResourceInventoryList) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceInventoryList) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceInventoryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceInventoryList(val *ResourceInventoryList) *NullableResourceInventoryList {
	return &NullableResourceInventoryList{value: val, isSet: true}
}

func (v NullableResourceInventoryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceInventoryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
