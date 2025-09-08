/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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

// LabelValueUpdate struct for LabelValueUpdate
type LabelValueUpdate struct {
	// Name of the label value
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelValueUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LabelValueUpdate) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.Metadata == nil {
		var ret LabelMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelValueUpdate) GetMetadataOk() (*LabelMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LabelValueUpdate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given LabelMetadata and assigns it to the Metadata field.
func (o *LabelValueUpdate) SetMetadata(v LabelMetadata) {
	o.Metadata = &v
}

func (o LabelValueUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LabelValueUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varLabelValueUpdate := _LabelValueUpdate{}

	err = json.Unmarshal(bytes, &varLabelValueUpdate)
	if err == nil {
		*o = LabelValueUpdate(varLabelValueUpdate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
