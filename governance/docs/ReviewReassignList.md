# ReviewReassignList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ReviewSparse**](ReviewSparse.md) | All the reviews that were reassigned  | 
**Links** | [**ReviewReassignListLinks**](ReviewReassignListLinks.md) |  | 

## Methods

### NewReviewReassignList

`func NewReviewReassignList(data []ReviewSparse, links ReviewReassignListLinks, ) *ReviewReassignList`

NewReviewReassignList instantiates a new ReviewReassignList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewReassignListWithDefaults

`func NewReviewReassignListWithDefaults() *ReviewReassignList`

NewReviewReassignListWithDefaults instantiates a new ReviewReassignList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReviewReassignList) GetData() []ReviewSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReviewReassignList) GetDataOk() (*[]ReviewSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReviewReassignList) SetData(v []ReviewSparse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ReviewReassignList) GetLinks() ReviewReassignListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewReassignList) GetLinksOk() (*ReviewReassignListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewReassignList) SetLinks(v ReviewReassignListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


