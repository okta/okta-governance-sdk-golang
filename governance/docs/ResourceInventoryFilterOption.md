# ResourceInventoryFilterOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the filter option. For owners, this is the Okta user ID. | 
**Name** | **string** | Display name of the filter option | 
**Logo** | Pointer to **string** | URL to a logo image. Only present for app filter options. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata properties for the filter option. Only present for certain filter types.  For the &#x60;labels&#x60; filter type, metadata contains: * &#x60;values&#x60;: array of label values, each with &#x60;id&#x60;, &#x60;name&#x60;, and &#x60;backgroundColor&#x60;   * &#x60;backgroundColor&#x60;: [&#x60;red&#x60;, &#x60;orange&#x60;, &#x60;yellow&#x60;, &#x60;green&#x60;, &#x60;blue&#x60;, &#x60;purple&#x60;, &#x60;teal&#x60;, &#x60;beige&#x60;, &#x60;gray&#x60;]  | [optional] 

## Methods

### NewResourceInventoryFilterOption

`func NewResourceInventoryFilterOption(id string, name string, ) *ResourceInventoryFilterOption`

NewResourceInventoryFilterOption instantiates a new ResourceInventoryFilterOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInventoryFilterOptionWithDefaults

`func NewResourceInventoryFilterOptionWithDefaults() *ResourceInventoryFilterOption`

NewResourceInventoryFilterOptionWithDefaults instantiates a new ResourceInventoryFilterOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceInventoryFilterOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceInventoryFilterOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceInventoryFilterOption) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ResourceInventoryFilterOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceInventoryFilterOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceInventoryFilterOption) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *ResourceInventoryFilterOption) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ResourceInventoryFilterOption) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ResourceInventoryFilterOption) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ResourceInventoryFilterOption) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourceInventoryFilterOption) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceInventoryFilterOption) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceInventoryFilterOption) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceInventoryFilterOption) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


