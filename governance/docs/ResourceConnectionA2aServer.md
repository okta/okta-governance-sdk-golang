# ResourceConnectionA2aServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the A2A server | 
**Name** | **string** | Display name of the agent-to-agent (A2A) server | 

## Methods

### NewResourceConnectionA2aServer

`func NewResourceConnectionA2aServer(orn string, name string, ) *ResourceConnectionA2aServer`

NewResourceConnectionA2aServer instantiates a new ResourceConnectionA2aServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConnectionA2aServerWithDefaults

`func NewResourceConnectionA2aServerWithDefaults() *ResourceConnectionA2aServer`

NewResourceConnectionA2aServerWithDefaults instantiates a new ResourceConnectionA2aServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *ResourceConnectionA2aServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceConnectionA2aServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceConnectionA2aServer) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *ResourceConnectionA2aServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceConnectionA2aServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceConnectionA2aServer) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


