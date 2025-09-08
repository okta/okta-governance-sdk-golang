# ListEntitlementBundles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EntitlementBundleFullWithEntitlements**](EntitlementBundleFullWithEntitlements.md) | All entitlement bundles on the current page | [optional] 
**Links** | Pointer to [**EntitlementBundleListLinks**](EntitlementBundleListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewListEntitlementBundles200Response

`func NewListEntitlementBundles200Response() *ListEntitlementBundles200Response`

NewListEntitlementBundles200Response instantiates a new ListEntitlementBundles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitlementBundles200ResponseWithDefaults

`func NewListEntitlementBundles200ResponseWithDefaults() *ListEntitlementBundles200Response`

NewListEntitlementBundles200ResponseWithDefaults instantiates a new ListEntitlementBundles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEntitlementBundles200Response) GetData() []EntitlementBundleFullWithEntitlements`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEntitlementBundles200Response) GetDataOk() (*[]EntitlementBundleFullWithEntitlements, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEntitlementBundles200Response) SetData(v []EntitlementBundleFullWithEntitlements)`

SetData sets Data field to given value.

### HasData

`func (o *ListEntitlementBundles200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ListEntitlementBundles200Response) GetLinks() EntitlementBundleListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListEntitlementBundles200Response) GetLinksOk() (*EntitlementBundleListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListEntitlementBundles200Response) SetLinks(v EntitlementBundleListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListEntitlementBundles200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ListEntitlementBundles200Response) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListEntitlementBundles200Response) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListEntitlementBundles200Response) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListEntitlementBundles200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


