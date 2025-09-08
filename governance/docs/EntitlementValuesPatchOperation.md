# EntitlementValuesPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**EntitlementValuePatchOp**](EntitlementValuePatchOp.md) |  | 
**Path** | **string** | The path of the property being updated. ex - &#x60;/values/{id}&#x60; for &#x60;REMOVE&#x60;, &#x60;REPLACE&#x60; and &#x60;/values/-&#x60; for &#x60;ADD&#x60; for entitlement value update. | 
**Value** | Pointer to [**EntitlementValuesPatchOperationValue**](EntitlementValuesPatchOperationValue.md) |  | [optional] 
**RefType** | **string** |  | 

## Methods

### NewEntitlementValuesPatchOperation

`func NewEntitlementValuesPatchOperation(op EntitlementValuePatchOp, path string, refType string, ) *EntitlementValuesPatchOperation`

NewEntitlementValuesPatchOperation instantiates a new EntitlementValuesPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValuesPatchOperationWithDefaults

`func NewEntitlementValuesPatchOperationWithDefaults() *EntitlementValuesPatchOperation`

NewEntitlementValuesPatchOperationWithDefaults instantiates a new EntitlementValuesPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *EntitlementValuesPatchOperation) GetOp() EntitlementValuePatchOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EntitlementValuesPatchOperation) GetOpOk() (*EntitlementValuePatchOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EntitlementValuesPatchOperation) SetOp(v EntitlementValuePatchOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *EntitlementValuesPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EntitlementValuesPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EntitlementValuesPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *EntitlementValuesPatchOperation) GetValue() EntitlementValuesPatchOperationValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementValuesPatchOperation) GetValueOk() (*EntitlementValuesPatchOperationValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementValuesPatchOperation) SetValue(v EntitlementValuesPatchOperationValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementValuesPatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRefType

`func (o *EntitlementValuesPatchOperation) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementValuesPatchOperation) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementValuesPatchOperation) SetRefType(v string)`

SetRefType sets RefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


