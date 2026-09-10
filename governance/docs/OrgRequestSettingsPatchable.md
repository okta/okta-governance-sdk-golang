# OrgRequestSettingsPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubprocessorsAcknowledged** | Pointer to **bool** | Indicates that Access Requests provisioning was triggered by the customer (such as in [Govern Okta admin roles](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;csh-governance-admin-roles)) | [optional] 
**ResourceCatalogVisibility** | Pointer to [**ResourceCatalogVisibility**](ResourceCatalogVisibility.md) |  | [optional] 
**Integrations** | Pointer to [**RequestIntegrations**](RequestIntegrations.md) |  | [optional] 

## Methods

### NewOrgRequestSettingsPatchable

`func NewOrgRequestSettingsPatchable() *OrgRequestSettingsPatchable`

NewOrgRequestSettingsPatchable instantiates a new OrgRequestSettingsPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgRequestSettingsPatchableWithDefaults

`func NewOrgRequestSettingsPatchableWithDefaults() *OrgRequestSettingsPatchable`

NewOrgRequestSettingsPatchableWithDefaults instantiates a new OrgRequestSettingsPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubprocessorsAcknowledged

`func (o *OrgRequestSettingsPatchable) GetSubprocessorsAcknowledged() bool`

GetSubprocessorsAcknowledged returns the SubprocessorsAcknowledged field if non-nil, zero value otherwise.

### GetSubprocessorsAcknowledgedOk

`func (o *OrgRequestSettingsPatchable) GetSubprocessorsAcknowledgedOk() (*bool, bool)`

GetSubprocessorsAcknowledgedOk returns a tuple with the SubprocessorsAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessorsAcknowledged

`func (o *OrgRequestSettingsPatchable) SetSubprocessorsAcknowledged(v bool)`

SetSubprocessorsAcknowledged sets SubprocessorsAcknowledged field to given value.

### HasSubprocessorsAcknowledged

`func (o *OrgRequestSettingsPatchable) HasSubprocessorsAcknowledged() bool`

HasSubprocessorsAcknowledged returns a boolean if a field has been set.

### GetResourceCatalogVisibility

`func (o *OrgRequestSettingsPatchable) GetResourceCatalogVisibility() ResourceCatalogVisibility`

GetResourceCatalogVisibility returns the ResourceCatalogVisibility field if non-nil, zero value otherwise.

### GetResourceCatalogVisibilityOk

`func (o *OrgRequestSettingsPatchable) GetResourceCatalogVisibilityOk() (*ResourceCatalogVisibility, bool)`

GetResourceCatalogVisibilityOk returns a tuple with the ResourceCatalogVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCatalogVisibility

`func (o *OrgRequestSettingsPatchable) SetResourceCatalogVisibility(v ResourceCatalogVisibility)`

SetResourceCatalogVisibility sets ResourceCatalogVisibility field to given value.

### HasResourceCatalogVisibility

`func (o *OrgRequestSettingsPatchable) HasResourceCatalogVisibility() bool`

HasResourceCatalogVisibility returns a boolean if a field has been set.

### GetIntegrations

`func (o *OrgRequestSettingsPatchable) GetIntegrations() RequestIntegrations`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *OrgRequestSettingsPatchable) GetIntegrationsOk() (*RequestIntegrations, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *OrgRequestSettingsPatchable) SetIntegrations(v RequestIntegrations)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *OrgRequestSettingsPatchable) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


