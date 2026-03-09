# CollectionResourceFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceProfile** | Pointer to [**ResourceProfile**](ResourceProfile.md) |  | [optional] 
**Entitlements** | Pointer to [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | [optional] 
**EntitlementValueCount** | Pointer to **int32** | The number of entitlements associated with this resource in the collection. Use the &#x60;include&#x60; query parameter to return this count in the response. | [optional] 
**ResourceOrn** | **string** | The ORN identifier for a specific app. Other resource types aren&#39;t supported.  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint for reference.  | 
**ResourceId** | Pointer to **string** | The Okta &#x60;app.id&#x60; of the resource.  See [List applications](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Application/#tag/Application/operation/listApplications) to retrieve app IDs.  | [optional] 
**Links** | [**CollectionResourceLinks**](CollectionResourceLinks.md) |  | 

## Methods

### NewCollectionResourceFull

`func NewCollectionResourceFull(resourceOrn string, links CollectionResourceLinks, ) *CollectionResourceFull`

NewCollectionResourceFull instantiates a new CollectionResourceFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceFullWithDefaults

`func NewCollectionResourceFullWithDefaults() *CollectionResourceFull`

NewCollectionResourceFullWithDefaults instantiates a new CollectionResourceFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceProfile

`func (o *CollectionResourceFull) GetResourceProfile() ResourceProfile`

GetResourceProfile returns the ResourceProfile field if non-nil, zero value otherwise.

### GetResourceProfileOk

`func (o *CollectionResourceFull) GetResourceProfileOk() (*ResourceProfile, bool)`

GetResourceProfileOk returns a tuple with the ResourceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProfile

`func (o *CollectionResourceFull) SetResourceProfile(v ResourceProfile)`

SetResourceProfile sets ResourceProfile field to given value.

### HasResourceProfile

`func (o *CollectionResourceFull) HasResourceProfile() bool`

HasResourceProfile returns a boolean if a field has been set.

### GetEntitlements

`func (o *CollectionResourceFull) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CollectionResourceFull) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CollectionResourceFull) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CollectionResourceFull) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEntitlementValueCount

`func (o *CollectionResourceFull) GetEntitlementValueCount() int32`

GetEntitlementValueCount returns the EntitlementValueCount field if non-nil, zero value otherwise.

### GetEntitlementValueCountOk

`func (o *CollectionResourceFull) GetEntitlementValueCountOk() (*int32, bool)`

GetEntitlementValueCountOk returns a tuple with the EntitlementValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValueCount

`func (o *CollectionResourceFull) SetEntitlementValueCount(v int32)`

SetEntitlementValueCount sets EntitlementValueCount field to given value.

### HasEntitlementValueCount

`func (o *CollectionResourceFull) HasEntitlementValueCount() bool`

HasEntitlementValueCount returns a boolean if a field has been set.

### GetResourceOrn

`func (o *CollectionResourceFull) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionResourceFull) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionResourceFull) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetResourceId

`func (o *CollectionResourceFull) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionResourceFull) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionResourceFull) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionResourceFull) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourceFull) GetLinks() CollectionResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourceFull) GetLinksOk() (*CollectionResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourceFull) SetLinks(v CollectionResourceLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


