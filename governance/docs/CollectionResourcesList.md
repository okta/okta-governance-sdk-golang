# CollectionResourcesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CollectionResourceFull**](CollectionResourceFull.md) |  | [optional] 
**Links** | Pointer to [**ListLinks**](ListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewCollectionResourcesList

`func NewCollectionResourcesList() *CollectionResourcesList`

NewCollectionResourcesList instantiates a new CollectionResourcesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourcesListWithDefaults

`func NewCollectionResourcesListWithDefaults() *CollectionResourcesList`

NewCollectionResourcesListWithDefaults instantiates a new CollectionResourcesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CollectionResourcesList) GetData() []CollectionResourceFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CollectionResourcesList) GetDataOk() (*[]CollectionResourceFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CollectionResourcesList) SetData(v []CollectionResourceFull)`

SetData sets Data field to given value.

### HasData

`func (o *CollectionResourcesList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourcesList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourcesList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourcesList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CollectionResourcesList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *CollectionResourcesList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CollectionResourcesList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CollectionResourcesList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CollectionResourcesList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


