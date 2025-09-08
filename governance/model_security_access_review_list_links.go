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

// SecurityAccessReviewListLinks Links available in security access review, access, or sub access list response
type SecurityAccessReviewListLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewListLinks SecurityAccessReviewListLinks

// NewSecurityAccessReviewListLinks instantiates a new SecurityAccessReviewListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewListLinks(self Link) *SecurityAccessReviewListLinks {
	this := SecurityAccessReviewListLinks{}
	this.Self = self
	return &this
}

// NewSecurityAccessReviewListLinksWithDefaults instantiates a new SecurityAccessReviewListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewListLinksWithDefaults() *SecurityAccessReviewListLinks {
	this := SecurityAccessReviewListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *SecurityAccessReviewListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *SecurityAccessReviewListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SecurityAccessReviewListLinks) GetNext() Link {
	if o == nil || o.Next == nil {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewListLinks) GetNextOk() (*Link, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SecurityAccessReviewListLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *SecurityAccessReviewListLinks) SetNext(v Link) {
	o.Next = &v
}

func (o SecurityAccessReviewListLinks) MarshalJSON() ([]byte, error) {
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

func (o *SecurityAccessReviewListLinks) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewListLinks := _SecurityAccessReviewListLinks{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewListLinks)
	if err == nil {
		*o = SecurityAccessReviewListLinks(varSecurityAccessReviewListLinks)
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

type NullableSecurityAccessReviewListLinks struct {
	value *SecurityAccessReviewListLinks
	isSet bool
}

func (v NullableSecurityAccessReviewListLinks) Get() *SecurityAccessReviewListLinks {
	return v.value
}

func (v *NullableSecurityAccessReviewListLinks) Set(val *SecurityAccessReviewListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewListLinks(val *SecurityAccessReviewListLinks) *NullableSecurityAccessReviewListLinks {
	return &NullableSecurityAccessReviewListLinks{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
