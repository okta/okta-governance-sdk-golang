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

// checks if the RequestTypesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypesList{}

// RequestTypesList struct for RequestTypesList
type RequestTypesList struct {
	// All Request Types on the current page
	Data                 []RequestTypeSparse    `json:"data,omitempty"`
	Links                *RequestTypesListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypesList RequestTypesList

// NewRequestTypesList instantiates a new RequestTypesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypesList() *RequestTypesList {
	this := RequestTypesList{}
	return &this
}

// NewRequestTypesListWithDefaults instantiates a new RequestTypesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypesListWithDefaults() *RequestTypesList {
	this := RequestTypesList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RequestTypesList) GetData() []RequestTypeSparse {
	if o == nil || IsNil(o.Data) {
		var ret []RequestTypeSparse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypesList) GetDataOk() ([]RequestTypeSparse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestTypesList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RequestTypeSparse and assigns it to the Data field.
func (o *RequestTypesList) SetData(v []RequestTypeSparse) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RequestTypesList) GetLinks() RequestTypesListLinks {
	if o == nil || IsNil(o.Links) {
		var ret RequestTypesListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypesList) GetLinksOk() (*RequestTypesListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RequestTypesList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RequestTypesListLinks and assigns it to the Links field.
func (o *RequestTypesList) SetLinks(v RequestTypesListLinks) {
	o.Links = &v
}

func (o RequestTypesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypesList) ToMap() (map[string]interface{}, error) {
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

func (o *RequestTypesList) UnmarshalJSON(data []byte) (err error) {
	varRequestTypesList := _RequestTypesList{}

	err = json.Unmarshal(data, &varRequestTypesList)

	if err != nil {
		return err
	}

	*o = RequestTypesList(varRequestTypesList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypesList struct {
	value *RequestTypesList
	isSet bool
}

func (v NullableRequestTypesList) Get() *RequestTypesList {
	return v.value
}

func (v *NullableRequestTypesList) Set(val *RequestTypesList) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypesList) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypesList(val *RequestTypesList) *NullableRequestTypesList {
	return &NullableRequestTypesList{value: val, isSet: true}
}

func (v NullableRequestTypesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
