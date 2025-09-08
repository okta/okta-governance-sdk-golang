# SecurityAccessReviewSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the security access review | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Status** | [**SecurityAccessReviewStatus**](SecurityAccessReviewStatus.md) |  | 
**Name** | **string** | The name of the security access review | 
**EndTime** | **time.Time** | The end time of the security access review | 
**ReviewerSettings** | [**SecurityAccessReviewReviewerSettings**](SecurityAccessReviewReviewerSettings.md) |  | 

## Methods

### NewSecurityAccessReviewSparse

`func NewSecurityAccessReviewSparse(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, status SecurityAccessReviewStatus, name string, endTime time.Time, reviewerSettings SecurityAccessReviewReviewerSettings, ) *SecurityAccessReviewSparse`

NewSecurityAccessReviewSparse instantiates a new SecurityAccessReviewSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSparseWithDefaults

`func NewSecurityAccessReviewSparseWithDefaults() *SecurityAccessReviewSparse`

NewSecurityAccessReviewSparseWithDefaults instantiates a new SecurityAccessReviewSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewSparse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewSparse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewSparse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *SecurityAccessReviewSparse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SecurityAccessReviewSparse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SecurityAccessReviewSparse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *SecurityAccessReviewSparse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SecurityAccessReviewSparse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SecurityAccessReviewSparse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *SecurityAccessReviewSparse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SecurityAccessReviewSparse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SecurityAccessReviewSparse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *SecurityAccessReviewSparse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *SecurityAccessReviewSparse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *SecurityAccessReviewSparse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *SecurityAccessReviewSparse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReviewSparse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReviewSparse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReviewSparse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityAccessReviewSparse) GetStatus() SecurityAccessReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityAccessReviewSparse) GetStatusOk() (*SecurityAccessReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityAccessReviewSparse) SetStatus(v SecurityAccessReviewStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *SecurityAccessReviewSparse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityAccessReviewSparse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityAccessReviewSparse) SetName(v string)`

SetName sets Name field to given value.


### GetEndTime

`func (o *SecurityAccessReviewSparse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SecurityAccessReviewSparse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SecurityAccessReviewSparse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetReviewerSettings

`func (o *SecurityAccessReviewSparse) GetReviewerSettings() SecurityAccessReviewReviewerSettings`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *SecurityAccessReviewSparse) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *SecurityAccessReviewSparse) SetReviewerSettings(v SecurityAccessReviewReviewerSettings)`

SetReviewerSettings sets ReviewerSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


