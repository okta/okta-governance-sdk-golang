# BulkReviewSubmissionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BulkReviewDecisionStatusType**](BulkReviewDecisionStatusType.md) |  | 
**EligibleReviewRecordCount** | **int32** | The number of eligible reviews in the submission | 
**JobId** | **string** | The &#x60;id&#x60; of the job | 
**Links** | Pointer to [**BulkReviewSubmissionLinks**](BulkReviewSubmissionLinks.md) |  | [optional] 

## Methods

### NewBulkReviewSubmissionDetails

`func NewBulkReviewSubmissionDetails(status BulkReviewDecisionStatusType, eligibleReviewRecordCount int32, jobId string, ) *BulkReviewSubmissionDetails`

NewBulkReviewSubmissionDetails instantiates a new BulkReviewSubmissionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReviewSubmissionDetailsWithDefaults

`func NewBulkReviewSubmissionDetailsWithDefaults() *BulkReviewSubmissionDetails`

NewBulkReviewSubmissionDetailsWithDefaults instantiates a new BulkReviewSubmissionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkReviewSubmissionDetails) GetStatus() BulkReviewDecisionStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkReviewSubmissionDetails) GetStatusOk() (*BulkReviewDecisionStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkReviewSubmissionDetails) SetStatus(v BulkReviewDecisionStatusType)`

SetStatus sets Status field to given value.


### GetEligibleReviewRecordCount

`func (o *BulkReviewSubmissionDetails) GetEligibleReviewRecordCount() int32`

GetEligibleReviewRecordCount returns the EligibleReviewRecordCount field if non-nil, zero value otherwise.

### GetEligibleReviewRecordCountOk

`func (o *BulkReviewSubmissionDetails) GetEligibleReviewRecordCountOk() (*int32, bool)`

GetEligibleReviewRecordCountOk returns a tuple with the EligibleReviewRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleReviewRecordCount

`func (o *BulkReviewSubmissionDetails) SetEligibleReviewRecordCount(v int32)`

SetEligibleReviewRecordCount sets EligibleReviewRecordCount field to given value.


### GetJobId

`func (o *BulkReviewSubmissionDetails) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BulkReviewSubmissionDetails) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BulkReviewSubmissionDetails) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetLinks

`func (o *BulkReviewSubmissionDetails) GetLinks() BulkReviewSubmissionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BulkReviewSubmissionDetails) GetLinksOk() (*BulkReviewSubmissionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BulkReviewSubmissionDetails) SetLinks(v BulkReviewSubmissionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BulkReviewSubmissionDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


