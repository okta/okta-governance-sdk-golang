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

// checks if the TargetGovernanceLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetGovernanceLabel{}

// TargetGovernanceLabel struct for TargetGovernanceLabel
type TargetGovernanceLabel struct {
	// The ID of a label
	LabelId string `json:"labelId"`
	// Name of the label
	Name string `json:"name"`
	// List of label values
	Values               []LabelValue `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _TargetGovernanceLabel TargetGovernanceLabel

// NewTargetGovernanceLabel instantiates a new TargetGovernanceLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetGovernanceLabel(labelId string, name string, values []LabelValue) *TargetGovernanceLabel {
	this := TargetGovernanceLabel{}
	this.LabelId = labelId
	this.Name = name
	this.Values = values
	return &this
}

// NewTargetGovernanceLabelWithDefaults instantiates a new TargetGovernanceLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetGovernanceLabelWithDefaults() *TargetGovernanceLabel {
	this := TargetGovernanceLabel{}
	return &this
}

// GetLabelId returns the LabelId field value
func (o *TargetGovernanceLabel) GetLabelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value
// and a boolean to check if the value has been set.
func (o *TargetGovernanceLabel) GetLabelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelId, true
}

// SetLabelId sets field value
func (o *TargetGovernanceLabel) SetLabelId(v string) {
	o.LabelId = v
}

// GetName returns the Name field value
func (o *TargetGovernanceLabel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TargetGovernanceLabel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TargetGovernanceLabel) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *TargetGovernanceLabel) GetValues() []LabelValue {
	if o == nil {
		var ret []LabelValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TargetGovernanceLabel) GetValuesOk() ([]LabelValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *TargetGovernanceLabel) SetValues(v []LabelValue) {
	o.Values = v
}

func (o TargetGovernanceLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetGovernanceLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelId"] = o.LabelId
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TargetGovernanceLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labelId",
		"name",
		"values",
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

	varTargetGovernanceLabel := _TargetGovernanceLabel{}

	err = json.Unmarshal(data, &varTargetGovernanceLabel)

	if err != nil {
		return err
	}

	*o = TargetGovernanceLabel(varTargetGovernanceLabel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "labelId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTargetGovernanceLabel struct {
	value *TargetGovernanceLabel
	isSet bool
}

func (v NullableTargetGovernanceLabel) Get() *TargetGovernanceLabel {
	return v.value
}

func (v *NullableTargetGovernanceLabel) Set(val *TargetGovernanceLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetGovernanceLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetGovernanceLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetGovernanceLabel(val *TargetGovernanceLabel) *NullableTargetGovernanceLabel {
	return &NullableTargetGovernanceLabel{value: val, isSet: true}
}

func (v NullableTargetGovernanceLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetGovernanceLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
