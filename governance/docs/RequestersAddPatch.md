# RequestersAddPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | Operation to add a new member to the existing requester set | 
**Path** | **string** | JSON Pointer path that specifies adding to the members array | 
**Value** | [**RequestersMember**](RequestersMember.md) |  | 

## Methods

### NewRequestersAddPatch

`func NewRequestersAddPatch(op string, path string, value RequestersMember, ) *RequestersAddPatch`

NewRequestersAddPatch instantiates a new RequestersAddPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestersAddPatchWithDefaults

`func NewRequestersAddPatchWithDefaults() *RequestersAddPatch`

NewRequestersAddPatchWithDefaults instantiates a new RequestersAddPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RequestersAddPatch) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RequestersAddPatch) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RequestersAddPatch) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *RequestersAddPatch) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestersAddPatch) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestersAddPatch) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RequestersAddPatch) GetValue() RequestersMember`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestersAddPatch) GetValueOk() (*RequestersMember, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestersAddPatch) SetValue(v RequestersMember)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


