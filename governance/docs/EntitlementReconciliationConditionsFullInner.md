# EntitlementReconciliationConditionsFullInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **string** | Identifies this condition as targeting an entitlement | 
**LabelId** | **string** | The ID of a label | 
**LabelName** | Pointer to **string** | The display name of the label, resolved by Okta on read | [optional] 
**Values** | Pointer to [**[]EntitlementReconciliationEntitlementConditionValueFull**](EntitlementReconciliationEntitlementConditionValueFull.md) | The values of the entitlement the condition is narrowed to. Absent or empty when the condition matches every value. | [optional] 
**EntitlementId** | **string** | The &#x60;id&#x60; property of an entitlement | 
**EntitlementName** | Pointer to **string** | The display name for an entitlement property | [optional] 

## Methods

### NewEntitlementReconciliationConditionsFullInner

`func NewEntitlementReconciliationConditionsFullInner(refType string, labelId string, entitlementId string, ) *EntitlementReconciliationConditionsFullInner`

NewEntitlementReconciliationConditionsFullInner instantiates a new EntitlementReconciliationConditionsFullInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationConditionsFullInnerWithDefaults

`func NewEntitlementReconciliationConditionsFullInnerWithDefaults() *EntitlementReconciliationConditionsFullInner`

NewEntitlementReconciliationConditionsFullInnerWithDefaults instantiates a new EntitlementReconciliationConditionsFullInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *EntitlementReconciliationConditionsFullInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementReconciliationConditionsFullInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementReconciliationConditionsFullInner) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetLabelId

`func (o *EntitlementReconciliationConditionsFullInner) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *EntitlementReconciliationConditionsFullInner) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *EntitlementReconciliationConditionsFullInner) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetLabelName

`func (o *EntitlementReconciliationConditionsFullInner) GetLabelName() string`

GetLabelName returns the LabelName field if non-nil, zero value otherwise.

### GetLabelNameOk

`func (o *EntitlementReconciliationConditionsFullInner) GetLabelNameOk() (*string, bool)`

GetLabelNameOk returns a tuple with the LabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelName

`func (o *EntitlementReconciliationConditionsFullInner) SetLabelName(v string)`

SetLabelName sets LabelName field to given value.

### HasLabelName

`func (o *EntitlementReconciliationConditionsFullInner) HasLabelName() bool`

HasLabelName returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementReconciliationConditionsFullInner) GetValues() []EntitlementReconciliationEntitlementConditionValueFull`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementReconciliationConditionsFullInner) GetValuesOk() (*[]EntitlementReconciliationEntitlementConditionValueFull, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementReconciliationConditionsFullInner) SetValues(v []EntitlementReconciliationEntitlementConditionValueFull)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementReconciliationConditionsFullInner) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetEntitlementId

`func (o *EntitlementReconciliationConditionsFullInner) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementReconciliationConditionsFullInner) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementReconciliationConditionsFullInner) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetEntitlementName

`func (o *EntitlementReconciliationConditionsFullInner) GetEntitlementName() string`

GetEntitlementName returns the EntitlementName field if non-nil, zero value otherwise.

### GetEntitlementNameOk

`func (o *EntitlementReconciliationConditionsFullInner) GetEntitlementNameOk() (*string, bool)`

GetEntitlementNameOk returns a tuple with the EntitlementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementName

`func (o *EntitlementReconciliationConditionsFullInner) SetEntitlementName(v string)`

SetEntitlementName sets EntitlementName field to given value.

### HasEntitlementName

`func (o *EntitlementReconciliationConditionsFullInner) HasEntitlementName() bool`

HasEntitlementName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


