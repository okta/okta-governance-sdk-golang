# AccessScopeSettingsCreatableAccessScopeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**EntitlementBundles** | [**[]EntitlementBundlesArrayCreatableInner**](EntitlementBundlesArrayCreatableInner.md) | Entitlement bundles that are made requestable | 
**Groups** | [**[]GroupsArrayCreatableInner**](GroupsArrayCreatableInner.md) |  | 

## Methods

### NewAccessScopeSettingsCreatableAccessScopeSettings

`func NewAccessScopeSettingsCreatableAccessScopeSettings(type_ string, entitlementBundles []EntitlementBundlesArrayCreatableInner, groups []GroupsArrayCreatableInner, ) *AccessScopeSettingsCreatableAccessScopeSettings`

NewAccessScopeSettingsCreatableAccessScopeSettings instantiates a new AccessScopeSettingsCreatableAccessScopeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessScopeSettingsCreatableAccessScopeSettingsWithDefaults

`func NewAccessScopeSettingsCreatableAccessScopeSettingsWithDefaults() *AccessScopeSettingsCreatableAccessScopeSettings`

NewAccessScopeSettingsCreatableAccessScopeSettingsWithDefaults instantiates a new AccessScopeSettingsCreatableAccessScopeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) SetType(v string)`

SetType sets Type field to given value.


### GetEntitlementBundles

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetEntitlementBundles() []EntitlementBundlesArrayCreatableInner`

GetEntitlementBundles returns the EntitlementBundles field if non-nil, zero value otherwise.

### GetEntitlementBundlesOk

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetEntitlementBundlesOk() (*[]EntitlementBundlesArrayCreatableInner, bool)`

GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundles

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) SetEntitlementBundles(v []EntitlementBundlesArrayCreatableInner)`

SetEntitlementBundles sets EntitlementBundles field to given value.


### GetGroups

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetGroups() []GroupsArrayCreatableInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetGroupsOk() (*[]GroupsArrayCreatableInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) SetGroups(v []GroupsArrayCreatableInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


