# CollectionResourceAppWithEntitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceConfiguration** | Pointer to [**CollectionResourceConfiguration**](CollectionResourceConfiguration.md) |  | [optional] 
**Entitlements** | Pointer to [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | [optional] 
**ResourceOrn** | **string** | The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.  | 
**ResourceId** | Pointer to **string** | The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as &#x60;GET&#x60;/&#x60;PUT&#x60;/&#x60;DELETE /v2/collections/{collectionId}/resources/{resourceId}&#x60;.  | [optional] 
**Links** | [**CollectionResourceLinksV2**](CollectionResourceLinksV2.md) |  | 

## Methods

### NewCollectionResourceAppWithEntitlements

`func NewCollectionResourceAppWithEntitlements(resourceOrn string, links CollectionResourceLinksV2, ) *CollectionResourceAppWithEntitlements`

NewCollectionResourceAppWithEntitlements instantiates a new CollectionResourceAppWithEntitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceAppWithEntitlementsWithDefaults

`func NewCollectionResourceAppWithEntitlementsWithDefaults() *CollectionResourceAppWithEntitlements`

NewCollectionResourceAppWithEntitlementsWithDefaults instantiates a new CollectionResourceAppWithEntitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceConfiguration

`func (o *CollectionResourceAppWithEntitlements) GetResourceConfiguration() CollectionResourceConfiguration`

GetResourceConfiguration returns the ResourceConfiguration field if non-nil, zero value otherwise.

### GetResourceConfigurationOk

`func (o *CollectionResourceAppWithEntitlements) GetResourceConfigurationOk() (*CollectionResourceConfiguration, bool)`

GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConfiguration

`func (o *CollectionResourceAppWithEntitlements) SetResourceConfiguration(v CollectionResourceConfiguration)`

SetResourceConfiguration sets ResourceConfiguration field to given value.

### HasResourceConfiguration

`func (o *CollectionResourceAppWithEntitlements) HasResourceConfiguration() bool`

HasResourceConfiguration returns a boolean if a field has been set.

### GetEntitlements

`func (o *CollectionResourceAppWithEntitlements) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CollectionResourceAppWithEntitlements) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CollectionResourceAppWithEntitlements) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CollectionResourceAppWithEntitlements) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetResourceOrn

`func (o *CollectionResourceAppWithEntitlements) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionResourceAppWithEntitlements) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionResourceAppWithEntitlements) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetResourceId

`func (o *CollectionResourceAppWithEntitlements) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionResourceAppWithEntitlements) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionResourceAppWithEntitlements) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionResourceAppWithEntitlements) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourceAppWithEntitlements) GetLinks() CollectionResourceLinksV2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourceAppWithEntitlements) GetLinksOk() (*CollectionResourceLinksV2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourceAppWithEntitlements) SetLinks(v CollectionResourceLinksV2)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


