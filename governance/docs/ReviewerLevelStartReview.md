# ReviewerLevelStartReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnDay** | **int32** | The day on which that reviewer level will start.  It will be &#x60;0&#x60; for &#x60;FIRST&#x60; reviewer level, as the first level will start when the campaign starts.  For &#x60;SECOND&#x60; reviewer level specify a value greater than &#x60;0&#x60;. This will indicate the day, the reviews will be moved to second level.  | [default to 0]
**When** | Pointer to [**ReviewerLowerLevelCondition**](ReviewerLowerLevelCondition.md) |  | [optional] 

## Methods

### NewReviewerLevelStartReview

`func NewReviewerLevelStartReview(onDay int32, ) *ReviewerLevelStartReview`

NewReviewerLevelStartReview instantiates a new ReviewerLevelStartReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerLevelStartReviewWithDefaults

`func NewReviewerLevelStartReviewWithDefaults() *ReviewerLevelStartReview`

NewReviewerLevelStartReviewWithDefaults instantiates a new ReviewerLevelStartReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnDay

`func (o *ReviewerLevelStartReview) GetOnDay() int32`

GetOnDay returns the OnDay field if non-nil, zero value otherwise.

### GetOnDayOk

`func (o *ReviewerLevelStartReview) GetOnDayOk() (*int32, bool)`

GetOnDayOk returns a tuple with the OnDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDay

`func (o *ReviewerLevelStartReview) SetOnDay(v int32)`

SetOnDay sets OnDay field to given value.


### GetWhen

`func (o *ReviewerLevelStartReview) GetWhen() ReviewerLowerLevelCondition`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ReviewerLevelStartReview) GetWhenOk() (*ReviewerLowerLevelCondition, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ReviewerLevelStartReview) SetWhen(v ReviewerLowerLevelCondition)`

SetWhen sets When field to given value.

### HasWhen

`func (o *ReviewerLevelStartReview) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


