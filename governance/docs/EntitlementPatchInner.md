# EntitlementPatchInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**EntitlementPatchOp**](EntitlementPatchOp.md) |  | 
**Path** | **string** | The path of the property being updated. ex - &#x60;/name&#x60; and &#x60;/description&#x60; for the entitlement update. | 
**Value** | Pointer to **string** | The value of the property being updated. | [optional] 
**RefType** | **string** |  | 

## Methods

### NewEntitlementPatchInner

`func NewEntitlementPatchInner(op EntitlementPatchOp, path string, refType string, ) *EntitlementPatchInner`

NewEntitlementPatchInner instantiates a new EntitlementPatchInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementPatchInnerWithDefaults

`func NewEntitlementPatchInnerWithDefaults() *EntitlementPatchInner`

NewEntitlementPatchInnerWithDefaults instantiates a new EntitlementPatchInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *EntitlementPatchInner) GetOp() EntitlementPatchOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EntitlementPatchInner) GetOpOk() (*EntitlementPatchOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EntitlementPatchInner) SetOp(v EntitlementPatchOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *EntitlementPatchInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EntitlementPatchInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EntitlementPatchInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *EntitlementPatchInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementPatchInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementPatchInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementPatchInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRefType

`func (o *EntitlementPatchInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementPatchInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementPatchInner) SetRefType(v string)`

SetRefType sets RefType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


