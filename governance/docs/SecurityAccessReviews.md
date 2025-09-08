# SecurityAccessReviews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecurityAccessReviewSparse**](SecurityAccessReviewSparse.md) | Security access reviews entries | [optional] 
**Links** | Pointer to [**SecurityAccessReviewListLinks**](SecurityAccessReviewListLinks.md) |  | [optional] 

## Methods

### NewSecurityAccessReviews

`func NewSecurityAccessReviews() *SecurityAccessReviews`

NewSecurityAccessReviews instantiates a new SecurityAccessReviews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewsWithDefaults

`func NewSecurityAccessReviewsWithDefaults() *SecurityAccessReviews`

NewSecurityAccessReviewsWithDefaults instantiates a new SecurityAccessReviews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityAccessReviews) GetData() []SecurityAccessReviewSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityAccessReviews) GetDataOk() (*[]SecurityAccessReviewSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityAccessReviews) SetData(v []SecurityAccessReviewSparse)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityAccessReviews) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityAccessReviews) GetLinks() SecurityAccessReviewListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReviews) GetLinksOk() (*SecurityAccessReviewListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReviews) SetLinks(v SecurityAccessReviewListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReviews) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


