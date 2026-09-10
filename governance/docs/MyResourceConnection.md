# MyResourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**AuthorizationServer** | [**CustomAuthorizationServer**](CustomAuthorizationServer.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the resource connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Secret** | [**ResourceConnectionVaultedSecret**](ResourceConnectionVaultedSecret.md) |  | 
**App** | [**ResourceConnectionAppInstance**](ResourceConnectionAppInstance.md) |  | 
**ServiceAccount** | [**ResourceConnectionServiceAccount**](ResourceConnectionServiceAccount.md) |  | 
**Resource** | [**StsAccessTokenResource**](StsAccessTokenResource.md) |  | 
**A2aServer** | [**ResourceConnectionA2aServer**](ResourceConnectionA2aServer.md) |  | 

## Methods

### NewMyResourceConnection

`func NewMyResourceConnection(connectionType string, authorizationServer CustomAuthorizationServer, secret ResourceConnectionVaultedSecret, app ResourceConnectionAppInstance, serviceAccount ResourceConnectionServiceAccount, resource StsAccessTokenResource, a2aServer ResourceConnectionA2aServer, ) *MyResourceConnection`

NewMyResourceConnection instantiates a new MyResourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyResourceConnectionWithDefaults

`func NewMyResourceConnectionWithDefaults() *MyResourceConnection`

NewMyResourceConnectionWithDefaults instantiates a new MyResourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *MyResourceConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *MyResourceConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *MyResourceConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetAuthorizationServer

`func (o *MyResourceConnection) GetAuthorizationServer() CustomAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *MyResourceConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *MyResourceConnection) SetAuthorizationServer(v CustomAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetId

`func (o *MyResourceConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyResourceConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyResourceConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MyResourceConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *MyResourceConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *MyResourceConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *MyResourceConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *MyResourceConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *MyResourceConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyResourceConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyResourceConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyResourceConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecret

`func (o *MyResourceConnection) GetSecret() ResourceConnectionVaultedSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MyResourceConnection) GetSecretOk() (*ResourceConnectionVaultedSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MyResourceConnection) SetSecret(v ResourceConnectionVaultedSecret)`

SetSecret sets Secret field to given value.


### GetApp

`func (o *MyResourceConnection) GetApp() ResourceConnectionAppInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *MyResourceConnection) GetAppOk() (*ResourceConnectionAppInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *MyResourceConnection) SetApp(v ResourceConnectionAppInstance)`

SetApp sets App field to given value.


### GetServiceAccount

`func (o *MyResourceConnection) GetServiceAccount() ResourceConnectionServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *MyResourceConnection) GetServiceAccountOk() (*ResourceConnectionServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *MyResourceConnection) SetServiceAccount(v ResourceConnectionServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetResource

`func (o *MyResourceConnection) GetResource() StsAccessTokenResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MyResourceConnection) GetResourceOk() (*StsAccessTokenResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MyResourceConnection) SetResource(v StsAccessTokenResource)`

SetResource sets Resource field to given value.


### GetA2aServer

`func (o *MyResourceConnection) GetA2aServer() ResourceConnectionA2aServer`

GetA2aServer returns the A2aServer field if non-nil, zero value otherwise.

### GetA2aServerOk

`func (o *MyResourceConnection) GetA2aServerOk() (*ResourceConnectionA2aServer, bool)`

GetA2aServerOk returns a tuple with the A2aServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2aServer

`func (o *MyResourceConnection) SetA2aServer(v ResourceConnectionA2aServer)`

SetA2aServer sets A2aServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


