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

// WritableStandardProperties struct for WritableStandardProperties
type WritableStandardProperties struct {
	// Writable unique key on Create. Not modifiable on update.
	Name *string `json:"name,omitempty"`
	// Human readable description.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableStandardProperties WritableStandardProperties

// NewWritableStandardProperties instantiates a new WritableStandardProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableStandardProperties() *WritableStandardProperties {
	this := WritableStandardProperties{}
	return &this
}

// NewWritableStandardPropertiesWithDefaults instantiates a new WritableStandardProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableStandardPropertiesWithDefaults() *WritableStandardProperties {
	this := WritableStandardProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WritableStandardProperties) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableStandardProperties) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WritableStandardProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WritableStandardProperties) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableStandardProperties) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableStandardProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableStandardProperties) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableStandardProperties) SetDescription(v string) {
	o.Description = &v
}

func (o WritableStandardProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WritableStandardProperties) UnmarshalJSON(bytes []byte) (err error) {
	varWritableStandardProperties := _WritableStandardProperties{}

	err = json.Unmarshal(bytes, &varWritableStandardProperties)
	if err == nil {
		*o = WritableStandardProperties(varWritableStandardProperties)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWritableStandardProperties struct {
	value *WritableStandardProperties
	isSet bool
}

func (v NullableWritableStandardProperties) Get() *WritableStandardProperties {
	return v.value
}

func (v *NullableWritableStandardProperties) Set(val *WritableStandardProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableStandardProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableStandardProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableStandardProperties(val *WritableStandardProperties) *NullableWritableStandardProperties {
	return &NullableWritableStandardProperties{value: val, isSet: true}
}

func (v NullableWritableStandardProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableStandardProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
