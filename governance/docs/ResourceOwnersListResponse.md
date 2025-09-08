# ResourceOwnersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResourceOwner**](ResourceOwner.md) | Resource owner details. | [optional] 
**Links** | Pointer to [**ResourceOwnersListLinks**](ResourceOwnersListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewResourceOwnersListResponse

`func NewResourceOwnersListResponse() *ResourceOwnersListResponse`

NewResourceOwnersListResponse instantiates a new ResourceOwnersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnersListResponseWithDefaults

`func NewResourceOwnersListResponseWithDefaults() *ResourceOwnersListResponse`

NewResourceOwnersListResponseWithDefaults instantiates a new ResourceOwnersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResourceOwnersListResponse) GetData() []ResourceOwner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceOwnersListResponse) GetDataOk() (*[]ResourceOwner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceOwnersListResponse) SetData(v []ResourceOwner)`

SetData sets Data field to given value.

### HasData

`func (o *ResourceOwnersListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceOwnersListResponse) GetLinks() ResourceOwnersListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceOwnersListResponse) GetLinksOk() (*ResourceOwnersListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceOwnersListResponse) SetLinks(v ResourceOwnersListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceOwnersListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourceOwnersListResponse) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceOwnersListResponse) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceOwnersListResponse) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceOwnersListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


