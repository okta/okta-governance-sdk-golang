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

// checks if the LabelValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelValue{}

// LabelValue struct for LabelValue
type LabelValue struct {
	// The ID of a label value
	LabelValueId string `json:"labelValueId"`
	// The label value
	Name                 string         `json:"name"`
	Metadata             *LabelMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LabelValue LabelValue

// NewLabelValue instantiates a new LabelValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelValue(labelValueId string, name string) *LabelValue {
	this := LabelValue{}
	this.LabelValueId = labelValueId
	this.Name = name
	return &this
}

// NewLabelValueWithDefaults instantiates a new LabelValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelValueWithDefaults() *LabelValue {
	this := LabelValue{}
	return &this
}

// GetLabelValueId returns the LabelValueId field value
func (o *LabelValue) GetLabelValueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelValueId
}

// GetLabelValueIdOk returns a tuple with the LabelValueId field value
// and a boolean to check if the value has been set.
func (o *LabelValue) GetLabelValueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelValueId, true
}

// SetLabelValueId sets field value
func (o *LabelValue) SetLabelValueId(v string) {
	o.LabelValueId = v
}

// GetName returns the Name field value
func (o *LabelValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelValue) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LabelValue) GetMetadata() LabelMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret LabelMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelValue) GetMetadataOk() (*LabelMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LabelValue) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given LabelMetadata and assigns it to the Metadata field.
func (o *LabelValue) SetMetadata(v LabelMetadata) {
	o.Metadata = &v
}

func (o LabelValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelValueId"] = o.LabelValueId
	toSerialize["name"] = o.Name
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LabelValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labelValueId",
		"name",
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

	varLabelValue := _LabelValue{}

	err = json.Unmarshal(data, &varLabelValue)

	if err != nil {
		return err
	}

	*o = LabelValue(varLabelValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "labelValueId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelValue struct {
	value *LabelValue
	isSet bool
}

func (v NullableLabelValue) Get() *LabelValue {
	return v.value
}

func (v *NullableLabelValue) Set(val *LabelValue) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelValue) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelValue(val *LabelValue) *NullableLabelValue {
	return &NullableLabelValue{value: val, isSet: true}
}

func (v NullableLabelValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
