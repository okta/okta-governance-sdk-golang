# EntryUsersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EntryUser**](EntryUser.md) | Catalog entry users list | [optional] 
**Links** | Pointer to [**EntryUserListLinks**](EntryUserListLinks.md) |  | [optional] 

## Methods

### NewEntryUsersList

`func NewEntryUsersList() *EntryUsersList`

NewEntryUsersList instantiates a new EntryUsersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryUsersListWithDefaults

`func NewEntryUsersListWithDefaults() *EntryUsersList`

NewEntryUsersListWithDefaults instantiates a new EntryUsersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EntryUsersList) GetData() []EntryUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EntryUsersList) GetDataOk() (*[]EntryUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EntryUsersList) SetData(v []EntryUser)`

SetData sets Data field to given value.

### HasData

`func (o *EntryUsersList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *EntryUsersList) GetLinks() EntryUserListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntryUsersList) GetLinksOk() (*EntryUserListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntryUsersList) SetLinks(v EntryUserListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntryUsersList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


