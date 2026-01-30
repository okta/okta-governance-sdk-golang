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

// checks if the RuleConflictCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleConflictCommon{}

// RuleConflictCommon struct for RuleConflictCommon
type RuleConflictCommon struct {
	// Unique identifier for the object
	Id *string `json:"id,omitempty"`
	// Description for the risk rule
	Description *string `json:"description,omitempty"`
	// Risk rule type
	Type                 *string           `json:"type,omitempty"`
	ConflictCriteria     *ConflictCriteria `json:"conflictCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleConflictCommon RuleConflictCommon

// NewRuleConflictCommon instantiates a new RuleConflictCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleConflictCommon() *RuleConflictCommon {
	this := RuleConflictCommon{}
	return &this
}

// NewRuleConflictCommonWithDefaults instantiates a new RuleConflictCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConflictCommonWithDefaults() *RuleConflictCommon {
	this := RuleConflictCommon{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleConflictCommon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCommon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleConflictCommon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleConflictCommon) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleConflictCommon) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCommon) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleConflictCommon) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleConflictCommon) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RuleConflictCommon) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCommon) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RuleConflictCommon) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RuleConflictCommon) SetType(v string) {
	o.Type = &v
}

// GetConflictCriteria returns the ConflictCriteria field value if set, zero value otherwise.
func (o *RuleConflictCommon) GetConflictCriteria() ConflictCriteria {
	if o == nil || IsNil(o.ConflictCriteria) {
		var ret ConflictCriteria
		return ret
	}
	return *o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCommon) GetConflictCriteriaOk() (*ConflictCriteria, bool) {
	if o == nil || IsNil(o.ConflictCriteria) {
		return nil, false
	}
	return o.ConflictCriteria, true
}

// HasConflictCriteria returns a boolean if a field has been set.
func (o *RuleConflictCommon) HasConflictCriteria() bool {
	if o != nil && !IsNil(o.ConflictCriteria) {
		return true
	}

	return false
}

// SetConflictCriteria gets a reference to the given ConflictCriteria and assigns it to the ConflictCriteria field.
func (o *RuleConflictCommon) SetConflictCriteria(v ConflictCriteria) {
	o.ConflictCriteria = &v
}

func (o RuleConflictCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleConflictCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ConflictCriteria) {
		toSerialize["conflictCriteria"] = o.ConflictCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleConflictCommon) UnmarshalJSON(data []byte) (err error) {
	varRuleConflictCommon := _RuleConflictCommon{}

	err = json.Unmarshal(data, &varRuleConflictCommon)

	if err != nil {
		return err
	}

	*o = RuleConflictCommon(varRuleConflictCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "conflictCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleConflictCommon struct {
	value *RuleConflictCommon
	isSet bool
}

func (v NullableRuleConflictCommon) Get() *RuleConflictCommon {
	return v.value
}

func (v *NullableRuleConflictCommon) Set(val *RuleConflictCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConflictCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConflictCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConflictCommon(val *RuleConflictCommon) *NullableRuleConflictCommon {
	return &NullableRuleConflictCommon{value: val, isSet: true}
}

func (v NullableRuleConflictCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleConflictCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
