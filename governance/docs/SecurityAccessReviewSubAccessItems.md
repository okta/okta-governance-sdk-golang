# SecurityAccessReviewSubAccessItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecurityAccessReviewSubAccessItem**](SecurityAccessReviewSubAccessItem.md) | List of access items in the security access review | [optional] 
**Links** | Pointer to [**SecurityAccessReviewListLinks**](SecurityAccessReviewListLinks.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewSubAccessItems

`func NewSecurityAccessReviewSubAccessItems() *SecurityAccessReviewSubAccessItems`

NewSecurityAccessReviewSubAccessItems instantiates a new SecurityAccessReviewSubAccessItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSubAccessItemsWithDefaults

`func NewSecurityAccessReviewSubAccessItemsWithDefaults() *SecurityAccessReviewSubAccessItems`

NewSecurityAccessReviewSubAccessItemsWithDefaults instantiates a new SecurityAccessReviewSubAccessItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityAccessReviewSubAccessItems) GetData() []SecurityAccessReviewSubAccessItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityAccessReviewSubAccessItems) GetDataOk() (*[]SecurityAccessReviewSubAccessItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityAccessReviewSubAccessItems) SetData(v []SecurityAccessReviewSubAccessItem)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityAccessReviewSubAccessItems) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityAccessReviewSubAccessItems) GetLinks() SecurityAccessReviewListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReviewSubAccessItems) GetLinksOk() (*SecurityAccessReviewListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReviewSubAccessItems) SetLinks(v SecurityAccessReviewListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReviewSubAccessItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


