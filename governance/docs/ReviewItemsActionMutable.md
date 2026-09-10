# ReviewItemsActionMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviews** | [**[]ReviewActionItem**](ReviewActionItem.md) | List of review actions  | 
**ReviewerLevel** | Pointer to [**ReviewerLevelType**](ReviewerLevelType.md) |  | [optional] 

## Methods

### NewReviewItemsActionMutable

`func NewReviewItemsActionMutable(reviews []ReviewActionItem, ) *ReviewItemsActionMutable`

NewReviewItemsActionMutable instantiates a new ReviewItemsActionMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewItemsActionMutableWithDefaults

`func NewReviewItemsActionMutableWithDefaults() *ReviewItemsActionMutable`

NewReviewItemsActionMutableWithDefaults instantiates a new ReviewItemsActionMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviews

`func (o *ReviewItemsActionMutable) GetReviews() []ReviewActionItem`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *ReviewItemsActionMutable) GetReviewsOk() (*[]ReviewActionItem, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *ReviewItemsActionMutable) SetReviews(v []ReviewActionItem)`

SetReviews sets Reviews field to given value.


### GetReviewerLevel

`func (o *ReviewItemsActionMutable) GetReviewerLevel() ReviewerLevelType`

GetReviewerLevel returns the ReviewerLevel field if non-nil, zero value otherwise.

### GetReviewerLevelOk

`func (o *ReviewItemsActionMutable) GetReviewerLevelOk() (*ReviewerLevelType, bool)`

GetReviewerLevelOk returns a tuple with the ReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerLevel

`func (o *ReviewItemsActionMutable) SetReviewerLevel(v ReviewerLevelType)`

SetReviewerLevel sets ReviewerLevel field to given value.

### HasReviewerLevel

`func (o *ReviewItemsActionMutable) HasReviewerLevel() bool`

HasReviewerLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


