/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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

// RequestConditionsList struct for RequestConditionsList
type RequestConditionsList struct {
	// All request conditions
	Data                 []RequestConditionSparse `json:"data"`
	Links                RequestConditionsLinks   `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionsList RequestConditionsList

// NewRequestConditionsList instantiates a new RequestConditionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionsList(data []RequestConditionSparse, links RequestConditionsLinks) *RequestConditionsList {
	this := RequestConditionsList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewRequestConditionsListWithDefaults instantiates a new RequestConditionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionsListWithDefaults() *RequestConditionsList {
	this := RequestConditionsList{}
	return &this
}

// GetData returns the Data field value
func (o *RequestConditionsList) GetData() []RequestConditionSparse {
	if o == nil {
		var ret []RequestConditionSparse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RequestConditionsList) GetDataOk() ([]RequestConditionSparse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *RequestConditionsList) SetData(v []RequestConditionSparse) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *RequestConditionsList) GetLinks() RequestConditionsLinks {
	if o == nil {
		var ret RequestConditionsLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestConditionsList) GetLinksOk() (*RequestConditionsLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestConditionsList) SetLinks(v RequestConditionsLinks) {
	o.Links = v
}

func (o RequestConditionsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestConditionsList) UnmarshalJSON(bytes []byte) (err error) {
	varRequestConditionsList := _RequestConditionsList{}

	err = json.Unmarshal(bytes, &varRequestConditionsList)
	if err == nil {
		*o = RequestConditionsList(varRequestConditionsList)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestConditionsList struct {
	value *RequestConditionsList
	isSet bool
}

func (v NullableRequestConditionsList) Get() *RequestConditionsList {
	return v.value
}

func (v *NullableRequestConditionsList) Set(val *RequestConditionsList) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionsList) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionsList(val *RequestConditionsList) *NullableRequestConditionsList {
	return &NullableRequestConditionsList{value: val, isSet: true}
}

func (v NullableRequestConditionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
