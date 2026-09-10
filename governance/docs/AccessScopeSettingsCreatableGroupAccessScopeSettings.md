# AccessScopeSettingsCreatableGroupAccessScopeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Request for access to groups | 
**Groups** | [**[]GroupsArrayRequesterSettingsCreatableInner**](GroupsArrayRequesterSettingsCreatableInner.md) | List of requestable groups. You can specify a maximum of 500 groups as the access scope for each condition.  &gt; **Note:** Both standard Okta groups and AD-sourced groups are supported in Access Requests. &gt; Standard Okta groups have the &#x60;okta:user_group&#x60; value, whereas AD-sourced groups have the &#x60;okta:windows_security_principal&#x60; value in their [&#x60;objectClass&#x60; property](https://developer.okta.com/docs/api/openapi/okta-management/management/tags/group/other/getgroup#other/getgroup/t&#x3D;response&amp;c&#x3D;200&amp;path&#x3D;objectclass). | 

## Methods

### NewAccessScopeSettingsCreatableGroupAccessScopeSettings

`func NewAccessScopeSettingsCreatableGroupAccessScopeSettings(type_ string, groups []GroupsArrayRequesterSettingsCreatableInner, ) *AccessScopeSettingsCreatableGroupAccessScopeSettings`

NewAccessScopeSettingsCreatableGroupAccessScopeSettings instantiates a new AccessScopeSettingsCreatableGroupAccessScopeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessScopeSettingsCreatableGroupAccessScopeSettingsWithDefaults

`func NewAccessScopeSettingsCreatableGroupAccessScopeSettingsWithDefaults() *AccessScopeSettingsCreatableGroupAccessScopeSettings`

NewAccessScopeSettingsCreatableGroupAccessScopeSettingsWithDefaults instantiates a new AccessScopeSettingsCreatableGroupAccessScopeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) SetType(v string)`

SetType sets Type field to given value.


### GetGroups

`func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetGroups() []GroupsArrayRequesterSettingsCreatableInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) GetGroupsOk() (*[]GroupsArrayRequesterSettingsCreatableInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AccessScopeSettingsCreatableGroupAccessScopeSettings) SetGroups(v []GroupsArrayRequesterSettingsCreatableInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


