# ResourceOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Principals** | Pointer to [**[]ResourceOwnerPrincipal**](ResourceOwnerPrincipal.md) |  | [optional] 
**Resource** | [**ResourceOwnerResource**](ResourceOwnerResource.md) |  | 

## Methods

### NewResourceOwner

`func NewResourceOwner(parentResourceOrn string, resource ResourceOwnerResource, ) *ResourceOwner`

NewResourceOwner instantiates a new ResourceOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnerWithDefaults

`func NewResourceOwnerWithDefaults() *ResourceOwner`

NewResourceOwnerWithDefaults instantiates a new ResourceOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentResourceOrn

`func (o *ResourceOwner) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *ResourceOwner) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *ResourceOwner) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.


### GetPrincipals

`func (o *ResourceOwner) GetPrincipals() []ResourceOwnerPrincipal`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *ResourceOwner) GetPrincipalsOk() (*[]ResourceOwnerPrincipal, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *ResourceOwner) SetPrincipals(v []ResourceOwnerPrincipal)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *ResourceOwner) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetResource

`func (o *ResourceOwner) GetResource() ResourceOwnerResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceOwner) GetResourceOk() (*ResourceOwnerResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceOwner) SetResource(v ResourceOwnerResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


