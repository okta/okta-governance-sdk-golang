# EntitlementReconciliationConfigWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**ReconciliationMode**](ReconciliationMode.md) |  | 
**Additive** | Pointer to [**EntitlementReconciliationDirectionWritable**](EntitlementReconciliationDirectionWritable.md) |  | [optional] 
**Subtractive** | Pointer to [**EntitlementReconciliationDirectionWritable**](EntitlementReconciliationDirectionWritable.md) |  | [optional] 

## Methods

### NewEntitlementReconciliationConfigWritable

`func NewEntitlementReconciliationConfigWritable(mode ReconciliationMode, ) *EntitlementReconciliationConfigWritable`

NewEntitlementReconciliationConfigWritable instantiates a new EntitlementReconciliationConfigWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationConfigWritableWithDefaults

`func NewEntitlementReconciliationConfigWritableWithDefaults() *EntitlementReconciliationConfigWritable`

NewEntitlementReconciliationConfigWritableWithDefaults instantiates a new EntitlementReconciliationConfigWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *EntitlementReconciliationConfigWritable) GetMode() ReconciliationMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EntitlementReconciliationConfigWritable) GetModeOk() (*ReconciliationMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EntitlementReconciliationConfigWritable) SetMode(v ReconciliationMode)`

SetMode sets Mode field to given value.


### GetAdditive

`func (o *EntitlementReconciliationConfigWritable) GetAdditive() EntitlementReconciliationDirectionWritable`

GetAdditive returns the Additive field if non-nil, zero value otherwise.

### GetAdditiveOk

`func (o *EntitlementReconciliationConfigWritable) GetAdditiveOk() (*EntitlementReconciliationDirectionWritable, bool)`

GetAdditiveOk returns a tuple with the Additive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditive

`func (o *EntitlementReconciliationConfigWritable) SetAdditive(v EntitlementReconciliationDirectionWritable)`

SetAdditive sets Additive field to given value.

### HasAdditive

`func (o *EntitlementReconciliationConfigWritable) HasAdditive() bool`

HasAdditive returns a boolean if a field has been set.

### GetSubtractive

`func (o *EntitlementReconciliationConfigWritable) GetSubtractive() EntitlementReconciliationDirectionWritable`

GetSubtractive returns the Subtractive field if non-nil, zero value otherwise.

### GetSubtractiveOk

`func (o *EntitlementReconciliationConfigWritable) GetSubtractiveOk() (*EntitlementReconciliationDirectionWritable, bool)`

GetSubtractiveOk returns a tuple with the Subtractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtractive

`func (o *EntitlementReconciliationConfigWritable) SetSubtractive(v EntitlementReconciliationDirectionWritable)`

SetSubtractive sets Subtractive field to given value.

### HasSubtractive

`func (o *EntitlementReconciliationConfigWritable) HasSubtractive() bool`

HasSubtractive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


