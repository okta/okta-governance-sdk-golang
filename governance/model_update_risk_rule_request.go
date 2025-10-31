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

// checks if the UpdateRiskRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRiskRuleRequest{}

// UpdateRiskRuleRequest struct for UpdateRiskRuleRequest
type UpdateRiskRuleRequest struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// Name of the resource risk rule
	Name NullableString `json:"name,omitempty"`
	// Additional information about the rule
	Notes NullableString `json:"notes,omitempty"`
	// Description of the risk rule
	Description          NullableString                    `json:"description,omitempty"`
	ConflictCriteria     NullableConflictCriteriaUpdatable `json:"conflictCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRiskRuleRequest UpdateRiskRuleRequest

// NewUpdateRiskRuleRequest instantiates a new UpdateRiskRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRiskRuleRequest(id string) *UpdateRiskRuleRequest {
	this := UpdateRiskRuleRequest{}
	this.Id = id
	return &this
}

// NewUpdateRiskRuleRequestWithDefaults instantiates a new UpdateRiskRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRiskRuleRequestWithDefaults() *UpdateRiskRuleRequest {
	this := UpdateRiskRuleRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateRiskRuleRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateRiskRuleRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateRiskRuleRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRiskRuleRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRiskRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRiskRuleRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateRiskRuleRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateRiskRuleRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateRiskRuleRequest) UnsetName() {
	o.Name.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRiskRuleRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRiskRuleRequest) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateRiskRuleRequest) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *UpdateRiskRuleRequest) SetNotes(v string) {
	o.Notes.Set(&v)
}

// SetNotesNil sets the value for Notes to be an explicit nil
func (o *UpdateRiskRuleRequest) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *UpdateRiskRuleRequest) UnsetNotes() {
	o.Notes.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRiskRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRiskRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRiskRuleRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateRiskRuleRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateRiskRuleRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateRiskRuleRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetConflictCriteria returns the ConflictCriteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRiskRuleRequest) GetConflictCriteria() ConflictCriteriaUpdatable {
	if o == nil || IsNil(o.ConflictCriteria.Get()) {
		var ret ConflictCriteriaUpdatable
		return ret
	}
	return *o.ConflictCriteria.Get()
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRiskRuleRequest) GetConflictCriteriaOk() (*ConflictCriteriaUpdatable, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConflictCriteria.Get(), o.ConflictCriteria.IsSet()
}

// HasConflictCriteria returns a boolean if a field has been set.
func (o *UpdateRiskRuleRequest) HasConflictCriteria() bool {
	if o != nil && o.ConflictCriteria.IsSet() {
		return true
	}

	return false
}

// SetConflictCriteria gets a reference to the given NullableConflictCriteriaUpdatable and assigns it to the ConflictCriteria field.
func (o *UpdateRiskRuleRequest) SetConflictCriteria(v ConflictCriteriaUpdatable) {
	o.ConflictCriteria.Set(&v)
}

// SetConflictCriteriaNil sets the value for ConflictCriteria to be an explicit nil
func (o *UpdateRiskRuleRequest) SetConflictCriteriaNil() {
	o.ConflictCriteria.Set(nil)
}

// UnsetConflictCriteria ensures that no value is present for ConflictCriteria, not even an explicit nil
func (o *UpdateRiskRuleRequest) UnsetConflictCriteria() {
	o.ConflictCriteria.Unset()
}

func (o UpdateRiskRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRiskRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ConflictCriteria.IsSet() {
		toSerialize["conflictCriteria"] = o.ConflictCriteria.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRiskRuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUpdateRiskRuleRequest := _UpdateRiskRuleRequest{}

	err = json.Unmarshal(data, &varUpdateRiskRuleRequest)

	if err != nil {
		return err
	}

	*o = UpdateRiskRuleRequest(varUpdateRiskRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "description")
		delete(additionalProperties, "conflictCriteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRiskRuleRequest struct {
	value *UpdateRiskRuleRequest
	isSet bool
}

func (v NullableUpdateRiskRuleRequest) Get() *UpdateRiskRuleRequest {
	return v.value
}

func (v *NullableUpdateRiskRuleRequest) Set(val *UpdateRiskRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRiskRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRiskRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRiskRuleRequest(val *UpdateRiskRuleRequest) *NullableUpdateRiskRuleRequest {
	return &NullableUpdateRiskRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateRiskRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRiskRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
