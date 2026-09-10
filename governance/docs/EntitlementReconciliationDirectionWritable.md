# EntitlementReconciliationDirectionWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ReconciliationDirectionAction**](ReconciliationDirectionAction.md) |  | 
**Conditions** | [**[]EntitlementReconciliationConditionsWritableInner**](EntitlementReconciliationConditionsWritableInner.md) | The drifts that a direction&#39;s &#x60;action&#x60; applies to.  Conditions combine with &#x60;OR&#x60;, so a drift that matches any one of them gets the action. A drift that matches none of them gets the inverse action.  An empty list means the action applies to every drift in the direction.  | 

## Methods

### NewEntitlementReconciliationDirectionWritable

`func NewEntitlementReconciliationDirectionWritable(action ReconciliationDirectionAction, conditions []EntitlementReconciliationConditionsWritableInner, ) *EntitlementReconciliationDirectionWritable`

NewEntitlementReconciliationDirectionWritable instantiates a new EntitlementReconciliationDirectionWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationDirectionWritableWithDefaults

`func NewEntitlementReconciliationDirectionWritableWithDefaults() *EntitlementReconciliationDirectionWritable`

NewEntitlementReconciliationDirectionWritableWithDefaults instantiates a new EntitlementReconciliationDirectionWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *EntitlementReconciliationDirectionWritable) GetAction() ReconciliationDirectionAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EntitlementReconciliationDirectionWritable) GetActionOk() (*ReconciliationDirectionAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EntitlementReconciliationDirectionWritable) SetAction(v ReconciliationDirectionAction)`

SetAction sets Action field to given value.


### GetConditions

`func (o *EntitlementReconciliationDirectionWritable) GetConditions() []EntitlementReconciliationConditionsWritableInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EntitlementReconciliationDirectionWritable) GetConditionsOk() (*[]EntitlementReconciliationConditionsWritableInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EntitlementReconciliationDirectionWritable) SetConditions(v []EntitlementReconciliationConditionsWritableInner)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


