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

// checks if the WritableStandardPropertiesRequestCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableStandardPropertiesRequestCondition{}

// WritableStandardPropertiesRequestCondition struct for WritableStandardPropertiesRequestCondition
type WritableStandardPropertiesRequestCondition struct {
	// Writable unique key on Create. Modifiable on update.
	Name *string `json:"name,omitempty"`
	// Human readable description.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableStandardPropertiesRequestCondition WritableStandardPropertiesRequestCondition

// NewWritableStandardPropertiesRequestCondition instantiates a new WritableStandardPropertiesRequestCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableStandardPropertiesRequestCondition() *WritableStandardPropertiesRequestCondition {
	this := WritableStandardPropertiesRequestCondition{}
	return &this
}

// NewWritableStandardPropertiesRequestConditionWithDefaults instantiates a new WritableStandardPropertiesRequestCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableStandardPropertiesRequestConditionWithDefaults() *WritableStandardPropertiesRequestCondition {
	this := WritableStandardPropertiesRequestCondition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WritableStandardPropertiesRequestCondition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableStandardPropertiesRequestCondition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WritableStandardPropertiesRequestCondition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WritableStandardPropertiesRequestCondition) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableStandardPropertiesRequestCondition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableStandardPropertiesRequestCondition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableStandardPropertiesRequestCondition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableStandardPropertiesRequestCondition) SetDescription(v string) {
	o.Description = &v
}

func (o WritableStandardPropertiesRequestCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableStandardPropertiesRequestCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableStandardPropertiesRequestCondition) UnmarshalJSON(data []byte) (err error) {
	varWritableStandardPropertiesRequestCondition := _WritableStandardPropertiesRequestCondition{}

	err = json.Unmarshal(data, &varWritableStandardPropertiesRequestCondition)

	if err != nil {
		return err
	}

	*o = WritableStandardPropertiesRequestCondition(varWritableStandardPropertiesRequestCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableStandardPropertiesRequestCondition struct {
	value *WritableStandardPropertiesRequestCondition
	isSet bool
}

func (v NullableWritableStandardPropertiesRequestCondition) Get() *WritableStandardPropertiesRequestCondition {
	return v.value
}

func (v *NullableWritableStandardPropertiesRequestCondition) Set(val *WritableStandardPropertiesRequestCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableStandardPropertiesRequestCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableStandardPropertiesRequestCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableStandardPropertiesRequestCondition(val *WritableStandardPropertiesRequestCondition) *NullableWritableStandardPropertiesRequestCondition {
	return &NullableWritableStandardPropertiesRequestCondition{value: val, isSet: true}
}

func (v NullableWritableStandardPropertiesRequestCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableStandardPropertiesRequestCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
