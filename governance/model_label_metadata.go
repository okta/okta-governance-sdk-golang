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
)

// checks if the LabelMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelMetadata{}

// LabelMetadata struct for LabelMetadata
type LabelMetadata struct {
	// (Optional) A map of key value pairs. Additional properties for the label. e.g {\\\"backgroundColor\\\": \\\"blue\\\", \\\"font\\\": \\\"Arial\\\"}
	AdditionalPropertiesField map[string]interface{} `json:"additionalProperties,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _LabelMetadata LabelMetadata

// NewLabelMetadata instantiates a new LabelMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelMetadata() *LabelMetadata {
	this := LabelMetadata{}
	return &this
}

// NewLabelMetadataWithDefaults instantiates a new LabelMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelMetadataWithDefaults() *LabelMetadata {
	this := LabelMetadata{}
	return &this
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *LabelMetadata) GetAdditionalPropertiesField() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelMetadata) GetAdditionalPropertiesFieldOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *LabelMetadata) HasAdditionalPropertiesField() bool {
	if o != nil && !IsNil(o.AdditionalPropertiesField) {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]interface{} and assigns it to the AdditionalPropertiesField field.
func (o *LabelMetadata) SetAdditionalPropertiesField(v map[string]interface{}) {
	o.AdditionalPropertiesField = v
}

func (o LabelMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalPropertiesField) {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LabelMetadata) UnmarshalJSON(data []byte) (err error) {
	varLabelMetadata := _LabelMetadata{}

	err = json.Unmarshal(data, &varLabelMetadata)

	if err != nil {
		return err
	}

	*o = LabelMetadata(varLabelMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additionalProperties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelMetadata struct {
	value *LabelMetadata
	isSet bool
}

func (v NullableLabelMetadata) Get() *LabelMetadata {
	return v.value
}

func (v *NullableLabelMetadata) Set(val *LabelMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelMetadata(val *LabelMetadata) *NullableLabelMetadata {
	return &NullableLabelMetadata{value: val, isSet: true}
}

func (v NullableLabelMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
