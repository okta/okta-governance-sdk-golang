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

// checks if the RiskRuleCriteriaValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRuleCriteriaValue{}

// RiskRuleCriteriaValue struct for RiskRuleCriteriaValue
type RiskRuleCriteriaValue struct {
	// The type of the criteria value
	Type *string `json:"type,omitempty"`
	// Collection of entitlements with associated values
	Value                []ReviewEntitlement `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskRuleCriteriaValue RiskRuleCriteriaValue

// NewRiskRuleCriteriaValue instantiates a new RiskRuleCriteriaValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRuleCriteriaValue() *RiskRuleCriteriaValue {
	this := RiskRuleCriteriaValue{}
	return &this
}

// NewRiskRuleCriteriaValueWithDefaults instantiates a new RiskRuleCriteriaValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleCriteriaValueWithDefaults() *RiskRuleCriteriaValue {
	this := RiskRuleCriteriaValue{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskRuleCriteriaValue) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleCriteriaValue) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskRuleCriteriaValue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskRuleCriteriaValue) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskRuleCriteriaValue) GetValue() []ReviewEntitlement {
	if o == nil || IsNil(o.Value) {
		var ret []ReviewEntitlement
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleCriteriaValue) GetValueOk() ([]ReviewEntitlement, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskRuleCriteriaValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []ReviewEntitlement and assigns it to the Value field.
func (o *RiskRuleCriteriaValue) SetValue(v []ReviewEntitlement) {
	o.Value = v
}

func (o RiskRuleCriteriaValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRuleCriteriaValue) ToMap() (map[string]interface{}, error) {
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

func (o *RiskRuleCriteriaValue) UnmarshalJSON(data []byte) (err error) {
	varRiskRuleCriteriaValue := _RiskRuleCriteriaValue{}

	err = json.Unmarshal(data, &varRiskRuleCriteriaValue)

	if err != nil {
		return err
	}

	*o = RiskRuleCriteriaValue(varRiskRuleCriteriaValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskRuleCriteriaValue struct {
	value *RiskRuleCriteriaValue
	isSet bool
}

func (v NullableRiskRuleCriteriaValue) Get() *RiskRuleCriteriaValue {
	return v.value
}

func (v *NullableRiskRuleCriteriaValue) Set(val *RiskRuleCriteriaValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRuleCriteriaValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRuleCriteriaValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRuleCriteriaValue(val *RiskRuleCriteriaValue) *NullableRiskRuleCriteriaValue {
	return &NullableRiskRuleCriteriaValue{value: val, isSet: true}
}

func (v NullableRiskRuleCriteriaValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRuleCriteriaValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
