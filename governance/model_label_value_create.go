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

// LabelValueCreate struct for LabelValueCreate
type LabelValueCreate struct {
	// Name of the label value
	Name                 string         `json:"name"`
	Metadata             *LabelMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LabelValueCreate LabelValueCreate

// NewLabelValueCreate instantiates a new LabelValueCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelValueCreate(name string) *LabelValueCreate {
	this := LabelValueCreate{}
	this.Name = name
	return &this
}

// NewLabelValueCreateWithDefaults instantiates a new LabelValueCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelValueCreateWithDefaults() *LabelValueCreate {
	this := LabelValueCreate{}
	return &this
}

// GetName returns the Name field value
func (o *LabelValueCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelValueCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelValueCreate) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LabelValueCreate) GetMetadata() LabelMetadata {
	if o == nil || o.Metadata == nil {
		var ret LabelMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelValueCreate) GetMetadataOk() (*LabelMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LabelValueCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given LabelMetadata and assigns it to the Metadata field.
func (o *LabelValueCreate) SetMetadata(v LabelMetadata) {
	o.Metadata = &v
}

func (o LabelValueCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

func (o *LabelValueCreate) UnmarshalJSON(bytes []byte) (err error) {
	varLabelValueCreate := _LabelValueCreate{}

	err = json.Unmarshal(bytes, &varLabelValueCreate)
	if err == nil {
		*o = LabelValueCreate(varLabelValueCreate)
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

type NullableLabelValueCreate struct {
	value *LabelValueCreate
	isSet bool
}

func (v NullableLabelValueCreate) Get() *LabelValueCreate {
	return v.value
}

func (v *NullableLabelValueCreate) Set(val *LabelValueCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelValueCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelValueCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelValueCreate(val *LabelValueCreate) *NullableLabelValueCreate {
	return &NullableLabelValueCreate{value: val, isSet: true}
}

func (v NullableLabelValueCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelValueCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
