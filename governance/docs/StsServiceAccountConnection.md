# StsServiceAccountConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**App** | [**ManagedConnectionAppInstance**](ManagedConnectionAppInstance.md) |  | 
**ServiceAccount** | [**ManagedConnectionServiceAccount**](ManagedConnectionServiceAccount.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the managed connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewStsServiceAccountConnection

`func NewStsServiceAccountConnection(connectionType string, app ManagedConnectionAppInstance, serviceAccount ManagedConnectionServiceAccount, ) *StsServiceAccountConnection`

NewStsServiceAccountConnection instantiates a new StsServiceAccountConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsServiceAccountConnectionWithDefaults

`func NewStsServiceAccountConnectionWithDefaults() *StsServiceAccountConnection`

NewStsServiceAccountConnectionWithDefaults instantiates a new StsServiceAccountConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *StsServiceAccountConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *StsServiceAccountConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *StsServiceAccountConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetApp

`func (o *StsServiceAccountConnection) GetApp() ManagedConnectionAppInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *StsServiceAccountConnection) GetAppOk() (*ManagedConnectionAppInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *StsServiceAccountConnection) SetApp(v ManagedConnectionAppInstance)`

SetApp sets App field to given value.


### GetServiceAccount

`func (o *StsServiceAccountConnection) GetServiceAccount() ManagedConnectionServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *StsServiceAccountConnection) GetServiceAccountOk() (*ManagedConnectionServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *StsServiceAccountConnection) SetServiceAccount(v ManagedConnectionServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetId

`func (o *StsServiceAccountConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StsServiceAccountConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StsServiceAccountConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StsServiceAccountConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *StsServiceAccountConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsServiceAccountConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsServiceAccountConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *StsServiceAccountConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *StsServiceAccountConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StsServiceAccountConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StsServiceAccountConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StsServiceAccountConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


