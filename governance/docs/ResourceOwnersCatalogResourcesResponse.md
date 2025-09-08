# ResourceOwnersCatalogResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Data** | Pointer to [**[]ResourceOwnerResource**](ResourceOwnerResource.md) | Resource owner details. | [optional] 
**Links** | [**ResourceOwnersListLinks**](ResourceOwnersListLinks.md) |  | 

## Methods

### NewResourceOwnersCatalogResourcesResponse

`func NewResourceOwnersCatalogResourcesResponse(parentResourceOrn string, links ResourceOwnersListLinks, ) *ResourceOwnersCatalogResourcesResponse`

NewResourceOwnersCatalogResourcesResponse instantiates a new ResourceOwnersCatalogResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnersCatalogResourcesResponseWithDefaults

`func NewResourceOwnersCatalogResourcesResponseWithDefaults() *ResourceOwnersCatalogResourcesResponse`

NewResourceOwnersCatalogResourcesResponseWithDefaults instantiates a new ResourceOwnersCatalogResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentResourceOrn

`func (o *ResourceOwnersCatalogResourcesResponse) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *ResourceOwnersCatalogResourcesResponse) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *ResourceOwnersCatalogResourcesResponse) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.


### GetData

`func (o *ResourceOwnersCatalogResourcesResponse) GetData() []ResourceOwnerResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceOwnersCatalogResourcesResponse) GetDataOk() (*[]ResourceOwnerResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceOwnersCatalogResourcesResponse) SetData(v []ResourceOwnerResource)`

SetData sets Data field to given value.

### HasData

`func (o *ResourceOwnersCatalogResourcesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceOwnersCatalogResourcesResponse) GetLinks() ResourceOwnersListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceOwnersCatalogResourcesResponse) GetLinksOk() (*ResourceOwnersListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceOwnersCatalogResourcesResponse) SetLinks(v ResourceOwnersListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


