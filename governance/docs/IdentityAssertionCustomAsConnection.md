# IdentityAssertionCustomAsConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**AuthorizationServer** | [**CustomAuthorizationServer**](CustomAuthorizationServer.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the managed connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewIdentityAssertionCustomAsConnection

`func NewIdentityAssertionCustomAsConnection(connectionType string, authorizationServer CustomAuthorizationServer, ) *IdentityAssertionCustomAsConnection`

NewIdentityAssertionCustomAsConnection instantiates a new IdentityAssertionCustomAsConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionCustomAsConnectionWithDefaults

`func NewIdentityAssertionCustomAsConnectionWithDefaults() *IdentityAssertionCustomAsConnection`

NewIdentityAssertionCustomAsConnectionWithDefaults instantiates a new IdentityAssertionCustomAsConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *IdentityAssertionCustomAsConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IdentityAssertionCustomAsConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IdentityAssertionCustomAsConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetAuthorizationServer

`func (o *IdentityAssertionCustomAsConnection) GetAuthorizationServer() CustomAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *IdentityAssertionCustomAsConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *IdentityAssertionCustomAsConnection) SetAuthorizationServer(v CustomAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetId

`func (o *IdentityAssertionCustomAsConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAssertionCustomAsConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAssertionCustomAsConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAssertionCustomAsConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *IdentityAssertionCustomAsConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *IdentityAssertionCustomAsConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *IdentityAssertionCustomAsConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *IdentityAssertionCustomAsConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityAssertionCustomAsConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityAssertionCustomAsConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityAssertionCustomAsConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityAssertionCustomAsConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


