# CollectionResourceCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to **int32** | The total number of entitlements included for this app resource in this collection | [optional] 
**PushGroups** | Pointer to **int32** | The total number of push groups associated with this app resource in this collection | [optional] 
**Labels** | Pointer to **int32** | The total number of governance labels associated with this resource. Only returned when &#x60;include&#x3D;labels&#x60; is specified. Not supported in LIST calls. | [optional] 
**RelatedApps** | Pointer to **int32** | The total number of related apps for this resource. Only returned when &#x60;include&#x3D;relatedApps&#x60; is specified. Not supported in LIST calls. | [optional] 

## Methods

### NewCollectionResourceCounts

`func NewCollectionResourceCounts() *CollectionResourceCounts`

NewCollectionResourceCounts instantiates a new CollectionResourceCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceCountsWithDefaults

`func NewCollectionResourceCountsWithDefaults() *CollectionResourceCounts`

NewCollectionResourceCountsWithDefaults instantiates a new CollectionResourceCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *CollectionResourceCounts) GetEntitlements() int32`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CollectionResourceCounts) GetEntitlementsOk() (*int32, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CollectionResourceCounts) SetEntitlements(v int32)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CollectionResourceCounts) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetPushGroups

`func (o *CollectionResourceCounts) GetPushGroups() int32`

GetPushGroups returns the PushGroups field if non-nil, zero value otherwise.

### GetPushGroupsOk

`func (o *CollectionResourceCounts) GetPushGroupsOk() (*int32, bool)`

GetPushGroupsOk returns a tuple with the PushGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushGroups

`func (o *CollectionResourceCounts) SetPushGroups(v int32)`

SetPushGroups sets PushGroups field to given value.

### HasPushGroups

`func (o *CollectionResourceCounts) HasPushGroups() bool`

HasPushGroups returns a boolean if a field has been set.

### GetLabels

`func (o *CollectionResourceCounts) GetLabels() int32`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CollectionResourceCounts) GetLabelsOk() (*int32, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CollectionResourceCounts) SetLabels(v int32)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CollectionResourceCounts) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRelatedApps

`func (o *CollectionResourceCounts) GetRelatedApps() int32`

GetRelatedApps returns the RelatedApps field if non-nil, zero value otherwise.

### GetRelatedAppsOk

`func (o *CollectionResourceCounts) GetRelatedAppsOk() (*int32, bool)`

GetRelatedAppsOk returns a tuple with the RelatedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedApps

`func (o *CollectionResourceCounts) SetRelatedApps(v int32)`

SetRelatedApps sets RelatedApps field to given value.

### HasRelatedApps

`func (o *CollectionResourceCounts) HasRelatedApps() bool`

HasRelatedApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


