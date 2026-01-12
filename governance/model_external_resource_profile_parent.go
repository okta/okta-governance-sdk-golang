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

// checks if the ExternalResourceProfileParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalResourceProfileParent{}

// ExternalResourceProfileParent A limited set of properties from the resource's parent profile
type ExternalResourceProfileParent struct {
	// Okta resource ID
	Id string `json:"id"`
	// The display name for the resource
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExternalResourceProfileParent ExternalResourceProfileParent

// NewExternalResourceProfileParent instantiates a new ExternalResourceProfileParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalResourceProfileParent(id string) *ExternalResourceProfileParent {
	this := ExternalResourceProfileParent{}
	this.Id = id
	return &this
}

// NewExternalResourceProfileParentWithDefaults instantiates a new ExternalResourceProfileParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalResourceProfileParentWithDefaults() *ExternalResourceProfileParent {
	this := ExternalResourceProfileParent{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalResourceProfileParent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfileParent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalResourceProfileParent) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalResourceProfileParent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfileParent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalResourceProfileParent) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalResourceProfileParent) SetName(v string) {
	o.Name = &v
}

func (o ExternalResourceProfileParent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalResourceProfileParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalResourceProfileParent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varExternalResourceProfileParent := _ExternalResourceProfileParent{}

	err = json.Unmarshal(data, &varExternalResourceProfileParent)

	if err != nil {
		return err
	}

	*o = ExternalResourceProfileParent(varExternalResourceProfileParent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalResourceProfileParent struct {
	value *ExternalResourceProfileParent
	isSet bool
}

func (v NullableExternalResourceProfileParent) Get() *ExternalResourceProfileParent {
	return v.value
}

func (v *NullableExternalResourceProfileParent) Set(val *ExternalResourceProfileParent) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalResourceProfileParent) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalResourceProfileParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalResourceProfileParent(val *ExternalResourceProfileParent) *NullableExternalResourceProfileParent {
	return &NullableExternalResourceProfileParent{value: val, isSet: true}
}

func (v NullableExternalResourceProfileParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalResourceProfileParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
