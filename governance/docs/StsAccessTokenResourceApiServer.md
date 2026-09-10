# StsAccessTokenResourceApiServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of resource for this &#x60;STS_ACCESS_TOKEN&#x60; connection | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the client auth settings | 
**Name** | **string** | Display name of the third-party API server | 

## Methods

### NewStsAccessTokenResourceApiServer

`func NewStsAccessTokenResourceApiServer(resourceType string, orn string, name string, ) *StsAccessTokenResourceApiServer`

NewStsAccessTokenResourceApiServer instantiates a new StsAccessTokenResourceApiServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsAccessTokenResourceApiServerWithDefaults

`func NewStsAccessTokenResourceApiServerWithDefaults() *StsAccessTokenResourceApiServer`

NewStsAccessTokenResourceApiServerWithDefaults instantiates a new StsAccessTokenResourceApiServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *StsAccessTokenResourceApiServer) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *StsAccessTokenResourceApiServer) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *StsAccessTokenResourceApiServer) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetOrn

`func (o *StsAccessTokenResourceApiServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsAccessTokenResourceApiServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsAccessTokenResourceApiServer) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *StsAccessTokenResourceApiServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StsAccessTokenResourceApiServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StsAccessTokenResourceApiServer) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


