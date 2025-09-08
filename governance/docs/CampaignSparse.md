# CampaignSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the campaign. Maintain some uniqueness when naming the campaign as it helps to identify and filter for campaigns when needed. | 
**Description** | Pointer to **string** | Human readable description. | [optional] 
**StartDate** | **time.Time** | The date on which the campaign is scheduled to start. Accepts date in ISO 8601 format. | 
**EndDate** | **time.Time** | The date on which the campaign is supposed to end. Accepts date in ISO 8601 format. | 
**ScheduleType** | [**ScheduleType**](ScheduleType.md) |  | 
**RecurringCampaignId** | Pointer to **NullableString** | ID of the recurring campaign if this campaign was created as part of a recurring schedule. | [optional] 
**ReviewerType** | [**CampaignReviewerType**](CampaignReviewerType.md) |  | 
**Links** | [**CampaignLinks**](CampaignLinks.md) |  | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Status** | [**CampaignStatus**](CampaignStatus.md) |  | 

## Methods

### NewCampaignSparse

`func NewCampaignSparse(name string, startDate time.Time, endDate time.Time, scheduleType ScheduleType, reviewerType CampaignReviewerType, links CampaignLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, status CampaignStatus, ) *CampaignSparse`

NewCampaignSparse instantiates a new CampaignSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSparseWithDefaults

`func NewCampaignSparseWithDefaults() *CampaignSparse`

NewCampaignSparseWithDefaults instantiates a new CampaignSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignSparse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSparse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignSparse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignSparse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignSparse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignSparse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignSparse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartDate

`func (o *CampaignSparse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CampaignSparse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CampaignSparse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CampaignSparse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CampaignSparse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CampaignSparse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetScheduleType

`func (o *CampaignSparse) GetScheduleType() ScheduleType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *CampaignSparse) GetScheduleTypeOk() (*ScheduleType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *CampaignSparse) SetScheduleType(v ScheduleType)`

SetScheduleType sets ScheduleType field to given value.


### GetRecurringCampaignId

`func (o *CampaignSparse) GetRecurringCampaignId() string`

GetRecurringCampaignId returns the RecurringCampaignId field if non-nil, zero value otherwise.

### GetRecurringCampaignIdOk

`func (o *CampaignSparse) GetRecurringCampaignIdOk() (*string, bool)`

GetRecurringCampaignIdOk returns a tuple with the RecurringCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCampaignId

`func (o *CampaignSparse) SetRecurringCampaignId(v string)`

SetRecurringCampaignId sets RecurringCampaignId field to given value.

### HasRecurringCampaignId

`func (o *CampaignSparse) HasRecurringCampaignId() bool`

HasRecurringCampaignId returns a boolean if a field has been set.

### SetRecurringCampaignIdNil

`func (o *CampaignSparse) SetRecurringCampaignIdNil(b bool)`

 SetRecurringCampaignIdNil sets the value for RecurringCampaignId to be an explicit nil

### UnsetRecurringCampaignId
`func (o *CampaignSparse) UnsetRecurringCampaignId()`

UnsetRecurringCampaignId ensures that no value is present for RecurringCampaignId, not even an explicit nil
### GetReviewerType

`func (o *CampaignSparse) GetReviewerType() CampaignReviewerType`

GetReviewerType returns the ReviewerType field if non-nil, zero value otherwise.

### GetReviewerTypeOk

`func (o *CampaignSparse) GetReviewerTypeOk() (*CampaignReviewerType, bool)`

GetReviewerTypeOk returns a tuple with the ReviewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerType

`func (o *CampaignSparse) SetReviewerType(v CampaignReviewerType)`

SetReviewerType sets ReviewerType field to given value.


### GetLinks

`func (o *CampaignSparse) GetLinks() CampaignLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignSparse) GetLinksOk() (*CampaignLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignSparse) SetLinks(v CampaignLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *CampaignSparse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSparse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSparse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *CampaignSparse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignSparse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CampaignSparse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *CampaignSparse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignSparse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignSparse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *CampaignSparse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CampaignSparse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CampaignSparse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *CampaignSparse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *CampaignSparse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *CampaignSparse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetStatus

`func (o *CampaignSparse) GetStatus() CampaignStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignSparse) GetStatusOk() (*CampaignStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignSparse) SetStatus(v CampaignStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


