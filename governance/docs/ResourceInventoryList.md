# ResourceInventoryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Resource**](Resource.md) | List of resources on the current page | 
**Links** | [**ListLinks**](ListLinks.md) |  | 
**Metadata** | [**ListMetadata**](ListMetadata.md) |  | 

## Methods

### NewResourceInventoryList

`func NewResourceInventoryList(data []Resource, links ListLinks, metadata ListMetadata, ) *ResourceInventoryList`

NewResourceInventoryList instantiates a new ResourceInventoryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInventoryListWithDefaults

`func NewResourceInventoryListWithDefaults() *ResourceInventoryList`

NewResourceInventoryListWithDefaults instantiates a new ResourceInventoryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResourceInventoryList) GetData() []Resource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceInventoryList) GetDataOk() (*[]Resource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceInventoryList) SetData(v []Resource)`

SetData sets Data field to given value.


### GetLinks

`func (o *ResourceInventoryList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceInventoryList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceInventoryList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *ResourceInventoryList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceInventoryList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceInventoryList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


