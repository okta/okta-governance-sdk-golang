# ReviewActionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewId** | **string** | Review ID | 
**Action** | [**ReviewAction**](ReviewAction.md) |  | 
**ReviewerId** | Pointer to **string** | The reassigned reviewer&#39;s Okta user ID Required when &#x60;action&#x60; is &#x60;REASSIGN&#x60;. All items with &#x60;action&#x60; set to &#x60;REASSIGN&#x60; in a single request must specify the same &#x60;reviewerId&#x60;.  | [optional] 
**Note** | Pointer to **NullableString** | Optional note to justify the action taken on this review | [optional] 

## Methods

### NewReviewActionItem

`func NewReviewActionItem(reviewId string, action ReviewAction, ) *ReviewActionItem`

NewReviewActionItem instantiates a new ReviewActionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewActionItemWithDefaults

`func NewReviewActionItemWithDefaults() *ReviewActionItem`

NewReviewActionItemWithDefaults instantiates a new ReviewActionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewId

`func (o *ReviewActionItem) GetReviewId() string`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *ReviewActionItem) GetReviewIdOk() (*string, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *ReviewActionItem) SetReviewId(v string)`

SetReviewId sets ReviewId field to given value.


### GetAction

`func (o *ReviewActionItem) GetAction() ReviewAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReviewActionItem) GetActionOk() (*ReviewAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReviewActionItem) SetAction(v ReviewAction)`

SetAction sets Action field to given value.


### GetReviewerId

`func (o *ReviewActionItem) GetReviewerId() string`

GetReviewerId returns the ReviewerId field if non-nil, zero value otherwise.

### GetReviewerIdOk

`func (o *ReviewActionItem) GetReviewerIdOk() (*string, bool)`

GetReviewerIdOk returns a tuple with the ReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerId

`func (o *ReviewActionItem) SetReviewerId(v string)`

SetReviewerId sets ReviewerId field to given value.

### HasReviewerId

`func (o *ReviewActionItem) HasReviewerId() bool`

HasReviewerId returns a boolean if a field has been set.

### GetNote

`func (o *ReviewActionItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReviewActionItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReviewActionItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReviewActionItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ReviewActionItem) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ReviewActionItem) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


