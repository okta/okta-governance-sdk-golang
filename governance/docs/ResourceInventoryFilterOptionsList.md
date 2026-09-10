# ResourceInventoryFilterOptionsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourceInventoryFilterOption**](ResourceInventoryFilterOption.md) | List of filtering options on the current page | 
**Links** | [**ListLinks**](ListLinks.md) |  | 
**Metadata** | [**ListMetadata**](ListMetadata.md) |  | 

## Methods

### NewResourceInventoryFilterOptionsList

`func NewResourceInventoryFilterOptionsList(data []ResourceInventoryFilterOption, links ListLinks, metadata ListMetadata, ) *ResourceInventoryFilterOptionsList`

NewResourceInventoryFilterOptionsList instantiates a new ResourceInventoryFilterOptionsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInventoryFilterOptionsListWithDefaults

`func NewResourceInventoryFilterOptionsListWithDefaults() *ResourceInventoryFilterOptionsList`

NewResourceInventoryFilterOptionsListWithDefaults instantiates a new ResourceInventoryFilterOptionsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResourceInventoryFilterOptionsList) GetData() []ResourceInventoryFilterOption`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceInventoryFilterOptionsList) GetDataOk() (*[]ResourceInventoryFilterOption, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceInventoryFilterOptionsList) SetData(v []ResourceInventoryFilterOption)`

SetData sets Data field to given value.


### GetLinks

`func (o *ResourceInventoryFilterOptionsList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceInventoryFilterOptionsList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceInventoryFilterOptionsList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *ResourceInventoryFilterOptionsList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceInventoryFilterOptionsList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceInventoryFilterOptionsList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


