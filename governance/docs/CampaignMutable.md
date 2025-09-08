# CampaignMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed. | 
**Description** | Pointer to **string** | Human readable description. | [optional] 
**CampaignType** | Pointer to [**CampaignType**](CampaignType.md) |  | [optional] 
**ScheduleSettings** | [**ScheduleSettingsMutable**](ScheduleSettingsMutable.md) |  | 
**ResourceSettings** | [**ResourceSettingsMutable**](ResourceSettingsMutable.md) |  | 
**PrincipalScopeSettings** | Pointer to [**PrincipalScopeSettingsMutable**](PrincipalScopeSettingsMutable.md) |  | [optional] 
**ReviewerSettings** | [**ReviewerSettingsMutable**](ReviewerSettingsMutable.md) |  | 
**NotificationSettings** | Pointer to [**NotificationSettings**](NotificationSettings.md) |  | [optional] 
**RemediationSettings** | [**RemediationSettings**](RemediationSettings.md) |  | 
**CampaignTier** | Pointer to [**CampaignTier**](CampaignTier.md) |  | [optional] 

## Methods

### NewCampaignMutable

`func NewCampaignMutable(name string, scheduleSettings ScheduleSettingsMutable, resourceSettings ResourceSettingsMutable, reviewerSettings ReviewerSettingsMutable, remediationSettings RemediationSettings, ) *CampaignMutable`

NewCampaignMutable instantiates a new CampaignMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMutableWithDefaults

`func NewCampaignMutableWithDefaults() *CampaignMutable`

NewCampaignMutableWithDefaults instantiates a new CampaignMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignMutable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignMutable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignMutable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignMutable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignMutable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignMutable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignMutable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCampaignType

`func (o *CampaignMutable) GetCampaignType() CampaignType`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CampaignMutable) GetCampaignTypeOk() (*CampaignType, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *CampaignMutable) SetCampaignType(v CampaignType)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *CampaignMutable) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetScheduleSettings

`func (o *CampaignMutable) GetScheduleSettings() ScheduleSettingsMutable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *CampaignMutable) GetScheduleSettingsOk() (*ScheduleSettingsMutable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *CampaignMutable) SetScheduleSettings(v ScheduleSettingsMutable)`

SetScheduleSettings sets ScheduleSettings field to given value.


### GetResourceSettings

`func (o *CampaignMutable) GetResourceSettings() ResourceSettingsMutable`

GetResourceSettings returns the ResourceSettings field if non-nil, zero value otherwise.

### GetResourceSettingsOk

`func (o *CampaignMutable) GetResourceSettingsOk() (*ResourceSettingsMutable, bool)`

GetResourceSettingsOk returns a tuple with the ResourceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettings

`func (o *CampaignMutable) SetResourceSettings(v ResourceSettingsMutable)`

SetResourceSettings sets ResourceSettings field to given value.


### GetPrincipalScopeSettings

`func (o *CampaignMutable) GetPrincipalScopeSettings() PrincipalScopeSettingsMutable`

GetPrincipalScopeSettings returns the PrincipalScopeSettings field if non-nil, zero value otherwise.

### GetPrincipalScopeSettingsOk

`func (o *CampaignMutable) GetPrincipalScopeSettingsOk() (*PrincipalScopeSettingsMutable, bool)`

GetPrincipalScopeSettingsOk returns a tuple with the PrincipalScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalScopeSettings

`func (o *CampaignMutable) SetPrincipalScopeSettings(v PrincipalScopeSettingsMutable)`

SetPrincipalScopeSettings sets PrincipalScopeSettings field to given value.

### HasPrincipalScopeSettings

`func (o *CampaignMutable) HasPrincipalScopeSettings() bool`

HasPrincipalScopeSettings returns a boolean if a field has been set.

### GetReviewerSettings

`func (o *CampaignMutable) GetReviewerSettings() ReviewerSettingsMutable`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *CampaignMutable) GetReviewerSettingsOk() (*ReviewerSettingsMutable, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *CampaignMutable) SetReviewerSettings(v ReviewerSettingsMutable)`

SetReviewerSettings sets ReviewerSettings field to given value.


### GetNotificationSettings

`func (o *CampaignMutable) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *CampaignMutable) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *CampaignMutable) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *CampaignMutable) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetRemediationSettings

`func (o *CampaignMutable) GetRemediationSettings() RemediationSettings`

GetRemediationSettings returns the RemediationSettings field if non-nil, zero value otherwise.

### GetRemediationSettingsOk

`func (o *CampaignMutable) GetRemediationSettingsOk() (*RemediationSettings, bool)`

GetRemediationSettingsOk returns a tuple with the RemediationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationSettings

`func (o *CampaignMutable) SetRemediationSettings(v RemediationSettings)`

SetRemediationSettings sets RemediationSettings field to given value.


### GetCampaignTier

`func (o *CampaignMutable) GetCampaignTier() CampaignTier`

GetCampaignTier returns the CampaignTier field if non-nil, zero value otherwise.

### GetCampaignTierOk

`func (o *CampaignMutable) GetCampaignTierOk() (*CampaignTier, bool)`

GetCampaignTierOk returns a tuple with the CampaignTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignTier

`func (o *CampaignMutable) SetCampaignTier(v CampaignTier)`

SetCampaignTier sets CampaignTier field to given value.

### HasCampaignTier

`func (o *CampaignMutable) HasCampaignTier() bool`

HasCampaignTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


