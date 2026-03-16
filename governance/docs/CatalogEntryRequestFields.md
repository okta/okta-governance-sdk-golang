# CatalogEntryRequestFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RequestField**](RequestField.md) | List of request fields | [optional] 
**Metadata** | Pointer to [**CatalogEntryRequestFieldsMetadata**](CatalogEntryRequestFieldsMetadata.md) |  | [optional] 

## Methods

### NewCatalogEntryRequestFields

`func NewCatalogEntryRequestFields() *CatalogEntryRequestFields`

NewCatalogEntryRequestFields instantiates a new CatalogEntryRequestFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogEntryRequestFieldsWithDefaults

`func NewCatalogEntryRequestFieldsWithDefaults() *CatalogEntryRequestFields`

NewCatalogEntryRequestFieldsWithDefaults instantiates a new CatalogEntryRequestFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CatalogEntryRequestFields) GetData() []RequestField`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogEntryRequestFields) GetDataOk() (*[]RequestField, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogEntryRequestFields) SetData(v []RequestField)`

SetData sets Data field to given value.

### HasData

`func (o *CatalogEntryRequestFields) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogEntryRequestFields) GetMetadata() CatalogEntryRequestFieldsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogEntryRequestFields) GetMetadataOk() (*CatalogEntryRequestFieldsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogEntryRequestFields) SetMetadata(v CatalogEntryRequestFieldsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogEntryRequestFields) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


