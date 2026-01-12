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

// checks if the GrantObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantObject{}

// GrantObject Unique identifier for the object
type GrantObject struct {
	// Unique identifier for the object
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantObject GrantObject

// NewGrantObject instantiates a new GrantObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantObject() *GrantObject {
	this := GrantObject{}
	return &this
}

// NewGrantObjectWithDefaults instantiates a new GrantObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantObjectWithDefaults() *GrantObject {
	this := GrantObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GrantObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GrantObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GrantObject) SetId(v string) {
	o.Id = &v
}

func (o GrantObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantObject) UnmarshalJSON(data []byte) (err error) {
	varGrantObject := _GrantObject{}

	err = json.Unmarshal(data, &varGrantObject)

	if err != nil {
		return err
	}

	*o = GrantObject(varGrantObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantObject struct {
	value *GrantObject
	isSet bool
}

func (v NullableGrantObject) Get() *GrantObject {
	return v.value
}

func (v *NullableGrantObject) Set(val *GrantObject) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantObject) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantObject(val *GrantObject) *NullableGrantObject {
	return &NullableGrantObject{value: val, isSet: true}
}

func (v NullableGrantObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
