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

// checks if the SecurityAccessReviewHistoryItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewHistoryItems{}

// SecurityAccessReviewHistoryItems struct for SecurityAccessReviewHistoryItems
type SecurityAccessReviewHistoryItems struct {
	// Lists all events occurred on a given security access review. The history covers all such events and actions.
	Data                 []SecurityAccessReviewHistoryItem `json:"data,omitempty"`
	Links                *SecurityAccessReviewListLinks    `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewHistoryItems SecurityAccessReviewHistoryItems

// NewSecurityAccessReviewHistoryItems instantiates a new SecurityAccessReviewHistoryItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewHistoryItems() *SecurityAccessReviewHistoryItems {
	this := SecurityAccessReviewHistoryItems{}
	return &this
}

// NewSecurityAccessReviewHistoryItemsWithDefaults instantiates a new SecurityAccessReviewHistoryItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewHistoryItemsWithDefaults() *SecurityAccessReviewHistoryItems {
	this := SecurityAccessReviewHistoryItems{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityAccessReviewHistoryItems) GetData() []SecurityAccessReviewHistoryItem {
	if o == nil || IsNil(o.Data) {
		var ret []SecurityAccessReviewHistoryItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItems) GetDataOk() ([]SecurityAccessReviewHistoryItem, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityAccessReviewHistoryItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecurityAccessReviewHistoryItem and assigns it to the Data field.
func (o *SecurityAccessReviewHistoryItems) SetData(v []SecurityAccessReviewHistoryItem) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityAccessReviewHistoryItems) GetLinks() SecurityAccessReviewListLinks {
	if o == nil || IsNil(o.Links) {
		var ret SecurityAccessReviewListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewHistoryItems) GetLinksOk() (*SecurityAccessReviewListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityAccessReviewHistoryItems) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SecurityAccessReviewListLinks and assigns it to the Links field.
func (o *SecurityAccessReviewHistoryItems) SetLinks(v SecurityAccessReviewListLinks) {
	o.Links = &v
}

func (o SecurityAccessReviewHistoryItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewHistoryItems) ToMap() (map[string]interface{}, error) {
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

func (o *SecurityAccessReviewHistoryItems) UnmarshalJSON(data []byte) (err error) {
	varSecurityAccessReviewHistoryItems := _SecurityAccessReviewHistoryItems{}

	err = json.Unmarshal(data, &varSecurityAccessReviewHistoryItems)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewHistoryItems(varSecurityAccessReviewHistoryItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewHistoryItems struct {
	value *SecurityAccessReviewHistoryItems
	isSet bool
}

func (v NullableSecurityAccessReviewHistoryItems) Get() *SecurityAccessReviewHistoryItems {
	return v.value
}

func (v *NullableSecurityAccessReviewHistoryItems) Set(val *SecurityAccessReviewHistoryItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewHistoryItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewHistoryItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewHistoryItems(val *SecurityAccessReviewHistoryItems) *NullableSecurityAccessReviewHistoryItems {
	return &NullableSecurityAccessReviewHistoryItems{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewHistoryItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewHistoryItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
