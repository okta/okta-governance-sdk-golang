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

// checks if the LabelValueUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelValueUpdate{}

// LabelValueUpdate The value of the updated label properties
type LabelValueUpdate struct {
	// The label value
	Name                 *string        `json:"name,omitempty"`
	Metadata             *LabelMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LabelValueUpdate LabelValueUpdate

// NewLabelValueUpdate instantiates a new LabelValueUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelValueUpdate() *LabelValueUpdate {
	this := LabelValueUpdate{}
	return &this
}

// NewLabelValueUpdateWithDefaults instantiates a new LabelValueUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelValueUpdateWithDefaults() *LabelValueUpdate {
	this := LabelValueUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LabelValueUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelValueUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LabelValueUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LabelValueUpdate) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LabelValueUpdate) GetMetadata() LabelMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret LabelMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelValueUpdate) GetMetadataOk() (*LabelMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LabelValueUpdate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given LabelMetadata and assigns it to the Metadata field.
func (o *LabelValueUpdate) SetMetadata(v LabelMetadata) {
	o.Metadata = &v
}

func (o LabelValueUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelValueUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LabelValueUpdate) UnmarshalJSON(data []byte) (err error) {
	varLabelValueUpdate := _LabelValueUpdate{}

	err = json.Unmarshal(data, &varLabelValueUpdate)

	if err != nil {
		return err
	}

	*o = LabelValueUpdate(varLabelValueUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelValueUpdate struct {
	value *LabelValueUpdate
	isSet bool
}

func (v NullableLabelValueUpdate) Get() *LabelValueUpdate {
	return v.value
}

func (v *NullableLabelValueUpdate) Set(val *LabelValueUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelValueUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelValueUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelValueUpdate(val *LabelValueUpdate) *NullableLabelValueUpdate {
	return &NullableLabelValueUpdate{value: val, isSet: true}
}

func (v NullableLabelValueUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelValueUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
