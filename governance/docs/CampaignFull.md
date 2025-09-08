# CampaignFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**CampaignLinks**](CampaignLinks.md) |  | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Name** | **string** | Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed. | 
**Description** | Pointer to **string** | Human readable description. | [optional] 
**CampaignType** | Pointer to [**CampaignType**](CampaignType.md) |  | [optional] 
**ScheduleSettings** | [**ScheduleSettingsReadOnly**](ScheduleSettingsReadOnly.md) |  | 
**ResourceSettings** | [**ResourceSettingsMutable**](ResourceSettingsMutable.md) |  | 
**PrincipalScopeSettings** | Pointer to [**PrincipalScopeSettingsMutable**](PrincipalScopeSettingsMutable.md) |  | [optional] 
**ReviewerSettings** | [**ReviewerSettingsMutable**](ReviewerSettingsMutable.md) |  | 
**NotificationSettings** | Pointer to [**NotificationSettings**](NotificationSettings.md) |  | [optional] 
**RemediationSettings** | [**RemediationSettings**](RemediationSettings.md) |  | 
**RecurringCampaignId** | Pointer to **NullableString** | ID of the recurring campaign if this campaign was created as part of a recurring schedule. | [optional] 
**Status** | [**CampaignStatus**](CampaignStatus.md) |  | 

## Methods

### NewCampaignFull

`func NewCampaignFull(links CampaignLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, name string, scheduleSettings ScheduleSettingsReadOnly, resourceSettings ResourceSettingsMutable, reviewerSettings ReviewerSettingsMutable, remediationSettings RemediationSettings, status CampaignStatus, ) *CampaignFull`

NewCampaignFull instantiates a new CampaignFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignFullWithDefaults

`func NewCampaignFullWithDefaults() *CampaignFull`

NewCampaignFullWithDefaults instantiates a new CampaignFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CampaignFull) GetLinks() CampaignLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignFull) GetLinksOk() (*CampaignLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignFull) SetLinks(v CampaignLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *CampaignFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *CampaignFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CampaignFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *CampaignFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *CampaignFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CampaignFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CampaignFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *CampaignFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *CampaignFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *CampaignFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetName

`func (o *CampaignFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignFull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCampaignType

`func (o *CampaignFull) GetCampaignType() CampaignType`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CampaignFull) GetCampaignTypeOk() (*CampaignType, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *CampaignFull) SetCampaignType(v CampaignType)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *CampaignFull) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetScheduleSettings

`func (o *CampaignFull) GetScheduleSettings() ScheduleSettingsReadOnly`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *CampaignFull) GetScheduleSettingsOk() (*ScheduleSettingsReadOnly, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *CampaignFull) SetScheduleSettings(v ScheduleSettingsReadOnly)`

SetScheduleSettings sets ScheduleSettings field to given value.


### GetResourceSettings

`func (o *CampaignFull) GetResourceSettings() ResourceSettingsMutable`

GetResourceSettings returns the ResourceSettings field if non-nil, zero value otherwise.

### GetResourceSettingsOk

`func (o *CampaignFull) GetResourceSettingsOk() (*ResourceSettingsMutable, bool)`

GetResourceSettingsOk returns a tuple with the ResourceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettings

`func (o *CampaignFull) SetResourceSettings(v ResourceSettingsMutable)`

SetResourceSettings sets ResourceSettings field to given value.


### GetPrincipalScopeSettings

`func (o *CampaignFull) GetPrincipalScopeSettings() PrincipalScopeSettingsMutable`

GetPrincipalScopeSettings returns the PrincipalScopeSettings field if non-nil, zero value otherwise.

### GetPrincipalScopeSettingsOk

`func (o *CampaignFull) GetPrincipalScopeSettingsOk() (*PrincipalScopeSettingsMutable, bool)`

GetPrincipalScopeSettingsOk returns a tuple with the PrincipalScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalScopeSettings

`func (o *CampaignFull) SetPrincipalScopeSettings(v PrincipalScopeSettingsMutable)`

SetPrincipalScopeSettings sets PrincipalScopeSettings field to given value.

### HasPrincipalScopeSettings

`func (o *CampaignFull) HasPrincipalScopeSettings() bool`

HasPrincipalScopeSettings returns a boolean if a field has been set.

### GetReviewerSettings

`func (o *CampaignFull) GetReviewerSettings() ReviewerSettingsMutable`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *CampaignFull) GetReviewerSettingsOk() (*ReviewerSettingsMutable, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *CampaignFull) SetReviewerSettings(v ReviewerSettingsMutable)`

SetReviewerSettings sets ReviewerSettings field to given value.


### GetNotificationSettings

`func (o *CampaignFull) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *CampaignFull) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *CampaignFull) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *CampaignFull) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetRemediationSettings

`func (o *CampaignFull) GetRemediationSettings() RemediationSettings`

GetRemediationSettings returns the RemediationSettings field if non-nil, zero value otherwise.

### GetRemediationSettingsOk

`func (o *CampaignFull) GetRemediationSettingsOk() (*RemediationSettings, bool)`

GetRemediationSettingsOk returns a tuple with the RemediationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationSettings

`func (o *CampaignFull) SetRemediationSettings(v RemediationSettings)`

SetRemediationSettings sets RemediationSettings field to given value.


### GetRecurringCampaignId

`func (o *CampaignFull) GetRecurringCampaignId() string`

GetRecurringCampaignId returns the RecurringCampaignId field if non-nil, zero value otherwise.

### GetRecurringCampaignIdOk

`func (o *CampaignFull) GetRecurringCampaignIdOk() (*string, bool)`

GetRecurringCampaignIdOk returns a tuple with the RecurringCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCampaignId

`func (o *CampaignFull) SetRecurringCampaignId(v string)`

SetRecurringCampaignId sets RecurringCampaignId field to given value.

### HasRecurringCampaignId

`func (o *CampaignFull) HasRecurringCampaignId() bool`

HasRecurringCampaignId returns a boolean if a field has been set.

### SetRecurringCampaignIdNil

`func (o *CampaignFull) SetRecurringCampaignIdNil(b bool)`

 SetRecurringCampaignIdNil sets the value for RecurringCampaignId to be an explicit nil

### UnsetRecurringCampaignId
`func (o *CampaignFull) UnsetRecurringCampaignId()`

UnsetRecurringCampaignId ensures that no value is present for RecurringCampaignId, not even an explicit nil
### GetStatus

`func (o *CampaignFull) GetStatus() CampaignStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignFull) GetStatusOk() (*CampaignStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignFull) SetStatus(v CampaignStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


