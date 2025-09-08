# RiskSettingsDefaultPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSubmissionType** | **string** |  | 
**ApprovalSequenceId** | **string** | The ID of the approval sequence | 
**AccessDurationSettings** | Pointer to [**NullableAccessDurationSettingsPatchable**](AccessDurationSettingsPatchable.md) |  | [optional] 

## Methods

### NewRiskSettingsDefaultPatchable

`func NewRiskSettingsDefaultPatchable(requestSubmissionType string, approvalSequenceId string, ) *RiskSettingsDefaultPatchable`

NewRiskSettingsDefaultPatchable instantiates a new RiskSettingsDefaultPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSettingsDefaultPatchableWithDefaults

`func NewRiskSettingsDefaultPatchableWithDefaults() *RiskSettingsDefaultPatchable`

NewRiskSettingsDefaultPatchableWithDefaults instantiates a new RiskSettingsDefaultPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSubmissionType

`func (o *RiskSettingsDefaultPatchable) GetRequestSubmissionType() string`

GetRequestSubmissionType returns the RequestSubmissionType field if non-nil, zero value otherwise.

### GetRequestSubmissionTypeOk

`func (o *RiskSettingsDefaultPatchable) GetRequestSubmissionTypeOk() (*string, bool)`

GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubmissionType

`func (o *RiskSettingsDefaultPatchable) SetRequestSubmissionType(v string)`

SetRequestSubmissionType sets RequestSubmissionType field to given value.


### GetApprovalSequenceId

`func (o *RiskSettingsDefaultPatchable) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RiskSettingsDefaultPatchable) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RiskSettingsDefaultPatchable) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.


### GetAccessDurationSettings

`func (o *RiskSettingsDefaultPatchable) GetAccessDurationSettings() AccessDurationSettingsPatchable`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RiskSettingsDefaultPatchable) GetAccessDurationSettingsOk() (*AccessDurationSettingsPatchable, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RiskSettingsDefaultPatchable) SetAccessDurationSettings(v AccessDurationSettingsPatchable)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RiskSettingsDefaultPatchable) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### SetAccessDurationSettingsNil

`func (o *RiskSettingsDefaultPatchable) SetAccessDurationSettingsNil(b bool)`

 SetAccessDurationSettingsNil sets the value for AccessDurationSettings to be an explicit nil

### UnsetAccessDurationSettings
`func (o *RiskSettingsDefaultPatchable) UnsetAccessDurationSettings()`

UnsetAccessDurationSettings ensures that no value is present for AccessDurationSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


