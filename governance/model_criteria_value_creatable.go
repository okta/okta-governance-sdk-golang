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

// checks if the CriteriaValueCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CriteriaValueCreatable{}

// CriteriaValueCreatable Criteria value
type CriteriaValueCreatable struct {
	// The type of the criteria value. Supported type: `ENTITLEMENTS`
	Type *string `json:"type,omitempty"`
	// Collection of entitlements and associated value identifiers
	Value                []EntitlementCreatable `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CriteriaValueCreatable CriteriaValueCreatable

// NewCriteriaValueCreatable instantiates a new CriteriaValueCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCriteriaValueCreatable() *CriteriaValueCreatable {
	this := CriteriaValueCreatable{}
	return &this
}

// NewCriteriaValueCreatableWithDefaults instantiates a new CriteriaValueCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCriteriaValueCreatableWithDefaults() *CriteriaValueCreatable {
	this := CriteriaValueCreatable{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CriteriaValueCreatable) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CriteriaValueCreatable) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CriteriaValueCreatable) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CriteriaValueCreatable) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CriteriaValueCreatable) GetValue() []EntitlementCreatable {
	if o == nil || IsNil(o.Value) {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CriteriaValueCreatable) GetValueOk() ([]EntitlementCreatable, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CriteriaValueCreatable) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []EntitlementCreatable and assigns it to the Value field.
func (o *CriteriaValueCreatable) SetValue(v []EntitlementCreatable) {
	o.Value = v
}

func (o CriteriaValueCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CriteriaValueCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CriteriaValueCreatable) UnmarshalJSON(data []byte) (err error) {
	varCriteriaValueCreatable := _CriteriaValueCreatable{}

	err = json.Unmarshal(data, &varCriteriaValueCreatable)

	if err != nil {
		return err
	}

	*o = CriteriaValueCreatable(varCriteriaValueCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCriteriaValueCreatable struct {
	value *CriteriaValueCreatable
	isSet bool
}

func (v NullableCriteriaValueCreatable) Get() *CriteriaValueCreatable {
	return v.value
}

func (v *NullableCriteriaValueCreatable) Set(val *CriteriaValueCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCriteriaValueCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCriteriaValueCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCriteriaValueCreatable(val *CriteriaValueCreatable) *NullableCriteriaValueCreatable {
	return &NullableCriteriaValueCreatable{value: val, isSet: true}
}

func (v NullableCriteriaValueCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCriteriaValueCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
