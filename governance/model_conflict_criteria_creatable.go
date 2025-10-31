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

// checks if the ConflictCriteriaCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConflictCriteriaCreatable{}

// ConflictCriteriaCreatable Conflict criteria for the risk rule
type ConflictCriteriaCreatable struct {
	// A conflict occurs when the logical AND evaluation of the two criteria is true
	And                  []CriteriaCreatable `json:"and,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConflictCriteriaCreatable ConflictCriteriaCreatable

// NewConflictCriteriaCreatable instantiates a new ConflictCriteriaCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictCriteriaCreatable() *ConflictCriteriaCreatable {
	this := ConflictCriteriaCreatable{}
	return &this
}

// NewConflictCriteriaCreatableWithDefaults instantiates a new ConflictCriteriaCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictCriteriaCreatableWithDefaults() *ConflictCriteriaCreatable {
	this := ConflictCriteriaCreatable{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *ConflictCriteriaCreatable) GetAnd() []CriteriaCreatable {
	if o == nil || IsNil(o.And) {
		var ret []CriteriaCreatable
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictCriteriaCreatable) GetAndOk() ([]CriteriaCreatable, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *ConflictCriteriaCreatable) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []CriteriaCreatable and assigns it to the And field.
func (o *ConflictCriteriaCreatable) SetAnd(v []CriteriaCreatable) {
	o.And = v
}

func (o ConflictCriteriaCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConflictCriteriaCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConflictCriteriaCreatable) UnmarshalJSON(data []byte) (err error) {
	varConflictCriteriaCreatable := _ConflictCriteriaCreatable{}

	err = json.Unmarshal(data, &varConflictCriteriaCreatable)

	if err != nil {
		return err
	}

	*o = ConflictCriteriaCreatable(varConflictCriteriaCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "and")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConflictCriteriaCreatable struct {
	value *ConflictCriteriaCreatable
	isSet bool
}

func (v NullableConflictCriteriaCreatable) Get() *ConflictCriteriaCreatable {
	return v.value
}

func (v *NullableConflictCriteriaCreatable) Set(val *ConflictCriteriaCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictCriteriaCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictCriteriaCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictCriteriaCreatable(val *ConflictCriteriaCreatable) *NullableConflictCriteriaCreatable {
	return &NullableConflictCriteriaCreatable{value: val, isSet: true}
}

func (v NullableConflictCriteriaCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictCriteriaCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
