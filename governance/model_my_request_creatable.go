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

// checks if the MyRequestCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyRequestCreatable{}

// MyRequestCreatable The properties expected in an initial create request
type MyRequestCreatable struct {
	// The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.
	RequesterFieldValues []RequestFieldValue `json:"requesterFieldValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyRequestCreatable MyRequestCreatable

// NewMyRequestCreatable instantiates a new MyRequestCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyRequestCreatable() *MyRequestCreatable {
	this := MyRequestCreatable{}
	return &this
}

// NewMyRequestCreatableWithDefaults instantiates a new MyRequestCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyRequestCreatableWithDefaults() *MyRequestCreatable {
	this := MyRequestCreatable{}
	return &this
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise.
func (o *MyRequestCreatable) GetRequesterFieldValues() []RequestFieldValue {
	if o == nil || IsNil(o.RequesterFieldValues) {
		var ret []RequestFieldValue
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyRequestCreatable) GetRequesterFieldValuesOk() ([]RequestFieldValue, bool) {
	if o == nil || IsNil(o.RequesterFieldValues) {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *MyRequestCreatable) HasRequesterFieldValues() bool {
	if o != nil && !IsNil(o.RequesterFieldValues) {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []RequestFieldValue and assigns it to the RequesterFieldValues field.
func (o *MyRequestCreatable) SetRequesterFieldValues(v []RequestFieldValue) {
	o.RequesterFieldValues = v
}

func (o MyRequestCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyRequestCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequesterFieldValues) {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyRequestCreatable) UnmarshalJSON(data []byte) (err error) {
	varMyRequestCreatable := _MyRequestCreatable{}

	err = json.Unmarshal(data, &varMyRequestCreatable)

	if err != nil {
		return err
	}

	*o = MyRequestCreatable(varMyRequestCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requesterFieldValues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyRequestCreatable struct {
	value *MyRequestCreatable
	isSet bool
}

func (v NullableMyRequestCreatable) Get() *MyRequestCreatable {
	return v.value
}

func (v *NullableMyRequestCreatable) Set(val *MyRequestCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableMyRequestCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableMyRequestCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyRequestCreatable(val *MyRequestCreatable) *NullableMyRequestCreatable {
	return &NullableMyRequestCreatable{value: val, isSet: true}
}

func (v NullableMyRequestCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyRequestCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
