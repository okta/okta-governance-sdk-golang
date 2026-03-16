# SecurityAccessReviewHistoryItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecurityAccessReviewHistoryItem**](SecurityAccessReviewHistoryItem.md) | Lists all events occurred on a given security access review. The history covers all such events and actions. | [optional] 
**Links** | Pointer to [**SecurityAccessReviewListLinks**](SecurityAccessReviewListLinks.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewHistoryItems

`func NewSecurityAccessReviewHistoryItems() *SecurityAccessReviewHistoryItems`

NewSecurityAccessReviewHistoryItems instantiates a new SecurityAccessReviewHistoryItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewHistoryItemsWithDefaults

`func NewSecurityAccessReviewHistoryItemsWithDefaults() *SecurityAccessReviewHistoryItems`

NewSecurityAccessReviewHistoryItemsWithDefaults instantiates a new SecurityAccessReviewHistoryItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityAccessReviewHistoryItems) GetData() []SecurityAccessReviewHistoryItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityAccessReviewHistoryItems) GetDataOk() (*[]SecurityAccessReviewHistoryItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityAccessReviewHistoryItems) SetData(v []SecurityAccessReviewHistoryItem)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityAccessReviewHistoryItems) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityAccessReviewHistoryItems) GetLinks() SecurityAccessReviewListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReviewHistoryItems) GetLinksOk() (*SecurityAccessReviewListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReviewHistoryItems) SetLinks(v SecurityAccessReviewListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReviewHistoryItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


