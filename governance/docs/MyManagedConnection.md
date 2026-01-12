# MyManagedConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**AuthorizationServer** | [**CustomAuthorizationServer**](CustomAuthorizationServer.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the managed connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Secret** | [**ManagedConnectionVaultedSecret**](ManagedConnectionVaultedSecret.md) |  | 
**App** | [**ManagedConnectionAppInstance**](ManagedConnectionAppInstance.md) |  | 
**ServiceAccount** | [**ManagedConnectionServiceAccount**](ManagedConnectionServiceAccount.md) |  | 

## Methods

### NewMyManagedConnection

`func NewMyManagedConnection(connectionType string, authorizationServer CustomAuthorizationServer, secret ManagedConnectionVaultedSecret, app ManagedConnectionAppInstance, serviceAccount ManagedConnectionServiceAccount, ) *MyManagedConnection`

NewMyManagedConnection instantiates a new MyManagedConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyManagedConnectionWithDefaults

`func NewMyManagedConnectionWithDefaults() *MyManagedConnection`

NewMyManagedConnectionWithDefaults instantiates a new MyManagedConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *MyManagedConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *MyManagedConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *MyManagedConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetAuthorizationServer

`func (o *MyManagedConnection) GetAuthorizationServer() CustomAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *MyManagedConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *MyManagedConnection) SetAuthorizationServer(v CustomAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetId

`func (o *MyManagedConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyManagedConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyManagedConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MyManagedConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *MyManagedConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *MyManagedConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *MyManagedConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *MyManagedConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *MyManagedConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyManagedConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyManagedConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyManagedConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecret

`func (o *MyManagedConnection) GetSecret() ManagedConnectionVaultedSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MyManagedConnection) GetSecretOk() (*ManagedConnectionVaultedSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MyManagedConnection) SetSecret(v ManagedConnectionVaultedSecret)`

SetSecret sets Secret field to given value.


### GetApp

`func (o *MyManagedConnection) GetApp() ManagedConnectionAppInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *MyManagedConnection) GetAppOk() (*ManagedConnectionAppInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *MyManagedConnection) SetApp(v ManagedConnectionAppInstance)`

SetApp sets App field to given value.


### GetServiceAccount

`func (o *MyManagedConnection) GetServiceAccount() ManagedConnectionServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *MyManagedConnection) GetServiceAccountOk() (*ManagedConnectionServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *MyManagedConnection) SetServiceAccount(v ManagedConnectionServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


