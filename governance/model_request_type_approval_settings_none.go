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

// checks if the RequestTypeApprovalSettingsNone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeApprovalSettingsNone{}

// RequestTypeApprovalSettingsNone Settings for no approvals
type RequestTypeApprovalSettingsNone struct {
	// No approvals are required for this request type.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalSettingsNone RequestTypeApprovalSettingsNone

// NewRequestTypeApprovalSettingsNone instantiates a new RequestTypeApprovalSettingsNone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalSettingsNone(type_ string) *RequestTypeApprovalSettingsNone {
	this := RequestTypeApprovalSettingsNone{}
	this.Type = type_
	return &this
}

// NewRequestTypeApprovalSettingsNoneWithDefaults instantiates a new RequestTypeApprovalSettingsNone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalSettingsNoneWithDefaults() *RequestTypeApprovalSettingsNone {
	this := RequestTypeApprovalSettingsNone{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeApprovalSettingsNone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalSettingsNone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeApprovalSettingsNone) SetType(v string) {
	o.Type = v
}

func (o RequestTypeApprovalSettingsNone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeApprovalSettingsNone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeApprovalSettingsNone) UnmarshalJSON(data []byte) (err error) {
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

	varRequestTypeApprovalSettingsNone := _RequestTypeApprovalSettingsNone{}

	err = json.Unmarshal(data, &varRequestTypeApprovalSettingsNone)

	if err != nil {
		return err
	}

	*o = RequestTypeApprovalSettingsNone(varRequestTypeApprovalSettingsNone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeApprovalSettingsNone struct {
	value *RequestTypeApprovalSettingsNone
	isSet bool
}

func (v NullableRequestTypeApprovalSettingsNone) Get() *RequestTypeApprovalSettingsNone {
	return v.value
}

func (v *NullableRequestTypeApprovalSettingsNone) Set(val *RequestTypeApprovalSettingsNone) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalSettingsNone) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalSettingsNone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalSettingsNone(val *RequestTypeApprovalSettingsNone) *NullableRequestTypeApprovalSettingsNone {
	return &NullableRequestTypeApprovalSettingsNone{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalSettingsNone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalSettingsNone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
