# OrgSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegates** | Pointer to [**OrgSettingsDelegates**](OrgSettingsDelegates.md) |  | [optional] 
**GovernanceAI** | Pointer to [**OrgSettingsGovernanceAI**](OrgSettingsGovernanceAI.md) |  | [optional] 
**Escalations** | Pointer to [**OrgSettingsEscalations**](OrgSettingsEscalations.md) |  | [optional] 

## Methods

### NewOrgSettings

`func NewOrgSettings() *OrgSettings`

NewOrgSettings instantiates a new OrgSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingsWithDefaults

`func NewOrgSettingsWithDefaults() *OrgSettings`

NewOrgSettingsWithDefaults instantiates a new OrgSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegates

`func (o *OrgSettings) GetDelegates() OrgSettingsDelegates`

GetDelegates returns the Delegates field if non-nil, zero value otherwise.

### GetDelegatesOk

`func (o *OrgSettings) GetDelegatesOk() (*OrgSettingsDelegates, bool)`

GetDelegatesOk returns a tuple with the Delegates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegates

`func (o *OrgSettings) SetDelegates(v OrgSettingsDelegates)`

SetDelegates sets Delegates field to given value.

### HasDelegates

`func (o *OrgSettings) HasDelegates() bool`

HasDelegates returns a boolean if a field has been set.

### GetGovernanceAI

`func (o *OrgSettings) GetGovernanceAI() OrgSettingsGovernanceAI`

GetGovernanceAI returns the GovernanceAI field if non-nil, zero value otherwise.

### GetGovernanceAIOk

`func (o *OrgSettings) GetGovernanceAIOk() (*OrgSettingsGovernanceAI, bool)`

GetGovernanceAIOk returns a tuple with the GovernanceAI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceAI

`func (o *OrgSettings) SetGovernanceAI(v OrgSettingsGovernanceAI)`

SetGovernanceAI sets GovernanceAI field to given value.

### HasGovernanceAI

`func (o *OrgSettings) HasGovernanceAI() bool`

HasGovernanceAI returns a boolean if a field has been set.

### GetEscalations

`func (o *OrgSettings) GetEscalations() OrgSettingsEscalations`

GetEscalations returns the Escalations field if non-nil, zero value otherwise.

### GetEscalationsOk

`func (o *OrgSettings) GetEscalationsOk() (*OrgSettingsEscalations, bool)`

GetEscalationsOk returns a tuple with the Escalations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalations

`func (o *OrgSettings) SetEscalations(v OrgSettingsEscalations)`

SetEscalations sets Escalations field to given value.

### HasEscalations

`func (o *OrgSettings) HasEscalations() bool`

HasEscalations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


