# GroupsRequesterSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Groups** | [**[]GroupsArrayFullInner**](GroupsArrayFullInner.md) | A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.  | 

## Methods

### NewGroupsRequesterSettings

`func NewGroupsRequesterSettings(type_ string, groups []GroupsArrayFullInner, ) *GroupsRequesterSettings`

NewGroupsRequesterSettings instantiates a new GroupsRequesterSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsRequesterSettingsWithDefaults

`func NewGroupsRequesterSettingsWithDefaults() *GroupsRequesterSettings`

NewGroupsRequesterSettingsWithDefaults instantiates a new GroupsRequesterSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GroupsRequesterSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupsRequesterSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupsRequesterSettings) SetType(v string)`

SetType sets Type field to given value.


### GetGroups

`func (o *GroupsRequesterSettings) GetGroups() []GroupsArrayFullInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupsRequesterSettings) GetGroupsOk() (*[]GroupsArrayFullInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupsRequesterSettings) SetGroups(v []GroupsArrayFullInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


