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

// ReviewReassignListLinks Links available in after reassigning reviews
type ReviewReassignListLinks struct {
	Self                 Link `json:"self"`
	Reviews              Link `json:"reviews"`
	AdditionalProperties map[string]interface{}
}

type _ReviewReassignListLinks ReviewReassignListLinks

// NewReviewReassignListLinks instantiates a new ReviewReassignListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewReassignListLinks(self Link, reviews Link) *ReviewReassignListLinks {
	this := ReviewReassignListLinks{}
	this.Self = self
	this.Reviews = reviews
	return &this
}

// NewReviewReassignListLinksWithDefaults instantiates a new ReviewReassignListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewReassignListLinksWithDefaults() *ReviewReassignListLinks {
	this := ReviewReassignListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ReviewReassignListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ReviewReassignListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ReviewReassignListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetReviews returns the Reviews field value
func (o *ReviewReassignListLinks) GetReviews() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value
// and a boolean to check if the value has been set.
func (o *ReviewReassignListLinks) GetReviewsOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reviews, true
}

// SetReviews sets field value
func (o *ReviewReassignListLinks) SetReviews(v Link) {
	o.Reviews = v
}

func (o ReviewReassignListLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if true {
		toSerialize["reviews"] = o.Reviews
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewReassignListLinks) UnmarshalJSON(bytes []byte) (err error) {
	varReviewReassignListLinks := _ReviewReassignListLinks{}

	err = json.Unmarshal(bytes, &varReviewReassignListLinks)
	if err == nil {
		*o = ReviewReassignListLinks(varReviewReassignListLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "reviews")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewReassignListLinks struct {
	value *ReviewReassignListLinks
	isSet bool
}

func (v NullableReviewReassignListLinks) Get() *ReviewReassignListLinks {
	return v.value
}

func (v *NullableReviewReassignListLinks) Set(val *ReviewReassignListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewReassignListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewReassignListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewReassignListLinks(val *ReviewReassignListLinks) *NullableReviewReassignListLinks {
	return &NullableReviewReassignListLinks{value: val, isSet: true}
}

func (v NullableReviewReassignListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewReassignListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
