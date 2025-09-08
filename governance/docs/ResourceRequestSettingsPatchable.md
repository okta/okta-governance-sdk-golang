# ResourceRequestSettingsPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestOnBehalfOfSettings** | Pointer to [**NullableRequestOnBehalfOfSettingsPatchable**](RequestOnBehalfOfSettingsPatchable.md) |  | [optional] 
**RiskSettings** | Pointer to [**RiskSettingsPatchable**](RiskSettingsPatchable.md) |  | [optional] 

## Methods

### NewResourceRequestSettingsPatchable

`func NewResourceRequestSettingsPatchable() *ResourceRequestSettingsPatchable`

NewResourceRequestSettingsPatchable instantiates a new ResourceRequestSettingsPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequestSettingsPatchableWithDefaults

`func NewResourceRequestSettingsPatchableWithDefaults() *ResourceRequestSettingsPatchable`

NewResourceRequestSettingsPatchableWithDefaults instantiates a new ResourceRequestSettingsPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestOnBehalfOfSettings

`func (o *ResourceRequestSettingsPatchable) GetRequestOnBehalfOfSettings() RequestOnBehalfOfSettingsPatchable`

GetRequestOnBehalfOfSettings returns the RequestOnBehalfOfSettings field if non-nil, zero value otherwise.

### GetRequestOnBehalfOfSettingsOk

`func (o *ResourceRequestSettingsPatchable) GetRequestOnBehalfOfSettingsOk() (*RequestOnBehalfOfSettingsPatchable, bool)`

GetRequestOnBehalfOfSettingsOk returns a tuple with the RequestOnBehalfOfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOnBehalfOfSettings

`func (o *ResourceRequestSettingsPatchable) SetRequestOnBehalfOfSettings(v RequestOnBehalfOfSettingsPatchable)`

SetRequestOnBehalfOfSettings sets RequestOnBehalfOfSettings field to given value.

### HasRequestOnBehalfOfSettings

`func (o *ResourceRequestSettingsPatchable) HasRequestOnBehalfOfSettings() bool`

HasRequestOnBehalfOfSettings returns a boolean if a field has been set.

### SetRequestOnBehalfOfSettingsNil

`func (o *ResourceRequestSettingsPatchable) SetRequestOnBehalfOfSettingsNil(b bool)`

 SetRequestOnBehalfOfSettingsNil sets the value for RequestOnBehalfOfSettings to be an explicit nil

### UnsetRequestOnBehalfOfSettings
`func (o *ResourceRequestSettingsPatchable) UnsetRequestOnBehalfOfSettings()`

UnsetRequestOnBehalfOfSettings ensures that no value is present for RequestOnBehalfOfSettings, not even an explicit nil
### GetRiskSettings

`func (o *ResourceRequestSettingsPatchable) GetRiskSettings() RiskSettingsPatchable`

GetRiskSettings returns the RiskSettings field if non-nil, zero value otherwise.

### GetRiskSettingsOk

`func (o *ResourceRequestSettingsPatchable) GetRiskSettingsOk() (*RiskSettingsPatchable, bool)`

GetRiskSettingsOk returns a tuple with the RiskSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSettings

`func (o *ResourceRequestSettingsPatchable) SetRiskSettings(v RiskSettingsPatchable)`

SetRiskSettings sets RiskSettings field to given value.

### HasRiskSettings

`func (o *ResourceRequestSettingsPatchable) HasRiskSettings() bool`

HasRiskSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


