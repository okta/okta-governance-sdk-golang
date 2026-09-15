# EntitlementReconciliationLabelConditionWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **string** | Identifies this condition as targeting a label | 
**LabelId** | **string** | The ID of a label | 
**Values** | Pointer to [**[]EntitlementReconciliationLabelConditionValueWritable**](EntitlementReconciliationLabelConditionValueWritable.md) | The values of the label to narrow to. Omit it, or send an empty array, to match every value of the label. | [optional] 

## Methods

### NewEntitlementReconciliationLabelConditionWritable

`func NewEntitlementReconciliationLabelConditionWritable(refType string, labelId string, ) *EntitlementReconciliationLabelConditionWritable`

NewEntitlementReconciliationLabelConditionWritable instantiates a new EntitlementReconciliationLabelConditionWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationLabelConditionWritableWithDefaults

`func NewEntitlementReconciliationLabelConditionWritableWithDefaults() *EntitlementReconciliationLabelConditionWritable`

NewEntitlementReconciliationLabelConditionWritableWithDefaults instantiates a new EntitlementReconciliationLabelConditionWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *EntitlementReconciliationLabelConditionWritable) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementReconciliationLabelConditionWritable) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementReconciliationLabelConditionWritable) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetLabelId

`func (o *EntitlementReconciliationLabelConditionWritable) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *EntitlementReconciliationLabelConditionWritable) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *EntitlementReconciliationLabelConditionWritable) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetValues

`func (o *EntitlementReconciliationLabelConditionWritable) GetValues() []EntitlementReconciliationLabelConditionValueWritable`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementReconciliationLabelConditionWritable) GetValuesOk() (*[]EntitlementReconciliationLabelConditionValueWritable, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementReconciliationLabelConditionWritable) SetValues(v []EntitlementReconciliationLabelConditionValueWritable)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementReconciliationLabelConditionWritable) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


