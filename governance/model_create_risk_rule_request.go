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
	"fmt"
)

// checks if the CreateRiskRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRiskRuleRequest{}

// CreateRiskRuleRequest struct for CreateRiskRuleRequest
type CreateRiskRuleRequest struct {
	// Name of the resource risk rule
	Name string `json:"name"`
	// Description of the risk rule
	Description *string `json:"description,omitempty"`
	// Additional information about the risk rule
	Notes *string `json:"notes,omitempty"`
	// Risk rule type
	Type string `json:"type"`
	// Resources that the risk rule applies to
	Resources            []RuleConflictResource    `json:"resources"`
	ConflictCriteria     ConflictCriteriaCreatable `json:"conflictCriteria"`
	AdditionalProperties map[string]interface{}
}

type _CreateRiskRuleRequest CreateRiskRuleRequest

// NewCreateRiskRuleRequest instantiates a new CreateRiskRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRiskRuleRequest(name string, type_ string, resources []RuleConflictResource, conflictCriteria ConflictCriteriaCreatable) *CreateRiskRuleRequest {
	this := CreateRiskRuleRequest{}
	this.Name = name
	this.Type = type_
	this.Resources = resources
	this.ConflictCriteria = conflictCriteria
	return &this
}

// NewCreateRiskRuleRequestWithDefaults instantiates a new CreateRiskRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRiskRuleRequestWithDefaults() *CreateRiskRuleRequest {
	this := CreateRiskRuleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRiskRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRiskRuleRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRiskRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRiskRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRiskRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRiskRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateRiskRuleRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRiskRuleRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateRiskRuleRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateRiskRuleRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetType returns the Type field value
func (o *CreateRiskRuleRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRuleRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateRiskRuleRequest) SetType(v string) {
	o.Type = v
}

// GetResources returns the Resources field value
func (o *CreateRiskRuleRequest) GetResources() []RuleConflictResource {
	if o == nil {
		var ret []RuleConflictResource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRuleRequest) GetResourcesOk() ([]RuleConflictResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *CreateRiskRuleRequest) SetResources(v []RuleConflictResource) {
	o.Resources = v
}

// GetConflictCriteria returns the ConflictCriteria field value
func (o *CreateRiskRuleRequest) GetConflictCriteria() ConflictCriteriaCreatable {
	if o == nil {
		var ret ConflictCriteriaCreatable
		return ret
	}

	return o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRuleRequest) GetConflictCriteriaOk() (*ConflictCriteriaCreatable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConflictCriteria, true
}

// SetConflictCriteria sets field value
func (o *CreateRiskRuleRequest) SetConflictCriteria(v ConflictCriteriaCreatable) {
	o.ConflictCriteria = v
}

func (o CreateRiskRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRiskRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["type"] = o.Type
	toSerialize["resources"] = o.Resources
	toSerialize["conflictCriteria"] = o.ConflictCriteria

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRiskRuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"resources",
		"conflictCriteria",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRiskRuleRequest := _CreateRiskRuleRequest{}

	err = json.Unmarshal(data, &varCreateRiskRuleRequest)

	if err != nil {
		return err
	}

	*o = CreateRiskRuleRequest(varCreateRiskRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "type")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "conflictCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRiskRuleRequest struct {
	value *CreateRiskRuleRequest
	isSet bool
}

func (v NullableCreateRiskRuleRequest) Get() *CreateRiskRuleRequest {
	return v.value
}

func (v *NullableCreateRiskRuleRequest) Set(val *CreateRiskRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRiskRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRiskRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRiskRuleRequest(val *CreateRiskRuleRequest) *NullableCreateRiskRuleRequest {
	return &NullableCreateRiskRuleRequest{value: val, isSet: true}
}

func (v NullableCreateRiskRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRiskRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
