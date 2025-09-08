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

// ReviewReassignList struct for ReviewReassignList
type ReviewReassignList struct {
	// All the reviews that were reassigned
	Data                 []ReviewSparse          `json:"data"`
	Links                ReviewReassignListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ReviewReassignList ReviewReassignList

// NewReviewReassignList instantiates a new ReviewReassignList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewReassignList(data []ReviewSparse, links ReviewReassignListLinks) *ReviewReassignList {
	this := ReviewReassignList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewReviewReassignListWithDefaults instantiates a new ReviewReassignList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewReassignListWithDefaults() *ReviewReassignList {
	this := ReviewReassignList{}
	return &this
}

// GetData returns the Data field value
func (o *ReviewReassignList) GetData() []ReviewSparse {
	if o == nil {
		var ret []ReviewSparse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReviewReassignList) GetDataOk() ([]ReviewSparse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ReviewReassignList) SetData(v []ReviewSparse) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ReviewReassignList) GetLinks() ReviewReassignListLinks {
	if o == nil {
		var ret ReviewReassignListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewReassignList) GetLinksOk() (*ReviewReassignListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewReassignList) SetLinks(v ReviewReassignListLinks) {
	o.Links = v
}

func (o ReviewReassignList) MarshalJSON() ([]byte, error) {
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

func (o *ReviewReassignList) UnmarshalJSON(bytes []byte) (err error) {
	varReviewReassignList := _ReviewReassignList{}

	err = json.Unmarshal(bytes, &varReviewReassignList)
	if err == nil {
		*o = ReviewReassignList(varReviewReassignList)
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

type NullableReviewReassignList struct {
	value *ReviewReassignList
	isSet bool
}

func (v NullableReviewReassignList) Get() *ReviewReassignList {
	return v.value
}

func (v *NullableReviewReassignList) Set(val *ReviewReassignList) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewReassignList) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewReassignList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewReassignList(val *ReviewReassignList) *NullableReviewReassignList {
	return &NullableReviewReassignList{value: val, isSet: true}
}

func (v NullableReviewReassignList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewReassignList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
