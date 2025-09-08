# PatchLabelValueOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**LabelValuePatchOp**](LabelValuePatchOp.md) |  | 
**Path** | **string** | The path of the property to update. For example:  * &#x60;/values/-&#x60; for the &#x60;ADD&#x60; operation * &#x60;/values/{id}&#x60; for the &#x60;REMOVE&#x60; operation *&#x60;/values/{id}/...&#x60; for the &#x60;REPLACE&#x60; operation  | 
**Value** | Pointer to [**LabelValueUpdate**](LabelValueUpdate.md) |  | [optional] 
**RefType** | **string** |  | 

## Methods

### NewPatchLabelValueOperation

`func NewPatchLabelValueOperation(op LabelValuePatchOp, path string, refType string, ) *PatchLabelValueOperation`

NewPatchLabelValueOperation instantiates a new PatchLabelValueOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLabelValueOperationWithDefaults

`func NewPatchLabelValueOperationWithDefaults() *PatchLabelValueOperation`

NewPatchLabelValueOperationWithDefaults instantiates a new PatchLabelValueOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchLabelValueOperation) GetOp() LabelValuePatchOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchLabelValueOperation) GetOpOk() (*LabelValuePatchOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchLabelValueOperation) SetOp(v LabelValuePatchOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *PatchLabelValueOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchLabelValueOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchLabelValueOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchLabelValueOperation) GetValue() LabelValueUpdate`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchLabelValueOperation) GetValueOk() (*LabelValueUpdate, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchLabelValueOperation) SetValue(v LabelValueUpdate)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchLabelValueOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRefType

`func (o *PatchLabelValueOperation) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PatchLabelValueOperation) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PatchLabelValueOperation) SetRefType(v string)`

SetRefType sets RefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


