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

// EntitlementsList struct for EntitlementsList
type EntitlementsList struct {
	// List of all entitlements matching the filter
	Data                 []EntitlementsListObject `json:"data"`
	Links                ListLinks                `json:"_links"`
	Metadata             ListMetadata             `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementsList EntitlementsList

// NewEntitlementsList instantiates a new EntitlementsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsList(data []EntitlementsListObject, links ListLinks, metadata ListMetadata) *EntitlementsList {
	this := EntitlementsList{}
	this.Data = data
	this.Links = links
	this.Metadata = metadata
	return &this
}

// NewEntitlementsListWithDefaults instantiates a new EntitlementsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsListWithDefaults() *EntitlementsList {
	this := EntitlementsList{}
	return &this
}

// GetData returns the Data field value
func (o *EntitlementsList) GetData() []EntitlementsListObject {
	if o == nil {
		var ret []EntitlementsListObject
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EntitlementsList) GetDataOk() ([]EntitlementsListObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EntitlementsList) SetData(v []EntitlementsListObject) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *EntitlementsList) GetLinks() ListLinks {
	if o == nil {
		var ret ListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EntitlementsList) GetLinksOk() (*ListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EntitlementsList) SetLinks(v ListLinks) {
	o.Links = v
}

// GetMetadata returns the Metadata field value
func (o *EntitlementsList) GetMetadata() ListMetadata {
	if o == nil {
		var ret ListMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *EntitlementsList) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *EntitlementsList) SetMetadata(v ListMetadata) {
	o.Metadata = v
}

func (o EntitlementsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementsList) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementsList := _EntitlementsList{}

	err = json.Unmarshal(bytes, &varEntitlementsList)
	if err == nil {
		*o = EntitlementsList(varEntitlementsList)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementsList struct {
	value *EntitlementsList
	isSet bool
}

func (v NullableEntitlementsList) Get() *EntitlementsList {
	return v.value
}

func (v *NullableEntitlementsList) Set(val *EntitlementsList) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsList) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsList(val *EntitlementsList) *NullableEntitlementsList {
	return &NullableEntitlementsList{value: val, isSet: true}
}

func (v NullableEntitlementsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
