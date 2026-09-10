# CollectionResourceFullV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceProfile** | Pointer to [**ResourceProfileV2**](ResourceProfileV2.md) |  | [optional] 
**ResourceConfiguration** | Pointer to [**CollectionResourceConfiguration**](CollectionResourceConfiguration.md) |  | [optional] 
**Entitlements** | Pointer to [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | [optional] 
**PushGroups** | Pointer to [**[]CollectionPushGroup**](CollectionPushGroup.md) | List of push groups associated with an app collection resource | [optional] 
**Labels** | Pointer to [**[]LabelValue**](LabelValue.md) | The governance labels associated with this resource | [optional] 
**RelatedApps** | Pointer to [**[]RelatedApp**](RelatedApp.md) | The apps related to this resource. Only populated for &#x60;GROUP&#x60; type resources. For &#x60;GROUP&#x60; resources with &#x60;hasPushMapping: true&#x60;, these are apps that this group is mapped to through push group. For &#x60;GROUP&#x60; resources with &#x60;hasPushMapping: false&#x60;, these are apps assigned to this group. | [optional] 
**ResourceOrn** | **string** | The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.  | 
**ResourceId** | Pointer to **string** | The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as &#x60;GET&#x60;/&#x60;PUT&#x60;/&#x60;DELETE /v2/collections/{collectionId}/resources/{resourceId}&#x60;.  | [optional] 
**Links** | [**CollectionResourceLinksV2**](CollectionResourceLinksV2.md) |  | 

## Methods

### NewCollectionResourceFullV2

`func NewCollectionResourceFullV2(resourceOrn string, links CollectionResourceLinksV2, ) *CollectionResourceFullV2`

NewCollectionResourceFullV2 instantiates a new CollectionResourceFullV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceFullV2WithDefaults

`func NewCollectionResourceFullV2WithDefaults() *CollectionResourceFullV2`

NewCollectionResourceFullV2WithDefaults instantiates a new CollectionResourceFullV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceProfile

`func (o *CollectionResourceFullV2) GetResourceProfile() ResourceProfileV2`

GetResourceProfile returns the ResourceProfile field if non-nil, zero value otherwise.

### GetResourceProfileOk

`func (o *CollectionResourceFullV2) GetResourceProfileOk() (*ResourceProfileV2, bool)`

GetResourceProfileOk returns a tuple with the ResourceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProfile

`func (o *CollectionResourceFullV2) SetResourceProfile(v ResourceProfileV2)`

SetResourceProfile sets ResourceProfile field to given value.

### HasResourceProfile

`func (o *CollectionResourceFullV2) HasResourceProfile() bool`

HasResourceProfile returns a boolean if a field has been set.

### GetResourceConfiguration

`func (o *CollectionResourceFullV2) GetResourceConfiguration() CollectionResourceConfiguration`

GetResourceConfiguration returns the ResourceConfiguration field if non-nil, zero value otherwise.

### GetResourceConfigurationOk

`func (o *CollectionResourceFullV2) GetResourceConfigurationOk() (*CollectionResourceConfiguration, bool)`

GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConfiguration

`func (o *CollectionResourceFullV2) SetResourceConfiguration(v CollectionResourceConfiguration)`

SetResourceConfiguration sets ResourceConfiguration field to given value.

### HasResourceConfiguration

`func (o *CollectionResourceFullV2) HasResourceConfiguration() bool`

HasResourceConfiguration returns a boolean if a field has been set.

### GetEntitlements

`func (o *CollectionResourceFullV2) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CollectionResourceFullV2) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CollectionResourceFullV2) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CollectionResourceFullV2) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetPushGroups

`func (o *CollectionResourceFullV2) GetPushGroups() []CollectionPushGroup`

GetPushGroups returns the PushGroups field if non-nil, zero value otherwise.

### GetPushGroupsOk

`func (o *CollectionResourceFullV2) GetPushGroupsOk() (*[]CollectionPushGroup, bool)`

GetPushGroupsOk returns a tuple with the PushGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushGroups

`func (o *CollectionResourceFullV2) SetPushGroups(v []CollectionPushGroup)`

SetPushGroups sets PushGroups field to given value.

### HasPushGroups

`func (o *CollectionResourceFullV2) HasPushGroups() bool`

HasPushGroups returns a boolean if a field has been set.

### GetLabels

`func (o *CollectionResourceFullV2) GetLabels() []LabelValue`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CollectionResourceFullV2) GetLabelsOk() (*[]LabelValue, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CollectionResourceFullV2) SetLabels(v []LabelValue)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CollectionResourceFullV2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRelatedApps

`func (o *CollectionResourceFullV2) GetRelatedApps() []RelatedApp`

GetRelatedApps returns the RelatedApps field if non-nil, zero value otherwise.

### GetRelatedAppsOk

`func (o *CollectionResourceFullV2) GetRelatedAppsOk() (*[]RelatedApp, bool)`

GetRelatedAppsOk returns a tuple with the RelatedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedApps

`func (o *CollectionResourceFullV2) SetRelatedApps(v []RelatedApp)`

SetRelatedApps sets RelatedApps field to given value.

### HasRelatedApps

`func (o *CollectionResourceFullV2) HasRelatedApps() bool`

HasRelatedApps returns a boolean if a field has been set.

### GetResourceOrn

`func (o *CollectionResourceFullV2) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionResourceFullV2) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionResourceFullV2) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetResourceId

`func (o *CollectionResourceFullV2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionResourceFullV2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionResourceFullV2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionResourceFullV2) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourceFullV2) GetLinks() CollectionResourceLinksV2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourceFullV2) GetLinksOk() (*CollectionResourceLinksV2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourceFullV2) SetLinks(v CollectionResourceLinksV2)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


