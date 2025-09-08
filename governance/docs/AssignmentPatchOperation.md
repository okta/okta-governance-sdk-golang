# AssignmentPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**AssignmentPatchOp**](AssignmentPatchOp.md) |  | 
**Path** | **string** | The path of the property that&#39;s updated. For example, &#x60;/expirationTime&#x60; and &#x60;/timeZone&#x60; for the assignment update. | 
**Value** | Pointer to **string** | The value of the property that&#39;s updated | [optional] 

## Methods

### NewAssignmentPatchOperation

`func NewAssignmentPatchOperation(op AssignmentPatchOp, path string, ) *AssignmentPatchOperation`

NewAssignmentPatchOperation instantiates a new AssignmentPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentPatchOperationWithDefaults

`func NewAssignmentPatchOperationWithDefaults() *AssignmentPatchOperation`

NewAssignmentPatchOperationWithDefaults instantiates a new AssignmentPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AssignmentPatchOperation) GetOp() AssignmentPatchOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AssignmentPatchOperation) GetOpOk() (*AssignmentPatchOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AssignmentPatchOperation) SetOp(v AssignmentPatchOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *AssignmentPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AssignmentPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AssignmentPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *AssignmentPatchOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AssignmentPatchOperation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AssignmentPatchOperation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AssignmentPatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


