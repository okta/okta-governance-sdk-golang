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

// checks if the GovernanceLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovernanceLabel{}

// GovernanceLabel struct for GovernanceLabel
type GovernanceLabel struct {
	// label value
	Value string `json:"value"`
	// hex color of the label
	Color                string `json:"color"`
	AdditionalProperties map[string]interface{}
}

type _GovernanceLabel GovernanceLabel

// NewGovernanceLabel instantiates a new GovernanceLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernanceLabel(value string, color string) *GovernanceLabel {
	this := GovernanceLabel{}
	this.Value = value
	this.Color = color
	return &this
}

// NewGovernanceLabelWithDefaults instantiates a new GovernanceLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernanceLabelWithDefaults() *GovernanceLabel {
	this := GovernanceLabel{}
	return &this
}

// GetValue returns the Value field value
func (o *GovernanceLabel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GovernanceLabel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GovernanceLabel) SetValue(v string) {
	o.Value = v
}

// GetColor returns the Color field value
func (o *GovernanceLabel) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *GovernanceLabel) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *GovernanceLabel) SetColor(v string) {
	o.Color = v
}

func (o GovernanceLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovernanceLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["color"] = o.Color

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GovernanceLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"color",
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

	varGovernanceLabel := _GovernanceLabel{}

	err = json.Unmarshal(data, &varGovernanceLabel)

	if err != nil {
		return err
	}

	*o = GovernanceLabel(varGovernanceLabel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGovernanceLabel struct {
	value *GovernanceLabel
	isSet bool
}

func (v NullableGovernanceLabel) Get() *GovernanceLabel {
	return v.value
}

func (v *NullableGovernanceLabel) Set(val *GovernanceLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernanceLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernanceLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernanceLabel(val *GovernanceLabel) *NullableGovernanceLabel {
	return &NullableGovernanceLabel{value: val, isSet: true}
}

func (v NullableGovernanceLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernanceLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
