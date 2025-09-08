# DelegateUsersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DelegateUser**](DelegateUser.md) | Delegate users list | [optional] 
**Links** | Pointer to [**LinkSelf**](LinkSelf.md) |  | [optional] 

## Methods

### NewDelegateUsersList

`func NewDelegateUsersList() *DelegateUsersList`

NewDelegateUsersList instantiates a new DelegateUsersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegateUsersListWithDefaults

`func NewDelegateUsersListWithDefaults() *DelegateUsersList`

NewDelegateUsersListWithDefaults instantiates a new DelegateUsersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DelegateUsersList) GetData() []DelegateUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DelegateUsersList) GetDataOk() (*[]DelegateUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DelegateUsersList) SetData(v []DelegateUser)`

SetData sets Data field to given value.

### HasData

`func (o *DelegateUsersList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DelegateUsersList) GetLinks() LinkSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DelegateUsersList) GetLinksOk() (*LinkSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DelegateUsersList) SetLinks(v LinkSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DelegateUsersList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


