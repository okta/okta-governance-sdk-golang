# StsAccessTokenResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of resource for this &#x60;STS_ACCESS_TOKEN&#x60; connection | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the client auth settings | 
**AppInstanceName** | **string** | Display name of the associated app instance | 
**Name** | **string** | Display name of the third-party API server | 

## Methods

### NewStsAccessTokenResource

`func NewStsAccessTokenResource(resourceType string, orn string, appInstanceName string, name string, ) *StsAccessTokenResource`

NewStsAccessTokenResource instantiates a new StsAccessTokenResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsAccessTokenResourceWithDefaults

`func NewStsAccessTokenResourceWithDefaults() *StsAccessTokenResource`

NewStsAccessTokenResourceWithDefaults instantiates a new StsAccessTokenResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *StsAccessTokenResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *StsAccessTokenResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *StsAccessTokenResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetOrn

`func (o *StsAccessTokenResource) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *StsAccessTokenResource) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *StsAccessTokenResource) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetAppInstanceName

`func (o *StsAccessTokenResource) GetAppInstanceName() string`

GetAppInstanceName returns the AppInstanceName field if non-nil, zero value otherwise.

### GetAppInstanceNameOk

`func (o *StsAccessTokenResource) GetAppInstanceNameOk() (*string, bool)`

GetAppInstanceNameOk returns a tuple with the AppInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceName

`func (o *StsAccessTokenResource) SetAppInstanceName(v string)`

SetAppInstanceName sets AppInstanceName field to given value.


### GetName

`func (o *StsAccessTokenResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StsAccessTokenResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StsAccessTokenResource) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


