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

// ReviewLinks Links available on a single review representation
type ReviewLinks struct {
	Self                 Link `json:"self"`
	ReassignReview       Link `json:"reassignReview"`
	AdditionalProperties map[string]interface{}
}

type _ReviewLinks ReviewLinks

// NewReviewLinks instantiates a new ReviewLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewLinks(self Link, reassignReview Link) *ReviewLinks {
	this := ReviewLinks{}
	this.Self = self
	this.ReassignReview = reassignReview
	return &this
}

// NewReviewLinksWithDefaults instantiates a new ReviewLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewLinksWithDefaults() *ReviewLinks {
	this := ReviewLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ReviewLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ReviewLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ReviewLinks) SetSelf(v Link) {
	o.Self = v
}

// GetReassignReview returns the ReassignReview field value
func (o *ReviewLinks) GetReassignReview() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.ReassignReview
}

// GetReassignReviewOk returns a tuple with the ReassignReview field value
// and a boolean to check if the value has been set.
func (o *ReviewLinks) GetReassignReviewOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReassignReview, true
}

// SetReassignReview sets field value
func (o *ReviewLinks) SetReassignReview(v Link) {
	o.ReassignReview = v
}

func (o ReviewLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if true {
		toSerialize["reassignReview"] = o.ReassignReview
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewLinks) UnmarshalJSON(bytes []byte) (err error) {
	varReviewLinks := _ReviewLinks{}

	err = json.Unmarshal(bytes, &varReviewLinks)
	if err == nil {
		*o = ReviewLinks(varReviewLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "reassignReview")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewLinks struct {
	value *ReviewLinks
	isSet bool
}

func (v NullableReviewLinks) Get() *ReviewLinks {
	return v.value
}

func (v *NullableReviewLinks) Set(val *ReviewLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewLinks(val *ReviewLinks) *NullableReviewLinks {
	return &NullableReviewLinks{value: val, isSet: true}
}

func (v NullableReviewLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
