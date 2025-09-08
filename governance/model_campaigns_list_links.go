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

// CampaignsListLinks Links available in campaign list response
type CampaignsListLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignsListLinks CampaignsListLinks

// NewCampaignsListLinks instantiates a new CampaignsListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignsListLinks(self Link) *CampaignsListLinks {
	this := CampaignsListLinks{}
	this.Self = self
	return &this
}

// NewCampaignsListLinksWithDefaults instantiates a new CampaignsListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignsListLinksWithDefaults() *CampaignsListLinks {
	this := CampaignsListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *CampaignsListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *CampaignsListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *CampaignsListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *CampaignsListLinks) GetNext() Link {
	if o == nil || o.Next == nil {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignsListLinks) GetNextOk() (*Link, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *CampaignsListLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *CampaignsListLinks) SetNext(v Link) {
	o.Next = &v
}

func (o CampaignsListLinks) MarshalJSON() ([]byte, error) {
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

func (o *CampaignsListLinks) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignsListLinks := _CampaignsListLinks{}

	err = json.Unmarshal(bytes, &varCampaignsListLinks)
	if err == nil {
		*o = CampaignsListLinks(varCampaignsListLinks)
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

type NullableCampaignsListLinks struct {
	value *CampaignsListLinks
	isSet bool
}

func (v NullableCampaignsListLinks) Get() *CampaignsListLinks {
	return v.value
}

func (v *NullableCampaignsListLinks) Set(val *CampaignsListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignsListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignsListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignsListLinks(val *CampaignsListLinks) *NullableCampaignsListLinks {
	return &NullableCampaignsListLinks{value: val, isSet: true}
}

func (v NullableCampaignsListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignsListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
