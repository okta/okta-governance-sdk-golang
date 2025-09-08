# RiskSettingsDefaultRestrictedDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSubmissionType** | **string** |  | 
**AccessDurationSettings** | Pointer to [**AccessDurationSettingsFull**](AccessDurationSettingsFull.md) |  | [optional] 
**Error** | Pointer to [**[]RiskSettingsError**](RiskSettingsError.md) |  | [optional] 

## Methods

### NewRiskSettingsDefaultRestrictedDetails

`func NewRiskSettingsDefaultRestrictedDetails(requestSubmissionType string, ) *RiskSettingsDefaultRestrictedDetails`

NewRiskSettingsDefaultRestrictedDetails instantiates a new RiskSettingsDefaultRestrictedDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSettingsDefaultRestrictedDetailsWithDefaults

`func NewRiskSettingsDefaultRestrictedDetailsWithDefaults() *RiskSettingsDefaultRestrictedDetails`

NewRiskSettingsDefaultRestrictedDetailsWithDefaults instantiates a new RiskSettingsDefaultRestrictedDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSubmissionType

`func (o *RiskSettingsDefaultRestrictedDetails) GetRequestSubmissionType() string`

GetRequestSubmissionType returns the RequestSubmissionType field if non-nil, zero value otherwise.

### GetRequestSubmissionTypeOk

`func (o *RiskSettingsDefaultRestrictedDetails) GetRequestSubmissionTypeOk() (*string, bool)`

GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubmissionType

`func (o *RiskSettingsDefaultRestrictedDetails) SetRequestSubmissionType(v string)`

SetRequestSubmissionType sets RequestSubmissionType field to given value.


### GetAccessDurationSettings

`func (o *RiskSettingsDefaultRestrictedDetails) GetAccessDurationSettings() AccessDurationSettingsFull`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RiskSettingsDefaultRestrictedDetails) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RiskSettingsDefaultRestrictedDetails) SetAccessDurationSettings(v AccessDurationSettingsFull)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RiskSettingsDefaultRestrictedDetails) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### GetError

`func (o *RiskSettingsDefaultRestrictedDetails) GetError() []RiskSettingsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RiskSettingsDefaultRestrictedDetails) GetErrorOk() (*[]RiskSettingsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RiskSettingsDefaultRestrictedDetails) SetError(v []RiskSettingsError)`

SetError sets Error field to given value.

### HasError

`func (o *RiskSettingsDefaultRestrictedDetails) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


