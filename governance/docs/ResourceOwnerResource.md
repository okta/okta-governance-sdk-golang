# ResourceOwnerResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the resource, which is a unique identifier for the resource in Okta. | 
**Type** | **string** | The resource type value from the orn. Examples:- apps, entitlement-bundles | 
**Orn** | **string** | The &#x60;id&#x60; of the resource in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an app, or a bundle. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Profile** | [**ExternalResourceProfile**](ExternalResourceProfile.md) |  | 

## Methods

### NewResourceOwnerResource

`func NewResourceOwnerResource(id string, type_ string, orn string, profile ExternalResourceProfile, ) *ResourceOwnerResource`

NewResourceOwnerResource instantiates a new ResourceOwnerResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnerResourceWithDefaults

`func NewResourceOwnerResourceWithDefaults() *ResourceOwnerResource`

NewResourceOwnerResourceWithDefaults instantiates a new ResourceOwnerResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceOwnerResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceOwnerResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceOwnerResource) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ResourceOwnerResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceOwnerResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceOwnerResource) SetType(v string)`

SetType sets Type field to given value.


### GetOrn

`func (o *ResourceOwnerResource) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceOwnerResource) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceOwnerResource) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetProfile

`func (o *ResourceOwnerResource) GetProfile() ExternalResourceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResourceOwnerResource) GetProfileOk() (*ExternalResourceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResourceOwnerResource) SetProfile(v ExternalResourceProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


