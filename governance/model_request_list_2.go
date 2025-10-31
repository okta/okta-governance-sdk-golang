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

// checks if the RequestList2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestList2{}

// RequestList2 struct for RequestList2
type RequestList2 struct {
	// All requests on the current page
	Data                 []RequestSparse2   `json:"data,omitempty"`
	Links                *RequestListLinks2 `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestList2 RequestList2

// NewRequestList2 instantiates a new RequestList2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestList2() *RequestList2 {
	this := RequestList2{}
	return &this
}

// NewRequestList2WithDefaults instantiates a new RequestList2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestList2WithDefaults() *RequestList2 {
	this := RequestList2{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RequestList2) GetData() []RequestSparse2 {
	if o == nil || IsNil(o.Data) {
		var ret []RequestSparse2
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestList2) GetDataOk() ([]RequestSparse2, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestList2) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RequestSparse2 and assigns it to the Data field.
func (o *RequestList2) SetData(v []RequestSparse2) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RequestList2) GetLinks() RequestListLinks2 {
	if o == nil || IsNil(o.Links) {
		var ret RequestListLinks2
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestList2) GetLinksOk() (*RequestListLinks2, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RequestList2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RequestListLinks2 and assigns it to the Links field.
func (o *RequestList2) SetLinks(v RequestListLinks2) {
	o.Links = &v
}

func (o RequestList2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestList2) ToMap() (map[string]interface{}, error) {
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

func (o *RequestList2) UnmarshalJSON(data []byte) (err error) {
	varRequestList2 := _RequestList2{}

	err = json.Unmarshal(data, &varRequestList2)

	if err != nil {
		return err
	}

	*o = RequestList2(varRequestList2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestList2 struct {
	value *RequestList2
	isSet bool
}

func (v NullableRequestList2) Get() *RequestList2 {
	return v.value
}

func (v *NullableRequestList2) Set(val *RequestList2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestList2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestList2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestList2(val *RequestList2) *NullableRequestList2 {
	return &NullableRequestList2{value: val, isSet: true}
}

func (v NullableRequestList2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestList2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
