# SecurityAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | Pointer to [**map[string]Link**](Link.md) | Links to related resources | [optional] 
**Status** | [**SecurityAccessReviewStatus**](SecurityAccessReviewStatus.md) |  | 
**Name** | **string** | The name of the security access review | 
**EndTime** | **time.Time** | The end time of the security access review | 
**ReviewerSettings** | [**SecurityAccessReviewReviewerSettingsResponse**](SecurityAccessReviewReviewerSettingsResponse.md) |  | 
**Summary** | Pointer to [**AiMessage**](AiMessage.md) |  | [optional] 

## Methods

### NewSecurityAccessReview

`func NewSecurityAccessReview(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, status SecurityAccessReviewStatus, name string, endTime time.Time, reviewerSettings SecurityAccessReviewReviewerSettingsResponse, ) *SecurityAccessReview`

NewSecurityAccessReview instantiates a new SecurityAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewWithDefaults

`func NewSecurityAccessReviewWithDefaults() *SecurityAccessReview`

NewSecurityAccessReviewWithDefaults instantiates a new SecurityAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReview) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *SecurityAccessReview) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SecurityAccessReview) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SecurityAccessReview) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *SecurityAccessReview) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SecurityAccessReview) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SecurityAccessReview) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *SecurityAccessReview) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SecurityAccessReview) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SecurityAccessReview) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *SecurityAccessReview) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *SecurityAccessReview) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *SecurityAccessReview) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *SecurityAccessReview) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReview) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReview) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReview) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityAccessReview) GetStatus() SecurityAccessReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityAccessReview) GetStatusOk() (*SecurityAccessReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityAccessReview) SetStatus(v SecurityAccessReviewStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *SecurityAccessReview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityAccessReview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityAccessReview) SetName(v string)`

SetName sets Name field to given value.


### GetEndTime

`func (o *SecurityAccessReview) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SecurityAccessReview) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SecurityAccessReview) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetReviewerSettings

`func (o *SecurityAccessReview) GetReviewerSettings() SecurityAccessReviewReviewerSettingsResponse`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *SecurityAccessReview) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettingsResponse, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *SecurityAccessReview) SetReviewerSettings(v SecurityAccessReviewReviewerSettingsResponse)`

SetReviewerSettings sets ReviewerSettings field to given value.


### GetSummary

`func (o *SecurityAccessReview) GetSummary() AiMessage`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SecurityAccessReview) GetSummaryOk() (*AiMessage, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SecurityAccessReview) SetSummary(v AiMessage)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SecurityAccessReview) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


