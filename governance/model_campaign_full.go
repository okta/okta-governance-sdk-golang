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

// CampaignFull Full representation of a Campaign resource
type CampaignFull struct {
	Links CampaignLinks `json:"_links"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed.
	Name string `json:"name"`
	// Human readable description.
	Description            *string                        `json:"description,omitempty"`
	CampaignType           *CampaignType                  `json:"campaignType,omitempty"`
	ScheduleSettings       ScheduleSettingsReadOnly       `json:"scheduleSettings"`
	ResourceSettings       ResourceSettingsMutable        `json:"resourceSettings"`
	PrincipalScopeSettings *PrincipalScopeSettingsMutable `json:"principalScopeSettings,omitempty"`
	ReviewerSettings       ReviewerSettingsMutable        `json:"reviewerSettings"`
	NotificationSettings   *NotificationSettings          `json:"notificationSettings,omitempty"`
	RemediationSettings    RemediationSettings            `json:"remediationSettings"`
	// ID of the recurring campaign if this campaign was created as part of a recurring schedule.
	RecurringCampaignId  NullableString `json:"recurringCampaignId,omitempty"`
	Status               CampaignStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _CampaignFull CampaignFull

// NewCampaignFull instantiates a new CampaignFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignFull(links CampaignLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, name string, scheduleSettings ScheduleSettingsReadOnly, resourceSettings ResourceSettingsMutable, reviewerSettings ReviewerSettingsMutable, remediationSettings RemediationSettings, status CampaignStatus) *CampaignFull {
	this := CampaignFull{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.Name = name
	this.ScheduleSettings = scheduleSettings
	this.ResourceSettings = resourceSettings
	this.ReviewerSettings = reviewerSettings
	this.RemediationSettings = remediationSettings
	this.Status = status
	return &this
}

// NewCampaignFullWithDefaults instantiates a new CampaignFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignFullWithDefaults() *CampaignFull {
	this := CampaignFull{}
	return &this
}

// GetLinks returns the Links field value
func (o *CampaignFull) GetLinks() CampaignLinks {
	if o == nil {
		var ret CampaignLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetLinksOk() (*CampaignLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CampaignFull) SetLinks(v CampaignLinks) {
	o.Links = v
}

// GetId returns the Id field value
func (o *CampaignFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CampaignFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *CampaignFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *CampaignFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *CampaignFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *CampaignFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *CampaignFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *CampaignFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetName returns the Name field value
func (o *CampaignFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignFull) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignFull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignFull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignFull) SetDescription(v string) {
	o.Description = &v
}

// GetCampaignType returns the CampaignType field value if set, zero value otherwise.
func (o *CampaignFull) GetCampaignType() CampaignType {
	if o == nil || o.CampaignType == nil {
		var ret CampaignType
		return ret
	}
	return *o.CampaignType
}

// GetCampaignTypeOk returns a tuple with the CampaignType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetCampaignTypeOk() (*CampaignType, bool) {
	if o == nil || o.CampaignType == nil {
		return nil, false
	}
	return o.CampaignType, true
}

// HasCampaignType returns a boolean if a field has been set.
func (o *CampaignFull) HasCampaignType() bool {
	if o != nil && o.CampaignType != nil {
		return true
	}

	return false
}

// SetCampaignType gets a reference to the given CampaignType and assigns it to the CampaignType field.
func (o *CampaignFull) SetCampaignType(v CampaignType) {
	o.CampaignType = &v
}

// GetScheduleSettings returns the ScheduleSettings field value
func (o *CampaignFull) GetScheduleSettings() ScheduleSettingsReadOnly {
	if o == nil {
		var ret ScheduleSettingsReadOnly
		return ret
	}

	return o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetScheduleSettingsOk() (*ScheduleSettingsReadOnly, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleSettings, true
}

// SetScheduleSettings sets field value
func (o *CampaignFull) SetScheduleSettings(v ScheduleSettingsReadOnly) {
	o.ScheduleSettings = v
}

// GetResourceSettings returns the ResourceSettings field value
func (o *CampaignFull) GetResourceSettings() ResourceSettingsMutable {
	if o == nil {
		var ret ResourceSettingsMutable
		return ret
	}

	return o.ResourceSettings
}

// GetResourceSettingsOk returns a tuple with the ResourceSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetResourceSettingsOk() (*ResourceSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSettings, true
}

// SetResourceSettings sets field value
func (o *CampaignFull) SetResourceSettings(v ResourceSettingsMutable) {
	o.ResourceSettings = v
}

// GetPrincipalScopeSettings returns the PrincipalScopeSettings field value if set, zero value otherwise.
func (o *CampaignFull) GetPrincipalScopeSettings() PrincipalScopeSettingsMutable {
	if o == nil || o.PrincipalScopeSettings == nil {
		var ret PrincipalScopeSettingsMutable
		return ret
	}
	return *o.PrincipalScopeSettings
}

// GetPrincipalScopeSettingsOk returns a tuple with the PrincipalScopeSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetPrincipalScopeSettingsOk() (*PrincipalScopeSettingsMutable, bool) {
	if o == nil || o.PrincipalScopeSettings == nil {
		return nil, false
	}
	return o.PrincipalScopeSettings, true
}

// HasPrincipalScopeSettings returns a boolean if a field has been set.
func (o *CampaignFull) HasPrincipalScopeSettings() bool {
	if o != nil && o.PrincipalScopeSettings != nil {
		return true
	}

	return false
}

// SetPrincipalScopeSettings gets a reference to the given PrincipalScopeSettingsMutable and assigns it to the PrincipalScopeSettings field.
func (o *CampaignFull) SetPrincipalScopeSettings(v PrincipalScopeSettingsMutable) {
	o.PrincipalScopeSettings = &v
}

// GetReviewerSettings returns the ReviewerSettings field value
func (o *CampaignFull) GetReviewerSettings() ReviewerSettingsMutable {
	if o == nil {
		var ret ReviewerSettingsMutable
		return ret
	}

	return o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetReviewerSettingsOk() (*ReviewerSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerSettings, true
}

// SetReviewerSettings sets field value
func (o *CampaignFull) SetReviewerSettings(v ReviewerSettingsMutable) {
	o.ReviewerSettings = v
}

// GetNotificationSettings returns the NotificationSettings field value if set, zero value otherwise.
func (o *CampaignFull) GetNotificationSettings() NotificationSettings {
	if o == nil || o.NotificationSettings == nil {
		var ret NotificationSettings
		return ret
	}
	return *o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetNotificationSettingsOk() (*NotificationSettings, bool) {
	if o == nil || o.NotificationSettings == nil {
		return nil, false
	}
	return o.NotificationSettings, true
}

// HasNotificationSettings returns a boolean if a field has been set.
func (o *CampaignFull) HasNotificationSettings() bool {
	if o != nil && o.NotificationSettings != nil {
		return true
	}

	return false
}

// SetNotificationSettings gets a reference to the given NotificationSettings and assigns it to the NotificationSettings field.
func (o *CampaignFull) SetNotificationSettings(v NotificationSettings) {
	o.NotificationSettings = &v
}

// GetRemediationSettings returns the RemediationSettings field value
func (o *CampaignFull) GetRemediationSettings() RemediationSettings {
	if o == nil {
		var ret RemediationSettings
		return ret
	}

	return o.RemediationSettings
}

// GetRemediationSettingsOk returns a tuple with the RemediationSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetRemediationSettingsOk() (*RemediationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationSettings, true
}

// SetRemediationSettings sets field value
func (o *CampaignFull) SetRemediationSettings(v RemediationSettings) {
	o.RemediationSettings = v
}

// GetRecurringCampaignId returns the RecurringCampaignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignFull) GetRecurringCampaignId() string {
	if o == nil || o.RecurringCampaignId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecurringCampaignId.Get()
}

// GetRecurringCampaignIdOk returns a tuple with the RecurringCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignFull) GetRecurringCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringCampaignId.Get(), o.RecurringCampaignId.IsSet()
}

// HasRecurringCampaignId returns a boolean if a field has been set.
func (o *CampaignFull) HasRecurringCampaignId() bool {
	if o != nil && o.RecurringCampaignId.IsSet() {
		return true
	}

	return false
}

// SetRecurringCampaignId gets a reference to the given NullableString and assigns it to the RecurringCampaignId field.
func (o *CampaignFull) SetRecurringCampaignId(v string) {
	o.RecurringCampaignId.Set(&v)
}

// SetRecurringCampaignIdNil sets the value for RecurringCampaignId to be an explicit nil
func (o *CampaignFull) SetRecurringCampaignIdNil() {
	o.RecurringCampaignId.Set(nil)
}

// UnsetRecurringCampaignId ensures that no value is present for RecurringCampaignId, not even an explicit nil
func (o *CampaignFull) UnsetRecurringCampaignId() {
	o.RecurringCampaignId.Unset()
}

// GetStatus returns the Status field value
func (o *CampaignFull) GetStatus() CampaignStatus {
	if o == nil {
		var ret CampaignStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CampaignFull) GetStatusOk() (*CampaignStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CampaignFull) SetStatus(v CampaignStatus) {
	o.Status = v
}

func (o CampaignFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CampaignType != nil {
		toSerialize["campaignType"] = o.CampaignType
	}
	if true {
		toSerialize["scheduleSettings"] = o.ScheduleSettings
	}
	if true {
		toSerialize["resourceSettings"] = o.ResourceSettings
	}
	if o.PrincipalScopeSettings != nil {
		toSerialize["principalScopeSettings"] = o.PrincipalScopeSettings
	}
	if true {
		toSerialize["reviewerSettings"] = o.ReviewerSettings
	}
	if o.NotificationSettings != nil {
		toSerialize["notificationSettings"] = o.NotificationSettings
	}
	if true {
		toSerialize["remediationSettings"] = o.RemediationSettings
	}
	if o.RecurringCampaignId.IsSet() {
		toSerialize["recurringCampaignId"] = o.RecurringCampaignId.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CampaignFull) UnmarshalJSON(bytes []byte) (err error) {
	varCampaignFull := _CampaignFull{}

	err = json.Unmarshal(bytes, &varCampaignFull)
	if err == nil {
		*o = CampaignFull(varCampaignFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "campaignType")
		delete(additionalProperties, "scheduleSettings")
		delete(additionalProperties, "resourceSettings")
		delete(additionalProperties, "principalScopeSettings")
		delete(additionalProperties, "reviewerSettings")
		delete(additionalProperties, "notificationSettings")
		delete(additionalProperties, "remediationSettings")
		delete(additionalProperties, "recurringCampaignId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCampaignFull struct {
	value *CampaignFull
	isSet bool
}

func (v NullableCampaignFull) Get() *CampaignFull {
	return v.value
}

func (v *NullableCampaignFull) Set(val *CampaignFull) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignFull) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignFull(val *CampaignFull) *NullableCampaignFull {
	return &NullableCampaignFull{value: val, isSet: true}
}

func (v NullableCampaignFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
