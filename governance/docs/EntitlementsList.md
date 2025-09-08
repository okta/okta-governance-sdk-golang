# EntitlementsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EntitlementsListObject**](EntitlementsListObject.md) | List of all entitlements matching the filter | 
**Links** | [**ListLinks**](ListLinks.md) |  | 
**Metadata** | [**ListMetadata**](ListMetadata.md) |  | 

## Methods

### NewEntitlementsList

`func NewEntitlementsList(data []EntitlementsListObject, links ListLinks, metadata ListMetadata, ) *EntitlementsList`

NewEntitlementsList instantiates a new EntitlementsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsListWithDefaults

`func NewEntitlementsListWithDefaults() *EntitlementsList`

NewEntitlementsListWithDefaults instantiates a new EntitlementsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EntitlementsList) GetData() []EntitlementsListObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EntitlementsList) GetDataOk() (*[]EntitlementsListObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EntitlementsList) SetData(v []EntitlementsListObject)`

SetData sets Data field to given value.


### GetLinks

`func (o *EntitlementsList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementsList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementsList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *EntitlementsList) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementsList) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementsList) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


