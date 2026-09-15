# CollectionResourcePropertiesReadOnlyV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as &#x60;GET&#x60;/&#x60;PUT&#x60;/&#x60;DELETE /v2/collections/{collectionId}/resources/{resourceId}&#x60;.  | [optional] 
**Links** | Pointer to [**CollectionResourceLinksV2**](CollectionResourceLinksV2.md) |  | [optional] 

## Methods

### NewCollectionResourcePropertiesReadOnlyV2

`func NewCollectionResourcePropertiesReadOnlyV2() *CollectionResourcePropertiesReadOnlyV2`

NewCollectionResourcePropertiesReadOnlyV2 instantiates a new CollectionResourcePropertiesReadOnlyV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourcePropertiesReadOnlyV2WithDefaults

`func NewCollectionResourcePropertiesReadOnlyV2WithDefaults() *CollectionResourcePropertiesReadOnlyV2`

NewCollectionResourcePropertiesReadOnlyV2WithDefaults instantiates a new CollectionResourcePropertiesReadOnlyV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *CollectionResourcePropertiesReadOnlyV2) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionResourcePropertiesReadOnlyV2) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionResourcePropertiesReadOnlyV2) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionResourcePropertiesReadOnlyV2) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourcePropertiesReadOnlyV2) GetLinks() CollectionResourceLinksV2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourcePropertiesReadOnlyV2) GetLinksOk() (*CollectionResourceLinksV2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourcePropertiesReadOnlyV2) SetLinks(v CollectionResourceLinksV2)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CollectionResourcePropertiesReadOnlyV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


