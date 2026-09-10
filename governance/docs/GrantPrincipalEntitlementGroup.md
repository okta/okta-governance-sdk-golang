# GrantPrincipalEntitlementGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceAssetOrn** | Pointer to **string** | The ORN of the resource asset to scope the entitlements to. Omit this field for app-level entitlements that aren&#39;t tied to a specific resource asset. | [optional] 
**EntitlementValueOrns** | **[]string** | List of entitlement value ORNs to grant within this scope | 

## Methods

### NewGrantPrincipalEntitlementGroup

`func NewGrantPrincipalEntitlementGroup(entitlementValueOrns []string, ) *GrantPrincipalEntitlementGroup`

NewGrantPrincipalEntitlementGroup instantiates a new GrantPrincipalEntitlementGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPrincipalEntitlementGroupWithDefaults

`func NewGrantPrincipalEntitlementGroupWithDefaults() *GrantPrincipalEntitlementGroup`

NewGrantPrincipalEntitlementGroupWithDefaults instantiates a new GrantPrincipalEntitlementGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceAssetOrn

`func (o *GrantPrincipalEntitlementGroup) GetResourceAssetOrn() string`

GetResourceAssetOrn returns the ResourceAssetOrn field if non-nil, zero value otherwise.

### GetResourceAssetOrnOk

`func (o *GrantPrincipalEntitlementGroup) GetResourceAssetOrnOk() (*string, bool)`

GetResourceAssetOrnOk returns a tuple with the ResourceAssetOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAssetOrn

`func (o *GrantPrincipalEntitlementGroup) SetResourceAssetOrn(v string)`

SetResourceAssetOrn sets ResourceAssetOrn field to given value.

### HasResourceAssetOrn

`func (o *GrantPrincipalEntitlementGroup) HasResourceAssetOrn() bool`

HasResourceAssetOrn returns a boolean if a field has been set.

### GetEntitlementValueOrns

`func (o *GrantPrincipalEntitlementGroup) GetEntitlementValueOrns() []string`

GetEntitlementValueOrns returns the EntitlementValueOrns field if non-nil, zero value otherwise.

### GetEntitlementValueOrnsOk

`func (o *GrantPrincipalEntitlementGroup) GetEntitlementValueOrnsOk() (*[]string, bool)`

GetEntitlementValueOrnsOk returns a tuple with the EntitlementValueOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValueOrns

`func (o *GrantPrincipalEntitlementGroup) SetEntitlementValueOrns(v []string)`

SetEntitlementValueOrns sets EntitlementValueOrns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


