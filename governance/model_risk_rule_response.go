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
	"time"
)

// RiskRuleResponse struct for RiskRuleResponse
type RiskRuleResponse struct {
	Links LinkSelf `json:"_links"`
	// Unique identifier for the object
	Id string `json:"id"`
	// Description for the risk rule
	Description *string `json:"description,omitempty"`
	// Risk rule type
	Type             string           `json:"type"`
	ConflictCriteria ConflictCriteria `json:"conflictCriteria"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// Name of the resource risk rule
	Name string `json:"name"`
	// Additional information about the risk rule
	Notes *string `json:"notes,omitempty"`
	// Resources that the risk rule applies to
	Resources []RuleConflictResource `json:"resources"`
	// Status of the risk rule
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _RiskRuleResponse RiskRuleResponse

// NewRiskRuleResponse instantiates a new RiskRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRuleResponse(links LinkSelf, id string, type_ string, conflictCriteria ConflictCriteria, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, name string, resources []RuleConflictResource, status string) *RiskRuleResponse {
	this := RiskRuleResponse{}
	this.Id = id
	this.Type = type_
	this.ConflictCriteria = conflictCriteria
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.Name = name
	this.Resources = resources
	this.Status = status
	return &this
}

// NewRiskRuleResponseWithDefaults instantiates a new RiskRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleResponseWithDefaults() *RiskRuleResponse {
	this := RiskRuleResponse{}
	return &this
}

// GetLinks returns the Links field value
func (o *RiskRuleResponse) GetLinks() LinkSelf {
	if o == nil {
		var ret LinkSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetLinksOk() (*LinkSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RiskRuleResponse) SetLinks(v LinkSelf) {
	o.Links = v
}

// GetId returns the Id field value
func (o *RiskRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RiskRuleResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskRuleResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskRuleResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskRuleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *RiskRuleResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskRuleResponse) SetType(v string) {
	o.Type = v
}

// GetConflictCriteria returns the ConflictCriteria field value
func (o *RiskRuleResponse) GetConflictCriteria() ConflictCriteria {
	if o == nil {
		var ret ConflictCriteria
		return ret
	}

	return o.ConflictCriteria
}

// GetConflictCriteriaOk returns a tuple with the ConflictCriteria field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetConflictCriteriaOk() (*ConflictCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConflictCriteria, true
}

// SetConflictCriteria sets field value
func (o *RiskRuleResponse) SetConflictCriteria(v ConflictCriteria) {
	o.ConflictCriteria = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RiskRuleResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RiskRuleResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RiskRuleResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RiskRuleResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RiskRuleResponse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RiskRuleResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RiskRuleResponse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RiskRuleResponse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetName returns the Name field value
func (o *RiskRuleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskRuleResponse) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *RiskRuleResponse) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *RiskRuleResponse) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *RiskRuleResponse) SetNotes(v string) {
	o.Notes = &v
}

// GetResources returns the Resources field value
func (o *RiskRuleResponse) GetResources() []RuleConflictResource {
	if o == nil {
		var ret []RuleConflictResource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetResourcesOk() ([]RuleConflictResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *RiskRuleResponse) SetResources(v []RuleConflictResource) {
	o.Resources = v
}

// GetStatus returns the Status field value
func (o *RiskRuleResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RiskRuleResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RiskRuleResponse) SetStatus(v string) {
	o.Status = v
}

func (o RiskRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["conflictCriteria"] = o.ConflictCriteria
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if true {
		toSerialize["resources"] = o.Resources
	}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRiskRuleResponse := _RiskRuleResponse{}

	err = json.Unmarshal(bytes, &varRiskRuleResponse)
	if err == nil {
		*o = RiskRuleResponse(varRiskRuleResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "conflictCriteria")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskRuleResponse struct {
	value *RiskRuleResponse
	isSet bool
}

func (v NullableRiskRuleResponse) Get() *RiskRuleResponse {
	return v.value
}

func (v *NullableRiskRuleResponse) Set(val *RiskRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRuleResponse(val *RiskRuleResponse) *NullableRiskRuleResponse {
	return &NullableRiskRuleResponse{value: val, isSet: true}
}

func (v NullableRiskRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
