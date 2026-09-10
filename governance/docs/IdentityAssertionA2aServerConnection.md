# IdentityAssertionA2aServerConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**A2aServer** | [**ResourceConnectionA2aServer**](ResourceConnectionA2aServer.md) |  | 
**AuthorizationServer** | [**CustomAuthorizationServer**](CustomAuthorizationServer.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the resource connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewIdentityAssertionA2aServerConnection

`func NewIdentityAssertionA2aServerConnection(connectionType string, a2aServer ResourceConnectionA2aServer, authorizationServer CustomAuthorizationServer, ) *IdentityAssertionA2aServerConnection`

NewIdentityAssertionA2aServerConnection instantiates a new IdentityAssertionA2aServerConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionA2aServerConnectionWithDefaults

`func NewIdentityAssertionA2aServerConnectionWithDefaults() *IdentityAssertionA2aServerConnection`

NewIdentityAssertionA2aServerConnectionWithDefaults instantiates a new IdentityAssertionA2aServerConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *IdentityAssertionA2aServerConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IdentityAssertionA2aServerConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IdentityAssertionA2aServerConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetA2aServer

`func (o *IdentityAssertionA2aServerConnection) GetA2aServer() ResourceConnectionA2aServer`

GetA2aServer returns the A2aServer field if non-nil, zero value otherwise.

### GetA2aServerOk

`func (o *IdentityAssertionA2aServerConnection) GetA2aServerOk() (*ResourceConnectionA2aServer, bool)`

GetA2aServerOk returns a tuple with the A2aServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2aServer

`func (o *IdentityAssertionA2aServerConnection) SetA2aServer(v ResourceConnectionA2aServer)`

SetA2aServer sets A2aServer field to given value.


### GetAuthorizationServer

`func (o *IdentityAssertionA2aServerConnection) GetAuthorizationServer() CustomAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *IdentityAssertionA2aServerConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *IdentityAssertionA2aServerConnection) SetAuthorizationServer(v CustomAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetId

`func (o *IdentityAssertionA2aServerConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAssertionA2aServerConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAssertionA2aServerConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAssertionA2aServerConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *IdentityAssertionA2aServerConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *IdentityAssertionA2aServerConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *IdentityAssertionA2aServerConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *IdentityAssertionA2aServerConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityAssertionA2aServerConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityAssertionA2aServerConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityAssertionA2aServerConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityAssertionA2aServerConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


