# ResourceCatalogVisibilityTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ResourceCatalogVisibilityType**](ResourceCatalogVisibilityType.md) |  | 
**Id** | **string** | The ID for the given target type (for example, a group ID for &#x60;GROUP&#x60;) | 

## Methods

### NewResourceCatalogVisibilityTarget

`func NewResourceCatalogVisibilityTarget(type_ ResourceCatalogVisibilityType, id string, ) *ResourceCatalogVisibilityTarget`

NewResourceCatalogVisibilityTarget instantiates a new ResourceCatalogVisibilityTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCatalogVisibilityTargetWithDefaults

`func NewResourceCatalogVisibilityTargetWithDefaults() *ResourceCatalogVisibilityTarget`

NewResourceCatalogVisibilityTargetWithDefaults instantiates a new ResourceCatalogVisibilityTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceCatalogVisibilityTarget) GetType() ResourceCatalogVisibilityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceCatalogVisibilityTarget) GetTypeOk() (*ResourceCatalogVisibilityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceCatalogVisibilityTarget) SetType(v ResourceCatalogVisibilityType)`

SetType sets Type field to given value.


### GetId

`func (o *ResourceCatalogVisibilityTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceCatalogVisibilityTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceCatalogVisibilityTarget) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


