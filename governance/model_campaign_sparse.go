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

// CampaignSparse Sparse representation of a Campaign resource.
type CampaignSparse struct {
	// Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed.
	Name string `json:"name"`
	// Human readable description.
	Description *string `json:"description,omitempty"`
	// The date on which the campaign is scheduled to start. Accepts date in ISO 8601 format.
	StartDate time.Time `json:"startDate"`
	// The date on which the campaign is supposed to end. Accepts date in ISO 8601 format.
	EndDate      time.Time    `json:"endDate"`
	ScheduleType ScheduleType `json:"scheduleType"`
	// ID of the recurring campaign if this campaign was created as part of a recurring schedule.
	RecurringCampaignId NullableString       `json:"recurringCampaignId,omitempty"`
	ReviewerType        CampaignReviewerType `json:"reviewerType"`
	Links               CampaignLinks        `json:"_links"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy        string         `json:"lastUpdatedBy"`
	Status               CampaignStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _CampaignSparse CampaignSparse

// NewCampaignSparse instantiates a new CampaignSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSparse(name string, startDate time.Time, endDate time.Time, scheduleType ScheduleType, reviewerType CampaignReviewerType, links CampaignLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, status CampaignStatus) *CampaignSparse {
	this := CampaignSparse{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.Status = status
	return &this
}

// NewCampaignSparseWithDefaults instantiates a new CampaignSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSparseWithDefaults() *CampaignSparse {
	this := CampaignSparse{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignSparse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignSparse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignSparse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignSparse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignSparse) SetDescription(v string) {
	o.Description = &v
}

// GetStartDate returns the StartDate field value
func (o *CampaignSparse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CampaignSparse) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *CampaignSparse) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *CampaignSparse) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetScheduleType returns the ScheduleType field value
func (o *CampaignSparse) GetScheduleType() ScheduleType {
	if o == nil {
		var ret ScheduleType
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetScheduleTypeOk() (*ScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *CampaignSparse) SetScheduleType(v ScheduleType) {
	o.ScheduleType = v
}

// GetRecurringCampaignId returns the RecurringCampaignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignSparse) GetRecurringCampaignId() string {
	if o == nil || o.RecurringCampaignId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecurringCampaignId.Get()
}

// GetRecurringCampaignIdOk returns a tuple with the RecurringCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignSparse) GetRecurringCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringCampaignId.Get(), o.RecurringCampaignId.IsSet()
}

// HasRecurringCampaignId returns a boolean if a field has been set.
func (o *CampaignSparse) HasRecurringCampaignId() bool {
	if o != nil && o.RecurringCampaignId.IsSet() {
		return true
	}

	return false
}

// SetRecurringCampaignId gets a reference to the given NullableString and assigns it to the RecurringCampaignId field.
func (o *CampaignSparse) SetRecurringCampaignId(v string) {
	o.RecurringCampaignId.Set(&v)
}

// SetRecurringCampaignIdNil sets the value for RecurringCampaignId to be an explicit nil
func (o *CampaignSparse) SetRecurringCampaignIdNil() {
	o.RecurringCampaignId.Set(nil)
}

// UnsetRecurringCampaignId ensures that no value is present for RecurringCampaignId, not even an explicit nil
func (o *CampaignSparse) UnsetRecurringCampaignId() {
	o.RecurringCampaignId.Unset()
}

// GetReviewerType returns the ReviewerType field value
func (o *CampaignSparse) GetReviewerType() CampaignReviewerType {
	if o == nil {
		var ret CampaignReviewerType
		return ret
	}

	return o.ReviewerType
}

// GetReviewerTypeOk returns a tuple with the ReviewerType field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetReviewerTypeOk() (*CampaignReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerType, true
}

// SetReviewerType sets field value
func (o *CampaignSparse) SetReviewerType(v CampaignReviewerType) {
	o.ReviewerType = v
}

// GetLinks returns the Links field value
func (o *CampaignSparse) GetLinks() CampaignLinks {
	if o == nil {
		var ret CampaignLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetLinksOk() (*CampaignLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CampaignSparse) SetLinks(v CampaignLinks) {
	o.Links = v
}

// GetId returns the Id field value
func (o *CampaignSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignSparse) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CampaignSparse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *CampaignSparse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *CampaignSparse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignSparse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *CampaignSparse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *CampaignSparse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *CampaignSparse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *CampaignSparse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetStatus returns the Status field value
func (o *CampaignSparse) GetStatus() CampaignStatus {
	if o == nil {
		var ret CampaignStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CampaignSparse) GetStatusOk() (*CampaignStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CampaignSparse) SetStatus(v CampaignStatus) {
	o.Status = v
}

func (o CampaignSparse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["endDate"] = o.EndDate
	}
	if true {
		toSerialize["scheduleType"] = o.ScheduleType
	}
	if o.RecurringCampaignId.IsSet() {
		toSerialize["recurringCampaignId"] = o.RecurringCampaignId.Get()
	}
	if true {
		toSerialize["reviewerType"] = o.ReviewerType
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["id"] = o.Id
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
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignSparse) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignSparse := _CampaignSparse{}

	err = json.Unmarshal(bytes, &varCampaignSparse)
	if err == nil {
		*o = CampaignSparse(varCampaignSparse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "scheduleType")
		delete(additionalProperties, "recurringCampaignId")
		delete(additionalProperties, "reviewerType")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCampaignSparse struct {
	value *CampaignSparse
	isSet bool
}

func (v NullableCampaignSparse) Get() *CampaignSparse {
	return v.value
}

func (v *NullableCampaignSparse) Set(val *CampaignSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSparse(val *CampaignSparse) *NullableCampaignSparse {
	return &NullableCampaignSparse{value: val, isSet: true}
}

func (v NullableCampaignSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
