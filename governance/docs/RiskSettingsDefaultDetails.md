# RiskSettingsDefaultDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSubmissionType** | **string** |  | 
**Error** | Pointer to [**[]RiskSettingsError**](RiskSettingsError.md) |  | [optional] 
**ApprovalSequenceId** | **string** | The ID of the approval sequence | 
**AccessDurationSettings** | Pointer to [**AccessDurationSettingsFull**](AccessDurationSettingsFull.md) |  | [optional] 

## Methods

### NewRiskSettingsDefaultDetails

`func NewRiskSettingsDefaultDetails(requestSubmissionType string, approvalSequenceId string, ) *RiskSettingsDefaultDetails`

NewRiskSettingsDefaultDetails instantiates a new RiskSettingsDefaultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSettingsDefaultDetailsWithDefaults

`func NewRiskSettingsDefaultDetailsWithDefaults() *RiskSettingsDefaultDetails`

NewRiskSettingsDefaultDetailsWithDefaults instantiates a new RiskSettingsDefaultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSubmissionType

`func (o *RiskSettingsDefaultDetails) GetRequestSubmissionType() string`

GetRequestSubmissionType returns the RequestSubmissionType field if non-nil, zero value otherwise.

### GetRequestSubmissionTypeOk

`func (o *RiskSettingsDefaultDetails) GetRequestSubmissionTypeOk() (*string, bool)`

GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubmissionType

`func (o *RiskSettingsDefaultDetails) SetRequestSubmissionType(v string)`

SetRequestSubmissionType sets RequestSubmissionType field to given value.


### GetError

`func (o *RiskSettingsDefaultDetails) GetError() []RiskSettingsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RiskSettingsDefaultDetails) GetErrorOk() (*[]RiskSettingsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RiskSettingsDefaultDetails) SetError(v []RiskSettingsError)`

SetError sets Error field to given value.

### HasError

`func (o *RiskSettingsDefaultDetails) HasError() bool`

HasError returns a boolean if a field has been set.

### GetApprovalSequenceId

`func (o *RiskSettingsDefaultDetails) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RiskSettingsDefaultDetails) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RiskSettingsDefaultDetails) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.


### GetAccessDurationSettings

`func (o *RiskSettingsDefaultDetails) GetAccessDurationSettings() AccessDurationSettingsFull`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RiskSettingsDefaultDetails) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RiskSettingsDefaultDetails) SetAccessDurationSettings(v AccessDurationSettingsFull)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RiskSettingsDefaultDetails) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


