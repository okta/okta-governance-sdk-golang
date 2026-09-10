# BulkReviewSubmissionJobDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The &#x60;id&#x60; of the job | 
**Status** | [**BulkReviewDecisionStatusType**](BulkReviewDecisionStatusType.md) |  | 
**Errors** | Pointer to [**[]ModelError**](ModelError.md) | A list of errors encountered while processing the bulk review submission job | [optional] 
**EligibleReviewRecordCount** | **int32** | The number of eligible reviews in the submission | 
**Links** | Pointer to [**BulkReviewSubmissionLinks**](BulkReviewSubmissionLinks.md) |  | [optional] 

## Methods

### NewBulkReviewSubmissionJobDetails

`func NewBulkReviewSubmissionJobDetails(id string, status BulkReviewDecisionStatusType, eligibleReviewRecordCount int32, ) *BulkReviewSubmissionJobDetails`

NewBulkReviewSubmissionJobDetails instantiates a new BulkReviewSubmissionJobDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReviewSubmissionJobDetailsWithDefaults

`func NewBulkReviewSubmissionJobDetailsWithDefaults() *BulkReviewSubmissionJobDetails`

NewBulkReviewSubmissionJobDetailsWithDefaults instantiates a new BulkReviewSubmissionJobDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkReviewSubmissionJobDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkReviewSubmissionJobDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkReviewSubmissionJobDetails) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BulkReviewSubmissionJobDetails) GetStatus() BulkReviewDecisionStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkReviewSubmissionJobDetails) GetStatusOk() (*BulkReviewDecisionStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkReviewSubmissionJobDetails) SetStatus(v BulkReviewDecisionStatusType)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *BulkReviewSubmissionJobDetails) GetErrors() []ModelError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkReviewSubmissionJobDetails) GetErrorsOk() (*[]ModelError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkReviewSubmissionJobDetails) SetErrors(v []ModelError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkReviewSubmissionJobDetails) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetEligibleReviewRecordCount

`func (o *BulkReviewSubmissionJobDetails) GetEligibleReviewRecordCount() int32`

GetEligibleReviewRecordCount returns the EligibleReviewRecordCount field if non-nil, zero value otherwise.

### GetEligibleReviewRecordCountOk

`func (o *BulkReviewSubmissionJobDetails) GetEligibleReviewRecordCountOk() (*int32, bool)`

GetEligibleReviewRecordCountOk returns a tuple with the EligibleReviewRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleReviewRecordCount

`func (o *BulkReviewSubmissionJobDetails) SetEligibleReviewRecordCount(v int32)`

SetEligibleReviewRecordCount sets EligibleReviewRecordCount field to given value.


### GetLinks

`func (o *BulkReviewSubmissionJobDetails) GetLinks() BulkReviewSubmissionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BulkReviewSubmissionJobDetails) GetLinksOk() (*BulkReviewSubmissionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BulkReviewSubmissionJobDetails) SetLinks(v BulkReviewSubmissionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BulkReviewSubmissionJobDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


