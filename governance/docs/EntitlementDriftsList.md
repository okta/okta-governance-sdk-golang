# EntitlementDriftsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EntitlementDriftFull**](EntitlementDriftFull.md) | The drifts on this page | [optional] 
**Links** | Pointer to [**EntitlementDriftListLinks**](EntitlementDriftListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewEntitlementDriftsList

`func NewEntitlementDriftsList() *EntitlementDriftsList`

NewEntitlementDriftsList instantiates a new EntitlementDriftsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDriftsListWithDefaults

`func NewEntitlementDriftsListWithDefaults() *EntitlementDriftsList`

NewEntitlementDriftsListWithDefaults instantiates a new EntitlementDriftsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EntitlementDriftsList) GetData() []EntitlementDriftFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EntitlementDriftsList) GetDataOk() (*[]EntitlementDriftFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EntitlementDriftsList) SetData(v []EntitlementDriftFull)`

SetData sets Data field to given value.

### HasData

`func (o *EntitlementDriftsList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *EntitlementDriftsList) GetLinks() EntitlementDriftListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementDriftsList) GetLinksOk() (*EntitlementDriftListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementDriftsList) SetLinks(v EntitlementDriftListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntitlementDriftsList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementDriftsList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementDriftsList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementDriftsList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementDriftsList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


