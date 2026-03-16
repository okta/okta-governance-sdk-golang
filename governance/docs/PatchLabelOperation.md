# PatchLabelOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**LabelPatchOp**](LabelPatchOp.md) |  | 
**Path** | **string** | The path of the property to update. Use the &#x60;/name&#x60; path to update the label category name. | 
**Value** | Pointer to **string** | The value of the updated property | [optional] 
**RefType** | **string** | The label property for the update operation | 

## Methods

### NewPatchLabelOperation

`func NewPatchLabelOperation(op LabelPatchOp, path string, refType string, ) *PatchLabelOperation`

NewPatchLabelOperation instantiates a new PatchLabelOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLabelOperationWithDefaults

`func NewPatchLabelOperationWithDefaults() *PatchLabelOperation`

NewPatchLabelOperationWithDefaults instantiates a new PatchLabelOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchLabelOperation) GetOp() LabelPatchOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchLabelOperation) GetOpOk() (*LabelPatchOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchLabelOperation) SetOp(v LabelPatchOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *PatchLabelOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchLabelOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchLabelOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchLabelOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchLabelOperation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchLabelOperation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchLabelOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRefType

`func (o *PatchLabelOperation) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PatchLabelOperation) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PatchLabelOperation) SetRefType(v string)`

SetRefType sets RefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


