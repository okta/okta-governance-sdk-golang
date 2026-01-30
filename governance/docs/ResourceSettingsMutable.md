# ResourceSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignResourceType**](CampaignResourceType.md) |  | 
**TargetResources** | Pointer to [**[]TargetResourcesRequestInner**](TargetResourcesRequestInner.md) |  Specific resources that are included in the access certification campaign:  * If &#x60;resourceSettings.targetResources.resourceType&#x60; is &#x60;APPLICATION&#x60; and the app is enabled with entitlement management, you can also review entitlements and entitlement bundles:     * Review all entitlements and bundles by setting &#x60;resourceSettings.targetResources.includeAllEntitlementsAndBundles&#x60; to &#x60;true&#x60; (&#x60;false&#x60; is set by default).     * Restrict the review to non-policy entitlement grants by setting &#x60;resourceSettings.onlyIncludeOutOfPolicyEntitlements&#x60; to &#x60;true&#x60; (&#x60;false&#x60; is set by default).     * If &#x60;resourceSettings.targetResources.includeAllEntitlementsAndBundles&#x60; is &#x60;false&#x60;, then you must specify a list of &#x60;resourceSettings.targetResources.entitlementBundles&#x60; and/or &#x60;resourceSettings.targetResources.entitlements&#x60;. * If &#x60;resourceSettings.type&#x60; is &#x60;OKTA_SERVICE_ACCOUNT&#x60;, then specify &#x60;OKTA_SERVICE_ACCOUNT&#x60; as &#x60;resourceSettings.targetResources.resourceType&#x60;, and &#x60;resourceId&#x60; as the ID of the Okta service account. * If &#x60;resourceSettings.type&#x60; is &#x60;APP_SERVICE_ACCOUNT&#x60;, then specify &#x60;APPLICATION&#x60; as the &#x60;resourceSettings.targetResources.resourceType&#x60;, &#x60;resourceSettings.targetResources.resourceId&#x60; as the ID of the Okta app instance, and add service account IDs into &#x60;resourceSettings.targetResources.appServiceAccounts&#x60;.  | [optional] 
**ExcludedResources** | Pointer to [**[]ResourceSettingsMutableExcludedResourcesInner**](ResourceSettingsMutableExcludedResourcesInner.md) | Only applicable if &#x60;campaignType&#x60; is &#x60;USER&#x60;.  A list of resources that are excluded from the review. | [optional] 
**IndividuallyAssignedAppsOnly** | Pointer to **bool** | Only applicable if &#x60;campaignType&#x60; is &#x60;USER&#x60;.  If &#x60;true&#x60;, only include individually assigned apps. | [optional] 
**IndividuallyAssignedGroupsOnly** | Pointer to **bool** | Only applicable if &#x60;campaignType&#x60; is &#x60;USER&#x60;.  If &#x60;true&#x60;, only include individually assigned groups. | [optional] 
**IncludeEntitlements** | Pointer to **bool** | Only applicable if &#x60;resourceSettings.targetResources.resourceType&#x60; is &#x60;APPLICATION&#x60; and entitlement management is enabled.  If &#x60;true&#x60;, include entitlements for the app. | [optional] [default to false]
**OnlyIncludeOutOfPolicyEntitlements** | Pointer to **bool** | Only applicable if &#x60;campaignType&#x60; is &#x60;USER&#x60;, &#x60;resourceSettings.targetResources.resourceType&#x60; is &#x60;APPLICATION&#x60;, and entitlement management is enabled:  If &#x60;true&#x60;, only include out-of-policy entitlements. | [optional] 
**IncludeAdminRoles** | Pointer to **bool** | Only applicable if &#x60;campaignType&#x60; is &#x60;USER&#x60;.  If &#x60;true&#x60;, include users assigned to admin roles in the campaign. | [optional] 
**IncludeAllOktaServiceAccounts** | Pointer to **bool** | Only applicable when &#x60;resourceSettings.type&#x60; is &#x60;OKTA_SERVICE_ACCOUNT&#x60;: * If &#x60;true&#x60;, all Okta service accounts in the org are included as target resources in the campaign. * If &#x60;false&#x60;, only the Okta service accounts IDs specified in &#x60;resourceSettings.targetResources.resourceId&#x60; are included in the campaign. | [optional] 

## Methods

### NewResourceSettingsMutable

`func NewResourceSettingsMutable(type_ CampaignResourceType, ) *ResourceSettingsMutable`

NewResourceSettingsMutable instantiates a new ResourceSettingsMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSettingsMutableWithDefaults

`func NewResourceSettingsMutableWithDefaults() *ResourceSettingsMutable`

NewResourceSettingsMutableWithDefaults instantiates a new ResourceSettingsMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceSettingsMutable) GetType() CampaignResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceSettingsMutable) GetTypeOk() (*CampaignResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceSettingsMutable) SetType(v CampaignResourceType)`

SetType sets Type field to given value.


### GetTargetResources

`func (o *ResourceSettingsMutable) GetTargetResources() []TargetResourcesRequestInner`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *ResourceSettingsMutable) GetTargetResourcesOk() (*[]TargetResourcesRequestInner, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *ResourceSettingsMutable) SetTargetResources(v []TargetResourcesRequestInner)`

SetTargetResources sets TargetResources field to given value.

### HasTargetResources

`func (o *ResourceSettingsMutable) HasTargetResources() bool`

HasTargetResources returns a boolean if a field has been set.

### GetExcludedResources

`func (o *ResourceSettingsMutable) GetExcludedResources() []ResourceSettingsMutableExcludedResourcesInner`

GetExcludedResources returns the ExcludedResources field if non-nil, zero value otherwise.

### GetExcludedResourcesOk

`func (o *ResourceSettingsMutable) GetExcludedResourcesOk() (*[]ResourceSettingsMutableExcludedResourcesInner, bool)`

GetExcludedResourcesOk returns a tuple with the ExcludedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedResources

`func (o *ResourceSettingsMutable) SetExcludedResources(v []ResourceSettingsMutableExcludedResourcesInner)`

SetExcludedResources sets ExcludedResources field to given value.

### HasExcludedResources

`func (o *ResourceSettingsMutable) HasExcludedResources() bool`

HasExcludedResources returns a boolean if a field has been set.

### GetIndividuallyAssignedAppsOnly

`func (o *ResourceSettingsMutable) GetIndividuallyAssignedAppsOnly() bool`

GetIndividuallyAssignedAppsOnly returns the IndividuallyAssignedAppsOnly field if non-nil, zero value otherwise.

### GetIndividuallyAssignedAppsOnlyOk

`func (o *ResourceSettingsMutable) GetIndividuallyAssignedAppsOnlyOk() (*bool, bool)`

GetIndividuallyAssignedAppsOnlyOk returns a tuple with the IndividuallyAssignedAppsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividuallyAssignedAppsOnly

`func (o *ResourceSettingsMutable) SetIndividuallyAssignedAppsOnly(v bool)`

SetIndividuallyAssignedAppsOnly sets IndividuallyAssignedAppsOnly field to given value.

### HasIndividuallyAssignedAppsOnly

`func (o *ResourceSettingsMutable) HasIndividuallyAssignedAppsOnly() bool`

HasIndividuallyAssignedAppsOnly returns a boolean if a field has been set.

### GetIndividuallyAssignedGroupsOnly

`func (o *ResourceSettingsMutable) GetIndividuallyAssignedGroupsOnly() bool`

GetIndividuallyAssignedGroupsOnly returns the IndividuallyAssignedGroupsOnly field if non-nil, zero value otherwise.

### GetIndividuallyAssignedGroupsOnlyOk

`func (o *ResourceSettingsMutable) GetIndividuallyAssignedGroupsOnlyOk() (*bool, bool)`

GetIndividuallyAssignedGroupsOnlyOk returns a tuple with the IndividuallyAssignedGroupsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividuallyAssignedGroupsOnly

`func (o *ResourceSettingsMutable) SetIndividuallyAssignedGroupsOnly(v bool)`

SetIndividuallyAssignedGroupsOnly sets IndividuallyAssignedGroupsOnly field to given value.

### HasIndividuallyAssignedGroupsOnly

`func (o *ResourceSettingsMutable) HasIndividuallyAssignedGroupsOnly() bool`

HasIndividuallyAssignedGroupsOnly returns a boolean if a field has been set.

### GetIncludeEntitlements

`func (o *ResourceSettingsMutable) GetIncludeEntitlements() bool`

GetIncludeEntitlements returns the IncludeEntitlements field if non-nil, zero value otherwise.

### GetIncludeEntitlementsOk

`func (o *ResourceSettingsMutable) GetIncludeEntitlementsOk() (*bool, bool)`

GetIncludeEntitlementsOk returns a tuple with the IncludeEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEntitlements

`func (o *ResourceSettingsMutable) SetIncludeEntitlements(v bool)`

SetIncludeEntitlements sets IncludeEntitlements field to given value.

### HasIncludeEntitlements

`func (o *ResourceSettingsMutable) HasIncludeEntitlements() bool`

HasIncludeEntitlements returns a boolean if a field has been set.

### GetOnlyIncludeOutOfPolicyEntitlements

`func (o *ResourceSettingsMutable) GetOnlyIncludeOutOfPolicyEntitlements() bool`

GetOnlyIncludeOutOfPolicyEntitlements returns the OnlyIncludeOutOfPolicyEntitlements field if non-nil, zero value otherwise.

### GetOnlyIncludeOutOfPolicyEntitlementsOk

`func (o *ResourceSettingsMutable) GetOnlyIncludeOutOfPolicyEntitlementsOk() (*bool, bool)`

GetOnlyIncludeOutOfPolicyEntitlementsOk returns a tuple with the OnlyIncludeOutOfPolicyEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIncludeOutOfPolicyEntitlements

`func (o *ResourceSettingsMutable) SetOnlyIncludeOutOfPolicyEntitlements(v bool)`

SetOnlyIncludeOutOfPolicyEntitlements sets OnlyIncludeOutOfPolicyEntitlements field to given value.

### HasOnlyIncludeOutOfPolicyEntitlements

`func (o *ResourceSettingsMutable) HasOnlyIncludeOutOfPolicyEntitlements() bool`

HasOnlyIncludeOutOfPolicyEntitlements returns a boolean if a field has been set.

### GetIncludeAdminRoles

`func (o *ResourceSettingsMutable) GetIncludeAdminRoles() bool`

GetIncludeAdminRoles returns the IncludeAdminRoles field if non-nil, zero value otherwise.

### GetIncludeAdminRolesOk

`func (o *ResourceSettingsMutable) GetIncludeAdminRolesOk() (*bool, bool)`

GetIncludeAdminRolesOk returns a tuple with the IncludeAdminRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAdminRoles

`func (o *ResourceSettingsMutable) SetIncludeAdminRoles(v bool)`

SetIncludeAdminRoles sets IncludeAdminRoles field to given value.

### HasIncludeAdminRoles

`func (o *ResourceSettingsMutable) HasIncludeAdminRoles() bool`

HasIncludeAdminRoles returns a boolean if a field has been set.

### GetIncludeAllOktaServiceAccounts

`func (o *ResourceSettingsMutable) GetIncludeAllOktaServiceAccounts() bool`

GetIncludeAllOktaServiceAccounts returns the IncludeAllOktaServiceAccounts field if non-nil, zero value otherwise.

### GetIncludeAllOktaServiceAccountsOk

`func (o *ResourceSettingsMutable) GetIncludeAllOktaServiceAccountsOk() (*bool, bool)`

GetIncludeAllOktaServiceAccountsOk returns a tuple with the IncludeAllOktaServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllOktaServiceAccounts

`func (o *ResourceSettingsMutable) SetIncludeAllOktaServiceAccounts(v bool)`

SetIncludeAllOktaServiceAccounts sets IncludeAllOktaServiceAccounts field to given value.

### HasIncludeAllOktaServiceAccounts

`func (o *ResourceSettingsMutable) HasIncludeAllOktaServiceAccounts() bool`

HasIncludeAllOktaServiceAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


