# MySettingsGetDelegateReadonlyDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The Okta user &#x60;id&#x60; | 
**Type** | [**PrincipalType**](PrincipalType.md) |  | 
**FirstName** | Pointer to **string** | The user&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address | [optional] 

## Methods

### NewMySettingsGetDelegateReadonlyDelegate

`func NewMySettingsGetDelegateReadonlyDelegate(externalId string, type_ PrincipalType, ) *MySettingsGetDelegateReadonlyDelegate`

NewMySettingsGetDelegateReadonlyDelegate instantiates a new MySettingsGetDelegateReadonlyDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySettingsGetDelegateReadonlyDelegateWithDefaults

`func NewMySettingsGetDelegateReadonlyDelegateWithDefaults() *MySettingsGetDelegateReadonlyDelegate`

NewMySettingsGetDelegateReadonlyDelegateWithDefaults instantiates a new MySettingsGetDelegateReadonlyDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *MySettingsGetDelegateReadonlyDelegate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MySettingsGetDelegateReadonlyDelegate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MySettingsGetDelegateReadonlyDelegate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *MySettingsGetDelegateReadonlyDelegate) GetType() PrincipalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MySettingsGetDelegateReadonlyDelegate) GetTypeOk() (*PrincipalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MySettingsGetDelegateReadonlyDelegate) SetType(v PrincipalType)`

SetType sets Type field to given value.


### GetFirstName

`func (o *MySettingsGetDelegateReadonlyDelegate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MySettingsGetDelegateReadonlyDelegate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MySettingsGetDelegateReadonlyDelegate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MySettingsGetDelegateReadonlyDelegate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MySettingsGetDelegateReadonlyDelegate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MySettingsGetDelegateReadonlyDelegate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MySettingsGetDelegateReadonlyDelegate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MySettingsGetDelegateReadonlyDelegate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *MySettingsGetDelegateReadonlyDelegate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MySettingsGetDelegateReadonlyDelegate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MySettingsGetDelegateReadonlyDelegate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MySettingsGetDelegateReadonlyDelegate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


