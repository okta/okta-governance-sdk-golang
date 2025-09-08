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

// SecurityAccessReviewSubAccessItems struct for SecurityAccessReviewSubAccessItems
type SecurityAccessReviewSubAccessItems struct {
	// List of access items in the security access review
	Data                 []SecurityAccessReviewSubAccessItem `json:"data,omitempty"`
	Links                *SecurityAccessReviewListLinks      `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewSubAccessItems SecurityAccessReviewSubAccessItems

// NewSecurityAccessReviewSubAccessItems instantiates a new SecurityAccessReviewSubAccessItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSubAccessItems() *SecurityAccessReviewSubAccessItems {
	this := SecurityAccessReviewSubAccessItems{}
	return &this
}

// NewSecurityAccessReviewSubAccessItemsWithDefaults instantiates a new SecurityAccessReviewSubAccessItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSubAccessItemsWithDefaults() *SecurityAccessReviewSubAccessItems {
	this := SecurityAccessReviewSubAccessItems{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItems) GetData() []SecurityAccessReviewSubAccessItem {
	if o == nil || o.Data == nil {
		var ret []SecurityAccessReviewSubAccessItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItems) GetDataOk() ([]SecurityAccessReviewSubAccessItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecurityAccessReviewSubAccessItem and assigns it to the Data field.
func (o *SecurityAccessReviewSubAccessItems) SetData(v []SecurityAccessReviewSubAccessItem) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItems) GetLinks() SecurityAccessReviewListLinks {
	if o == nil || o.Links == nil {
		var ret SecurityAccessReviewListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItems) GetLinksOk() (*SecurityAccessReviewListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SecurityAccessReviewListLinks and assigns it to the Links field.
func (o *SecurityAccessReviewSubAccessItems) SetLinks(v SecurityAccessReviewListLinks) {
	o.Links = &v
}

func (o SecurityAccessReviewSubAccessItems) MarshalJSON() ([]byte, error) {
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

func (o *SecurityAccessReviewSubAccessItems) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewSubAccessItems := _SecurityAccessReviewSubAccessItems{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewSubAccessItems)
	if err == nil {
		*o = SecurityAccessReviewSubAccessItems(varSecurityAccessReviewSubAccessItems)
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

type NullableSecurityAccessReviewSubAccessItems struct {
	value *SecurityAccessReviewSubAccessItems
	isSet bool
}

func (v NullableSecurityAccessReviewSubAccessItems) Get() *SecurityAccessReviewSubAccessItems {
	return v.value
}

func (v *NullableSecurityAccessReviewSubAccessItems) Set(val *SecurityAccessReviewSubAccessItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSubAccessItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSubAccessItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSubAccessItems(val *SecurityAccessReviewSubAccessItems) *NullableSecurityAccessReviewSubAccessItems {
	return &NullableSecurityAccessReviewSubAccessItems{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSubAccessItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSubAccessItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
