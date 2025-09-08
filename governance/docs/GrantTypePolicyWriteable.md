# GrantTypePolicyWriteable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | Base grant type for assigning entitlements driven by policy | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**TargetPrincipal** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Action** | Pointer to [**GrantAction**](GrantAction.md) |  | [optional] [default to GRANTACTION_ALLOW]
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]

## Methods

### NewGrantTypePolicyWriteable

`func NewGrantTypePolicyWriteable(grantType string, target TargetResource, targetPrincipal TargetPrincipal, ) *GrantTypePolicyWriteable`

NewGrantTypePolicyWriteable instantiates a new GrantTypePolicyWriteable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantTypePolicyWriteableWithDefaults

`func NewGrantTypePolicyWriteableWithDefaults() *GrantTypePolicyWriteable`

NewGrantTypePolicyWriteableWithDefaults instantiates a new GrantTypePolicyWriteable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GrantTypePolicyWriteable) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrantTypePolicyWriteable) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrantTypePolicyWriteable) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetTarget

`func (o *GrantTypePolicyWriteable) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GrantTypePolicyWriteable) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GrantTypePolicyWriteable) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetTargetPrincipal

`func (o *GrantTypePolicyWriteable) GetTargetPrincipal() TargetPrincipal`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *GrantTypePolicyWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *GrantTypePolicyWriteable) SetTargetPrincipal(v TargetPrincipal)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetScheduleSettings

`func (o *GrantTypePolicyWriteable) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GrantTypePolicyWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GrantTypePolicyWriteable) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *GrantTypePolicyWriteable) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetAction

`func (o *GrantTypePolicyWriteable) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GrantTypePolicyWriteable) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GrantTypePolicyWriteable) SetAction(v GrantAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *GrantTypePolicyWriteable) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *GrantTypePolicyWriteable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantTypePolicyWriteable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantTypePolicyWriteable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GrantTypePolicyWriteable) HasActor() bool`

HasActor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


