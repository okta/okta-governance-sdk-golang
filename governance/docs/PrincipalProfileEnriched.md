# PrincipalProfileEnriched

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Okta user &#x60;id&#x60; | 
**Email** | Pointer to **string** | The Okta user&#39;s email address | [optional] 
**FirstName** | Pointer to **string** | The Okta user&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The Okta user&#39;s last name | [optional] 
**Login** | Pointer to **string** | The Okta user&#39;s login | [optional] 
**Status** | [**PrincipalProfileStatus**](PrincipalProfileStatus.md) |  | 
**Type** | [**PrincipalProfileType**](PrincipalProfileType.md) |  | 

## Methods

### NewPrincipalProfileEnriched

`func NewPrincipalProfileEnriched(id string, status PrincipalProfileStatus, type_ PrincipalProfileType, ) *PrincipalProfileEnriched`

NewPrincipalProfileEnriched instantiates a new PrincipalProfileEnriched object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalProfileEnrichedWithDefaults

`func NewPrincipalProfileEnrichedWithDefaults() *PrincipalProfileEnriched`

NewPrincipalProfileEnrichedWithDefaults instantiates a new PrincipalProfileEnriched object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalProfileEnriched) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalProfileEnriched) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalProfileEnriched) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *PrincipalProfileEnriched) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrincipalProfileEnriched) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrincipalProfileEnriched) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PrincipalProfileEnriched) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *PrincipalProfileEnriched) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PrincipalProfileEnriched) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PrincipalProfileEnriched) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PrincipalProfileEnriched) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PrincipalProfileEnriched) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PrincipalProfileEnriched) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PrincipalProfileEnriched) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PrincipalProfileEnriched) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLogin

`func (o *PrincipalProfileEnriched) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PrincipalProfileEnriched) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PrincipalProfileEnriched) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *PrincipalProfileEnriched) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalProfileEnriched) GetStatus() PrincipalProfileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalProfileEnriched) GetStatusOk() (*PrincipalProfileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalProfileEnriched) SetStatus(v PrincipalProfileStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *PrincipalProfileEnriched) GetType() PrincipalProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrincipalProfileEnriched) GetTypeOk() (*PrincipalProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrincipalProfileEnriched) SetType(v PrincipalProfileType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


