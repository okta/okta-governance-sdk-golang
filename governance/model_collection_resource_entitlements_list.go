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

// checks if the CollectionResourceEntitlementsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceEntitlementsList{}

// CollectionResourceEntitlementsList struct for CollectionResourceEntitlementsList
type CollectionResourceEntitlementsList struct {
	Data                 []EntitlementFull `json:"data,omitempty"`
	Links                *ListLinks        `json:"_links,omitempty"`
	Metadata             *ListMetadata     `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceEntitlementsList CollectionResourceEntitlementsList

// NewCollectionResourceEntitlementsList instantiates a new CollectionResourceEntitlementsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceEntitlementsList() *CollectionResourceEntitlementsList {
	this := CollectionResourceEntitlementsList{}
	return &this
}

// NewCollectionResourceEntitlementsListWithDefaults instantiates a new CollectionResourceEntitlementsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceEntitlementsListWithDefaults() *CollectionResourceEntitlementsList {
	this := CollectionResourceEntitlementsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CollectionResourceEntitlementsList) GetData() []EntitlementFull {
	if o == nil || IsNil(o.Data) {
		var ret []EntitlementFull
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceEntitlementsList) GetDataOk() ([]EntitlementFull, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CollectionResourceEntitlementsList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EntitlementFull and assigns it to the Data field.
func (o *CollectionResourceEntitlementsList) SetData(v []EntitlementFull) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionResourceEntitlementsList) GetLinks() ListLinks {
	if o == nil || IsNil(o.Links) {
		var ret ListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceEntitlementsList) GetLinksOk() (*ListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionResourceEntitlementsList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListLinks and assigns it to the Links field.
func (o *CollectionResourceEntitlementsList) SetLinks(v ListLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CollectionResourceEntitlementsList) GetMetadata() ListMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceEntitlementsList) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CollectionResourceEntitlementsList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *CollectionResourceEntitlementsList) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o CollectionResourceEntitlementsList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceEntitlementsList) ToMap() (map[string]interface{}, error) {
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

func (o *CollectionResourceEntitlementsList) UnmarshalJSON(data []byte) (err error) {
	varCollectionResourceEntitlementsList := _CollectionResourceEntitlementsList{}

	err = json.Unmarshal(data, &varCollectionResourceEntitlementsList)

	if err != nil {
		return err
	}

	*o = CollectionResourceEntitlementsList(varCollectionResourceEntitlementsList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceEntitlementsList struct {
	value *CollectionResourceEntitlementsList
	isSet bool
}

func (v NullableCollectionResourceEntitlementsList) Get() *CollectionResourceEntitlementsList {
	return v.value
}

func (v *NullableCollectionResourceEntitlementsList) Set(val *CollectionResourceEntitlementsList) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceEntitlementsList) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceEntitlementsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceEntitlementsList(val *CollectionResourceEntitlementsList) *NullableCollectionResourceEntitlementsList {
	return &NullableCollectionResourceEntitlementsList{value: val, isSet: true}
}

func (v NullableCollectionResourceEntitlementsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceEntitlementsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
