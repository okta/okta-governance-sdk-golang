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

// ReviewListLinks Links available in review list response
type ReviewListLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewListLinks ReviewListLinks

// NewReviewListLinks instantiates a new ReviewListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewListLinks(self Link) *ReviewListLinks {
	this := ReviewListLinks{}
	this.Self = self
	return &this
}

// NewReviewListLinksWithDefaults instantiates a new ReviewListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewListLinksWithDefaults() *ReviewListLinks {
	this := ReviewListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ReviewListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ReviewListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ReviewListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ReviewListLinks) GetNext() Link {
	if o == nil || o.Next == nil {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewListLinks) GetNextOk() (*Link, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ReviewListLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *ReviewListLinks) SetNext(v Link) {
	o.Next = &v
}

func (o ReviewListLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewListLinks) UnmarshalJSON(bytes []byte) (err error) {
	varReviewListLinks := _ReviewListLinks{}

	err = json.Unmarshal(bytes, &varReviewListLinks)
	if err == nil {
		*o = ReviewListLinks(varReviewListLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewListLinks struct {
	value *ReviewListLinks
	isSet bool
}

func (v NullableReviewListLinks) Get() *ReviewListLinks {
	return v.value
}

func (v *NullableReviewListLinks) Set(val *ReviewListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewListLinks(val *ReviewListLinks) *NullableReviewListLinks {
	return &NullableReviewListLinks{value: val, isSet: true}
}

func (v NullableReviewListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
