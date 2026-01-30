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

// checks if the RequestTypeApprovalSettingsCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeApprovalSettingsCustom{}

// RequestTypeApprovalSettingsCustom Approval settings which are not represented in the API accurately. Often the case when lastUpdatedSource is 'WEB'.
type RequestTypeApprovalSettingsCustom struct {
	// When custom modifications to a Request Type cannot be modeled with the SERIAL request type.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalSettingsCustom RequestTypeApprovalSettingsCustom

// NewRequestTypeApprovalSettingsCustom instantiates a new RequestTypeApprovalSettingsCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalSettingsCustom(type_ string) *RequestTypeApprovalSettingsCustom {
	this := RequestTypeApprovalSettingsCustom{}
	this.Type = type_
	return &this
}

// NewRequestTypeApprovalSettingsCustomWithDefaults instantiates a new RequestTypeApprovalSettingsCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalSettingsCustomWithDefaults() *RequestTypeApprovalSettingsCustom {
	this := RequestTypeApprovalSettingsCustom{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeApprovalSettingsCustom) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalSettingsCustom) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeApprovalSettingsCustom) SetType(v string) {
	o.Type = v
}

func (o RequestTypeApprovalSettingsCustom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeApprovalSettingsCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeApprovalSettingsCustom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varRequestTypeApprovalSettingsCustom := _RequestTypeApprovalSettingsCustom{}

	err = json.Unmarshal(data, &varRequestTypeApprovalSettingsCustom)

	if err != nil {
		return err
	}

	*o = RequestTypeApprovalSettingsCustom(varRequestTypeApprovalSettingsCustom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeApprovalSettingsCustom struct {
	value *RequestTypeApprovalSettingsCustom
	isSet bool
}

func (v NullableRequestTypeApprovalSettingsCustom) Get() *RequestTypeApprovalSettingsCustom {
	return v.value
}

func (v *NullableRequestTypeApprovalSettingsCustom) Set(val *RequestTypeApprovalSettingsCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalSettingsCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalSettingsCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalSettingsCustom(val *RequestTypeApprovalSettingsCustom) *NullableRequestTypeApprovalSettingsCustom {
	return &NullableRequestTypeApprovalSettingsCustom{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalSettingsCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalSettingsCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
