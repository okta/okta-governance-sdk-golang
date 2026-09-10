# OrgRequestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubprocessorsAcknowledged** | **bool** | Indicates that Access Requests provisioning was triggered by the customer (such as in [Govern Okta admin roles](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;csh-governance-admin-roles)) | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) |  | 
**RequestExperiences** | [**[]RequestExperience**](RequestExperience.md) | The supported request experiences in the org | 
**LongTimePastProvisioned** | **bool** | Indicates that Access Request provisioning was triggered more than 10 minutes ago | 
**ResourceCatalogVisibility** | Pointer to [**ResourceCatalogVisibility**](ResourceCatalogVisibility.md) |  | [optional] 
**Integrations** | Pointer to [**RequestIntegrations**](RequestIntegrations.md) |  | [optional] 

## Methods

### NewOrgRequestSettings

`func NewOrgRequestSettings(subprocessorsAcknowledged bool, provisioningStatus ProvisioningStatus, requestExperiences []RequestExperience, longTimePastProvisioned bool, ) *OrgRequestSettings`

NewOrgRequestSettings instantiates a new OrgRequestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgRequestSettingsWithDefaults

`func NewOrgRequestSettingsWithDefaults() *OrgRequestSettings`

NewOrgRequestSettingsWithDefaults instantiates a new OrgRequestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubprocessorsAcknowledged

`func (o *OrgRequestSettings) GetSubprocessorsAcknowledged() bool`

GetSubprocessorsAcknowledged returns the SubprocessorsAcknowledged field if non-nil, zero value otherwise.

### GetSubprocessorsAcknowledgedOk

`func (o *OrgRequestSettings) GetSubprocessorsAcknowledgedOk() (*bool, bool)`

GetSubprocessorsAcknowledgedOk returns a tuple with the SubprocessorsAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessorsAcknowledged

`func (o *OrgRequestSettings) SetSubprocessorsAcknowledged(v bool)`

SetSubprocessorsAcknowledged sets SubprocessorsAcknowledged field to given value.


### GetProvisioningStatus

`func (o *OrgRequestSettings) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *OrgRequestSettings) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *OrgRequestSettings) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetRequestExperiences

`func (o *OrgRequestSettings) GetRequestExperiences() []RequestExperience`

GetRequestExperiences returns the RequestExperiences field if non-nil, zero value otherwise.

### GetRequestExperiencesOk

`func (o *OrgRequestSettings) GetRequestExperiencesOk() (*[]RequestExperience, bool)`

GetRequestExperiencesOk returns a tuple with the RequestExperiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestExperiences

`func (o *OrgRequestSettings) SetRequestExperiences(v []RequestExperience)`

SetRequestExperiences sets RequestExperiences field to given value.


### GetLongTimePastProvisioned

`func (o *OrgRequestSettings) GetLongTimePastProvisioned() bool`

GetLongTimePastProvisioned returns the LongTimePastProvisioned field if non-nil, zero value otherwise.

### GetLongTimePastProvisionedOk

`func (o *OrgRequestSettings) GetLongTimePastProvisionedOk() (*bool, bool)`

GetLongTimePastProvisionedOk returns a tuple with the LongTimePastProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTimePastProvisioned

`func (o *OrgRequestSettings) SetLongTimePastProvisioned(v bool)`

SetLongTimePastProvisioned sets LongTimePastProvisioned field to given value.


### GetResourceCatalogVisibility

`func (o *OrgRequestSettings) GetResourceCatalogVisibility() ResourceCatalogVisibility`

GetResourceCatalogVisibility returns the ResourceCatalogVisibility field if non-nil, zero value otherwise.

### GetResourceCatalogVisibilityOk

`func (o *OrgRequestSettings) GetResourceCatalogVisibilityOk() (*ResourceCatalogVisibility, bool)`

GetResourceCatalogVisibilityOk returns a tuple with the ResourceCatalogVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCatalogVisibility

`func (o *OrgRequestSettings) SetResourceCatalogVisibility(v ResourceCatalogVisibility)`

SetResourceCatalogVisibility sets ResourceCatalogVisibility field to given value.

### HasResourceCatalogVisibility

`func (o *OrgRequestSettings) HasResourceCatalogVisibility() bool`

HasResourceCatalogVisibility returns a boolean if a field has been set.

### GetIntegrations

`func (o *OrgRequestSettings) GetIntegrations() RequestIntegrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *OrgRequestSettings) GetIntegrationsOk() (*RequestIntegrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *OrgRequestSettings) SetIntegrations(v RequestIntegrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *OrgRequestSettings) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


