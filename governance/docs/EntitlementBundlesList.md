# EntitlementBundlesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EntitlementBundleFull**](EntitlementBundleFull.md) | All entitlement bundles on the current page | [optional] 
**Links** | Pointer to [**EntitlementBundleListLinks**](EntitlementBundleListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewEntitlementBundlesList

`func NewEntitlementBundlesList() *EntitlementBundlesList`

NewEntitlementBundlesList instantiates a new EntitlementBundlesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementBundlesListWithDefaults

`func NewEntitlementBundlesListWithDefaults() *EntitlementBundlesList`

NewEntitlementBundlesListWithDefaults instantiates a new EntitlementBundlesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EntitlementBundlesList) GetData() []EntitlementBundleFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EntitlementBundlesList) GetDataOk() (*[]EntitlementBundleFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EntitlementBundlesList) SetData(v []EntitlementBundleFull)`

SetData sets Data field to given value.

### HasData

`func (o *EntitlementBundlesList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *EntitlementBundlesList) GetLinks() EntitlementBundleListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementBundlesList) GetLinksOk() (*EntitlementBundleListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementBundlesList) SetLinks(v EntitlementBundleListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntitlementBundlesList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementBundlesList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementBundlesList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementBundlesList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementBundlesList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


