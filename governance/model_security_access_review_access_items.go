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

// SecurityAccessReviewAccessItems struct for SecurityAccessReviewAccessItems
type SecurityAccessReviewAccessItems struct {
	// List of access items in the security access review
	Data                 []SecurityAccessReviewAccessItem `json:"data,omitempty"`
	Links                *SecurityAccessReviewListLinks   `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewAccessItems SecurityAccessReviewAccessItems

// NewSecurityAccessReviewAccessItems instantiates a new SecurityAccessReviewAccessItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewAccessItems() *SecurityAccessReviewAccessItems {
	this := SecurityAccessReviewAccessItems{}
	return &this
}

// NewSecurityAccessReviewAccessItemsWithDefaults instantiates a new SecurityAccessReviewAccessItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewAccessItemsWithDefaults() *SecurityAccessReviewAccessItems {
	this := SecurityAccessReviewAccessItems{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItems) GetData() []SecurityAccessReviewAccessItem {
	if o == nil || o.Data == nil {
		var ret []SecurityAccessReviewAccessItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItems) GetDataOk() ([]SecurityAccessReviewAccessItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecurityAccessReviewAccessItem and assigns it to the Data field.
func (o *SecurityAccessReviewAccessItems) SetData(v []SecurityAccessReviewAccessItem) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItems) GetLinks() SecurityAccessReviewListLinks {
	if o == nil || o.Links == nil {
		var ret SecurityAccessReviewListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItems) GetLinksOk() (*SecurityAccessReviewListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SecurityAccessReviewListLinks and assigns it to the Links field.
func (o *SecurityAccessReviewAccessItems) SetLinks(v SecurityAccessReviewListLinks) {
	o.Links = &v
}

func (o SecurityAccessReviewAccessItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewAccessItems) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewAccessItems := _SecurityAccessReviewAccessItems{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewAccessItems)
	if err == nil {
		*o = SecurityAccessReviewAccessItems(varSecurityAccessReviewAccessItems)
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

type NullableSecurityAccessReviewAccessItems struct {
	value *SecurityAccessReviewAccessItems
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItems) Get() *SecurityAccessReviewAccessItems {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItems) Set(val *SecurityAccessReviewAccessItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItems(val *SecurityAccessReviewAccessItems) *NullableSecurityAccessReviewAccessItems {
	return &NullableSecurityAccessReviewAccessItems{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
