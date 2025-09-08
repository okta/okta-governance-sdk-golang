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

// ListResourceLabels struct for ListResourceLabels
type ListResourceLabels struct {
	// Resources with labels
	Data                 []ResourceLabel `json:"data"`
	Links                ListLinks       `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ListResourceLabels ListResourceLabels

// NewListResourceLabels instantiates a new ListResourceLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResourceLabels(data []ResourceLabel, links ListLinks) *ListResourceLabels {
	this := ListResourceLabels{}
	this.Data = data
	this.Links = links
	return &this
}

// NewListResourceLabelsWithDefaults instantiates a new ListResourceLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResourceLabelsWithDefaults() *ListResourceLabels {
	this := ListResourceLabels{}
	return &this
}

// GetData returns the Data field value
func (o *ListResourceLabels) GetData() []ResourceLabel {
	if o == nil {
		var ret []ResourceLabel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResourceLabels) GetDataOk() ([]ResourceLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResourceLabels) SetData(v []ResourceLabel) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ListResourceLabels) GetLinks() ListLinks {
	if o == nil {
		var ret ListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListResourceLabels) GetLinksOk() (*ListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ListResourceLabels) SetLinks(v ListLinks) {
	o.Links = v
}

func (o ListResourceLabels) MarshalJSON() ([]byte, error) {
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

func (o *ListResourceLabels) UnmarshalJSON(bytes []byte) (err error) {
	varListResourceLabels := _ListResourceLabels{}

	err = json.Unmarshal(bytes, &varListResourceLabels)
	if err == nil {
		*o = ListResourceLabels(varListResourceLabels)
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

type NullableListResourceLabels struct {
	value *ListResourceLabels
	isSet bool
}

func (v NullableListResourceLabels) Get() *ListResourceLabels {
	return v.value
}

func (v *NullableListResourceLabels) Set(val *ListResourceLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableListResourceLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableListResourceLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResourceLabels(val *ListResourceLabels) *NullableListResourceLabels {
	return &NullableListResourceLabels{value: val, isSet: true}
}

func (v NullableListResourceLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResourceLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
