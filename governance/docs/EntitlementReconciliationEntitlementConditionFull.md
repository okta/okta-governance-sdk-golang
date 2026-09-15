# EntitlementReconciliationEntitlementConditionFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **string** | Identifies this condition as targeting an entitlement | 
**EntitlementId** | **string** | The &#x60;id&#x60; property of an entitlement | 
**EntitlementName** | Pointer to **string** | The display name for an entitlement property | [optional] 
**Values** | Pointer to [**[]EntitlementReconciliationEntitlementConditionValueFull**](EntitlementReconciliationEntitlementConditionValueFull.md) | The values of the entitlement the condition is narrowed to. Absent or empty when the condition matches every value. | [optional] 

## Methods

### NewEntitlementReconciliationEntitlementConditionFull

`func NewEntitlementReconciliationEntitlementConditionFull(refType string, entitlementId string, ) *EntitlementReconciliationEntitlementConditionFull`

NewEntitlementReconciliationEntitlementConditionFull instantiates a new EntitlementReconciliationEntitlementConditionFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationEntitlementConditionFullWithDefaults

`func NewEntitlementReconciliationEntitlementConditionFullWithDefaults() *EntitlementReconciliationEntitlementConditionFull`

NewEntitlementReconciliationEntitlementConditionFullWithDefaults instantiates a new EntitlementReconciliationEntitlementConditionFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *EntitlementReconciliationEntitlementConditionFull) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *EntitlementReconciliationEntitlementConditionFull) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *EntitlementReconciliationEntitlementConditionFull) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetEntitlementId

`func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementReconciliationEntitlementConditionFull) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetEntitlementName

`func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementName() string`

GetEntitlementName returns the EntitlementName field if non-nil, zero value otherwise.

### GetEntitlementNameOk

`func (o *EntitlementReconciliationEntitlementConditionFull) GetEntitlementNameOk() (*string, bool)`

GetEntitlementNameOk returns a tuple with the EntitlementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementName

`func (o *EntitlementReconciliationEntitlementConditionFull) SetEntitlementName(v string)`

SetEntitlementName sets EntitlementName field to given value.

### HasEntitlementName

`func (o *EntitlementReconciliationEntitlementConditionFull) HasEntitlementName() bool`

HasEntitlementName returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementReconciliationEntitlementConditionFull) GetValues() []EntitlementReconciliationEntitlementConditionValueFull`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementReconciliationEntitlementConditionFull) GetValuesOk() (*[]EntitlementReconciliationEntitlementConditionValueFull, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementReconciliationEntitlementConditionFull) SetValues(v []EntitlementReconciliationEntitlementConditionValueFull)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementReconciliationEntitlementConditionFull) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


