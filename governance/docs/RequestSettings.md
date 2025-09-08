# RequestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidAccessScopeSettings** | [**[]ValidAccessDetail**](ValidAccessDetail.md) | Access scope settings that are eligible to be added to a request condition for the specified resource | 
**ValidRequesterSettings** | [**[]ValidRequesterSetting**](ValidRequesterSetting.md) | Request scope settings that are eligible to be added to a request condition for the specified resource | 
**ValidAccessDurationSettings** | [**ValidAccessDurationSettingsDetails**](ValidAccessDurationSettingsDetails.md) |  | 
**RequestOnBehalfOfSettings** | Pointer to [**RequestOnBehalfOfSettingsDetails**](RequestOnBehalfOfSettingsDetails.md) |  | [optional] 
**ValidRiskSettings** | Pointer to [**ValidRiskSettingsDetails**](ValidRiskSettingsDetails.md) |  | [optional] 
**RiskSettings** | Pointer to [**RiskSettingsDetails**](RiskSettingsDetails.md) |  | [optional] 

## Methods

### NewRequestSettings

`func NewRequestSettings(validAccessScopeSettings []ValidAccessDetail, validRequesterSettings []ValidRequesterSetting, validAccessDurationSettings ValidAccessDurationSettingsDetails, ) *RequestSettings`

NewRequestSettings instantiates a new RequestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSettingsWithDefaults

`func NewRequestSettingsWithDefaults() *RequestSettings`

NewRequestSettingsWithDefaults instantiates a new RequestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidAccessScopeSettings

`func (o *RequestSettings) GetValidAccessScopeSettings() []ValidAccessDetail`

GetValidAccessScopeSettings returns the ValidAccessScopeSettings field if non-nil, zero value otherwise.

### GetValidAccessScopeSettingsOk

`func (o *RequestSettings) GetValidAccessScopeSettingsOk() (*[]ValidAccessDetail, bool)`

GetValidAccessScopeSettingsOk returns a tuple with the ValidAccessScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAccessScopeSettings

`func (o *RequestSettings) SetValidAccessScopeSettings(v []ValidAccessDetail)`

SetValidAccessScopeSettings sets ValidAccessScopeSettings field to given value.


### GetValidRequesterSettings

`func (o *RequestSettings) GetValidRequesterSettings() []ValidRequesterSetting`

GetValidRequesterSettings returns the ValidRequesterSettings field if non-nil, zero value otherwise.

### GetValidRequesterSettingsOk

`func (o *RequestSettings) GetValidRequesterSettingsOk() (*[]ValidRequesterSetting, bool)`

GetValidRequesterSettingsOk returns a tuple with the ValidRequesterSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRequesterSettings

`func (o *RequestSettings) SetValidRequesterSettings(v []ValidRequesterSetting)`

SetValidRequesterSettings sets ValidRequesterSettings field to given value.


### GetValidAccessDurationSettings

`func (o *RequestSettings) GetValidAccessDurationSettings() ValidAccessDurationSettingsDetails`

GetValidAccessDurationSettings returns the ValidAccessDurationSettings field if non-nil, zero value otherwise.

### GetValidAccessDurationSettingsOk

`func (o *RequestSettings) GetValidAccessDurationSettingsOk() (*ValidAccessDurationSettingsDetails, bool)`

GetValidAccessDurationSettingsOk returns a tuple with the ValidAccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAccessDurationSettings

`func (o *RequestSettings) SetValidAccessDurationSettings(v ValidAccessDurationSettingsDetails)`

SetValidAccessDurationSettings sets ValidAccessDurationSettings field to given value.


### GetRequestOnBehalfOfSettings

`func (o *RequestSettings) GetRequestOnBehalfOfSettings() RequestOnBehalfOfSettingsDetails`

GetRequestOnBehalfOfSettings returns the RequestOnBehalfOfSettings field if non-nil, zero value otherwise.

### GetRequestOnBehalfOfSettingsOk

`func (o *RequestSettings) GetRequestOnBehalfOfSettingsOk() (*RequestOnBehalfOfSettingsDetails, bool)`

GetRequestOnBehalfOfSettingsOk returns a tuple with the RequestOnBehalfOfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOnBehalfOfSettings

`func (o *RequestSettings) SetRequestOnBehalfOfSettings(v RequestOnBehalfOfSettingsDetails)`

SetRequestOnBehalfOfSettings sets RequestOnBehalfOfSettings field to given value.

### HasRequestOnBehalfOfSettings

`func (o *RequestSettings) HasRequestOnBehalfOfSettings() bool`

HasRequestOnBehalfOfSettings returns a boolean if a field has been set.

### GetValidRiskSettings

`func (o *RequestSettings) GetValidRiskSettings() ValidRiskSettingsDetails`

GetValidRiskSettings returns the ValidRiskSettings field if non-nil, zero value otherwise.

### GetValidRiskSettingsOk

`func (o *RequestSettings) GetValidRiskSettingsOk() (*ValidRiskSettingsDetails, bool)`

GetValidRiskSettingsOk returns a tuple with the ValidRiskSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRiskSettings

`func (o *RequestSettings) SetValidRiskSettings(v ValidRiskSettingsDetails)`

SetValidRiskSettings sets ValidRiskSettings field to given value.

### HasValidRiskSettings

`func (o *RequestSettings) HasValidRiskSettings() bool`

HasValidRiskSettings returns a boolean if a field has been set.

### GetRiskSettings

`func (o *RequestSettings) GetRiskSettings() RiskSettingsDetails`

GetRiskSettings returns the RiskSettings field if non-nil, zero value otherwise.

### GetRiskSettingsOk

`func (o *RequestSettings) GetRiskSettingsOk() (*RiskSettingsDetails, bool)`

GetRiskSettingsOk returns a tuple with the RiskSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSettings

`func (o *RequestSettings) SetRiskSettings(v RiskSettingsDetails)`

SetRiskSettings sets RiskSettings field to given value.

### HasRiskSettings

`func (o *RequestSettings) HasRiskSettings() bool`

HasRiskSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


