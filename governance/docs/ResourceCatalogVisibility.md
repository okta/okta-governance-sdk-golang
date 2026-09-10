# ResourceCatalogVisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | **bool** | Indicates whether users can access the resource catalog:   * If &#x60;false&#x60;, no users can access the resource catalog.   * If &#x60;true&#x60; and &#x60;onlyFor&#x60; isn&#39;t specified, all users in the org can access the resource catalog. | 
**OnlyFor** | Pointer to [**[]ResourceCatalogVisibilityTarget**](ResourceCatalogVisibilityTarget.md) | Specific user targets for resource catalog visibility: * If this array is specified, only the specified targets can access the resource catalog. * If this array is null and &#x60;visible&#x60; is &#x60;true&#x60;, all users in the org can access the resource catalog. | [optional] 

## Methods

### NewResourceCatalogVisibility

`func NewResourceCatalogVisibility(visible bool, ) *ResourceCatalogVisibility`

NewResourceCatalogVisibility instantiates a new ResourceCatalogVisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCatalogVisibilityWithDefaults

`func NewResourceCatalogVisibilityWithDefaults() *ResourceCatalogVisibility`

NewResourceCatalogVisibilityWithDefaults instantiates a new ResourceCatalogVisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *ResourceCatalogVisibility) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ResourceCatalogVisibility) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ResourceCatalogVisibility) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetOnlyFor

`func (o *ResourceCatalogVisibility) GetOnlyFor() []ResourceCatalogVisibilityTarget`

GetOnlyFor returns the OnlyFor field if non-nil, zero value otherwise.

### GetOnlyForOk

`func (o *ResourceCatalogVisibility) GetOnlyForOk() (*[]ResourceCatalogVisibilityTarget, bool)`

GetOnlyForOk returns a tuple with the OnlyFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFor

`func (o *ResourceCatalogVisibility) SetOnlyFor(v []ResourceCatalogVisibilityTarget)`

SetOnlyFor sets OnlyFor field to given value.

### HasOnlyFor

`func (o *ResourceCatalogVisibility) HasOnlyFor() bool`

HasOnlyFor returns a boolean if a field has been set.

### SetOnlyForNil

`func (o *ResourceCatalogVisibility) SetOnlyForNil(b bool)`

 SetOnlyForNil sets the value for OnlyFor to be an explicit nil

### UnsetOnlyFor
`func (o *ResourceCatalogVisibility) UnsetOnlyFor()`

UnsetOnlyFor ensures that no value is present for OnlyFor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


