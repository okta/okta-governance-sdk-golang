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

// CampaignReadOnlyFields Read-only attributes
type CampaignReadOnlyFields struct {
	Status               CampaignStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _CampaignReadOnlyFields CampaignReadOnlyFields

// NewCampaignReadOnlyFields instantiates a new CampaignReadOnlyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReadOnlyFields(status CampaignStatus) *CampaignReadOnlyFields {
	this := CampaignReadOnlyFields{}
	this.Status = status
	return &this
}

// NewCampaignReadOnlyFieldsWithDefaults instantiates a new CampaignReadOnlyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReadOnlyFieldsWithDefaults() *CampaignReadOnlyFields {
	this := CampaignReadOnlyFields{}
	return &this
}

// GetStatus returns the Status field value
func (o *CampaignReadOnlyFields) GetStatus() CampaignStatus {
	if o == nil {
		var ret CampaignStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CampaignReadOnlyFields) GetStatusOk() (*CampaignStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CampaignReadOnlyFields) SetStatus(v CampaignStatus) {
	o.Status = v
}

func (o CampaignReadOnlyFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignReadOnlyFields) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignReadOnlyFields := _CampaignReadOnlyFields{}

	err = json.Unmarshal(bytes, &varCampaignReadOnlyFields)
	if err == nil {
		*o = CampaignReadOnlyFields(varCampaignReadOnlyFields)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCampaignReadOnlyFields struct {
	value *CampaignReadOnlyFields
	isSet bool
}

func (v NullableCampaignReadOnlyFields) Get() *CampaignReadOnlyFields {
	return v.value
}

func (v *NullableCampaignReadOnlyFields) Set(val *CampaignReadOnlyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReadOnlyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReadOnlyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReadOnlyFields(val *CampaignReadOnlyFields) *NullableCampaignReadOnlyFields {
	return &NullableCampaignReadOnlyFields{value: val, isSet: true}
}

func (v NullableCampaignReadOnlyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReadOnlyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
