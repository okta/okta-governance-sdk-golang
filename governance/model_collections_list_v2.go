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

// checks if the CollectionsListV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionsListV2{}

// CollectionsListV2 struct for CollectionsListV2
type CollectionsListV2 struct {
	Data                 []CollectionFullWithFilterContextV2 `json:"data,omitempty"`
	Links                *ListLinks                          `json:"_links,omitempty"`
	Metadata             *ListMetadata                       `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionsListV2 CollectionsListV2

// NewCollectionsListV2 instantiates a new CollectionsListV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionsListV2() *CollectionsListV2 {
	this := CollectionsListV2{}
	return &this
}

// NewCollectionsListV2WithDefaults instantiates a new CollectionsListV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionsListV2WithDefaults() *CollectionsListV2 {
	this := CollectionsListV2{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CollectionsListV2) GetData() []CollectionFullWithFilterContextV2 {
	if o == nil || IsNil(o.Data) {
		var ret []CollectionFullWithFilterContextV2
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionsListV2) GetDataOk() ([]CollectionFullWithFilterContextV2, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CollectionsListV2) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CollectionFullWithFilterContextV2 and assigns it to the Data field.
func (o *CollectionsListV2) SetData(v []CollectionFullWithFilterContextV2) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionsListV2) GetLinks() ListLinks {
	if o == nil || IsNil(o.Links) {
		var ret ListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionsListV2) GetLinksOk() (*ListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionsListV2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListLinks and assigns it to the Links field.
func (o *CollectionsListV2) SetLinks(v ListLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CollectionsListV2) GetMetadata() ListMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionsListV2) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CollectionsListV2) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *CollectionsListV2) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o CollectionsListV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionsListV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionsListV2) UnmarshalJSON(data []byte) (err error) {
	varCollectionsListV2 := _CollectionsListV2{}

	err = json.Unmarshal(data, &varCollectionsListV2)

	if err != nil {
		return err
	}

	*o = CollectionsListV2(varCollectionsListV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionsListV2 struct {
	value *CollectionsListV2
	isSet bool
}

func (v NullableCollectionsListV2) Get() *CollectionsListV2 {
	return v.value
}

func (v *NullableCollectionsListV2) Set(val *CollectionsListV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionsListV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionsListV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionsListV2(val *CollectionsListV2) *NullableCollectionsListV2 {
	return &NullableCollectionsListV2{value: val, isSet: true}
}

func (v NullableCollectionsListV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionsListV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
