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

// RiskRuleConflicts struct for RiskRuleConflicts
type RiskRuleConflicts struct {
	// Unique identifier of the risk rule
	Id *string `json:"id,omitempty"`
	// The type of the rule
	Type                 *string                            `json:"type,omitempty"`
	ConflictCriteria     *RiskRuleConflictsConflictCriteria `json:"conflictCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskRuleConflicts RiskRuleConflicts

// NewRiskRuleConflicts instantiates a new RiskRuleConflicts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRuleConflicts() *RiskRuleConflicts {
	this := RiskRuleConflicts{}
	return &this
}

// NewRiskRuleConflictsWithDefaults instantiates a new RiskRuleConflicts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleConflictsWithDefaults() *RiskRuleConflicts {
	this := RiskRuleConflicts{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskRuleConflicts) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflicts) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskRuleConflicts) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskRuleConflicts) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskRuleConflicts) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflicts) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskRuleConflicts) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskRuleConflicts) SetType(v string) {
	o.Type = &v
}

// GetConflictCriteria returns the ConflictCriteria field value if set, zero value otherwise.
func (o *RiskRuleConflicts) GetConflictCriteria() RiskRuleConflictsConflictCriteria {
	if o == nil || o.ConflictCriteria == nil {
		var ret RiskRuleConflictsConflictCriteria
		return ret
	}
	return *o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflicts) GetConflictCriteriaOk() (*RiskRuleConflictsConflictCriteria, bool) {
	if o == nil || o.ConflictCriteria == nil {
		return nil, false
	}
	return o.ConflictCriteria, true
}

// HasConflictCriteria returns a boolean if a field has been set.
func (o *RiskRuleConflicts) HasConflictCriteria() bool {
	if o != nil && o.ConflictCriteria != nil {
		return true
	}

	return false
}

// SetConflictCriteria gets a reference to the given RiskRuleConflictsConflictCriteria and assigns it to the ConflictCriteria field.
func (o *RiskRuleConflicts) SetConflictCriteria(v RiskRuleConflictsConflictCriteria) {
	o.ConflictCriteria = &v
}

func (o RiskRuleConflicts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ConflictCriteria != nil {
		toSerialize["conflictCriteria"] = o.ConflictCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskRuleConflicts) UnmarshalJSON(bytes []byte) (err error) {
	varRiskRuleConflicts := _RiskRuleConflicts{}

	err = json.Unmarshal(bytes, &varRiskRuleConflicts)
	if err == nil {
		*o = RiskRuleConflicts(varRiskRuleConflicts)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "conflictCriteria")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskRuleConflicts struct {
	value *RiskRuleConflicts
	isSet bool
}

func (v NullableRiskRuleConflicts) Get() *RiskRuleConflicts {
	return v.value
}

func (v *NullableRiskRuleConflicts) Set(val *RiskRuleConflicts) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRuleConflicts) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRuleConflicts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRuleConflicts(val *RiskRuleConflicts) *NullableRiskRuleConflicts {
	return &NullableRiskRuleConflicts{value: val, isSet: true}
}

func (v NullableRiskRuleConflicts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRuleConflicts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
