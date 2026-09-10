# AccessScopeSettingsCreatableAccessScopeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Request for resource access | 
**EntitlementBundles** | [**[]EntitlementBundlesArrayCreatableInner**](EntitlementBundlesArrayCreatableInner.md) | Entitlement bundles that are made requestable | 
**Groups** | [**[]GroupsArrayRequesterSettingsCreatableInner**](GroupsArrayRequesterSettingsCreatableInner.md) | List of requestable groups. You can specify a maximum of 500 groups as the access scope for each condition.  &gt; **Note:** Both standard Okta groups and AD-sourced groups are supported in Access Requests. &gt; Standard Okta groups have the &#x60;okta:user_group&#x60; value, whereas AD-sourced groups have the &#x60;okta:windows_security_principal&#x60; value in their [&#x60;objectClass&#x60; property](https://developer.okta.com/docs/api/openapi/okta-management/management/tags/group/other/getgroup#other/getgroup/t&#x3D;response&amp;c&#x3D;200&amp;path&#x3D;objectclass). | 

## Methods

### NewAccessScopeSettingsCreatableAccessScopeSettings

`func NewAccessScopeSettingsCreatableAccessScopeSettings(type_ string, entitlementBundles []EntitlementBundlesArrayCreatableInner, groups []GroupsArrayRequesterSettingsCreatableInner, ) *AccessScopeSettingsCreatableAccessScopeSettings`

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

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetGroups() []GroupsArrayRequesterSettingsCreatableInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) GetGroupsOk() (*[]GroupsArrayRequesterSettingsCreatableInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AccessScopeSettingsCreatableAccessScopeSettings) SetGroups(v []GroupsArrayRequesterSettingsCreatableInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


