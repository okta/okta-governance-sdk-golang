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
	"fmt"
)

// checks if the ReviewList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewList{}

// ReviewList struct for ReviewList
type ReviewList struct {
	// Sparse representation of reviews
	Data                 []ReviewSparse  `json:"data"`
	Links                ReviewListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ReviewList ReviewList

// NewReviewList instantiates a new ReviewList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewList(data []ReviewSparse, links ReviewListLinks) *ReviewList {
	this := ReviewList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewReviewListWithDefaults instantiates a new ReviewList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewListWithDefaults() *ReviewList {
	this := ReviewList{}
	return &this
}

// GetData returns the Data field value
func (o *ReviewList) GetData() []ReviewSparse {
	if o == nil {
		var ret []ReviewSparse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReviewList) GetDataOk() ([]ReviewSparse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ReviewList) SetData(v []ReviewSparse) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ReviewList) GetLinks() ReviewListLinks {
	if o == nil {
		var ret ReviewListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewList) GetLinksOk() (*ReviewListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewList) SetLinks(v ReviewListLinks) {
	o.Links = v
}

func (o ReviewList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReviewList := _ReviewList{}

	err = json.Unmarshal(data, &varReviewList)

	if err != nil {
		return err
	}

	*o = ReviewList(varReviewList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewList struct {
	value *ReviewList
	isSet bool
}

func (v NullableReviewList) Get() *ReviewList {
	return v.value
}

func (v *NullableReviewList) Set(val *ReviewList) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewList) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewList(val *ReviewList) *NullableReviewList {
	return &NullableReviewList{value: val, isSet: true}
}

func (v NullableReviewList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
