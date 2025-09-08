# EntitlementValuesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EntitlementValueWithParent**](EntitlementValueWithParent.md) | List of entitlement values matching the filter | [optional] 
**Links** | Pointer to [**ListLinks**](ListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewEntitlementValuesList

`func NewEntitlementValuesList() *EntitlementValuesList`

NewEntitlementValuesList instantiates a new EntitlementValuesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValuesListWithDefaults

`func NewEntitlementValuesListWithDefaults() *EntitlementValuesList`

NewEntitlementValuesListWithDefaults instantiates a new EntitlementValuesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EntitlementValuesList) GetData() []EntitlementValueWithParent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EntitlementValuesList) GetDataOk() (*[]EntitlementValueWithParent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EntitlementValuesList) SetData(v []EntitlementValueWithParent)`

SetData sets Data field to given value.

### HasData

`func (o *EntitlementValuesList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *EntitlementValuesList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementValuesList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementValuesList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntitlementValuesList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementValuesList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementValuesList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementValuesList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementValuesList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


