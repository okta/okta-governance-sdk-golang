# RequestersReplacePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | Operation to replace all existing members with a new set | 
**Path** | **string** | JSON Pointer path that specifies replacing the entire members array | 
**Value** | [**[]RequestersMember**](RequestersMember.md) |  | 

## Methods

### NewRequestersReplacePatch

`func NewRequestersReplacePatch(op string, path string, value []RequestersMember, ) *RequestersReplacePatch`

NewRequestersReplacePatch instantiates a new RequestersReplacePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestersReplacePatchWithDefaults

`func NewRequestersReplacePatchWithDefaults() *RequestersReplacePatch`

NewRequestersReplacePatchWithDefaults instantiates a new RequestersReplacePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RequestersReplacePatch) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RequestersReplacePatch) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RequestersReplacePatch) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *RequestersReplacePatch) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestersReplacePatch) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestersReplacePatch) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RequestersReplacePatch) GetValue() []RequestersMember`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestersReplacePatch) GetValueOk() (*[]RequestersMember, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestersReplacePatch) SetValue(v []RequestersMember)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


