# EntitlementReconciliationConditionsWritableInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **string** | Identifies this condition as targeting an entitlement | 
**LabelId** | **string** | The ID of a label | 
**Values** | Pointer to [**[]EntitlementReconciliationEntitlementConditionValueWritable**](EntitlementReconciliationEntitlementConditionValueWritable.md) | The values of the entitlement to narrow to. Omit it, or send an empty array, to match every value of the entitlement. | [optional] 
**EntitlementId** | **string** | The &#x60;id&#x60; property of an entitlement | 

## Methods

### NewEntitlementReconciliationConditionsWritableInner

`func NewEntitlementReconciliationConditionsWritableInner(refType string, labelId string, entitlementId string, ) *EntitlementReconciliationConditionsWritableInner`

NewEntitlementReconciliationConditionsWritableInner instantiates a new EntitlementReconciliationConditionsWritableInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationConditionsWritableInnerWithDefaults

`func NewEntitlementReconciliationConditionsWritableInnerWithDefaults() *EntitlementReconciliationConditionsWritableInner`

NewEntitlementReconciliationConditionsWritableInnerWithDefaults instantiates a new EntitlementReconciliationConditionsWritableInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *EntitlementReconciliationConditionsWritableInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementReconciliationConditionsWritableInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementReconciliationConditionsWritableInner) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetLabelId

`func (o *EntitlementReconciliationConditionsWritableInner) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *EntitlementReconciliationConditionsWritableInner) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *EntitlementReconciliationConditionsWritableInner) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetValues

`func (o *EntitlementReconciliationConditionsWritableInner) GetValues() []EntitlementReconciliationEntitlementConditionValueWritable`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementReconciliationConditionsWritableInner) GetValuesOk() (*[]EntitlementReconciliationEntitlementConditionValueWritable, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementReconciliationConditionsWritableInner) SetValues(v []EntitlementReconciliationEntitlementConditionValueWritable)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementReconciliationConditionsWritableInner) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetEntitlementId

`func (o *EntitlementReconciliationConditionsWritableInner) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementReconciliationConditionsWritableInner) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementReconciliationConditionsWritableInner) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


