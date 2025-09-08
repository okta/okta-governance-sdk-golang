# GrantDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Grant &#x60;id&#x60; | [optional] 
**Metadata** | Pointer to [**GrantMetadata**](GrantMetadata.md) |  | [optional] 
**Links** | Pointer to [**GrantDetailsLinks**](GrantDetailsLinks.md) |  | [optional] 

## Methods

### NewGrantDetails

`func NewGrantDetails() *GrantDetails`

NewGrantDetails instantiates a new GrantDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantDetailsWithDefaults

`func NewGrantDetailsWithDefaults() *GrantDetails`

NewGrantDetailsWithDefaults instantiates a new GrantDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GrantDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GrantDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *GrantDetails) GetMetadata() GrantMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GrantDetails) GetMetadataOk() (*GrantMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GrantDetails) SetMetadata(v GrantMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GrantDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLinks

`func (o *GrantDetails) GetLinks() GrantDetailsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GrantDetails) GetLinksOk() (*GrantDetailsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GrantDetails) SetLinks(v GrantDetailsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GrantDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


