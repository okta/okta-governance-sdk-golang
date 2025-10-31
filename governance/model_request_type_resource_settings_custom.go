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

// checks if the RequestTypeResourceSettingsCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeResourceSettingsCustom{}

// RequestTypeResourceSettingsCustom Resource settings which cannot be represented in the API accurately. Often the case when lastUpdatedSource is 'WEB'.
type RequestTypeResourceSettingsCustom struct {
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeResourceSettingsCustom RequestTypeResourceSettingsCustom

// NewRequestTypeResourceSettingsCustom instantiates a new RequestTypeResourceSettingsCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeResourceSettingsCustom(type_ string) *RequestTypeResourceSettingsCustom {
	this := RequestTypeResourceSettingsCustom{}
	this.Type = type_
	return &this
}

// NewRequestTypeResourceSettingsCustomWithDefaults instantiates a new RequestTypeResourceSettingsCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeResourceSettingsCustomWithDefaults() *RequestTypeResourceSettingsCustom {
	this := RequestTypeResourceSettingsCustom{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeResourceSettingsCustom) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsCustom) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeResourceSettingsCustom) SetType(v string) {
	o.Type = v
}

func (o RequestTypeResourceSettingsCustom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeResourceSettingsCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeResourceSettingsCustom) UnmarshalJSON(data []byte) (err error) {
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

	varRequestTypeResourceSettingsCustom := _RequestTypeResourceSettingsCustom{}

	err = json.Unmarshal(data, &varRequestTypeResourceSettingsCustom)

	if err != nil {
		return err
	}

	*o = RequestTypeResourceSettingsCustom(varRequestTypeResourceSettingsCustom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeResourceSettingsCustom struct {
	value *RequestTypeResourceSettingsCustom
	isSet bool
}

func (v NullableRequestTypeResourceSettingsCustom) Get() *RequestTypeResourceSettingsCustom {
	return v.value
}

func (v *NullableRequestTypeResourceSettingsCustom) Set(val *RequestTypeResourceSettingsCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeResourceSettingsCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeResourceSettingsCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeResourceSettingsCustom(val *RequestTypeResourceSettingsCustom) *NullableRequestTypeResourceSettingsCustom {
	return &NullableRequestTypeResourceSettingsCustom{value: val, isSet: true}
}

func (v NullableRequestTypeResourceSettingsCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeResourceSettingsCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
