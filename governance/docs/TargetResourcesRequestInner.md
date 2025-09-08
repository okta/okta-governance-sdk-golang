# TargetResourcesRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The resource ID that is being reviewed. The &#x60;resourceId&#x60; can have a different value based on the &#x60;resourceType&#x60;. When the &#x60;resourceType &#x3D; GROUP&#x60;, the value is a group ID. When the &#x60;resourceType &#x3D; APPLICATION&#x60;, the value is an application ID.  | 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**IncludeAllEntitlementsAndBundles** | Pointer to **bool** | Include all entitlements and entitlement bundles for this application. Only applicable if the &#x60;resourcetype &#x3D; APPLICATION&#x60; and Entitlement Management is enabled. | [optional] 
**EntitlementBundles** | Pointer to [**[]EntitlementBundlesInner**](EntitlementBundlesInner.md) | An array of entitlement bundles associated with &#x60;resourceId&#x60; that should be chosen as target when creating reviews. Only applicable if the &#x60;resourceType &#x3D; APPLICATION&#x60; and Entitlement Management is enabled. | [optional] 
**Entitlements** | Pointer to [**[]EntitlementsInner**](EntitlementsInner.md) | An array of entitlements associated with &#x60;resourceId&#x60; that should be chosen as target when creating reviews. Only applicable if &#x60;resourceType &#x3D; APPLICATION&#x60; and Entitlement Management is enabled. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


