# GrantTypeBundleWriteable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | Additive grant type for entitlement bundle. | 
**EntitlementBundleId** | **string** | The entitlement bundle &#x60;id&#x60;  | 
**TargetPrincipal** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Action** | Pointer to [**GrantAction**](GrantAction.md) |  | [optional] [default to GRANTACTION_ALLOW]
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]

## Methods

### NewGrantTypeBundleWriteable

`func NewGrantTypeBundleWriteable(grantType string, entitlementBundleId string, targetPrincipal TargetPrincipal, ) *GrantTypeBundleWriteable`

NewGrantTypeBundleWriteable instantiates a new GrantTypeBundleWriteable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantTypeBundleWriteableWithDefaults

`func NewGrantTypeBundleWriteableWithDefaults() *GrantTypeBundleWriteable`

NewGrantTypeBundleWriteableWithDefaults instantiates a new GrantTypeBundleWriteable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GrantTypeBundleWriteable) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrantTypeBundleWriteable) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrantTypeBundleWriteable) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetEntitlementBundleId

`func (o *GrantTypeBundleWriteable) GetEntitlementBundleId() string`

GetEntitlementBundleId returns the EntitlementBundleId field if non-nil, zero value otherwise.

### GetEntitlementBundleIdOk

`func (o *GrantTypeBundleWriteable) GetEntitlementBundleIdOk() (*string, bool)`

GetEntitlementBundleIdOk returns a tuple with the EntitlementBundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundleId

`func (o *GrantTypeBundleWriteable) SetEntitlementBundleId(v string)`

SetEntitlementBundleId sets EntitlementBundleId field to given value.


### GetTargetPrincipal

`func (o *GrantTypeBundleWriteable) GetTargetPrincipal() TargetPrincipal`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *GrantTypeBundleWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *GrantTypeBundleWriteable) SetTargetPrincipal(v TargetPrincipal)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetScheduleSettings

`func (o *GrantTypeBundleWriteable) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GrantTypeBundleWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GrantTypeBundleWriteable) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *GrantTypeBundleWriteable) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetAction

`func (o *GrantTypeBundleWriteable) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GrantTypeBundleWriteable) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GrantTypeBundleWriteable) SetAction(v GrantAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *GrantTypeBundleWriteable) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *GrantTypeBundleWriteable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantTypeBundleWriteable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantTypeBundleWriteable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GrantTypeBundleWriteable) HasActor() bool`

HasActor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


