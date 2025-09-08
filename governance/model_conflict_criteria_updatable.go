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

// ConflictCriteriaUpdatable Conflict criteria for the risk rule
type ConflictCriteriaUpdatable struct {
	// A conflict occurs when the logical AND evaluation of the two criteria is true
	And                  []CriteriaCreatable `json:"and,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConflictCriteriaUpdatable ConflictCriteriaUpdatable

// NewConflictCriteriaUpdatable instantiates a new ConflictCriteriaUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictCriteriaUpdatable() *ConflictCriteriaUpdatable {
	this := ConflictCriteriaUpdatable{}
	return &this
}

// NewConflictCriteriaUpdatableWithDefaults instantiates a new ConflictCriteriaUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictCriteriaUpdatableWithDefaults() *ConflictCriteriaUpdatable {
	this := ConflictCriteriaUpdatable{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *ConflictCriteriaUpdatable) GetAnd() []CriteriaCreatable {
	if o == nil || o.And == nil {
		var ret []CriteriaCreatable
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictCriteriaUpdatable) GetAndOk() ([]CriteriaCreatable, bool) {
	if o == nil || o.And == nil {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *ConflictCriteriaUpdatable) HasAnd() bool {
	if o != nil && o.And != nil {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []CriteriaCreatable and assigns it to the And field.
func (o *ConflictCriteriaUpdatable) SetAnd(v []CriteriaCreatable) {
	o.And = v
}

func (o ConflictCriteriaUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.And != nil {
		toSerialize["and"] = o.And
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConflictCriteriaUpdatable) UnmarshalJSON(bytes []byte) (err error) {
	varConflictCriteriaUpdatable := _ConflictCriteriaUpdatable{}

	err = json.Unmarshal(bytes, &varConflictCriteriaUpdatable)
	if err == nil {
		*o = ConflictCriteriaUpdatable(varConflictCriteriaUpdatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "and")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableConflictCriteriaUpdatable struct {
	value *ConflictCriteriaUpdatable
	isSet bool
}

func (v NullableConflictCriteriaUpdatable) Get() *ConflictCriteriaUpdatable {
	return v.value
}

func (v *NullableConflictCriteriaUpdatable) Set(val *ConflictCriteriaUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictCriteriaUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictCriteriaUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictCriteriaUpdatable(val *ConflictCriteriaUpdatable) *NullableConflictCriteriaUpdatable {
	return &NullableConflictCriteriaUpdatable{value: val, isSet: true}
}

func (v NullableConflictCriteriaUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictCriteriaUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
