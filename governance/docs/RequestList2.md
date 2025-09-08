# RequestList2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RequestSparse2**](RequestSparse2.md) | All requests on the current page | [optional] 
**Links** | Pointer to [**RequestListLinks2**](RequestListLinks2.md) |  | [optional] 

## Methods

### NewRequestList2

`func NewRequestList2() *RequestList2`

NewRequestList2 instantiates a new RequestList2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestList2WithDefaults

`func NewRequestList2WithDefaults() *RequestList2`

NewRequestList2WithDefaults instantiates a new RequestList2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RequestList2) GetData() []RequestSparse2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestList2) GetDataOk() (*[]RequestSparse2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestList2) SetData(v []RequestSparse2)`

SetData sets Data field to given value.

### HasData

`func (o *RequestList2) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *RequestList2) GetLinks() RequestListLinks2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestList2) GetLinksOk() (*RequestListLinks2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestList2) SetLinks(v RequestListLinks2)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RequestList2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


