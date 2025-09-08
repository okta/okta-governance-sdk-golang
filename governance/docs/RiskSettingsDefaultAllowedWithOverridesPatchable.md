# RiskSettingsDefaultAllowedWithOverridesPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSubmissionType** | **string** |  | 
**ApprovalSequenceId** | **string** | The ID of the approval sequence | 
**AccessDurationSettings** | Pointer to [**NullableAccessDurationSettingsPatchable**](AccessDurationSettingsPatchable.md) |  | [optional] 

## Methods

### NewRiskSettingsDefaultAllowedWithOverridesPatchable

`func NewRiskSettingsDefaultAllowedWithOverridesPatchable(requestSubmissionType string, approvalSequenceId string, ) *RiskSettingsDefaultAllowedWithOverridesPatchable`

NewRiskSettingsDefaultAllowedWithOverridesPatchable instantiates a new RiskSettingsDefaultAllowedWithOverridesPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSettingsDefaultAllowedWithOverridesPatchableWithDefaults

`func NewRiskSettingsDefaultAllowedWithOverridesPatchableWithDefaults() *RiskSettingsDefaultAllowedWithOverridesPatchable`

NewRiskSettingsDefaultAllowedWithOverridesPatchableWithDefaults instantiates a new RiskSettingsDefaultAllowedWithOverridesPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSubmissionType

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetRequestSubmissionType() string`

GetRequestSubmissionType returns the RequestSubmissionType field if non-nil, zero value otherwise.

### GetRequestSubmissionTypeOk

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetRequestSubmissionTypeOk() (*string, bool)`

GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubmissionType

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetRequestSubmissionType(v string)`

SetRequestSubmissionType sets RequestSubmissionType field to given value.


### GetApprovalSequenceId

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.


### GetAccessDurationSettings

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetAccessDurationSettings() AccessDurationSettingsPatchable`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) GetAccessDurationSettingsOk() (*AccessDurationSettingsPatchable, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetAccessDurationSettings(v AccessDurationSettingsPatchable)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### SetAccessDurationSettingsNil

`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) SetAccessDurationSettingsNil(b bool)`

 SetAccessDurationSettingsNil sets the value for AccessDurationSettings to be an explicit nil

### UnsetAccessDurationSettings
`func (o *RiskSettingsDefaultAllowedWithOverridesPatchable) UnsetAccessDurationSettings()`

UnsetAccessDurationSettings ensures that no value is present for AccessDurationSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


