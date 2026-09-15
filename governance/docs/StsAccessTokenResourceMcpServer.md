# StsAccessTokenResourceMcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of resource for this &#x60;STS_ACCESS_TOKEN&#x60; connection | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the client auth settings | 
**Name** | **string** | Display name of the third-party MCP server | 

## Methods

### NewStsAccessTokenResourceMcpServer

`func NewStsAccessTokenResourceMcpServer(resourceType string, orn string, name string, ) *StsAccessTokenResourceMcpServer`

NewStsAccessTokenResourceMcpServer instantiates a new StsAccessTokenResourceMcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsAccessTokenResourceMcpServerWithDefaults

`func NewStsAccessTokenResourceMcpServerWithDefaults() *StsAccessTokenResourceMcpServer`

NewStsAccessTokenResourceMcpServerWithDefaults instantiates a new StsAccessTokenResourceMcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *StsAccessTokenResourceMcpServer) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *StsAccessTokenResourceMcpServer) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *StsAccessTokenResourceMcpServer) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetOrn

`func (o *StsAccessTokenResourceMcpServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsAccessTokenResourceMcpServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsAccessTokenResourceMcpServer) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *StsAccessTokenResourceMcpServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StsAccessTokenResourceMcpServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StsAccessTokenResourceMcpServer) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


