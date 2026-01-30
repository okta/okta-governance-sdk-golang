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

// checks if the SecurityAccessReviews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviews{}

// SecurityAccessReviews struct for SecurityAccessReviews
type SecurityAccessReviews struct {
	// Security access reviews entries
	Data                 []SecurityAccessReviewSparse   `json:"data,omitempty"`
	Links                *SecurityAccessReviewListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviews SecurityAccessReviews

// NewSecurityAccessReviews instantiates a new SecurityAccessReviews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviews() *SecurityAccessReviews {
	this := SecurityAccessReviews{}
	return &this
}

// NewSecurityAccessReviewsWithDefaults instantiates a new SecurityAccessReviews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewsWithDefaults() *SecurityAccessReviews {
	this := SecurityAccessReviews{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityAccessReviews) GetData() []SecurityAccessReviewSparse {
	if o == nil || IsNil(o.Data) {
		var ret []SecurityAccessReviewSparse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviews) GetDataOk() ([]SecurityAccessReviewSparse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityAccessReviews) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecurityAccessReviewSparse and assigns it to the Data field.
func (o *SecurityAccessReviews) SetData(v []SecurityAccessReviewSparse) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReviews) GetLinks() SecurityAccessReviewListLinks {
	if o == nil || IsNil(o.Links) {
		var ret SecurityAccessReviewListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviews) GetLinksOk() (*SecurityAccessReviewListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviews) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SecurityAccessReviewListLinks and assigns it to the Links field.
func (o *SecurityAccessReviews) SetLinks(v SecurityAccessReviewListLinks) {
	o.Links = &v
}

func (o SecurityAccessReviews) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviews) ToMap() (map[string]interface{}, error) {
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

func (o *SecurityAccessReviews) UnmarshalJSON(data []byte) (err error) {
	varSecurityAccessReviews := _SecurityAccessReviews{}

	err = json.Unmarshal(data, &varSecurityAccessReviews)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviews(varSecurityAccessReviews)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviews struct {
	value *SecurityAccessReviews
	isSet bool
}

func (v NullableSecurityAccessReviews) Get() *SecurityAccessReviews {
	return v.value
}

func (v *NullableSecurityAccessReviews) Set(val *SecurityAccessReviews) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviews) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviews(val *SecurityAccessReviews) *NullableSecurityAccessReviews {
	return &NullableSecurityAccessReviews{value: val, isSet: true}
}

func (v NullableSecurityAccessReviews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
