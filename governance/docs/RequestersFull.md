# RequestersFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Members** | [**[]RequestersMember**](RequestersMember.md) |  | 

## Methods

### NewRequestersFull

`func NewRequestersFull(id string, members []RequestersMember, ) *RequestersFull`

NewRequestersFull instantiates a new RequestersFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestersFullWithDefaults

`func NewRequestersFullWithDefaults() *RequestersFull`

NewRequestersFullWithDefaults instantiates a new RequestersFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestersFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestersFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestersFull) SetId(v string)`

SetId sets Id field to given value.


### GetMembers

`func (o *RequestersFull) GetMembers() []RequestersMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RequestersFull) GetMembersOk() (*[]RequestersMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *RequestersFull) SetMembers(v []RequestersMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


