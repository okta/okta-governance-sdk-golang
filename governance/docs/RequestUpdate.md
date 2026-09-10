# RequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**RequestStatusPatchable**](RequestStatusPatchable.md) |  | 

## Methods

### NewRequestUpdate

`func NewRequestUpdate(status RequestStatusPatchable, ) *RequestUpdate`

NewRequestUpdate instantiates a new RequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpdateWithDefaults

`func NewRequestUpdateWithDefaults() *RequestUpdate`

NewRequestUpdateWithDefaults instantiates a new RequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestUpdate) GetStatus() RequestStatusPatchable`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestUpdate) GetStatusOk() (*RequestStatusPatchable, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestUpdate) SetStatus(v RequestStatusPatchable)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


