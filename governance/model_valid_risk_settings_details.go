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

// checks if the ValidRiskSettingsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidRiskSettingsDetails{}

// ValidRiskSettingsDetails Risk settings that are eligible for the specified resource
type ValidRiskSettingsDetails struct {
	// Risk settings that are eligible for the specified resource
	SupportedTypes       []ValidRiskSettingsType `json:"supportedTypes"`
	AdditionalProperties map[string]interface{}
}

type _ValidRiskSettingsDetails ValidRiskSettingsDetails

// NewValidRiskSettingsDetails instantiates a new ValidRiskSettingsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidRiskSettingsDetails(supportedTypes []ValidRiskSettingsType) *ValidRiskSettingsDetails {
	this := ValidRiskSettingsDetails{}
	this.SupportedTypes = supportedTypes
	return &this
}

// NewValidRiskSettingsDetailsWithDefaults instantiates a new ValidRiskSettingsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidRiskSettingsDetailsWithDefaults() *ValidRiskSettingsDetails {
	this := ValidRiskSettingsDetails{}
	return &this
}

// GetSupportedTypes returns the SupportedTypes field value
func (o *ValidRiskSettingsDetails) GetSupportedTypes() []ValidRiskSettingsType {
	if o == nil {
		var ret []ValidRiskSettingsType
		return ret
	}

	return o.SupportedTypes
}

// GetSupportedTypesOk returns a tuple with the SupportedTypes field value
// and a boolean to check if the value has been set.
func (o *ValidRiskSettingsDetails) GetSupportedTypesOk() ([]ValidRiskSettingsType, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedTypes, true
}

// SetSupportedTypes sets field value
func (o *ValidRiskSettingsDetails) SetSupportedTypes(v []ValidRiskSettingsType) {
	o.SupportedTypes = v
}

func (o ValidRiskSettingsDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidRiskSettingsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supportedTypes"] = o.SupportedTypes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidRiskSettingsDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"supportedTypes",
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

	varValidRiskSettingsDetails := _ValidRiskSettingsDetails{}

	err = json.Unmarshal(data, &varValidRiskSettingsDetails)

	if err != nil {
		return err
	}

	*o = ValidRiskSettingsDetails(varValidRiskSettingsDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "supportedTypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidRiskSettingsDetails struct {
	value *ValidRiskSettingsDetails
	isSet bool
}

func (v NullableValidRiskSettingsDetails) Get() *ValidRiskSettingsDetails {
	return v.value
}

func (v *NullableValidRiskSettingsDetails) Set(val *ValidRiskSettingsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableValidRiskSettingsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableValidRiskSettingsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidRiskSettingsDetails(val *ValidRiskSettingsDetails) *NullableValidRiskSettingsDetails {
	return &NullableValidRiskSettingsDetails{value: val, isSet: true}
}

func (v NullableValidRiskSettingsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidRiskSettingsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
