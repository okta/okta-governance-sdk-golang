# RequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RequestSparse**](RequestSparse.md) | All requests on the current page | [optional] 
**Links** | Pointer to [**RequestListLinks**](RequestListLinks.md) |  | [optional] 

## Methods

### NewRequestList

`func NewRequestList() *RequestList`

NewRequestList instantiates a new RequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestListWithDefaults

`func NewRequestListWithDefaults() *RequestList`

NewRequestListWithDefaults instantiates a new RequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RequestList) GetData() []RequestSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestList) GetDataOk() (*[]RequestSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestList) SetData(v []RequestSparse)`

SetData sets Data field to given value.

### HasData

`func (o *RequestList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *RequestList) GetLinks() RequestListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestList) GetLinksOk() (*RequestListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestList) SetLinks(v RequestListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RequestList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


