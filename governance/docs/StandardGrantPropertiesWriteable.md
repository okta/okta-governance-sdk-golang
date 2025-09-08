# StandardGrantPropertiesWriteable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetPrincipal** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Action** | Pointer to [**GrantAction**](GrantAction.md) |  | [optional] [default to GRANTACTION_ALLOW]
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]

## Methods

### NewStandardGrantPropertiesWriteable

`func NewStandardGrantPropertiesWriteable(targetPrincipal TargetPrincipal, ) *StandardGrantPropertiesWriteable`

NewStandardGrantPropertiesWriteable instantiates a new StandardGrantPropertiesWriteable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardGrantPropertiesWriteableWithDefaults

`func NewStandardGrantPropertiesWriteableWithDefaults() *StandardGrantPropertiesWriteable`

NewStandardGrantPropertiesWriteableWithDefaults instantiates a new StandardGrantPropertiesWriteable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetPrincipal

`func (o *StandardGrantPropertiesWriteable) GetTargetPrincipal() TargetPrincipal`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *StandardGrantPropertiesWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *StandardGrantPropertiesWriteable) SetTargetPrincipal(v TargetPrincipal)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetScheduleSettings

`func (o *StandardGrantPropertiesWriteable) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *StandardGrantPropertiesWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *StandardGrantPropertiesWriteable) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *StandardGrantPropertiesWriteable) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetAction

`func (o *StandardGrantPropertiesWriteable) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StandardGrantPropertiesWriteable) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StandardGrantPropertiesWriteable) SetAction(v GrantAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *StandardGrantPropertiesWriteable) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *StandardGrantPropertiesWriteable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *StandardGrantPropertiesWriteable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *StandardGrantPropertiesWriteable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *StandardGrantPropertiesWriteable) HasActor() bool`

HasActor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


