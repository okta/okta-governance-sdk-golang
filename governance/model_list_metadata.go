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

// checks if the ListMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetadata{}

// ListMetadata Metadata for the list response
type ListMetadata struct {
	// Total number of objects
	Total                *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListMetadata ListMetadata

// NewListMetadata instantiates a new ListMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetadata() *ListMetadata {
	this := ListMetadata{}
	return &this
}

// NewListMetadataWithDefaults instantiates a new ListMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetadataWithDefaults() *ListMetadata {
	this := ListMetadata{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ListMetadata) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetadata) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ListMetadata) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ListMetadata) SetTotal(v int32) {
	o.Total = &v
}

func (o ListMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListMetadata) UnmarshalJSON(data []byte) (err error) {
	varListMetadata := _ListMetadata{}

	err = json.Unmarshal(data, &varListMetadata)

	if err != nil {
		return err
	}

	*o = ListMetadata(varListMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListMetadata struct {
	value *ListMetadata
	isSet bool
}

func (v NullableListMetadata) Get() *ListMetadata {
	return v.value
}

func (v *NullableListMetadata) Set(val *ListMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetadata(val *ListMetadata) *NullableListMetadata {
	return &NullableListMetadata{value: val, isSet: true}
}

func (v NullableListMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
