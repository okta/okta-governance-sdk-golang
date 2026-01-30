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

// checks if the CampaignMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMutable{}

// CampaignMutable There are two `campaignType` types.   RESOURCE: Campaigns created focused on certain resources USER: Campaigns created focused on certain users
type CampaignMutable struct {
	// Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed.
	Name string `json:"name"`
	// Human readable description.
	Description            *string                        `json:"description,omitempty"`
	CampaignType           *CampaignType                  `json:"campaignType,omitempty"`
	ScheduleSettings       ScheduleSettingsMutable        `json:"scheduleSettings"`
	ResourceSettings       ResourceSettingsMutable        `json:"resourceSettings"`
	PrincipalScopeSettings *PrincipalScopeSettingsMutable `json:"principalScopeSettings,omitempty"`
	ReviewerSettings       ReviewerSettingsMutable        `json:"reviewerSettings"`
	NotificationSettings   *NotificationSettings          `json:"notificationSettings,omitempty"`
	RemediationSettings    RemediationSettings            `json:"remediationSettings"`
	CampaignTier           *CampaignTier                  `json:"campaignTier,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _CampaignMutable CampaignMutable

// NewCampaignMutable instantiates a new CampaignMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMutable(name string, scheduleSettings ScheduleSettingsMutable, resourceSettings ResourceSettingsMutable, reviewerSettings ReviewerSettingsMutable, remediationSettings RemediationSettings) *CampaignMutable {
	this := CampaignMutable{}
	this.Name = name
	this.ScheduleSettings = scheduleSettings
	this.ResourceSettings = resourceSettings
	this.ReviewerSettings = reviewerSettings
	this.RemediationSettings = remediationSettings
	return &this
}

// NewCampaignMutableWithDefaults instantiates a new CampaignMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMutableWithDefaults() *CampaignMutable {
	this := CampaignMutable{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignMutable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignMutable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignMutable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignMutable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignMutable) SetDescription(v string) {
	o.Description = &v
}

// GetCampaignType returns the CampaignType field value if set, zero value otherwise.
func (o *CampaignMutable) GetCampaignType() CampaignType {
	if o == nil || IsNil(o.CampaignType) {
		var ret CampaignType
		return ret
	}
	return *o.CampaignType
}

// GetCampaignTypeOk returns a tuple with the CampaignType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetCampaignTypeOk() (*CampaignType, bool) {
	if o == nil || IsNil(o.CampaignType) {
		return nil, false
	}
	return o.CampaignType, true
}

// HasCampaignType returns a boolean if a field has been set.
func (o *CampaignMutable) HasCampaignType() bool {
	if o != nil && !IsNil(o.CampaignType) {
		return true
	}

	return false
}

// SetCampaignType gets a reference to the given CampaignType and assigns it to the CampaignType field.
func (o *CampaignMutable) SetCampaignType(v CampaignType) {
	o.CampaignType = &v
}

// GetScheduleSettings returns the ScheduleSettings field value
func (o *CampaignMutable) GetScheduleSettings() ScheduleSettingsMutable {
	if o == nil {
		var ret ScheduleSettingsMutable
		return ret
	}

	return o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetScheduleSettingsOk() (*ScheduleSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleSettings, true
}

// SetScheduleSettings sets field value
func (o *CampaignMutable) SetScheduleSettings(v ScheduleSettingsMutable) {
	o.ScheduleSettings = v
}

// GetResourceSettings returns the ResourceSettings field value
func (o *CampaignMutable) GetResourceSettings() ResourceSettingsMutable {
	if o == nil {
		var ret ResourceSettingsMutable
		return ret
	}

	return o.ResourceSettings
}

// GetResourceSettingsOk returns a tuple with the ResourceSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetResourceSettingsOk() (*ResourceSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSettings, true
}

// SetResourceSettings sets field value
func (o *CampaignMutable) SetResourceSettings(v ResourceSettingsMutable) {
	o.ResourceSettings = v
}

// GetPrincipalScopeSettings returns the PrincipalScopeSettings field value if set, zero value otherwise.
func (o *CampaignMutable) GetPrincipalScopeSettings() PrincipalScopeSettingsMutable {
	if o == nil || IsNil(o.PrincipalScopeSettings) {
		var ret PrincipalScopeSettingsMutable
		return ret
	}
	return *o.PrincipalScopeSettings
}

// GetPrincipalScopeSettingsOk returns a tuple with the PrincipalScopeSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetPrincipalScopeSettingsOk() (*PrincipalScopeSettingsMutable, bool) {
	if o == nil || IsNil(o.PrincipalScopeSettings) {
		return nil, false
	}
	return o.PrincipalScopeSettings, true
}

// HasPrincipalScopeSettings returns a boolean if a field has been set.
func (o *CampaignMutable) HasPrincipalScopeSettings() bool {
	if o != nil && !IsNil(o.PrincipalScopeSettings) {
		return true
	}

	return false
}

// SetPrincipalScopeSettings gets a reference to the given PrincipalScopeSettingsMutable and assigns it to the PrincipalScopeSettings field.
func (o *CampaignMutable) SetPrincipalScopeSettings(v PrincipalScopeSettingsMutable) {
	o.PrincipalScopeSettings = &v
}

// GetReviewerSettings returns the ReviewerSettings field value
func (o *CampaignMutable) GetReviewerSettings() ReviewerSettingsMutable {
	if o == nil {
		var ret ReviewerSettingsMutable
		return ret
	}

	return o.ReviewerSettings
}

// GetReviewerSettingsOk returns a tuple with the ReviewerSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetReviewerSettingsOk() (*ReviewerSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerSettings, true
}

// SetReviewerSettings sets field value
func (o *CampaignMutable) SetReviewerSettings(v ReviewerSettingsMutable) {
	o.ReviewerSettings = v
}

// GetNotificationSettings returns the NotificationSettings field value if set, zero value otherwise.
func (o *CampaignMutable) GetNotificationSettings() NotificationSettings {
	if o == nil || IsNil(o.NotificationSettings) {
		var ret NotificationSettings
		return ret
	}
	return *o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetNotificationSettingsOk() (*NotificationSettings, bool) {
	if o == nil || IsNil(o.NotificationSettings) {
		return nil, false
	}
	return o.NotificationSettings, true
}

// HasNotificationSettings returns a boolean if a field has been set.
func (o *CampaignMutable) HasNotificationSettings() bool {
	if o != nil && !IsNil(o.NotificationSettings) {
		return true
	}

	return false
}

// SetNotificationSettings gets a reference to the given NotificationSettings and assigns it to the NotificationSettings field.
func (o *CampaignMutable) SetNotificationSettings(v NotificationSettings) {
	o.NotificationSettings = &v
}

// GetRemediationSettings returns the RemediationSettings field value
func (o *CampaignMutable) GetRemediationSettings() RemediationSettings {
	if o == nil {
		var ret RemediationSettings
		return ret
	}

	return o.RemediationSettings
}

// GetRemediationSettingsOk returns a tuple with the RemediationSettings field value
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetRemediationSettingsOk() (*RemediationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationSettings, true
}

// SetRemediationSettings sets field value
func (o *CampaignMutable) SetRemediationSettings(v RemediationSettings) {
	o.RemediationSettings = v
}

// GetCampaignTier returns the CampaignTier field value if set, zero value otherwise.
func (o *CampaignMutable) GetCampaignTier() CampaignTier {
	if o == nil || IsNil(o.CampaignTier) {
		var ret CampaignTier
		return ret
	}
	return *o.CampaignTier
}

// GetCampaignTierOk returns a tuple with the CampaignTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMutable) GetCampaignTierOk() (*CampaignTier, bool) {
	if o == nil || IsNil(o.CampaignTier) {
		return nil, false
	}
	return o.CampaignTier, true
}

// HasCampaignTier returns a boolean if a field has been set.
func (o *CampaignMutable) HasCampaignTier() bool {
	if o != nil && !IsNil(o.CampaignTier) {
		return true
	}

	return false
}

// SetCampaignTier gets a reference to the given CampaignTier and assigns it to the CampaignTier field.
func (o *CampaignMutable) SetCampaignTier(v CampaignTier) {
	o.CampaignTier = &v
}

func (o CampaignMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignMutable) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CampaignTier) {
		toSerialize["campaignTier"] = o.CampaignTier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignMutable) UnmarshalJSON(data []byte) (err error) {
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

	varCampaignMutable := _CampaignMutable{}

	err = json.Unmarshal(data, &varCampaignMutable)

	if err != nil {
		return err
	}

	*o = CampaignMutable(varCampaignMutable)

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
		delete(additionalProperties, "campaignTier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignMutable struct {
	value *CampaignMutable
	isSet bool
}

func (v NullableCampaignMutable) Get() *CampaignMutable {
	return v.value
}

func (v *NullableCampaignMutable) Set(val *CampaignMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMutable(val *CampaignMutable) *NullableCampaignMutable {
	return &NullableCampaignMutable{value: val, isSet: true}
}

func (v NullableCampaignMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
