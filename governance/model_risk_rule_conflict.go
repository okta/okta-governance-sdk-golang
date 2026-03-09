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
	"time"
)

// checks if the RiskRuleConflict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRuleConflict{}

// RiskRuleConflict struct for RiskRuleConflict
type RiskRuleConflict struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// Description for the risk rule
	Description *string `json:"description,omitempty"`
	// Risk rule type
	Type             *string           `json:"type,omitempty"`
	ConflictCriteria *ConflictCriteria `json:"conflictCriteria,omitempty"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// Links to related resources
	Links *map[string]Link `json:"_links,omitempty"`
	// Name of the resource risk rule
	Name *string `json:"name,omitempty"`
	// Additional information about the risk rule
	Notes *string `json:"notes,omitempty"`
	// Resources that the risk rule applies to
	Resources []RuleConflictResource `json:"resources,omitempty"`
	// Status of the risk rule
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskRuleConflict RiskRuleConflict

// NewRiskRuleConflict instantiates a new RiskRuleConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRuleConflict(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *RiskRuleConflict {
	this := RiskRuleConflict{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	return &this
}

// NewRiskRuleConflictWithDefaults instantiates a new RiskRuleConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleConflictWithDefaults() *RiskRuleConflict {
	this := RiskRuleConflict{}
	return &this
}

// GetId returns the Id field value
func (o *RiskRuleConflict) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RiskRuleConflict) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskRuleConflict) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskRuleConflict) SetType(v string) {
	o.Type = &v
}

// GetConflictCriteria returns the ConflictCriteria field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetConflictCriteria() ConflictCriteria {
	if o == nil || IsNil(o.ConflictCriteria) {
		var ret ConflictCriteria
		return ret
	}
	return *o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetConflictCriteriaOk() (*ConflictCriteria, bool) {
	if o == nil || IsNil(o.ConflictCriteria) {
		return nil, false
	}
	return o.ConflictCriteria, true
}

// HasConflictCriteria returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasConflictCriteria() bool {
	if o != nil && !IsNil(o.ConflictCriteria) {
		return true
	}

	return false
}

// SetConflictCriteria gets a reference to the given ConflictCriteria and assigns it to the ConflictCriteria field.
func (o *RiskRuleConflict) SetConflictCriteria(v ConflictCriteria) {
	o.ConflictCriteria = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RiskRuleConflict) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RiskRuleConflict) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RiskRuleConflict) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RiskRuleConflict) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RiskRuleConflict) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RiskRuleConflict) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RiskRuleConflict) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RiskRuleConflict) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetLinks() map[string]Link {
	if o == nil || IsNil(o.Links) {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *RiskRuleConflict) SetLinks(v map[string]Link) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RiskRuleConflict) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *RiskRuleConflict) SetNotes(v string) {
	o.Notes = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetResources() []RuleConflictResource {
	if o == nil || IsNil(o.Resources) {
		var ret []RuleConflictResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetResourcesOk() ([]RuleConflictResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []RuleConflictResource and assigns it to the Resources field.
func (o *RiskRuleConflict) SetResources(v []RuleConflictResource) {
	o.Resources = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RiskRuleConflict) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleConflict) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RiskRuleConflict) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RiskRuleConflict) SetStatus(v string) {
	o.Status = &v
}

func (o RiskRuleConflict) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRuleConflict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ConflictCriteria) {
		toSerialize["conflictCriteria"] = o.ConflictCriteria
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskRuleConflict) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
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

	varRiskRuleConflict := _RiskRuleConflict{}

	err = json.Unmarshal(data, &varRiskRuleConflict)

	if err != nil {
		return err
	}

	*o = RiskRuleConflict(varRiskRuleConflict)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "conflictCriteria")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskRuleConflict struct {
	value *RiskRuleConflict
	isSet bool
}

func (v NullableRiskRuleConflict) Get() *RiskRuleConflict {
	return v.value
}

func (v *NullableRiskRuleConflict) Set(val *RiskRuleConflict) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRuleConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRuleConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRuleConflict(val *RiskRuleConflict) *NullableRiskRuleConflict {
	return &NullableRiskRuleConflict{value: val, isSet: true}
}

func (v NullableRiskRuleConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRuleConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
