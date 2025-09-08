# GrantTypeCustomWriteable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | Base grant type for assigning specific entitlement and respective value(s). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Entitlements** | Pointer to [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | [optional] 
**TargetPrincipal** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Action** | Pointer to [**GrantAction**](GrantAction.md) |  | [optional] [default to GRANTACTION_ALLOW]
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]

## Methods

### NewGrantTypeCustomWriteable

`func NewGrantTypeCustomWriteable(grantType string, target TargetResource, targetPrincipal TargetPrincipal, ) *GrantTypeCustomWriteable`

NewGrantTypeCustomWriteable instantiates a new GrantTypeCustomWriteable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantTypeCustomWriteableWithDefaults

`func NewGrantTypeCustomWriteableWithDefaults() *GrantTypeCustomWriteable`

NewGrantTypeCustomWriteableWithDefaults instantiates a new GrantTypeCustomWriteable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GrantTypeCustomWriteable) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrantTypeCustomWriteable) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrantTypeCustomWriteable) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetTarget

`func (o *GrantTypeCustomWriteable) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GrantTypeCustomWriteable) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GrantTypeCustomWriteable) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetEntitlements

`func (o *GrantTypeCustomWriteable) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantTypeCustomWriteable) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantTypeCustomWriteable) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *GrantTypeCustomWriteable) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetTargetPrincipal

`func (o *GrantTypeCustomWriteable) GetTargetPrincipal() TargetPrincipal`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *GrantTypeCustomWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *GrantTypeCustomWriteable) SetTargetPrincipal(v TargetPrincipal)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetScheduleSettings

`func (o *GrantTypeCustomWriteable) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GrantTypeCustomWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GrantTypeCustomWriteable) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *GrantTypeCustomWriteable) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetAction

`func (o *GrantTypeCustomWriteable) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GrantTypeCustomWriteable) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GrantTypeCustomWriteable) SetAction(v GrantAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *GrantTypeCustomWriteable) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *GrantTypeCustomWriteable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantTypeCustomWriteable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantTypeCustomWriteable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GrantTypeCustomWriteable) HasActor() bool`

HasActor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


