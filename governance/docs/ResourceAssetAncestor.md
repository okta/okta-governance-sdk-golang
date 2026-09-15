# ResourceAssetAncestor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the asset | 
**Orn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**Name** | **string** | The display name for a resource asset | 
**Type** | [**ResourceAssetTypeSparse**](ResourceAssetTypeSparse.md) |  | 
**ParentId** | Pointer to **NullableString** | The ID of this ancestor&#39;s direct parent, or null if it is a root asset | [optional] 
**Links** | Pointer to [**ResourceAssetLinks**](ResourceAssetLinks.md) |  | [optional] 

## Methods

### NewResourceAssetAncestor

`func NewResourceAssetAncestor(id string, orn string, name string, type_ ResourceAssetTypeSparse, ) *ResourceAssetAncestor`

NewResourceAssetAncestor instantiates a new ResourceAssetAncestor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAssetAncestorWithDefaults

`func NewResourceAssetAncestorWithDefaults() *ResourceAssetAncestor`

NewResourceAssetAncestorWithDefaults instantiates a new ResourceAssetAncestor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceAssetAncestor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceAssetAncestor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceAssetAncestor) SetId(v string)`

SetId sets Id field to given value.


### GetOrn

`func (o *ResourceAssetAncestor) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceAssetAncestor) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceAssetAncestor) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *ResourceAssetAncestor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAssetAncestor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAssetAncestor) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResourceAssetAncestor) GetType() ResourceAssetTypeSparse`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceAssetAncestor) GetTypeOk() (*ResourceAssetTypeSparse, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceAssetAncestor) SetType(v ResourceAssetTypeSparse)`

SetType sets Type field to given value.


### GetParentId

`func (o *ResourceAssetAncestor) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ResourceAssetAncestor) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ResourceAssetAncestor) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ResourceAssetAncestor) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ResourceAssetAncestor) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ResourceAssetAncestor) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetLinks

`func (o *ResourceAssetAncestor) GetLinks() ResourceAssetLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceAssetAncestor) GetLinksOk() (*ResourceAssetLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceAssetAncestor) SetLinks(v ResourceAssetLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceAssetAncestor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


