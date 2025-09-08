# DelegateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Okta user &#x60;id&#x60; | 
**FirstName** | Pointer to **string** | The user&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name | [optional] 
**Email** | **string** | The user&#39;s email address | 

## Methods

### NewDelegateUser

`func NewDelegateUser(id string, email string, ) *DelegateUser`

NewDelegateUser instantiates a new DelegateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegateUserWithDefaults

`func NewDelegateUserWithDefaults() *DelegateUser`

NewDelegateUserWithDefaults instantiates a new DelegateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelegateUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegateUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegateUser) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *DelegateUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DelegateUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DelegateUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DelegateUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *DelegateUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DelegateUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DelegateUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DelegateUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *DelegateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DelegateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DelegateUser) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


