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

// checks if the RiskRuleConflictsConflictCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRuleConflictsConflictCriteria{}

// RiskRuleConflictsConflictCriteria struct for RiskRuleConflictsConflictCriteria
type RiskRuleConflictsConflictCriteria struct {
	And                  []RiskRuleCriteria `json:"and,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskRuleConflictsConflictCriteria RiskRuleConflictsConflictCriteria

// NewRiskRuleConflictsConflictCriteria instantiates a new RiskRuleConflictsConflictCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRuleConflictsConflictCriteria() *RiskRuleConflictsConflictCriteria {
	this := RiskRuleConflictsConflictCriteria{}
	return &this
}

// NewRiskRuleConflictsConflictCriteriaWithDefaults instantiates a new RiskRuleConflictsConflictCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleConflictsConflictCriteriaWithDefaults() *RiskRuleConflictsConflictCriteria {
	this := RiskRuleConflictsConflictCriteria{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *RiskRuleConflictsConflictCriteria) GetAnd() []RiskRuleCriteria {
	if o == nil || IsNil(o.And) {
		var ret []RiskRuleCriteria
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflictsConflictCriteria) GetAndOk() ([]RiskRuleCriteria, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *RiskRuleConflictsConflictCriteria) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []RiskRuleCriteria and assigns it to the And field.
func (o *RiskRuleConflictsConflictCriteria) SetAnd(v []RiskRuleCriteria) {
	o.And = v
}

func (o RiskRuleConflictsConflictCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRuleConflictsConflictCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskRuleConflictsConflictCriteria) UnmarshalJSON(data []byte) (err error) {
	varRiskRuleConflictsConflictCriteria := _RiskRuleConflictsConflictCriteria{}

	err = json.Unmarshal(data, &varRiskRuleConflictsConflictCriteria)

	if err != nil {
		return err
	}

	*o = RiskRuleConflictsConflictCriteria(varRiskRuleConflictsConflictCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "and")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskRuleConflictsConflictCriteria struct {
	value *RiskRuleConflictsConflictCriteria
	isSet bool
}

func (v NullableRiskRuleConflictsConflictCriteria) Get() *RiskRuleConflictsConflictCriteria {
	return v.value
}

func (v *NullableRiskRuleConflictsConflictCriteria) Set(val *RiskRuleConflictsConflictCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRuleConflictsConflictCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRuleConflictsConflictCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRuleConflictsConflictCriteria(val *RiskRuleConflictsConflictCriteria) *NullableRiskRuleConflictsConflictCriteria {
	return &NullableRiskRuleConflictsConflictCriteria{value: val, isSet: true}
}

func (v NullableRiskRuleConflictsConflictCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRuleConflictsConflictCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
