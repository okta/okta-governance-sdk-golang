# RequesterSettingsCreatableRequesterSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Groups** | [**[]GroupsArrayCreatableInner**](GroupsArrayCreatableInner.md) | List of requestable groups  &gt; **Note:** Both standard Okta groups and AD-sourced groups are supported in Access Requests. &gt; Standard Okta groups have the &#x60;okta:user_group&#x60; value, whereas AD-sourced groups have the &#x60;okta:windows_security_principal&#x60; value in their &#x60;objectClass&#x60; property. | 
**Teams** | [**[]TeamsArrayCreatableInner**](TeamsArrayCreatableInner.md) |  | 

## Methods

### NewRequesterSettingsCreatableRequesterSettings

`func NewRequesterSettingsCreatableRequesterSettings(type_ string, groups []GroupsArrayCreatableInner, teams []TeamsArrayCreatableInner, ) *RequesterSettingsCreatableRequesterSettings`

NewRequesterSettingsCreatableRequesterSettings instantiates a new RequesterSettingsCreatableRequesterSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequesterSettingsCreatableRequesterSettingsWithDefaults

`func NewRequesterSettingsCreatableRequesterSettingsWithDefaults() *RequesterSettingsCreatableRequesterSettings`

NewRequesterSettingsCreatableRequesterSettingsWithDefaults instantiates a new RequesterSettingsCreatableRequesterSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequesterSettingsCreatableRequesterSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequesterSettingsCreatableRequesterSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequesterSettingsCreatableRequesterSettings) SetType(v string)`

SetType sets Type field to given value.


### GetGroups

`func (o *RequesterSettingsCreatableRequesterSettings) GetGroups() []GroupsArrayCreatableInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RequesterSettingsCreatableRequesterSettings) GetGroupsOk() (*[]GroupsArrayCreatableInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RequesterSettingsCreatableRequesterSettings) SetGroups(v []GroupsArrayCreatableInner)`

SetGroups sets Groups field to given value.


### GetTeams

`func (o *RequesterSettingsCreatableRequesterSettings) GetTeams() []TeamsArrayCreatableInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *RequesterSettingsCreatableRequesterSettings) GetTeamsOk() (*[]TeamsArrayCreatableInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *RequesterSettingsCreatableRequesterSettings) SetTeams(v []TeamsArrayCreatableInner)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


