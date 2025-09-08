# RequestTypeLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateRequest** | Pointer to [**Link**](Link.md) | Only available for request type with status of ACTIVE | [optional] 
**Requests** | [**Link**](Link.md) |  | 
**Self** | [**Link**](Link.md) |  | 

## Methods

### NewRequestTypeLinks

`func NewRequestTypeLinks(requests Link, self Link, ) *RequestTypeLinks`

NewRequestTypeLinks instantiates a new RequestTypeLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeLinksWithDefaults

`func NewRequestTypeLinksWithDefaults() *RequestTypeLinks`

NewRequestTypeLinksWithDefaults instantiates a new RequestTypeLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateRequest

`func (o *RequestTypeLinks) GetCreateRequest() Link`

GetCreateRequest returns the CreateRequest field if non-nil, zero value otherwise.

### GetCreateRequestOk

`func (o *RequestTypeLinks) GetCreateRequestOk() (*Link, bool)`

GetCreateRequestOk returns a tuple with the CreateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateRequest

`func (o *RequestTypeLinks) SetCreateRequest(v Link)`

SetCreateRequest sets CreateRequest field to given value.

### HasCreateRequest

`func (o *RequestTypeLinks) HasCreateRequest() bool`

HasCreateRequest returns a boolean if a field has been set.

### GetRequests

`func (o *RequestTypeLinks) GetRequests() Link`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestTypeLinks) GetRequestsOk() (*Link, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestTypeLinks) SetRequests(v Link)`

SetRequests sets Requests field to given value.


### GetSelf

`func (o *RequestTypeLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RequestTypeLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RequestTypeLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


