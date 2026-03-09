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

// checks if the TeamsArrayFullInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamsArrayFullInner{}

// TeamsArrayFullInner struct for TeamsArrayFullInner
type TeamsArrayFullInner struct {
	// The access request team `id`.
	Id                   string `json:"id" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	AdditionalProperties map[string]interface{}
}

type _TeamsArrayFullInner TeamsArrayFullInner

// NewTeamsArrayFullInner instantiates a new TeamsArrayFullInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsArrayFullInner(id string) *TeamsArrayFullInner {
	this := TeamsArrayFullInner{}
	this.Id = id
	return &this
}

// NewTeamsArrayFullInnerWithDefaults instantiates a new TeamsArrayFullInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsArrayFullInnerWithDefaults() *TeamsArrayFullInner {
	this := TeamsArrayFullInner{}
	return &this
}

// GetId returns the Id field value
func (o *TeamsArrayFullInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamsArrayFullInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamsArrayFullInner) SetId(v string) {
	o.Id = v
}

func (o TeamsArrayFullInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamsArrayFullInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TeamsArrayFullInner) UnmarshalJSON(data []byte) (err error) {
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

	varTeamsArrayFullInner := _TeamsArrayFullInner{}

	err = json.Unmarshal(data, &varTeamsArrayFullInner)

	if err != nil {
		return err
	}

	*o = TeamsArrayFullInner(varTeamsArrayFullInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTeamsArrayFullInner struct {
	value *TeamsArrayFullInner
	isSet bool
}

func (v NullableTeamsArrayFullInner) Get() *TeamsArrayFullInner {
	return v.value
}

func (v *NullableTeamsArrayFullInner) Set(val *TeamsArrayFullInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsArrayFullInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsArrayFullInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsArrayFullInner(val *TeamsArrayFullInner) *NullableTeamsArrayFullInner {
	return &NullableTeamsArrayFullInner{value: val, isSet: true}
}

func (v NullableTeamsArrayFullInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsArrayFullInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
