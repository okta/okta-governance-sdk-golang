# StsAccessTokenConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**Resource** | [**StsAccessTokenResource**](StsAccessTokenResource.md) |  | 
**Id** | Pointer to **string** | Unique identifier for the resource connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewStsAccessTokenConnection

`func NewStsAccessTokenConnection(connectionType string, resource StsAccessTokenResource, ) *StsAccessTokenConnection`

NewStsAccessTokenConnection instantiates a new StsAccessTokenConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsAccessTokenConnectionWithDefaults

`func NewStsAccessTokenConnectionWithDefaults() *StsAccessTokenConnection`

NewStsAccessTokenConnectionWithDefaults instantiates a new StsAccessTokenConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *StsAccessTokenConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *StsAccessTokenConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *StsAccessTokenConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetResource

`func (o *StsAccessTokenConnection) GetResource() StsAccessTokenResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *StsAccessTokenConnection) GetResourceOk() (*StsAccessTokenResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *StsAccessTokenConnection) SetResource(v StsAccessTokenResource)`

SetResource sets Resource field to given value.


### GetId

`func (o *StsAccessTokenConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StsAccessTokenConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StsAccessTokenConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StsAccessTokenConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *StsAccessTokenConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsAccessTokenConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsAccessTokenConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *StsAccessTokenConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *StsAccessTokenConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StsAccessTokenConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StsAccessTokenConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StsAccessTokenConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


