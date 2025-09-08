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

// RiskRuleCriteria struct for RiskRuleCriteria
type RiskRuleCriteria struct {
	// The name of the risk rule criteria
	Name *string `json:"name,omitempty"`
	// Operation to be performed on the criteria value
	Operation            *string                `json:"operation,omitempty"`
	Value                *RiskRuleCriteriaValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskRuleCriteria RiskRuleCriteria

// NewRiskRuleCriteria instantiates a new RiskRuleCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRuleCriteria() *RiskRuleCriteria {
	this := RiskRuleCriteria{}
	return &this
}

// NewRiskRuleCriteriaWithDefaults instantiates a new RiskRuleCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleCriteriaWithDefaults() *RiskRuleCriteria {
	this := RiskRuleCriteria{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RiskRuleCriteria) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleCriteria) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RiskRuleCriteria) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RiskRuleCriteria) SetName(v string) {
	o.Name = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RiskRuleCriteria) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleCriteria) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RiskRuleCriteria) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *RiskRuleCriteria) SetOperation(v string) {
	o.Operation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskRuleCriteria) GetValue() RiskRuleCriteriaValue {
	if o == nil || o.Value == nil {
		var ret RiskRuleCriteriaValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleCriteria) GetValueOk() (*RiskRuleCriteriaValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskRuleCriteria) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given RiskRuleCriteriaValue and assigns it to the Value field.
func (o *RiskRuleCriteria) SetValue(v RiskRuleCriteriaValue) {
	o.Value = &v
}

func (o RiskRuleCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskRuleCriteria) UnmarshalJSON(bytes []byte) (err error) {
	varRiskRuleCriteria := _RiskRuleCriteria{}

	err = json.Unmarshal(bytes, &varRiskRuleCriteria)
	if err == nil {
		*o = RiskRuleCriteria(varRiskRuleCriteria)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskRuleCriteria struct {
	value *RiskRuleCriteria
	isSet bool
}

func (v NullableRiskRuleCriteria) Get() *RiskRuleCriteria {
	return v.value
}

func (v *NullableRiskRuleCriteria) Set(val *RiskRuleCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRuleCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRuleCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRuleCriteria(val *RiskRuleCriteria) *NullableRiskRuleCriteria {
	return &NullableRiskRuleCriteria{value: val, isSet: true}
}

func (v NullableRiskRuleCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRuleCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
