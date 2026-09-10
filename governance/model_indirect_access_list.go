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

// checks if the IndirectAccessList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndirectAccessList{}

// IndirectAccessList Paginated list of indirect access
type IndirectAccessList struct {
	// Array of indirect access
	Data                 []IndirectAccess `json:"data,omitempty"`
	Links                *ListLinks       `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndirectAccessList IndirectAccessList

// NewIndirectAccessList instantiates a new IndirectAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndirectAccessList() *IndirectAccessList {
	this := IndirectAccessList{}
	return &this
}

// NewIndirectAccessListWithDefaults instantiates a new IndirectAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndirectAccessListWithDefaults() *IndirectAccessList {
	this := IndirectAccessList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IndirectAccessList) GetData() []IndirectAccess {
	if o == nil || IsNil(o.Data) {
		var ret []IndirectAccess
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectAccessList) GetDataOk() ([]IndirectAccess, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IndirectAccessList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []IndirectAccess and assigns it to the Data field.
func (o *IndirectAccessList) SetData(v []IndirectAccess) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IndirectAccessList) GetLinks() ListLinks {
	if o == nil || IsNil(o.Links) {
		var ret ListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectAccessList) GetLinksOk() (*ListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IndirectAccessList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListLinks and assigns it to the Links field.
func (o *IndirectAccessList) SetLinks(v ListLinks) {
	o.Links = &v
}

func (o IndirectAccessList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndirectAccessList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndirectAccessList) UnmarshalJSON(data []byte) (err error) {
	varIndirectAccessList := _IndirectAccessList{}

	err = json.Unmarshal(data, &varIndirectAccessList)

	if err != nil {
		return err
	}

	*o = IndirectAccessList(varIndirectAccessList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndirectAccessList struct {
	value *IndirectAccessList
	isSet bool
}

func (v NullableIndirectAccessList) Get() *IndirectAccessList {
	return v.value
}

func (v *NullableIndirectAccessList) Set(val *IndirectAccessList) {
	v.value = val
	v.isSet = true
}

func (v NullableIndirectAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullableIndirectAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndirectAccessList(val *IndirectAccessList) *NullableIndirectAccessList {
	return &NullableIndirectAccessList{value: val, isSet: true}
}

func (v NullableIndirectAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndirectAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
