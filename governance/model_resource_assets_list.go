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

// checks if the ResourceAssetsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetsList{}

// ResourceAssetsList struct for ResourceAssetsList
type ResourceAssetsList struct {
	// List of all resource assets matching the filter
	Data []ResourceAssetWithHierarchyContext `json:"data"`
	// A map of all unique ancestor assets for items in `data`, keyed by asset ID. Only present if ?include=ancestors is specified. Each ancestor also includes its own `parentId` to allow full chain traversal to the root.
	Ancestors            *map[string]ResourceAssetAncestor `json:"ancestors,omitempty"`
	Links                ResourceAssetsLinks               `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetsList ResourceAssetsList

// NewResourceAssetsList instantiates a new ResourceAssetsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetsList(data []ResourceAssetWithHierarchyContext, links ResourceAssetsLinks) *ResourceAssetsList {
	this := ResourceAssetsList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewResourceAssetsListWithDefaults instantiates a new ResourceAssetsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetsListWithDefaults() *ResourceAssetsList {
	this := ResourceAssetsList{}
	return &this
}

// GetData returns the Data field value
func (o *ResourceAssetsList) GetData() []ResourceAssetWithHierarchyContext {
	if o == nil {
		var ret []ResourceAssetWithHierarchyContext
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetsList) GetDataOk() ([]ResourceAssetWithHierarchyContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ResourceAssetsList) SetData(v []ResourceAssetWithHierarchyContext) {
	o.Data = v
}

// GetAncestors returns the Ancestors field value if set, zero value otherwise.
func (o *ResourceAssetsList) GetAncestors() map[string]ResourceAssetAncestor {
	if o == nil || IsNil(o.Ancestors) {
		var ret map[string]ResourceAssetAncestor
		return ret
	}
	return *o.Ancestors
}

// GetAncestorsOk returns a tuple with the Ancestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetsList) GetAncestorsOk() (*map[string]ResourceAssetAncestor, bool) {
	if o == nil || IsNil(o.Ancestors) {
		return nil, false
	}
	return o.Ancestors, true
}

// HasAncestors returns a boolean if a field has been set.
func (o *ResourceAssetsList) HasAncestors() bool {
	if o != nil && !IsNil(o.Ancestors) {
		return true
	}

	return false
}

// SetAncestors gets a reference to the given map[string]ResourceAssetAncestor and assigns it to the Ancestors field.
func (o *ResourceAssetsList) SetAncestors(v map[string]ResourceAssetAncestor) {
	o.Ancestors = &v
}

// GetLinks returns the Links field value
func (o *ResourceAssetsList) GetLinks() ResourceAssetsLinks {
	if o == nil {
		var ret ResourceAssetsLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetsList) GetLinksOk() (*ResourceAssetsLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ResourceAssetsList) SetLinks(v ResourceAssetsLinks) {
	o.Links = v
}

func (o ResourceAssetsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Ancestors) {
		toSerialize["ancestors"] = o.Ancestors
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetsList) UnmarshalJSON(data []byte) (err error) {
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

	varResourceAssetsList := _ResourceAssetsList{}

	err = json.Unmarshal(data, &varResourceAssetsList)

	if err != nil {
		return err
	}

	*o = ResourceAssetsList(varResourceAssetsList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "ancestors")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetsList struct {
	value *ResourceAssetsList
	isSet bool
}

func (v NullableResourceAssetsList) Get() *ResourceAssetsList {
	return v.value
}

func (v *NullableResourceAssetsList) Set(val *ResourceAssetsList) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetsList) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetsList(val *ResourceAssetsList) *NullableResourceAssetsList {
	return &NullableResourceAssetsList{value: val, isSet: true}
}

func (v NullableResourceAssetsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
