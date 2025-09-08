# RequestTypesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RequestTypeSparse**](RequestTypeSparse.md) | All Request Types on the current page | [optional] 
**Links** | Pointer to [**RequestTypesListLinks**](RequestTypesListLinks.md) |  | [optional] 

## Methods

### NewRequestTypesList

`func NewRequestTypesList() *RequestTypesList`

NewRequestTypesList instantiates a new RequestTypesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypesListWithDefaults

`func NewRequestTypesListWithDefaults() *RequestTypesList`

NewRequestTypesListWithDefaults instantiates a new RequestTypesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RequestTypesList) GetData() []RequestTypeSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestTypesList) GetDataOk() (*[]RequestTypeSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestTypesList) SetData(v []RequestTypeSparse)`

SetData sets Data field to given value.

### HasData

`func (o *RequestTypesList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *RequestTypesList) GetLinks() RequestTypesListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestTypesList) GetLinksOk() (*RequestTypesListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestTypesList) SetLinks(v RequestTypesListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RequestTypesList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


