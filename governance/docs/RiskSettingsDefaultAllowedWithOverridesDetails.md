# RiskSettingsDefaultAllowedWithOverridesDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSubmissionType** | **string** |  | 
**ApprovalSequenceId** | **string** | The ID of the approval sequence | 
**AccessDurationSettings** | Pointer to [**AccessDurationSettingsFull**](AccessDurationSettingsFull.md) |  | [optional] 
**Error** | Pointer to [**[]RiskSettingsError**](RiskSettingsError.md) |  | [optional] 

## Methods

### NewRiskSettingsDefaultAllowedWithOverridesDetails

`func NewRiskSettingsDefaultAllowedWithOverridesDetails(requestSubmissionType string, approvalSequenceId string, ) *RiskSettingsDefaultAllowedWithOverridesDetails`

NewRiskSettingsDefaultAllowedWithOverridesDetails instantiates a new RiskSettingsDefaultAllowedWithOverridesDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSettingsDefaultAllowedWithOverridesDetailsWithDefaults

`func NewRiskSettingsDefaultAllowedWithOverridesDetailsWithDefaults() *RiskSettingsDefaultAllowedWithOverridesDetails`

NewRiskSettingsDefaultAllowedWithOverridesDetailsWithDefaults instantiates a new RiskSettingsDefaultAllowedWithOverridesDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSubmissionType

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetRequestSubmissionType() string`

GetRequestSubmissionType returns the RequestSubmissionType field if non-nil, zero value otherwise.

### GetRequestSubmissionTypeOk

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetRequestSubmissionTypeOk() (*string, bool)`

GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubmissionType

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetRequestSubmissionType(v string)`

SetRequestSubmissionType sets RequestSubmissionType field to given value.


### GetApprovalSequenceId

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.


### GetAccessDurationSettings

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetAccessDurationSettings() AccessDurationSettingsFull`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetAccessDurationSettings(v AccessDurationSettingsFull)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### GetError

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetError() []RiskSettingsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) GetErrorOk() (*[]RiskSettingsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) SetError(v []RiskSettingsError)`

SetError sets Error field to given value.

### HasError

`func (o *RiskSettingsDefaultAllowedWithOverridesDetails) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


