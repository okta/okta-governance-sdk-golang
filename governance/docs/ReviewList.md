# ReviewList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ReviewSparse**](ReviewSparse.md) | Sparse representation of reviews  | 
**Links** | [**ReviewListLinks**](ReviewListLinks.md) |  | 

## Methods

### NewReviewList

`func NewReviewList(data []ReviewSparse, links ReviewListLinks, ) *ReviewList`

NewReviewList instantiates a new ReviewList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewListWithDefaults

`func NewReviewListWithDefaults() *ReviewList`

NewReviewListWithDefaults instantiates a new ReviewList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReviewList) GetData() []ReviewSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReviewList) GetDataOk() (*[]ReviewSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReviewList) SetData(v []ReviewSparse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ReviewList) GetLinks() ReviewListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewList) GetLinksOk() (*ReviewListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewList) SetLinks(v ReviewListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


