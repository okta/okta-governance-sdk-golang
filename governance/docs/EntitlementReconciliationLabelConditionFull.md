# EntitlementReconciliationLabelConditionFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **string** | Identifies this condition as targeting a label | 
**LabelId** | **string** | The ID of a label | 
**LabelName** | Pointer to **string** | The display name of the label, resolved by Okta on read | [optional] 
**Values** | Pointer to [**[]EntitlementReconciliationLabelConditionValueFull**](EntitlementReconciliationLabelConditionValueFull.md) | The values of the label the condition is narrowed to. Absent or empty when the condition matches every value. | [optional] 

## Methods

### NewEntitlementReconciliationLabelConditionFull

`func NewEntitlementReconciliationLabelConditionFull(refType string, labelId string, ) *EntitlementReconciliationLabelConditionFull`

NewEntitlementReconciliationLabelConditionFull instantiates a new EntitlementReconciliationLabelConditionFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationLabelConditionFullWithDefaults

`func NewEntitlementReconciliationLabelConditionFullWithDefaults() *EntitlementReconciliationLabelConditionFull`

NewEntitlementReconciliationLabelConditionFullWithDefaults instantiates a new EntitlementReconciliationLabelConditionFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *EntitlementReconciliationLabelConditionFull) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementReconciliationLabelConditionFull) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementReconciliationLabelConditionFull) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetLabelId

`func (o *EntitlementReconciliationLabelConditionFull) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *EntitlementReconciliationLabelConditionFull) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *EntitlementReconciliationLabelConditionFull) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetLabelName

`func (o *EntitlementReconciliationLabelConditionFull) GetLabelName() string`

GetLabelName returns the LabelName field if non-nil, zero value otherwise.

### GetLabelNameOk

`func (o *EntitlementReconciliationLabelConditionFull) GetLabelNameOk() (*string, bool)`

GetLabelNameOk returns a tuple with the LabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelName

`func (o *EntitlementReconciliationLabelConditionFull) SetLabelName(v string)`

SetLabelName sets LabelName field to given value.

### HasLabelName

`func (o *EntitlementReconciliationLabelConditionFull) HasLabelName() bool`

HasLabelName returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementReconciliationLabelConditionFull) GetValues() []EntitlementReconciliationLabelConditionValueFull`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementReconciliationLabelConditionFull) GetValuesOk() (*[]EntitlementReconciliationLabelConditionValueFull, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementReconciliationLabelConditionFull) SetValues(v []EntitlementReconciliationLabelConditionValueFull)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementReconciliationLabelConditionFull) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


