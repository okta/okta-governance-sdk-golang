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

// CampaignsList struct for CampaignsList
type CampaignsList struct {
	// Sparse representation of a Campaign resource.
	Data                 []CampaignSparse   `json:"data"`
	Links                CampaignsListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _CampaignsList CampaignsList

// NewCampaignsList instantiates a new CampaignsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignsList(data []CampaignSparse, links CampaignsListLinks) *CampaignsList {
	this := CampaignsList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCampaignsListWithDefaults instantiates a new CampaignsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignsListWithDefaults() *CampaignsList {
	this := CampaignsList{}
	return &this
}

// GetData returns the Data field value
func (o *CampaignsList) GetData() []CampaignSparse {
	if o == nil {
		var ret []CampaignSparse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CampaignsList) GetDataOk() ([]CampaignSparse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CampaignsList) SetData(v []CampaignSparse) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *CampaignsList) GetLinks() CampaignsListLinks {
	if o == nil {
		var ret CampaignsListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CampaignsList) GetLinksOk() (*CampaignsListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CampaignsList) SetLinks(v CampaignsListLinks) {
	o.Links = v
}

func (o CampaignsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignsList) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignsList := _CampaignsList{}

	err = json.Unmarshal(bytes, &varCampaignsList)
	if err == nil {
		*o = CampaignsList(varCampaignsList)
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

type NullableCampaignsList struct {
	value *CampaignsList
	isSet bool
}

func (v NullableCampaignsList) Get() *CampaignsList {
	return v.value
}

func (v *NullableCampaignsList) Set(val *CampaignsList) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignsList) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignsList(val *CampaignsList) *NullableCampaignsList {
	return &NullableCampaignsList{value: val, isSet: true}
}

func (v NullableCampaignsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
