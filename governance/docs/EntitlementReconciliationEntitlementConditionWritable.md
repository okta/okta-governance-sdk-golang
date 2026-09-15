# EntitlementReconciliationEntitlementConditionWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **string** | Identifies this condition as targeting an entitlement | 
**EntitlementId** | **string** | The &#x60;id&#x60; property of an entitlement | 
**Values** | Pointer to [**[]EntitlementReconciliationEntitlementConditionValueWritable**](EntitlementReconciliationEntitlementConditionValueWritable.md) | The values of the entitlement to narrow to. Omit it, or send an empty array, to match every value of the entitlement. | [optional] 

## Methods

### NewEntitlementReconciliationEntitlementConditionWritable

`func NewEntitlementReconciliationEntitlementConditionWritable(refType string, entitlementId string, ) *EntitlementReconciliationEntitlementConditionWritable`

NewEntitlementReconciliationEntitlementConditionWritable instantiates a new EntitlementReconciliationEntitlementConditionWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationEntitlementConditionWritableWithDefaults

`func NewEntitlementReconciliationEntitlementConditionWritableWithDefaults() *EntitlementReconciliationEntitlementConditionWritable`

NewEntitlementReconciliationEntitlementConditionWritableWithDefaults instantiates a new EntitlementReconciliationEntitlementConditionWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *EntitlementReconciliationEntitlementConditionWritable) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementReconciliationEntitlementConditionWritable) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementReconciliationEntitlementConditionWritable) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetEntitlementId

`func (o *EntitlementReconciliationEntitlementConditionWritable) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementReconciliationEntitlementConditionWritable) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementReconciliationEntitlementConditionWritable) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetValues

`func (o *EntitlementReconciliationEntitlementConditionWritable) GetValues() []EntitlementReconciliationEntitlementConditionValueWritable`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementReconciliationEntitlementConditionWritable) GetValuesOk() (*[]EntitlementReconciliationEntitlementConditionValueWritable, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementReconciliationEntitlementConditionWritable) SetValues(v []EntitlementReconciliationEntitlementConditionValueWritable)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementReconciliationEntitlementConditionWritable) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


