# IndirectAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grant** | [**GrantDetails**](GrantDetails.md) |  | 
**ScopedEntitlements** | [**[]ScopedEntitlementItem**](ScopedEntitlementItem.md) | Entitlements obtained through this grant grouped by resource asset scope | 

## Methods

### NewIndirectAccess

`func NewIndirectAccess(grant GrantDetails, scopedEntitlements []ScopedEntitlementItem, ) *IndirectAccess`

NewIndirectAccess instantiates a new IndirectAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndirectAccessWithDefaults

`func NewIndirectAccessWithDefaults() *IndirectAccess`

NewIndirectAccessWithDefaults instantiates a new IndirectAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrant

`func (o *IndirectAccess) GetGrant() GrantDetails`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *IndirectAccess) GetGrantOk() (*GrantDetails, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *IndirectAccess) SetGrant(v GrantDetails)`

SetGrant sets Grant field to given value.


### GetScopedEntitlements

`func (o *IndirectAccess) GetScopedEntitlements() []ScopedEntitlementItem`

GetScopedEntitlements returns the ScopedEntitlements field if non-nil, zero value otherwise.

### GetScopedEntitlementsOk

`func (o *IndirectAccess) GetScopedEntitlementsOk() (*[]ScopedEntitlementItem, bool)`

GetScopedEntitlementsOk returns a tuple with the ScopedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedEntitlements

`func (o *IndirectAccess) SetScopedEntitlements(v []ScopedEntitlementItem)`

SetScopedEntitlements sets ScopedEntitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


