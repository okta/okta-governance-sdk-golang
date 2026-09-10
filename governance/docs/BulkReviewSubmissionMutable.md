# BulkReviewSubmissionMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | [**BulkReviewSubmissionCriteriaMutable**](BulkReviewSubmissionCriteriaMutable.md) |  | 
**ReviewDecision** | [**BulkReviewDecision**](BulkReviewDecision.md) |  | 
**ReviewerLevel** | Pointer to [**ReviewerLevelType**](ReviewerLevelType.md) |  | [optional] 
**Notes** | Pointer to **string** | The reviewer&#39;s note to include with the bulk decision. This note is added to all reviews that match the bulk-review criteria. | [optional] 

## Methods

### NewBulkReviewSubmissionMutable

`func NewBulkReviewSubmissionMutable(criteria BulkReviewSubmissionCriteriaMutable, reviewDecision BulkReviewDecision, ) *BulkReviewSubmissionMutable`

NewBulkReviewSubmissionMutable instantiates a new BulkReviewSubmissionMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReviewSubmissionMutableWithDefaults

`func NewBulkReviewSubmissionMutableWithDefaults() *BulkReviewSubmissionMutable`

NewBulkReviewSubmissionMutableWithDefaults instantiates a new BulkReviewSubmissionMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *BulkReviewSubmissionMutable) GetCriteria() BulkReviewSubmissionCriteriaMutable`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *BulkReviewSubmissionMutable) GetCriteriaOk() (*BulkReviewSubmissionCriteriaMutable, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *BulkReviewSubmissionMutable) SetCriteria(v BulkReviewSubmissionCriteriaMutable)`

SetCriteria sets Criteria field to given value.


### GetReviewDecision

`func (o *BulkReviewSubmissionMutable) GetReviewDecision() BulkReviewDecision`

GetReviewDecision returns the ReviewDecision field if non-nil, zero value otherwise.

### GetReviewDecisionOk

`func (o *BulkReviewSubmissionMutable) GetReviewDecisionOk() (*BulkReviewDecision, bool)`

GetReviewDecisionOk returns a tuple with the ReviewDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewDecision

`func (o *BulkReviewSubmissionMutable) SetReviewDecision(v BulkReviewDecision)`

SetReviewDecision sets ReviewDecision field to given value.


### GetReviewerLevel

`func (o *BulkReviewSubmissionMutable) GetReviewerLevel() ReviewerLevelType`

GetReviewerLevel returns the ReviewerLevel field if non-nil, zero value otherwise.

### GetReviewerLevelOk

`func (o *BulkReviewSubmissionMutable) GetReviewerLevelOk() (*ReviewerLevelType, bool)`

GetReviewerLevelOk returns a tuple with the ReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerLevel

`func (o *BulkReviewSubmissionMutable) SetReviewerLevel(v ReviewerLevelType)`

SetReviewerLevel sets ReviewerLevel field to given value.

### HasReviewerLevel

`func (o *BulkReviewSubmissionMutable) HasReviewerLevel() bool`

HasReviewerLevel returns a boolean if a field has been set.

### GetNotes

`func (o *BulkReviewSubmissionMutable) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BulkReviewSubmissionMutable) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BulkReviewSubmissionMutable) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BulkReviewSubmissionMutable) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


