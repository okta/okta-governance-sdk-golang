# CampaignDetailsReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed. | 
**Description** | Pointer to **string** | Human readable description | [optional] 
**CampaignType** | Pointer to [**CampaignType**](CampaignType.md) |  | [optional] [default to CAMPAIGNTYPE_RESOURCE]
**ScheduleSettings** | [**ScheduleSettingsReadOnly**](ScheduleSettingsReadOnly.md) |  | 
**ResourceSettings** | [**ResourceSettingsMutable**](ResourceSettingsMutable.md) |  | 
**PrincipalScopeSettings** | Pointer to [**PrincipalScopeSettingsMutable**](PrincipalScopeSettingsMutable.md) |  | [optional] 
**ReviewerSettings** | [**ReviewerSettingsMutable**](ReviewerSettingsMutable.md) |  | 
**NotificationSettings** | Pointer to [**NotificationSettings**](NotificationSettings.md) |  | [optional] 
**RemediationSettings** | [**RemediationSettings**](RemediationSettings.md) |  | 
**RecurringCampaignId** | Pointer to **NullableString** | ID of the recurring campaign if this campaign was created as part of a recurring schedule. | [optional] 
**ReportingSettings** | Pointer to [**ReportingSettingsMutable**](ReportingSettingsMutable.md) |  | [optional] 

## Methods

### NewCampaignDetailsReadOnly

`func NewCampaignDetailsReadOnly(name string, scheduleSettings ScheduleSettingsReadOnly, resourceSettings ResourceSettingsMutable, reviewerSettings ReviewerSettingsMutable, remediationSettings RemediationSettings, ) *CampaignDetailsReadOnly`

NewCampaignDetailsReadOnly instantiates a new CampaignDetailsReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDetailsReadOnlyWithDefaults

`func NewCampaignDetailsReadOnlyWithDefaults() *CampaignDetailsReadOnly`

NewCampaignDetailsReadOnlyWithDefaults instantiates a new CampaignDetailsReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignDetailsReadOnly) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignDetailsReadOnly) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignDetailsReadOnly) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignDetailsReadOnly) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignDetailsReadOnly) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignDetailsReadOnly) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignDetailsReadOnly) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCampaignType

`func (o *CampaignDetailsReadOnly) GetCampaignType() CampaignType`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CampaignDetailsReadOnly) GetCampaignTypeOk() (*CampaignType, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *CampaignDetailsReadOnly) SetCampaignType(v CampaignType)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *CampaignDetailsReadOnly) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetScheduleSettings

`func (o *CampaignDetailsReadOnly) GetScheduleSettings() ScheduleSettingsReadOnly`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *CampaignDetailsReadOnly) GetScheduleSettingsOk() (*ScheduleSettingsReadOnly, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *CampaignDetailsReadOnly) SetScheduleSettings(v ScheduleSettingsReadOnly)`

SetScheduleSettings sets ScheduleSettings field to given value.


### GetResourceSettings

`func (o *CampaignDetailsReadOnly) GetResourceSettings() ResourceSettingsMutable`

GetResourceSettings returns the ResourceSettings field if non-nil, zero value otherwise.

### GetResourceSettingsOk

`func (o *CampaignDetailsReadOnly) GetResourceSettingsOk() (*ResourceSettingsMutable, bool)`

GetResourceSettingsOk returns a tuple with the ResourceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettings

`func (o *CampaignDetailsReadOnly) SetResourceSettings(v ResourceSettingsMutable)`

SetResourceSettings sets ResourceSettings field to given value.


### GetPrincipalScopeSettings

`func (o *CampaignDetailsReadOnly) GetPrincipalScopeSettings() PrincipalScopeSettingsMutable`

GetPrincipalScopeSettings returns the PrincipalScopeSettings field if non-nil, zero value otherwise.

### GetPrincipalScopeSettingsOk

`func (o *CampaignDetailsReadOnly) GetPrincipalScopeSettingsOk() (*PrincipalScopeSettingsMutable, bool)`

GetPrincipalScopeSettingsOk returns a tuple with the PrincipalScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalScopeSettings

`func (o *CampaignDetailsReadOnly) SetPrincipalScopeSettings(v PrincipalScopeSettingsMutable)`

SetPrincipalScopeSettings sets PrincipalScopeSettings field to given value.

### HasPrincipalScopeSettings

`func (o *CampaignDetailsReadOnly) HasPrincipalScopeSettings() bool`

HasPrincipalScopeSettings returns a boolean if a field has been set.

### GetReviewerSettings

`func (o *CampaignDetailsReadOnly) GetReviewerSettings() ReviewerSettingsMutable`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *CampaignDetailsReadOnly) GetReviewerSettingsOk() (*ReviewerSettingsMutable, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *CampaignDetailsReadOnly) SetReviewerSettings(v ReviewerSettingsMutable)`

SetReviewerSettings sets ReviewerSettings field to given value.


### GetNotificationSettings

`func (o *CampaignDetailsReadOnly) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *CampaignDetailsReadOnly) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *CampaignDetailsReadOnly) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *CampaignDetailsReadOnly) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetRemediationSettings

`func (o *CampaignDetailsReadOnly) GetRemediationSettings() RemediationSettings`

GetRemediationSettings returns the RemediationSettings field if non-nil, zero value otherwise.

### GetRemediationSettingsOk

`func (o *CampaignDetailsReadOnly) GetRemediationSettingsOk() (*RemediationSettings, bool)`

GetRemediationSettingsOk returns a tuple with the RemediationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationSettings

`func (o *CampaignDetailsReadOnly) SetRemediationSettings(v RemediationSettings)`

SetRemediationSettings sets RemediationSettings field to given value.


### GetRecurringCampaignId

`func (o *CampaignDetailsReadOnly) GetRecurringCampaignId() string`

GetRecurringCampaignId returns the RecurringCampaignId field if non-nil, zero value otherwise.

### GetRecurringCampaignIdOk

`func (o *CampaignDetailsReadOnly) GetRecurringCampaignIdOk() (*string, bool)`

GetRecurringCampaignIdOk returns a tuple with the RecurringCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCampaignId

`func (o *CampaignDetailsReadOnly) SetRecurringCampaignId(v string)`

SetRecurringCampaignId sets RecurringCampaignId field to given value.

### HasRecurringCampaignId

`func (o *CampaignDetailsReadOnly) HasRecurringCampaignId() bool`

HasRecurringCampaignId returns a boolean if a field has been set.

### SetRecurringCampaignIdNil

`func (o *CampaignDetailsReadOnly) SetRecurringCampaignIdNil(b bool)`

 SetRecurringCampaignIdNil sets the value for RecurringCampaignId to be an explicit nil

### UnsetRecurringCampaignId
`func (o *CampaignDetailsReadOnly) UnsetRecurringCampaignId()`

UnsetRecurringCampaignId ensures that no value is present for RecurringCampaignId, not even an explicit nil
### GetReportingSettings

`func (o *CampaignDetailsReadOnly) GetReportingSettings() ReportingSettingsMutable`

GetReportingSettings returns the ReportingSettings field if non-nil, zero value otherwise.

### GetReportingSettingsOk

`func (o *CampaignDetailsReadOnly) GetReportingSettingsOk() (*ReportingSettingsMutable, bool)`

GetReportingSettingsOk returns a tuple with the ReportingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingSettings

`func (o *CampaignDetailsReadOnly) SetReportingSettings(v ReportingSettingsMutable)`

SetReportingSettings sets ReportingSettings field to given value.

### HasReportingSettings

`func (o *CampaignDetailsReadOnly) HasReportingSettings() bool`

HasReportingSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


