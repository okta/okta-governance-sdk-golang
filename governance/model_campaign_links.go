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

// CampaignLinks Links available on a single campaign representation.
type CampaignLinks struct {
	LaunchCampaign       Link `json:"launchCampaign"`
	EndCampaign          Link `json:"endCampaign"`
	Reviews              Link `json:"reviews"`
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _CampaignLinks CampaignLinks

// NewCampaignLinks instantiates a new CampaignLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignLinks(launchCampaign Link, endCampaign Link, reviews Link, self Link) *CampaignLinks {
	this := CampaignLinks{}
	this.LaunchCampaign = launchCampaign
	this.EndCampaign = endCampaign
	this.Reviews = reviews
	this.Self = self
	return &this
}

// NewCampaignLinksWithDefaults instantiates a new CampaignLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignLinksWithDefaults() *CampaignLinks {
	this := CampaignLinks{}
	return &this
}

// GetLaunchCampaign returns the LaunchCampaign field value
func (o *CampaignLinks) GetLaunchCampaign() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.LaunchCampaign
}

// GetLaunchCampaignOk returns a tuple with the LaunchCampaign field value
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetLaunchCampaignOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LaunchCampaign, true
}

// SetLaunchCampaign sets field value
func (o *CampaignLinks) SetLaunchCampaign(v Link) {
	o.LaunchCampaign = v
}

// GetEndCampaign returns the EndCampaign field value
func (o *CampaignLinks) GetEndCampaign() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.EndCampaign
}

// GetEndCampaignOk returns a tuple with the EndCampaign field value
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetEndCampaignOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndCampaign, true
}

// SetEndCampaign sets field value
func (o *CampaignLinks) SetEndCampaign(v Link) {
	o.EndCampaign = v
}

// GetReviews returns the Reviews field value
func (o *CampaignLinks) GetReviews() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetReviewsOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reviews, true
}

// SetReviews sets field value
func (o *CampaignLinks) SetReviews(v Link) {
	o.Reviews = v
}

// GetSelf returns the Self field value
func (o *CampaignLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *CampaignLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *CampaignLinks) SetSelf(v Link) {
	o.Self = v
}

func (o CampaignLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["launchCampaign"] = o.LaunchCampaign
	}
	if true {
		toSerialize["endCampaign"] = o.EndCampaign
	}
	if true {
		toSerialize["reviews"] = o.Reviews
	}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignLinks) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignLinks := _CampaignLinks{}

	err = json.Unmarshal(bytes, &varCampaignLinks)
	if err == nil {
		*o = CampaignLinks(varCampaignLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "launchCampaign")
		delete(additionalProperties, "endCampaign")
		delete(additionalProperties, "reviews")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCampaignLinks struct {
	value *CampaignLinks
	isSet bool
}

func (v NullableCampaignLinks) Get() *CampaignLinks {
	return v.value
}

func (v *NullableCampaignLinks) Set(val *CampaignLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignLinks(val *CampaignLinks) *NullableCampaignLinks {
	return &NullableCampaignLinks{value: val, isSet: true}
}

func (v NullableCampaignLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
