# StsAccessTokenResourceAppInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of resource for this &#x60;STS_ACCESS_TOKEN&#x60; connection | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the client auth settings | 
**AppInstanceName** | **string** | Display name of the associated app instance | 

## Methods

### NewStsAccessTokenResourceAppInstance

`func NewStsAccessTokenResourceAppInstance(resourceType string, orn string, appInstanceName string, ) *StsAccessTokenResourceAppInstance`

NewStsAccessTokenResourceAppInstance instantiates a new StsAccessTokenResourceAppInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsAccessTokenResourceAppInstanceWithDefaults

`func NewStsAccessTokenResourceAppInstanceWithDefaults() *StsAccessTokenResourceAppInstance`

NewStsAccessTokenResourceAppInstanceWithDefaults instantiates a new StsAccessTokenResourceAppInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *StsAccessTokenResourceAppInstance) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *StsAccessTokenResourceAppInstance) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *StsAccessTokenResourceAppInstance) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetOrn

`func (o *StsAccessTokenResourceAppInstance) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsAccessTokenResourceAppInstance) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsAccessTokenResourceAppInstance) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetAppInstanceName

`func (o *StsAccessTokenResourceAppInstance) GetAppInstanceName() string`

GetAppInstanceName returns the AppInstanceName field if non-nil, zero value otherwise.

### GetAppInstanceNameOk

`func (o *StsAccessTokenResourceAppInstance) GetAppInstanceNameOk() (*string, bool)`

GetAppInstanceNameOk returns a tuple with the AppInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceName

`func (o *StsAccessTokenResourceAppInstance) SetAppInstanceName(v string)`

SetAppInstanceName sets AppInstanceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


