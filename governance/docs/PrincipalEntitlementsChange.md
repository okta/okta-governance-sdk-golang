# PrincipalEntitlementsChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementsChanged** | Pointer to [**[]EntitlementChangedFull**](EntitlementChangedFull.md) | List of changed entitlements | [optional] 
**ResourceOrn** | Pointer to **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 
**Resource** | Pointer to [**TargetResource**](TargetResource.md) |  | [optional] 
**PrincipalOrn** | Pointer to **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | [optional] 
**Principal** | Pointer to [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | [optional] 
**Links** | Pointer to [**PrincipalEntitlementsChangeLinks**](PrincipalEntitlementsChangeLinks.md) |  | [optional] 

## Methods

### NewPrincipalEntitlementsChange

`func NewPrincipalEntitlementsChange() *PrincipalEntitlementsChange`

NewPrincipalEntitlementsChange instantiates a new PrincipalEntitlementsChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementsChangeWithDefaults

`func NewPrincipalEntitlementsChangeWithDefaults() *PrincipalEntitlementsChange`

NewPrincipalEntitlementsChangeWithDefaults instantiates a new PrincipalEntitlementsChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementsChanged

`func (o *PrincipalEntitlementsChange) GetEntitlementsChanged() []EntitlementChangedFull`

GetEntitlementsChanged returns the EntitlementsChanged field if non-nil, zero value otherwise.

### GetEntitlementsChangedOk

`func (o *PrincipalEntitlementsChange) GetEntitlementsChangedOk() (*[]EntitlementChangedFull, bool)`

GetEntitlementsChangedOk returns a tuple with the EntitlementsChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsChanged

`func (o *PrincipalEntitlementsChange) SetEntitlementsChanged(v []EntitlementChangedFull)`

SetEntitlementsChanged sets EntitlementsChanged field to given value.

### HasEntitlementsChanged

`func (o *PrincipalEntitlementsChange) HasEntitlementsChanged() bool`

HasEntitlementsChanged returns a boolean if a field has been set.

### GetResourceOrn

`func (o *PrincipalEntitlementsChange) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *PrincipalEntitlementsChange) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *PrincipalEntitlementsChange) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.

### HasResourceOrn

`func (o *PrincipalEntitlementsChange) HasResourceOrn() bool`

HasResourceOrn returns a boolean if a field has been set.

### GetResource

`func (o *PrincipalEntitlementsChange) GetResource() TargetResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PrincipalEntitlementsChange) GetResourceOk() (*TargetResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PrincipalEntitlementsChange) SetResource(v TargetResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PrincipalEntitlementsChange) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPrincipalOrn

`func (o *PrincipalEntitlementsChange) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *PrincipalEntitlementsChange) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *PrincipalEntitlementsChange) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.

### HasPrincipalOrn

`func (o *PrincipalEntitlementsChange) HasPrincipalOrn() bool`

HasPrincipalOrn returns a boolean if a field has been set.

### GetPrincipal

`func (o *PrincipalEntitlementsChange) GetPrincipal() TargetPrincipalFull`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PrincipalEntitlementsChange) GetPrincipalOk() (*TargetPrincipalFull, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PrincipalEntitlementsChange) SetPrincipal(v TargetPrincipalFull)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *PrincipalEntitlementsChange) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetLinks

`func (o *PrincipalEntitlementsChange) GetLinks() PrincipalEntitlementsChangeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrincipalEntitlementsChange) GetLinksOk() (*PrincipalEntitlementsChangeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrincipalEntitlementsChange) SetLinks(v PrincipalEntitlementsChangeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PrincipalEntitlementsChange) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


