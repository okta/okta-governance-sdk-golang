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
	"fmt"
)

// checks if the CampaignReadOnlyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignReadOnlyFields{}

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignReadOnlyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignReadOnlyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCampaignReadOnlyFields := _CampaignReadOnlyFields{}

	err = json.Unmarshal(data, &varCampaignReadOnlyFields)

	if err != nil {
		return err
	}

	*o = CampaignReadOnlyFields(varCampaignReadOnlyFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
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
