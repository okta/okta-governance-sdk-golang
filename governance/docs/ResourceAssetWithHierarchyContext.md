# ResourceAssetWithHierarchyContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | Pointer to [**ResourceAssetLinks**](ResourceAssetLinks.md) |  | [optional] 
**Orn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**Name** | **string** | The display name for a resource asset | 
**Description** | Pointer to **string** | The description of a resource asset | [optional] 
**Type** | [**ResourceAssetTypeSparse**](ResourceAssetTypeSparse.md) |  | 
**ExternalId** | **string** | The external ID of a resource asset | 
**HasChildren** | Pointer to **bool** | Whether this asset has any children in the hierarchy. Present only if ?include&#x3D;children_info is specified | [optional] 
**ParentId** | Pointer to **NullableString** | The ID of the direct parent asset, or null if this is a root asset. Present only if ?include&#x3D;ancestors is specified | [optional] 

## Methods

### NewResourceAssetWithHierarchyContext

`func NewResourceAssetWithHierarchyContext(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, orn string, name string, type_ ResourceAssetTypeSparse, externalId string, ) *ResourceAssetWithHierarchyContext`

NewResourceAssetWithHierarchyContext instantiates a new ResourceAssetWithHierarchyContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAssetWithHierarchyContextWithDefaults

`func NewResourceAssetWithHierarchyContextWithDefaults() *ResourceAssetWithHierarchyContext`

NewResourceAssetWithHierarchyContextWithDefaults instantiates a new ResourceAssetWithHierarchyContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceAssetWithHierarchyContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceAssetWithHierarchyContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceAssetWithHierarchyContext) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ResourceAssetWithHierarchyContext) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ResourceAssetWithHierarchyContext) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ResourceAssetWithHierarchyContext) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *ResourceAssetWithHierarchyContext) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResourceAssetWithHierarchyContext) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResourceAssetWithHierarchyContext) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *ResourceAssetWithHierarchyContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ResourceAssetWithHierarchyContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ResourceAssetWithHierarchyContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *ResourceAssetWithHierarchyContext) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ResourceAssetWithHierarchyContext) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ResourceAssetWithHierarchyContext) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *ResourceAssetWithHierarchyContext) GetLinks() ResourceAssetLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceAssetWithHierarchyContext) GetLinksOk() (*ResourceAssetLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceAssetWithHierarchyContext) SetLinks(v ResourceAssetLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceAssetWithHierarchyContext) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrn

`func (o *ResourceAssetWithHierarchyContext) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceAssetWithHierarchyContext) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceAssetWithHierarchyContext) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *ResourceAssetWithHierarchyContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAssetWithHierarchyContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAssetWithHierarchyContext) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ResourceAssetWithHierarchyContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceAssetWithHierarchyContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceAssetWithHierarchyContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceAssetWithHierarchyContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ResourceAssetWithHierarchyContext) GetType() ResourceAssetTypeSparse`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceAssetWithHierarchyContext) GetTypeOk() (*ResourceAssetTypeSparse, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceAssetWithHierarchyContext) SetType(v ResourceAssetTypeSparse)`

SetType sets Type field to given value.


### GetExternalId

`func (o *ResourceAssetWithHierarchyContext) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ResourceAssetWithHierarchyContext) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ResourceAssetWithHierarchyContext) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetHasChildren

`func (o *ResourceAssetWithHierarchyContext) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *ResourceAssetWithHierarchyContext) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *ResourceAssetWithHierarchyContext) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *ResourceAssetWithHierarchyContext) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetParentId

`func (o *ResourceAssetWithHierarchyContext) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ResourceAssetWithHierarchyContext) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ResourceAssetWithHierarchyContext) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ResourceAssetWithHierarchyContext) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ResourceAssetWithHierarchyContext) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ResourceAssetWithHierarchyContext) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


