# CollectionResourcesListV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CollectionResourceFullV2**](CollectionResourceFullV2.md) |  | [optional] 
**Links** | Pointer to [**ListLinks**](ListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewCollectionResourcesListV2

`func NewCollectionResourcesListV2() *CollectionResourcesListV2`

NewCollectionResourcesListV2 instantiates a new CollectionResourcesListV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourcesListV2WithDefaults

`func NewCollectionResourcesListV2WithDefaults() *CollectionResourcesListV2`

NewCollectionResourcesListV2WithDefaults instantiates a new CollectionResourcesListV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CollectionResourcesListV2) GetData() []CollectionResourceFullV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CollectionResourcesListV2) GetDataOk() (*[]CollectionResourceFullV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CollectionResourcesListV2) SetData(v []CollectionResourceFullV2)`

SetData sets Data field to given value.

### HasData

`func (o *CollectionResourcesListV2) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourcesListV2) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourcesListV2) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourcesListV2) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CollectionResourcesListV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *CollectionResourcesListV2) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CollectionResourcesListV2) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CollectionResourcesListV2) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CollectionResourcesListV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


