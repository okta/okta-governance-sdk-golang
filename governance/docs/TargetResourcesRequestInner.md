# TargetResourcesRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The resource ID to review. The &#x60;resourceId&#x60; value depends on the &#x60;resourceType&#x60; object. | 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**IncludeAllEntitlementsAndBundles** | Pointer to **bool** | Only applicable if the &#x60;resourceType&#x60; is &#x60;APPLICATION&#x60; and entitlement management is enabled.  If &#x60;true&#x60;, include all entitlements and entitlement bundles for the app. | [optional] 
**EntitlementBundles** | Pointer to [**[]EntitlementBundlesInner**](EntitlementBundlesInner.md) | Only applicable if the &#x60;resourceType&#x60; is &#x60;APPLICATION&#x60; and entitlement management is enabled.  A list of entitlement bundles (associated with the app specified in &#x60;resourceId&#x60;) selected as targets for review. | [optional] 
**Entitlements** | Pointer to [**[]EntitlementsInner**](EntitlementsInner.md) | Only applicable if &#x60;resourceType&#x60; is &#x60;APPLICATION&#x60; and entitlement management is enabled. A list of entitlements (associated with the app specified in &#x60;resourceId&#x60;) selected as targets for review. | [optional] 
**EntitlementValueScopeByGovernanceLabelValues** | Pointer to [**[]GovernanceLabelValuesInner**](GovernanceLabelValuesInner.md) | &lt;x-lifecycle class&#x3D;\&quot;beta\&quot;&gt;&lt;/x-lifecycle&gt;&lt;br&gt; Only applicable when &#x60;targetResources.resourceType&#x60; is &#x60;APPLICATION&#x60; or &#x60;GOVERNANCE_LABEL_VALUE&#x60;.  A list of governance label values assigned to resources to target for the review. | [optional] 
**IncludeBundlesHavingSameEntitlementValues** | Pointer to **bool** | &lt;x-lifecycle class&#x3D;\&quot;beta\&quot;&gt;&lt;/x-lifecycle&gt;&lt;br&gt; If &#x60;true&#x60;, entitlements bundles are also included as additional targets to review. Only entitlement bundles with values that are targeted through the &#x60;entitlementValueScopeByGovernanceLabelValues&#x60; list are included.  | [optional] 
**EntitlementBundleScopeByGovernanceLabelValues** | Pointer to [**[]GovernanceLabelValuesInner**](GovernanceLabelValuesInner.md) | &lt;x-lifecycle class&#x3D;\&quot;beta\&quot;&gt;&lt;/x-lifecycle&gt;&lt;br&gt; Only applicable when &#x60;targetResources.resourceType&#x60; is &#x60;APPLICATION&#x60; or &#x60;GOVERNANCE_LABEL_VALUE&#x60;.  A list of governance label values assigned to resources to target for the review. | [optional] 
**AppServiceAccounts** | Pointer to [**[]AppServiceAccountsInner**](AppServiceAccountsInner.md) | Only applicable if the &#x60;resourceType&#x60; is &#x60;APPLICATION&#x60;.  A list of app service accounts (associated with the app specified in &#x60;resourceId&#x60;) selected as targets for review. | [optional] 
**IncludeAllAppServiceAccounts** | Pointer to **bool** | Only applicable if &#x60;resourceSettings.type&#x60; is &#x60;APP_SERVICE_ACCOUNT&#x60; and &#x60;resourceSettings.targetResources.resourceType&#x60; is &#x60;APPLICATION&#x60;:  * If true, includes all SaaS app service accounts associated with the app specified in &#x60;resourceSettings.targetResources.resourceId&#x60;. * If false, includes only the SaaS app service accounts specified in the list of &#x60;resourceSettings.targetResources.appServiceAccounts&#x60;. The list of service account IDs in &#x60;resourceSettings.targetResources.appServiceAccounts&#x60; must be associated with the app in &#x60;resourceSettings.targetResources.resourceId&#x60;. | [optional] 

## Methods

### NewTargetResourcesRequestInner

`func NewTargetResourcesRequestInner(resourceId string, ) *TargetResourcesRequestInner`

NewTargetResourcesRequestInner instantiates a new TargetResourcesRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetResourcesRequestInnerWithDefaults

`func NewTargetResourcesRequestInnerWithDefaults() *TargetResourcesRequestInner`

NewTargetResourcesRequestInnerWithDefaults instantiates a new TargetResourcesRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *TargetResourcesRequestInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TargetResourcesRequestInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TargetResourcesRequestInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *TargetResourcesRequestInner) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TargetResourcesRequestInner) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TargetResourcesRequestInner) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TargetResourcesRequestInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetIncludeAllEntitlementsAndBundles

`func (o *TargetResourcesRequestInner) GetIncludeAllEntitlementsAndBundles() bool`

GetIncludeAllEntitlementsAndBundles returns the IncludeAllEntitlementsAndBundles field if non-nil, zero value otherwise.

### GetIncludeAllEntitlementsAndBundlesOk

`func (o *TargetResourcesRequestInner) GetIncludeAllEntitlementsAndBundlesOk() (*bool, bool)`

GetIncludeAllEntitlementsAndBundlesOk returns a tuple with the IncludeAllEntitlementsAndBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllEntitlementsAndBundles

`func (o *TargetResourcesRequestInner) SetIncludeAllEntitlementsAndBundles(v bool)`

SetIncludeAllEntitlementsAndBundles sets IncludeAllEntitlementsAndBundles field to given value.

### HasIncludeAllEntitlementsAndBundles

`func (o *TargetResourcesRequestInner) HasIncludeAllEntitlementsAndBundles() bool`

HasIncludeAllEntitlementsAndBundles returns a boolean if a field has been set.

### GetEntitlementBundles

`func (o *TargetResourcesRequestInner) GetEntitlementBundles() []EntitlementBundlesInner`

GetEntitlementBundles returns the EntitlementBundles field if non-nil, zero value otherwise.

### GetEntitlementBundlesOk

`func (o *TargetResourcesRequestInner) GetEntitlementBundlesOk() (*[]EntitlementBundlesInner, bool)`

GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundles

`func (o *TargetResourcesRequestInner) SetEntitlementBundles(v []EntitlementBundlesInner)`

SetEntitlementBundles sets EntitlementBundles field to given value.

### HasEntitlementBundles

`func (o *TargetResourcesRequestInner) HasEntitlementBundles() bool`

HasEntitlementBundles returns a boolean if a field has been set.

### GetEntitlements

`func (o *TargetResourcesRequestInner) GetEntitlements() []EntitlementsInner`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *TargetResourcesRequestInner) GetEntitlementsOk() (*[]EntitlementsInner, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *TargetResourcesRequestInner) SetEntitlements(v []EntitlementsInner)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *TargetResourcesRequestInner) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEntitlementValueScopeByGovernanceLabelValues

`func (o *TargetResourcesRequestInner) GetEntitlementValueScopeByGovernanceLabelValues() []GovernanceLabelValuesInner`

GetEntitlementValueScopeByGovernanceLabelValues returns the EntitlementValueScopeByGovernanceLabelValues field if non-nil, zero value otherwise.

### GetEntitlementValueScopeByGovernanceLabelValuesOk

`func (o *TargetResourcesRequestInner) GetEntitlementValueScopeByGovernanceLabelValuesOk() (*[]GovernanceLabelValuesInner, bool)`

GetEntitlementValueScopeByGovernanceLabelValuesOk returns a tuple with the EntitlementValueScopeByGovernanceLabelValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValueScopeByGovernanceLabelValues

`func (o *TargetResourcesRequestInner) SetEntitlementValueScopeByGovernanceLabelValues(v []GovernanceLabelValuesInner)`

SetEntitlementValueScopeByGovernanceLabelValues sets EntitlementValueScopeByGovernanceLabelValues field to given value.

### HasEntitlementValueScopeByGovernanceLabelValues

`func (o *TargetResourcesRequestInner) HasEntitlementValueScopeByGovernanceLabelValues() bool`

HasEntitlementValueScopeByGovernanceLabelValues returns a boolean if a field has been set.

### GetIncludeBundlesHavingSameEntitlementValues

`func (o *TargetResourcesRequestInner) GetIncludeBundlesHavingSameEntitlementValues() bool`

GetIncludeBundlesHavingSameEntitlementValues returns the IncludeBundlesHavingSameEntitlementValues field if non-nil, zero value otherwise.

### GetIncludeBundlesHavingSameEntitlementValuesOk

`func (o *TargetResourcesRequestInner) GetIncludeBundlesHavingSameEntitlementValuesOk() (*bool, bool)`

GetIncludeBundlesHavingSameEntitlementValuesOk returns a tuple with the IncludeBundlesHavingSameEntitlementValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBundlesHavingSameEntitlementValues

`func (o *TargetResourcesRequestInner) SetIncludeBundlesHavingSameEntitlementValues(v bool)`

SetIncludeBundlesHavingSameEntitlementValues sets IncludeBundlesHavingSameEntitlementValues field to given value.

### HasIncludeBundlesHavingSameEntitlementValues

`func (o *TargetResourcesRequestInner) HasIncludeBundlesHavingSameEntitlementValues() bool`

HasIncludeBundlesHavingSameEntitlementValues returns a boolean if a field has been set.

### GetEntitlementBundleScopeByGovernanceLabelValues

`func (o *TargetResourcesRequestInner) GetEntitlementBundleScopeByGovernanceLabelValues() []GovernanceLabelValuesInner`

GetEntitlementBundleScopeByGovernanceLabelValues returns the EntitlementBundleScopeByGovernanceLabelValues field if non-nil, zero value otherwise.

### GetEntitlementBundleScopeByGovernanceLabelValuesOk

`func (o *TargetResourcesRequestInner) GetEntitlementBundleScopeByGovernanceLabelValuesOk() (*[]GovernanceLabelValuesInner, bool)`

GetEntitlementBundleScopeByGovernanceLabelValuesOk returns a tuple with the EntitlementBundleScopeByGovernanceLabelValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundleScopeByGovernanceLabelValues

`func (o *TargetResourcesRequestInner) SetEntitlementBundleScopeByGovernanceLabelValues(v []GovernanceLabelValuesInner)`

SetEntitlementBundleScopeByGovernanceLabelValues sets EntitlementBundleScopeByGovernanceLabelValues field to given value.

### HasEntitlementBundleScopeByGovernanceLabelValues

`func (o *TargetResourcesRequestInner) HasEntitlementBundleScopeByGovernanceLabelValues() bool`

HasEntitlementBundleScopeByGovernanceLabelValues returns a boolean if a field has been set.

### GetAppServiceAccounts

`func (o *TargetResourcesRequestInner) GetAppServiceAccounts() []AppServiceAccountsInner`

GetAppServiceAccounts returns the AppServiceAccounts field if non-nil, zero value otherwise.

### GetAppServiceAccountsOk

`func (o *TargetResourcesRequestInner) GetAppServiceAccountsOk() (*[]AppServiceAccountsInner, bool)`

GetAppServiceAccountsOk returns a tuple with the AppServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceAccounts

`func (o *TargetResourcesRequestInner) SetAppServiceAccounts(v []AppServiceAccountsInner)`

SetAppServiceAccounts sets AppServiceAccounts field to given value.

### HasAppServiceAccounts

`func (o *TargetResourcesRequestInner) HasAppServiceAccounts() bool`

HasAppServiceAccounts returns a boolean if a field has been set.

### GetIncludeAllAppServiceAccounts

`func (o *TargetResourcesRequestInner) GetIncludeAllAppServiceAccounts() bool`

GetIncludeAllAppServiceAccounts returns the IncludeAllAppServiceAccounts field if non-nil, zero value otherwise.

### GetIncludeAllAppServiceAccountsOk

`func (o *TargetResourcesRequestInner) GetIncludeAllAppServiceAccountsOk() (*bool, bool)`

GetIncludeAllAppServiceAccountsOk returns a tuple with the IncludeAllAppServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllAppServiceAccounts

`func (o *TargetResourcesRequestInner) SetIncludeAllAppServiceAccounts(v bool)`

SetIncludeAllAppServiceAccounts sets IncludeAllAppServiceAccounts field to given value.

### HasIncludeAllAppServiceAccounts

`func (o *TargetResourcesRequestInner) HasIncludeAllAppServiceAccounts() bool`

HasIncludeAllAppServiceAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


