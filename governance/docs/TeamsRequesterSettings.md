# TeamsRequesterSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Teams** | [**[]TeamsArrayFullInner**](TeamsArrayFullInner.md) | A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.  | 

## Methods

### NewTeamsRequesterSettings

`func NewTeamsRequesterSettings(type_ string, teams []TeamsArrayFullInner, ) *TeamsRequesterSettings`

NewTeamsRequesterSettings instantiates a new TeamsRequesterSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsRequesterSettingsWithDefaults

`func NewTeamsRequesterSettingsWithDefaults() *TeamsRequesterSettings`

NewTeamsRequesterSettingsWithDefaults instantiates a new TeamsRequesterSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TeamsRequesterSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamsRequesterSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamsRequesterSettings) SetType(v string)`

SetType sets Type field to given value.


### GetTeams

`func (o *TeamsRequesterSettings) GetTeams() []TeamsArrayFullInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TeamsRequesterSettings) GetTeamsOk() (*[]TeamsArrayFullInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TeamsRequesterSettings) SetTeams(v []TeamsArrayFullInner)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


