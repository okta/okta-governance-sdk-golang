# GrantTypeEntitlementWriteable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | Additive grant type for assigning specific entitlement and respective value(s). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Entitlements** | [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | 
**TargetPrincipal** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Action** | Pointer to [**GrantAction**](GrantAction.md) |  | [optional] [default to GRANTACTION_ALLOW]
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]

## Methods

### NewGrantTypeEntitlementWriteable

`func NewGrantTypeEntitlementWriteable(grantType string, target TargetResource, entitlements []EntitlementCreatable, targetPrincipal TargetPrincipal, ) *GrantTypeEntitlementWriteable`

NewGrantTypeEntitlementWriteable instantiates a new GrantTypeEntitlementWriteable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantTypeEntitlementWriteableWithDefaults

`func NewGrantTypeEntitlementWriteableWithDefaults() *GrantTypeEntitlementWriteable`

NewGrantTypeEntitlementWriteableWithDefaults instantiates a new GrantTypeEntitlementWriteable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GrantTypeEntitlementWriteable) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrantTypeEntitlementWriteable) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrantTypeEntitlementWriteable) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetTarget

`func (o *GrantTypeEntitlementWriteable) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GrantTypeEntitlementWriteable) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GrantTypeEntitlementWriteable) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetEntitlements

`func (o *GrantTypeEntitlementWriteable) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantTypeEntitlementWriteable) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantTypeEntitlementWriteable) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.


### GetTargetPrincipal

`func (o *GrantTypeEntitlementWriteable) GetTargetPrincipal() TargetPrincipal`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *GrantTypeEntitlementWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *GrantTypeEntitlementWriteable) SetTargetPrincipal(v TargetPrincipal)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetScheduleSettings

`func (o *GrantTypeEntitlementWriteable) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GrantTypeEntitlementWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GrantTypeEntitlementWriteable) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *GrantTypeEntitlementWriteable) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetAction

`func (o *GrantTypeEntitlementWriteable) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GrantTypeEntitlementWriteable) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GrantTypeEntitlementWriteable) SetAction(v GrantAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *GrantTypeEntitlementWriteable) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *GrantTypeEntitlementWriteable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantTypeEntitlementWriteable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantTypeEntitlementWriteable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GrantTypeEntitlementWriteable) HasActor() bool`

HasActor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


