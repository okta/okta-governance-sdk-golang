# RequestersMembersPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | The operation to perform on the requesters set | 
**Path** | **string** | JSON Pointer path to the target location | 
**Value** | [**RequestersMember**](RequestersMember.md) |  | 

## Methods

### NewRequestersMembersPatchOperation

`func NewRequestersMembersPatchOperation(op string, path string, value RequestersMember, ) *RequestersMembersPatchOperation`

NewRequestersMembersPatchOperation instantiates a new RequestersMembersPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestersMembersPatchOperationWithDefaults

`func NewRequestersMembersPatchOperationWithDefaults() *RequestersMembersPatchOperation`

NewRequestersMembersPatchOperationWithDefaults instantiates a new RequestersMembersPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RequestersMembersPatchOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RequestersMembersPatchOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RequestersMembersPatchOperation) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *RequestersMembersPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestersMembersPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestersMembersPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RequestersMembersPatchOperation) GetValue() RequestersMember`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestersMembersPatchOperation) GetValueOk() (*RequestersMember, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestersMembersPatchOperation) SetValue(v RequestersMember)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


