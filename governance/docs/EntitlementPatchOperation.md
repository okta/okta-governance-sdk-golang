# EntitlementPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**EntitlementPatchOp**](EntitlementPatchOp.md) |  | 
**Path** | **string** | The path of the property being updated. ex - &#x60;/name&#x60; and &#x60;/description&#x60; for the entitlement update. | 
**Value** | Pointer to **string** | The value of the property being updated. | [optional] 
**RefType** | **string** |  | 

## Methods

### NewEntitlementPatchOperation

`func NewEntitlementPatchOperation(op EntitlementPatchOp, path string, refType string, ) *EntitlementPatchOperation`

NewEntitlementPatchOperation instantiates a new EntitlementPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementPatchOperationWithDefaults

`func NewEntitlementPatchOperationWithDefaults() *EntitlementPatchOperation`

NewEntitlementPatchOperationWithDefaults instantiates a new EntitlementPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *EntitlementPatchOperation) GetOp() EntitlementPatchOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EntitlementPatchOperation) GetOpOk() (*EntitlementPatchOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EntitlementPatchOperation) SetOp(v EntitlementPatchOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *EntitlementPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EntitlementPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EntitlementPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *EntitlementPatchOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementPatchOperation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementPatchOperation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementPatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRefType

`func (o *EntitlementPatchOperation) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementPatchOperation) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementPatchOperation) SetRefType(v string)`

SetRefType sets RefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


