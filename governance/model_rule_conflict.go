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

// RuleConflict struct for RuleConflict
type RuleConflict struct {
	// The Okta user `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	PrincipalOrn *string `json:"principalOrn,omitempty"`
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn *string `json:"resourceOrn,omitempty"`
	// Unique identifier for rule object
	RuleId *string `json:"ruleId,omitempty"`
	// The name of a resource rule causing the conflict
	RuleName *string `json:"ruleName,omitempty"`
	// Unique identifier for the object
	Id *string `json:"id,omitempty"`
	// Description for the risk rule
	Description *string `json:"description,omitempty"`
	// Risk rule type
	Type                 *string           `json:"type,omitempty"`
	ConflictCriteria     *ConflictCriteria `json:"conflictCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleConflict RuleConflict

// NewRuleConflict instantiates a new RuleConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleConflict() *RuleConflict {
	this := RuleConflict{}
	return &this
}

// NewRuleConflictWithDefaults instantiates a new RuleConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConflictWithDefaults() *RuleConflict {
	this := RuleConflict{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value if set, zero value otherwise.
func (o *RuleConflict) GetPrincipalOrn() string {
	if o == nil || o.PrincipalOrn == nil {
		var ret string
		return ret
	}
	return *o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetPrincipalOrnOk() (*string, bool) {
	if o == nil || o.PrincipalOrn == nil {
		return nil, false
	}
	return o.PrincipalOrn, true
}

// HasPrincipalOrn returns a boolean if a field has been set.
func (o *RuleConflict) HasPrincipalOrn() bool {
	if o != nil && o.PrincipalOrn != nil {
		return true
	}

	return false
}

// SetPrincipalOrn gets a reference to the given string and assigns it to the PrincipalOrn field.
func (o *RuleConflict) SetPrincipalOrn(v string) {
	o.PrincipalOrn = &v
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *RuleConflict) GetResourceOrn() string {
	if o == nil || o.ResourceOrn == nil {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetResourceOrnOk() (*string, bool) {
	if o == nil || o.ResourceOrn == nil {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *RuleConflict) HasResourceOrn() bool {
	if o != nil && o.ResourceOrn != nil {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *RuleConflict) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *RuleConflict) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *RuleConflict) HasRuleId() bool {
	if o != nil && o.RuleId != nil {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *RuleConflict) SetRuleId(v string) {
	o.RuleId = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *RuleConflict) GetRuleName() string {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetRuleNameOk() (*string, bool) {
	if o == nil || o.RuleName == nil {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *RuleConflict) HasRuleName() bool {
	if o != nil && o.RuleName != nil {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *RuleConflict) SetRuleName(v string) {
	o.RuleName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleConflict) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleConflict) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleConflict) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleConflict) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleConflict) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleConflict) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RuleConflict) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RuleConflict) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RuleConflict) SetType(v string) {
	o.Type = &v
}

// GetConflictCriteria returns the ConflictCriteria field value if set, zero value otherwise.
func (o *RuleConflict) GetConflictCriteria() ConflictCriteria {
	if o == nil || o.ConflictCriteria == nil {
		var ret ConflictCriteria
		return ret
	}
	return *o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflict) GetConflictCriteriaOk() (*ConflictCriteria, bool) {
	if o == nil || o.ConflictCriteria == nil {
		return nil, false
	}
	return o.ConflictCriteria, true
}

// HasConflictCriteria returns a boolean if a field has been set.
func (o *RuleConflict) HasConflictCriteria() bool {
	if o != nil && o.ConflictCriteria != nil {
		return true
	}

	return false
}

// SetConflictCriteria gets a reference to the given ConflictCriteria and assigns it to the ConflictCriteria field.
func (o *RuleConflict) SetConflictCriteria(v ConflictCriteria) {
	o.ConflictCriteria = &v
}

func (o RuleConflict) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrincipalOrn != nil {
		toSerialize["principalOrn"] = o.PrincipalOrn
	}
	if o.ResourceOrn != nil {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if o.RuleId != nil {
		toSerialize["ruleId"] = o.RuleId
	}
	if o.RuleName != nil {
		toSerialize["ruleName"] = o.RuleName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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

func (o *RuleConflict) UnmarshalJSON(bytes []byte) (err error) {
	varRuleConflict := _RuleConflict{}

	err = json.Unmarshal(bytes, &varRuleConflict)
	if err == nil {
		*o = RuleConflict(varRuleConflict)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "ruleId")
		delete(additionalProperties, "ruleName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "conflictCriteria")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRuleConflict struct {
	value *RuleConflict
	isSet bool
}

func (v NullableRuleConflict) Get() *RuleConflict {
	return v.value
}

func (v *NullableRuleConflict) Set(val *RuleConflict) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConflict(val *RuleConflict) *NullableRuleConflict {
	return &NullableRuleConflict{value: val, isSet: true}
}

func (v NullableRuleConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
