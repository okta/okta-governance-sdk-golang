# ReviewsReassign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewerId** | **string** | The Okta user &#x60;id&#x60; of the new reviewer | 
**ReviewIds** | **[]string** | A list of reviews (review &#x60;id&#x60; values) that are reassigned to the new reviewer. | 
**ReviewerLevel** | Pointer to [**ReviewerLevelType**](ReviewerLevelType.md) |  | [optional] 
**Note** | **string** | A note to justify the reassignment decision for the specified review(s). | 

## Methods

### NewReviewsReassign

`func NewReviewsReassign(reviewerId string, reviewIds []string, note string, ) *ReviewsReassign`

NewReviewsReassign instantiates a new ReviewsReassign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewsReassignWithDefaults

`func NewReviewsReassignWithDefaults() *ReviewsReassign`

NewReviewsReassignWithDefaults instantiates a new ReviewsReassign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewerId

`func (o *ReviewsReassign) GetReviewerId() string`

GetReviewerId returns the ReviewerId field if non-nil, zero value otherwise.

### GetReviewerIdOk

`func (o *ReviewsReassign) GetReviewerIdOk() (*string, bool)`

GetReviewerIdOk returns a tuple with the ReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerId

`func (o *ReviewsReassign) SetReviewerId(v string)`

SetReviewerId sets ReviewerId field to given value.


### GetReviewIds

`func (o *ReviewsReassign) GetReviewIds() []string`

GetReviewIds returns the ReviewIds field if non-nil, zero value otherwise.

### GetReviewIdsOk

`func (o *ReviewsReassign) GetReviewIdsOk() (*[]string, bool)`

GetReviewIdsOk returns a tuple with the ReviewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewIds

`func (o *ReviewsReassign) SetReviewIds(v []string)`

SetReviewIds sets ReviewIds field to given value.


### GetReviewerLevel

`func (o *ReviewsReassign) GetReviewerLevel() ReviewerLevelType`

GetReviewerLevel returns the ReviewerLevel field if non-nil, zero value otherwise.

### GetReviewerLevelOk

`func (o *ReviewsReassign) GetReviewerLevelOk() (*ReviewerLevelType, bool)`

GetReviewerLevelOk returns a tuple with the ReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerLevel

`func (o *ReviewsReassign) SetReviewerLevel(v ReviewerLevelType)`

SetReviewerLevel sets ReviewerLevel field to given value.

### HasReviewerLevel

`func (o *ReviewsReassign) HasReviewerLevel() bool`

HasReviewerLevel returns a boolean if a field has been set.

### GetNote

`func (o *ReviewsReassign) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReviewsReassign) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReviewsReassign) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


