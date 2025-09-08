# PrincipalProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Okta user &#x60;id&#x60; | 
**Email** | **string** | The Okta user&#39;s email address | 
**FirstName** | Pointer to **string** | The Okta user&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The Okta user&#39;s last name | [optional] 
**Login** | Pointer to **string** | The Okta user&#39;s login | [optional] 
**Status** | [**PrincipalProfileStatus**](PrincipalProfileStatus.md) |  | 

## Methods

### NewPrincipalProfile

`func NewPrincipalProfile(id string, email string, status PrincipalProfileStatus, ) *PrincipalProfile`

NewPrincipalProfile instantiates a new PrincipalProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalProfileWithDefaults

`func NewPrincipalProfileWithDefaults() *PrincipalProfile`

NewPrincipalProfileWithDefaults instantiates a new PrincipalProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalProfile) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *PrincipalProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrincipalProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrincipalProfile) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *PrincipalProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PrincipalProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PrincipalProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PrincipalProfile) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PrincipalProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PrincipalProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PrincipalProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PrincipalProfile) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLogin

`func (o *PrincipalProfile) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PrincipalProfile) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PrincipalProfile) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *PrincipalProfile) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalProfile) GetStatus() PrincipalProfileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalProfile) GetStatusOk() (*PrincipalProfileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalProfile) SetStatus(v PrincipalProfileStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


