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

// checks if the RuleConflictCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleConflictCreatable{}

// RuleConflictCreatable struct for RuleConflictCreatable
type RuleConflictCreatable struct {
	// Name of the resource risk rule
	Name *string `json:"name,omitempty"`
	// Description of the risk rule
	Description *string `json:"description,omitempty"`
	// Additional information about the risk rule
	Notes *string `json:"notes,omitempty"`
	// Risk rule type
	Type *string `json:"type,omitempty"`
	// Resources that the risk rule applies to
	Resources            []RuleConflictResource     `json:"resources,omitempty"`
	ConflictCriteria     *ConflictCriteriaCreatable `json:"conflictCriteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleConflictCreatable RuleConflictCreatable

// NewRuleConflictCreatable instantiates a new RuleConflictCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleConflictCreatable() *RuleConflictCreatable {
	this := RuleConflictCreatable{}
	return &this
}

// NewRuleConflictCreatableWithDefaults instantiates a new RuleConflictCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConflictCreatableWithDefaults() *RuleConflictCreatable {
	this := RuleConflictCreatable{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleConflictCreatable) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCreatable) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleConflictCreatable) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleConflictCreatable) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleConflictCreatable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCreatable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleConflictCreatable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleConflictCreatable) SetDescription(v string) {
	o.Description = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *RuleConflictCreatable) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCreatable) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *RuleConflictCreatable) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *RuleConflictCreatable) SetNotes(v string) {
	o.Notes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RuleConflictCreatable) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCreatable) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RuleConflictCreatable) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RuleConflictCreatable) SetType(v string) {
	o.Type = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *RuleConflictCreatable) GetResources() []RuleConflictResource {
	if o == nil || IsNil(o.Resources) {
		var ret []RuleConflictResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCreatable) GetResourcesOk() ([]RuleConflictResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *RuleConflictCreatable) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []RuleConflictResource and assigns it to the Resources field.
func (o *RuleConflictCreatable) SetResources(v []RuleConflictResource) {
	o.Resources = v
}

// GetConflictCriteria returns the ConflictCriteria field value if set, zero value otherwise.
func (o *RuleConflictCreatable) GetConflictCriteria() ConflictCriteriaCreatable {
	if o == nil || IsNil(o.ConflictCriteria) {
		var ret ConflictCriteriaCreatable
		return ret
	}
	return *o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictCreatable) GetConflictCriteriaOk() (*ConflictCriteriaCreatable, bool) {
	if o == nil || IsNil(o.ConflictCriteria) {
		return nil, false
	}
	return o.ConflictCriteria, true
}

// HasConflictCriteria returns a boolean if a field has been set.
func (o *RuleConflictCreatable) HasConflictCriteria() bool {
	if o != nil && !IsNil(o.ConflictCriteria) {
		return true
	}

	return false
}

// SetConflictCriteria gets a reference to the given ConflictCriteriaCreatable and assigns it to the ConflictCriteria field.
func (o *RuleConflictCreatable) SetConflictCriteria(v ConflictCriteriaCreatable) {
	o.ConflictCriteria = &v
}

func (o RuleConflictCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleConflictCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.ConflictCriteria) {
		toSerialize["conflictCriteria"] = o.ConflictCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleConflictCreatable) UnmarshalJSON(data []byte) (err error) {
	varRuleConflictCreatable := _RuleConflictCreatable{}

	err = json.Unmarshal(data, &varRuleConflictCreatable)

	if err != nil {
		return err
	}

	*o = RuleConflictCreatable(varRuleConflictCreatable)

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

type NullableRuleConflictCreatable struct {
	value *RuleConflictCreatable
	isSet bool
}

func (v NullableRuleConflictCreatable) Get() *RuleConflictCreatable {
	return v.value
}

func (v *NullableRuleConflictCreatable) Set(val *RuleConflictCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConflictCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConflictCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConflictCreatable(val *RuleConflictCreatable) *NullableRuleConflictCreatable {
	return &NullableRuleConflictCreatable{value: val, isSet: true}
}

func (v NullableRuleConflictCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleConflictCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
