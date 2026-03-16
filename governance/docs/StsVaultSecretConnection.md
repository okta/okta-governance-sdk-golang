# StsVaultSecretConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**Secret** | [**ManagedConnectionVaultedSecret**](ManagedConnectionVaultedSecret.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the managed connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewStsVaultSecretConnection

`func NewStsVaultSecretConnection(connectionType string, secret ManagedConnectionVaultedSecret, ) *StsVaultSecretConnection`

NewStsVaultSecretConnection instantiates a new StsVaultSecretConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsVaultSecretConnectionWithDefaults

`func NewStsVaultSecretConnectionWithDefaults() *StsVaultSecretConnection`

NewStsVaultSecretConnectionWithDefaults instantiates a new StsVaultSecretConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *StsVaultSecretConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *StsVaultSecretConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *StsVaultSecretConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetSecret

`func (o *StsVaultSecretConnection) GetSecret() ManagedConnectionVaultedSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *StsVaultSecretConnection) GetSecretOk() (*ManagedConnectionVaultedSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *StsVaultSecretConnection) SetSecret(v ManagedConnectionVaultedSecret)`

SetSecret sets Secret field to given value.


### GetId

`func (o *StsVaultSecretConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StsVaultSecretConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StsVaultSecretConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StsVaultSecretConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *StsVaultSecretConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsVaultSecretConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsVaultSecretConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *StsVaultSecretConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *StsVaultSecretConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StsVaultSecretConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StsVaultSecretConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StsVaultSecretConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


