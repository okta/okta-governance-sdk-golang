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

// checks if the RequestList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestList{}

// RequestList struct for RequestList
type RequestList struct {
	// All requests on the current page
	Data                 []RequestSparse   `json:"data,omitempty"`
	Links                *RequestListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestList RequestList

// NewRequestList instantiates a new RequestList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestList() *RequestList {
	this := RequestList{}
	return &this
}

// NewRequestListWithDefaults instantiates a new RequestList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestListWithDefaults() *RequestList {
	this := RequestList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RequestList) GetData() []RequestSparse {
	if o == nil || IsNil(o.Data) {
		var ret []RequestSparse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestList) GetDataOk() ([]RequestSparse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RequestSparse and assigns it to the Data field.
func (o *RequestList) SetData(v []RequestSparse) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RequestList) GetLinks() RequestListLinks {
	if o == nil || IsNil(o.Links) {
		var ret RequestListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestList) GetLinksOk() (*RequestListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RequestList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RequestListLinks and assigns it to the Links field.
func (o *RequestList) SetLinks(v RequestListLinks) {
	o.Links = &v
}

func (o RequestList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestList) ToMap() (map[string]interface{}, error) {
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

func (o *RequestList) UnmarshalJSON(data []byte) (err error) {
	varRequestList := _RequestList{}

	err = json.Unmarshal(data, &varRequestList)

	if err != nil {
		return err
	}

	*o = RequestList(varRequestList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestList struct {
	value *RequestList
	isSet bool
}

func (v NullableRequestList) Get() *RequestList {
	return v.value
}

func (v *NullableRequestList) Set(val *RequestList) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestList) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestList(val *RequestList) *NullableRequestList {
	return &NullableRequestList{value: val, isSet: true}
}

func (v NullableRequestList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
