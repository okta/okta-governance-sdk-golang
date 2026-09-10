# ResourceAssetsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourceAssetWithHierarchyContext**](ResourceAssetWithHierarchyContext.md) | List of all resource assets matching the filter | 
**Ancestors** | Pointer to [**map[string]ResourceAssetAncestor**](ResourceAssetAncestor.md) | A map of all unique ancestor assets for items in &#x60;data&#x60;, keyed by asset ID. Only present if ?include&#x3D;ancestors is specified. Each ancestor also includes its own &#x60;parentId&#x60; to allow full chain traversal to the root.  | [optional] 
**Links** | [**ResourceAssetsLinks**](ResourceAssetsLinks.md) |  | 

## Methods

### NewResourceAssetsList

`func NewResourceAssetsList(data []ResourceAssetWithHierarchyContext, links ResourceAssetsLinks, ) *ResourceAssetsList`

NewResourceAssetsList instantiates a new ResourceAssetsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAssetsListWithDefaults

`func NewResourceAssetsListWithDefaults() *ResourceAssetsList`

NewResourceAssetsListWithDefaults instantiates a new ResourceAssetsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResourceAssetsList) GetData() []ResourceAssetWithHierarchyContext`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceAssetsList) GetDataOk() (*[]ResourceAssetWithHierarchyContext, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceAssetsList) SetData(v []ResourceAssetWithHierarchyContext)`

SetData sets Data field to given value.


### GetAncestors

`func (o *ResourceAssetsList) GetAncestors() map[string]ResourceAssetAncestor`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ResourceAssetsList) GetAncestorsOk() (*map[string]ResourceAssetAncestor, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ResourceAssetsList) SetAncestors(v map[string]ResourceAssetAncestor)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ResourceAssetsList) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceAssetsList) GetLinks() ResourceAssetsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceAssetsList) GetLinksOk() (*ResourceAssetsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceAssetsList) SetLinks(v ResourceAssetsLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


