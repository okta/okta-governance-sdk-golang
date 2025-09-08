# PrincipalEntitlementsHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOrn** | Pointer to **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 
**Resource** | Pointer to [**TargetResource**](TargetResource.md) |  | [optional] 
**PrincipalOrn** | Pointer to **string** | The Okta user &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 
**Principal** | Pointer to [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | [optional] 
**EntitlementHistory** | Pointer to [**[]EntitlementHistoryRecord**](EntitlementHistoryRecord.md) | Principal entitlements history list | [optional] 
**Links** | Pointer to [**ListLinks**](ListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewPrincipalEntitlementsHistory

`func NewPrincipalEntitlementsHistory() *PrincipalEntitlementsHistory`

NewPrincipalEntitlementsHistory instantiates a new PrincipalEntitlementsHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementsHistoryWithDefaults

`func NewPrincipalEntitlementsHistoryWithDefaults() *PrincipalEntitlementsHistory`

NewPrincipalEntitlementsHistoryWithDefaults instantiates a new PrincipalEntitlementsHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOrn

`func (o *PrincipalEntitlementsHistory) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *PrincipalEntitlementsHistory) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *PrincipalEntitlementsHistory) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.

### HasResourceOrn

`func (o *PrincipalEntitlementsHistory) HasResourceOrn() bool`

HasResourceOrn returns a boolean if a field has been set.

### GetResource

`func (o *PrincipalEntitlementsHistory) GetResource() TargetResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PrincipalEntitlementsHistory) GetResourceOk() (*TargetResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PrincipalEntitlementsHistory) SetResource(v TargetResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PrincipalEntitlementsHistory) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPrincipalOrn

`func (o *PrincipalEntitlementsHistory) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *PrincipalEntitlementsHistory) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *PrincipalEntitlementsHistory) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.

### HasPrincipalOrn

`func (o *PrincipalEntitlementsHistory) HasPrincipalOrn() bool`

HasPrincipalOrn returns a boolean if a field has been set.

### GetPrincipal

`func (o *PrincipalEntitlementsHistory) GetPrincipal() TargetPrincipalFull`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PrincipalEntitlementsHistory) GetPrincipalOk() (*TargetPrincipalFull, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PrincipalEntitlementsHistory) SetPrincipal(v TargetPrincipalFull)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *PrincipalEntitlementsHistory) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetEntitlementHistory

`func (o *PrincipalEntitlementsHistory) GetEntitlementHistory() []EntitlementHistoryRecord`

GetEntitlementHistory returns the EntitlementHistory field if non-nil, zero value otherwise.

### GetEntitlementHistoryOk

`func (o *PrincipalEntitlementsHistory) GetEntitlementHistoryOk() (*[]EntitlementHistoryRecord, bool)`

GetEntitlementHistoryOk returns a tuple with the EntitlementHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementHistory

`func (o *PrincipalEntitlementsHistory) SetEntitlementHistory(v []EntitlementHistoryRecord)`

SetEntitlementHistory sets EntitlementHistory field to given value.

### HasEntitlementHistory

`func (o *PrincipalEntitlementsHistory) HasEntitlementHistory() bool`

HasEntitlementHistory returns a boolean if a field has been set.

### GetLinks

`func (o *PrincipalEntitlementsHistory) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrincipalEntitlementsHistory) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrincipalEntitlementsHistory) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PrincipalEntitlementsHistory) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *PrincipalEntitlementsHistory) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PrincipalEntitlementsHistory) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PrincipalEntitlementsHistory) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PrincipalEntitlementsHistory) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


