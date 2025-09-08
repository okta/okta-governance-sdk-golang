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

// CampaignEndSkipRemediation Boolean variable to skip remediation when campaign end operation is requested.
type CampaignEndSkipRemediation struct {
	// Set the Boolean variable to `true` to skip remediation in cases where `remediationSetting.noResponse=DENY`.
	SkipRemediation      *bool `json:"skipRemediation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignEndSkipRemediation CampaignEndSkipRemediation

// NewCampaignEndSkipRemediation instantiates a new CampaignEndSkipRemediation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignEndSkipRemediation() *CampaignEndSkipRemediation {
	this := CampaignEndSkipRemediation{}
	return &this
}

// NewCampaignEndSkipRemediationWithDefaults instantiates a new CampaignEndSkipRemediation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignEndSkipRemediationWithDefaults() *CampaignEndSkipRemediation {
	this := CampaignEndSkipRemediation{}
	return &this
}

// GetSkipRemediation returns the SkipRemediation field value if set, zero value otherwise.
func (o *CampaignEndSkipRemediation) GetSkipRemediation() bool {
	if o == nil || o.SkipRemediation == nil {
		var ret bool
		return ret
	}
	return *o.SkipRemediation
}

// GetSkipRemediationOk returns a tuple with the SkipRemediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignEndSkipRemediation) GetSkipRemediationOk() (*bool, bool) {
	if o == nil || o.SkipRemediation == nil {
		return nil, false
	}
	return o.SkipRemediation, true
}

// HasSkipRemediation returns a boolean if a field has been set.
func (o *CampaignEndSkipRemediation) HasSkipRemediation() bool {
	if o != nil && o.SkipRemediation != nil {
		return true
	}

	return false
}

// SetSkipRemediation gets a reference to the given bool and assigns it to the SkipRemediation field.
func (o *CampaignEndSkipRemediation) SetSkipRemediation(v bool) {
	o.SkipRemediation = &v
}

func (o CampaignEndSkipRemediation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkipRemediation != nil {
		toSerialize["skipRemediation"] = o.SkipRemediation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignEndSkipRemediation) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignEndSkipRemediation := _CampaignEndSkipRemediation{}

	err = json.Unmarshal(bytes, &varCampaignEndSkipRemediation)
	if err == nil {
		*o = CampaignEndSkipRemediation(varCampaignEndSkipRemediation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "skipRemediation")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCampaignEndSkipRemediation struct {
	value *CampaignEndSkipRemediation
	isSet bool
}

func (v NullableCampaignEndSkipRemediation) Get() *CampaignEndSkipRemediation {
	return v.value
}

func (v *NullableCampaignEndSkipRemediation) Set(val *CampaignEndSkipRemediation) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignEndSkipRemediation) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignEndSkipRemediation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignEndSkipRemediation(val *CampaignEndSkipRemediation) *NullableCampaignEndSkipRemediation {
	return &NullableCampaignEndSkipRemediation{value: val, isSet: true}
}

func (v NullableCampaignEndSkipRemediation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignEndSkipRemediation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
