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

// checks if the CampaignDetailsReadOnly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignDetailsReadOnly{}

// CampaignDetailsReadOnly struct for CampaignDetailsReadOnly
type CampaignDetailsReadOnly struct {
	// Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed.
	Name string `json:"name"`
	// Human readable description
	Description            *string                        `json:"description,omitempty"`
	CampaignType           *CampaignType                  `json:"campaignType,omitempty"`
	ScheduleSettings       ScheduleSettingsReadOnly       `json:"scheduleSettings"`
	ResourceSettings       ResourceSettingsMutable        `json:"resourceSettings"`
	PrincipalScopeSettings *PrincipalScopeSettingsMutable `json:"principalScopeSettings,omitempty"`
	ReviewerSettings       ReviewerSettingsMutable        `json:"reviewerSettings"`
	NotificationSettings   *NotificationSettings          `json:"notificationSettings,omitempty"`
	RemediationSettings    RemediationSettings            `json:"remediationSettings"`
	// ID of the recurring campaign if this campaign was created as part of a recurring schedule.
	RecurringCampaignId  NullableString            `json:"recurringCampaignId,omitempty"`
	ReportingSettings    *ReportingSettingsMutable `json:"reportingSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignDetailsReadOnly CampaignDetailsReadOnly

// NewCampaignDetailsReadOnly instantiates a new CampaignDetailsReadOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDetailsReadOnly(name string, scheduleSettings ScheduleSettingsReadOnly, resourceSettings ResourceSettingsMutable, reviewerSettings ReviewerSettingsMutable, remediationSettings RemediationSettings) *CampaignDetailsReadOnly {
	this := CampaignDetailsReadOnly{}
	this.Name = name
	var campaignType CampaignType = CAMPAIGNTYPE_RESOURCE
	this.CampaignType = &campaignType
	this.ScheduleSettings = scheduleSettings
	this.ResourceSettings = resourceSettings
	this.ReviewerSettings = reviewerSettings
	this.RemediationSettings = remediationSettings
	return &this
}

// NewCampaignDetailsReadOnlyWithDefaults instantiates a new CampaignDetailsReadOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDetailsReadOnlyWithDefaults() *CampaignDetailsReadOnly {
	this := CampaignDetailsReadOnly{}
	var campaignType CampaignType = CAMPAIGNTYPE_RESOURCE
	this.CampaignType = &campaignType
	return &this
}

// GetName returns the Name field value
func (o *CampaignDetailsReadOnly) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignDetailsReadOnly) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignDetailsReadOnly) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignDetailsReadOnly) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignDetailsReadOnly) SetDescription(v string) {
	o.Description = &v
}

// GetCampaignType returns the CampaignType field value if set, zero value otherwise.
func (o *CampaignDetailsReadOnly) GetCampaignType() CampaignType {
	if o == nil || IsNil(o.CampaignType) {
		var ret CampaignType
		return ret
	}
	return *o.CampaignType
}

// GetCampaignTypeOk returns a tuple with the CampaignType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetCampaignTypeOk() (*CampaignType, bool) {
	if o == nil || IsNil(o.CampaignType) {
		return nil, false
	}
	return o.CampaignType, true
}

// HasCampaignType returns a boolean if a field has been set.
func (o *CampaignDetailsReadOnly) HasCampaignType() bool {
	if o != nil && !IsNil(o.CampaignType) {
		return true
	}

	return false
}

// SetCampaignType gets a reference to the given CampaignType and assigns it to the CampaignType field.
func (o *CampaignDetailsReadOnly) SetCampaignType(v CampaignType) {
	o.CampaignType = &v
}

// GetScheduleSettings returns the ScheduleSettings field value
func (o *CampaignDetailsReadOnly) GetScheduleSettings() ScheduleSettingsReadOnly {
	if o == nil {
		var ret ScheduleSettingsReadOnly
		return ret
	}

	return o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetScheduleSettingsOk() (*ScheduleSettingsReadOnly, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleSettings, true
}

// SetScheduleSettings sets field value
func (o *CampaignDetailsReadOnly) SetScheduleSettings(v ScheduleSettingsReadOnly) {
	o.ScheduleSettings = v
}

// GetResourceSettings returns the ResourceSettings field value
func (o *CampaignDetailsReadOnly) GetResourceSettings() ResourceSettingsMutable {
	if o == nil {
		var ret ResourceSettingsMutable
		return ret
	}

	return o.ResourceSettings
}

// GetResourceSettingsOk returns a tuple with the ResourceSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetResourceSettingsOk() (*ResourceSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSettings, true
}

// SetResourceSettings sets field value
func (o *CampaignDetailsReadOnly) SetResourceSettings(v ResourceSettingsMutable) {
	o.ResourceSettings = v
}

// GetPrincipalScopeSettings returns the PrincipalScopeSettings field value if set, zero value otherwise.
func (o *CampaignDetailsReadOnly) GetPrincipalScopeSettings() PrincipalScopeSettingsMutable {
	if o == nil || IsNil(o.PrincipalScopeSettings) {
		var ret PrincipalScopeSettingsMutable
		return ret
	}
	return *o.PrincipalScopeSettings
}

// GetPrincipalScopeSettingsOk returns a tuple with the PrincipalScopeSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetPrincipalScopeSettingsOk() (*PrincipalScopeSettingsMutable, bool) {
	if o == nil || IsNil(o.PrincipalScopeSettings) {
		return nil, false
	}
	return o.PrincipalScopeSettings, true
}

// HasPrincipalScopeSettings returns a boolean if a field has been set.
func (o *CampaignDetailsReadOnly) HasPrincipalScopeSettings() bool {
	if o != nil && !IsNil(o.PrincipalScopeSettings) {
		return true
	}

	return false
}

// SetPrincipalScopeSettings gets a reference to the given PrincipalScopeSettingsMutable and assigns it to the PrincipalScopeSettings field.
func (o *CampaignDetailsReadOnly) SetPrincipalScopeSettings(v PrincipalScopeSettingsMutable) {
	o.PrincipalScopeSettings = &v
}

// GetReviewerSettings returns the ReviewerSettings field value
func (o *CampaignDetailsReadOnly) GetReviewerSettings() ReviewerSettingsMutable {
	if o == nil {
		var ret ReviewerSettingsMutable
		return ret
	}

	return o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetReviewerSettingsOk() (*ReviewerSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerSettings, true
}

// SetReviewerSettings sets field value
func (o *CampaignDetailsReadOnly) SetReviewerSettings(v ReviewerSettingsMutable) {
	o.ReviewerSettings = v
}

// GetNotificationSettings returns the NotificationSettings field value if set, zero value otherwise.
func (o *CampaignDetailsReadOnly) GetNotificationSettings() NotificationSettings {
	if o == nil || IsNil(o.NotificationSettings) {
		var ret NotificationSettings
		return ret
	}
	return *o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetNotificationSettingsOk() (*NotificationSettings, bool) {
	if o == nil || IsNil(o.NotificationSettings) {
		return nil, false
	}
	return o.NotificationSettings, true
}

// HasNotificationSettings returns a boolean if a field has been set.
func (o *CampaignDetailsReadOnly) HasNotificationSettings() bool {
	if o != nil && !IsNil(o.NotificationSettings) {
		return true
	}

	return false
}

// SetNotificationSettings gets a reference to the given NotificationSettings and assigns it to the NotificationSettings field.
func (o *CampaignDetailsReadOnly) SetNotificationSettings(v NotificationSettings) {
	o.NotificationSettings = &v
}

// GetRemediationSettings returns the RemediationSettings field value
func (o *CampaignDetailsReadOnly) GetRemediationSettings() RemediationSettings {
	if o == nil {
		var ret RemediationSettings
		return ret
	}

	return o.RemediationSettings
}

// GetRemediationSettingsOk returns a tuple with the RemediationSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetRemediationSettingsOk() (*RemediationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationSettings, true
}

// SetRemediationSettings sets field value
func (o *CampaignDetailsReadOnly) SetRemediationSettings(v RemediationSettings) {
	o.RemediationSettings = v
}

// GetRecurringCampaignId returns the RecurringCampaignId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDetailsReadOnly) GetRecurringCampaignId() string {
	if o == nil || IsNil(o.RecurringCampaignId.Get()) {
		var ret string
		return ret
	}
	return *o.RecurringCampaignId.Get()
}

// GetRecurringCampaignIdOk returns a tuple with the RecurringCampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDetailsReadOnly) GetRecurringCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringCampaignId.Get(), o.RecurringCampaignId.IsSet()
}

// HasRecurringCampaignId returns a boolean if a field has been set.
func (o *CampaignDetailsReadOnly) HasRecurringCampaignId() bool {
	if o != nil && o.RecurringCampaignId.IsSet() {
		return true
	}

	return false
}

// SetRecurringCampaignId gets a reference to the given NullableString and assigns it to the RecurringCampaignId field.
func (o *CampaignDetailsReadOnly) SetRecurringCampaignId(v string) {
	o.RecurringCampaignId.Set(&v)
}

// SetRecurringCampaignIdNil sets the value for RecurringCampaignId to be an explicit nil
func (o *CampaignDetailsReadOnly) SetRecurringCampaignIdNil() {
	o.RecurringCampaignId.Set(nil)
}

// UnsetRecurringCampaignId ensures that no value is present for RecurringCampaignId, not even an explicit nil
func (o *CampaignDetailsReadOnly) UnsetRecurringCampaignId() {
	o.RecurringCampaignId.Unset()
}

// GetReportingSettings returns the ReportingSettings field value if set, zero value otherwise.
func (o *CampaignDetailsReadOnly) GetReportingSettings() ReportingSettingsMutable {
	if o == nil || IsNil(o.ReportingSettings) {
		var ret ReportingSettingsMutable
		return ret
	}
	return *o.ReportingSettings
}

// GetReportingSettingsOk returns a tuple with the ReportingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetailsReadOnly) GetReportingSettingsOk() (*ReportingSettingsMutable, bool) {
	if o == nil || IsNil(o.ReportingSettings) {
		return nil, false
	}
	return o.ReportingSettings, true
}

// HasReportingSettings returns a boolean if a field has been set.
func (o *CampaignDetailsReadOnly) HasReportingSettings() bool {
	if o != nil && !IsNil(o.ReportingSettings) {
		return true
	}

	return false
}

// SetReportingSettings gets a reference to the given ReportingSettingsMutable and assigns it to the ReportingSettings field.
func (o *CampaignDetailsReadOnly) SetReportingSettings(v ReportingSettingsMutable) {
	o.ReportingSettings = &v
}

func (o CampaignDetailsReadOnly) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignDetailsReadOnly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CampaignType) {
		toSerialize["campaignType"] = o.CampaignType
	}
	toSerialize["scheduleSettings"] = o.ScheduleSettings
	toSerialize["resourceSettings"] = o.ResourceSettings
	if !IsNil(o.PrincipalScopeSettings) {
		toSerialize["principalScopeSettings"] = o.PrincipalScopeSettings
	}
	toSerialize["reviewerSettings"] = o.ReviewerSettings
	if !IsNil(o.NotificationSettings) {
		toSerialize["notificationSettings"] = o.NotificationSettings
	}
	toSerialize["remediationSettings"] = o.RemediationSettings
	if o.RecurringCampaignId.IsSet() {
		toSerialize["recurringCampaignId"] = o.RecurringCampaignId.Get()
	}
	if !IsNil(o.ReportingSettings) {
		toSerialize["reportingSettings"] = o.ReportingSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignDetailsReadOnly) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scheduleSettings",
		"resourceSettings",
		"reviewerSettings",
		"remediationSettings",
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

	varCampaignDetailsReadOnly := _CampaignDetailsReadOnly{}

	err = json.Unmarshal(data, &varCampaignDetailsReadOnly)

	if err != nil {
		return err
	}

	*o = CampaignDetailsReadOnly(varCampaignDetailsReadOnly)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
		delete(additionalProperties, "reportingSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignDetailsReadOnly struct {
	value *CampaignDetailsReadOnly
	isSet bool
}

func (v NullableCampaignDetailsReadOnly) Get() *CampaignDetailsReadOnly {
	return v.value
}

func (v *NullableCampaignDetailsReadOnly) Set(val *CampaignDetailsReadOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDetailsReadOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDetailsReadOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDetailsReadOnly(val *CampaignDetailsReadOnly) *NullableCampaignDetailsReadOnly {
	return &NullableCampaignDetailsReadOnly{value: val, isSet: true}
}

func (v NullableCampaignDetailsReadOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDetailsReadOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
