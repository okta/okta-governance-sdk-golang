# ResourceInventoryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | **int32** | The number of app resources matching the search query | [readonly] 
**Groups** | **int32** | The number of group resources matching the search query | [readonly] 
**EntitlementValues** | **int32** | The number of entitlement value resources matching the search query | [readonly] 
**EntitlementBundles** | **int32** | The number of entitlement bundle resources matching the search query | [readonly] 
**Collections** | **int32** | The number of collection resources matching the search query | [readonly] 

## Methods

### NewResourceInventoryStats

`func NewResourceInventoryStats(apps int32, groups int32, entitlementValues int32, entitlementBundles int32, collections int32, ) *ResourceInventoryStats`

NewResourceInventoryStats instantiates a new ResourceInventoryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInventoryStatsWithDefaults

`func NewResourceInventoryStatsWithDefaults() *ResourceInventoryStats`

NewResourceInventoryStatsWithDefaults instantiates a new ResourceInventoryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *ResourceInventoryStats) GetApps() int32`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ResourceInventoryStats) GetAppsOk() (*int32, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ResourceInventoryStats) SetApps(v int32)`

SetApps sets Apps field to given value.


### GetGroups

`func (o *ResourceInventoryStats) GetGroups() int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ResourceInventoryStats) GetGroupsOk() (*int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ResourceInventoryStats) SetGroups(v int32)`

SetGroups sets Groups field to given value.


### GetEntitlementValues

`func (o *ResourceInventoryStats) GetEntitlementValues() int32`

GetEntitlementValues returns the EntitlementValues field if non-nil, zero value otherwise.

### GetEntitlementValuesOk

`func (o *ResourceInventoryStats) GetEntitlementValuesOk() (*int32, bool)`

GetEntitlementValuesOk returns a tuple with the EntitlementValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValues

`func (o *ResourceInventoryStats) SetEntitlementValues(v int32)`

SetEntitlementValues sets EntitlementValues field to given value.


### GetEntitlementBundles

`func (o *ResourceInventoryStats) GetEntitlementBundles() int32`

GetEntitlementBundles returns the EntitlementBundles field if non-nil, zero value otherwise.

### GetEntitlementBundlesOk

`func (o *ResourceInventoryStats) GetEntitlementBundlesOk() (*int32, bool)`

GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundles

`func (o *ResourceInventoryStats) SetEntitlementBundles(v int32)`

SetEntitlementBundles sets EntitlementBundles field to given value.


### GetCollections

`func (o *ResourceInventoryStats) GetCollections() int32`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ResourceInventoryStats) GetCollectionsOk() (*int32, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ResourceInventoryStats) SetCollections(v int32)`

SetCollections sets Collections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


