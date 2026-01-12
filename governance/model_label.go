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

// checks if the Label type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Label{}

// Label struct for Label
type Label struct {
	// The ID of a label
	LabelId string `json:"labelId"`
	// Key name of the label
	Name string `json:"name"`
	// List of label values
	Values               []LabelValue `json:"values"`
	Links                LinkSelf     `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _Label Label

// NewLabel instantiates a new Label object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel(labelId string, name string, values []LabelValue, links LinkSelf) *Label {
	this := Label{}
	this.LabelId = labelId
	this.Name = name
	this.Values = values
	this.Links = links
	return &this
}

// NewLabelWithDefaults instantiates a new Label object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelWithDefaults() *Label {
	this := Label{}
	return &this
}

// GetLabelId returns the LabelId field value
func (o *Label) GetLabelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value
// and a boolean to check if the value has been set.
func (o *Label) GetLabelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelId, true
}

// SetLabelId sets field value
func (o *Label) SetLabelId(v string) {
	o.LabelId = v
}

// GetName returns the Name field value
func (o *Label) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Label) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Label) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *Label) GetValues() []LabelValue {
	if o == nil {
		var ret []LabelValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Label) GetValuesOk() ([]LabelValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *Label) SetValues(v []LabelValue) {
	o.Values = v
}

// GetLinks returns the Links field value
func (o *Label) GetLinks() LinkSelf {
	if o == nil {
		var ret LinkSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Label) GetLinksOk() (*LinkSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *Label) SetLinks(v LinkSelf) {
	o.Links = v
}

func (o Label) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Label) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelId"] = o.LabelId
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Label) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labelId",
		"name",
		"values",
		"_links",
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

	varLabel := _Label{}

	err = json.Unmarshal(data, &varLabel)

	if err != nil {
		return err
	}

	*o = Label(varLabel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "labelId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabel struct {
	value *Label
	isSet bool
}

func (v NullableLabel) Get() *Label {
	return v.value
}

func (v *NullableLabel) Set(val *Label) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel(val *Label) *NullableLabel {
	return &NullableLabel{value: val, isSet: true}
}

func (v NullableLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
