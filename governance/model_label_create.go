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

// checks if the LabelCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelCreate{}

// LabelCreate struct for LabelCreate
type LabelCreate struct {
	// Key name of the label
	Name string `json:"name"`
	// List of label values
	Values               []LabelValueCreate `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _LabelCreate LabelCreate

// NewLabelCreate instantiates a new LabelCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelCreate(name string, values []LabelValueCreate) *LabelCreate {
	this := LabelCreate{}
	this.Name = name
	this.Values = values
	return &this
}

// NewLabelCreateWithDefaults instantiates a new LabelCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelCreateWithDefaults() *LabelCreate {
	this := LabelCreate{}
	return &this
}

// GetName returns the Name field value
func (o *LabelCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelCreate) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *LabelCreate) GetValues() []LabelValueCreate {
	if o == nil {
		var ret []LabelValueCreate
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *LabelCreate) GetValuesOk() ([]LabelValueCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *LabelCreate) SetValues(v []LabelValueCreate) {
	o.Values = v
}

func (o LabelCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LabelCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varLabelCreate := _LabelCreate{}

	err = json.Unmarshal(data, &varLabelCreate)

	if err != nil {
		return err
	}

	*o = LabelCreate(varLabelCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelCreate struct {
	value *LabelCreate
	isSet bool
}

func (v NullableLabelCreate) Get() *LabelCreate {
	return v.value
}

func (v *NullableLabelCreate) Set(val *LabelCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelCreate(val *LabelCreate) *NullableLabelCreate {
	return &NullableLabelCreate{value: val, isSet: true}
}

func (v NullableLabelCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
