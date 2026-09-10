# ScopedEntitlementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceAssetId** | **string** | The identifier of the resource asset associated with this scoped entitlement | 
**ResourceAsset** | [**ScopedEntitlementResourceAsset**](ScopedEntitlementResourceAsset.md) |  | 
**Entitlements** | [**[]GrantedEntitlements**](GrantedEntitlements.md) | Entitlements associated with this resource asset | 

## Methods

### NewScopedEntitlementItem

`func NewScopedEntitlementItem(resourceAssetId string, resourceAsset ScopedEntitlementResourceAsset, entitlements []GrantedEntitlements, ) *ScopedEntitlementItem`

NewScopedEntitlementItem instantiates a new ScopedEntitlementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedEntitlementItemWithDefaults

`func NewScopedEntitlementItemWithDefaults() *ScopedEntitlementItem`

NewScopedEntitlementItemWithDefaults instantiates a new ScopedEntitlementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceAssetId

`func (o *ScopedEntitlementItem) GetResourceAssetId() string`

GetResourceAssetId returns the ResourceAssetId field if non-nil, zero value otherwise.

### GetResourceAssetIdOk

`func (o *ScopedEntitlementItem) GetResourceAssetIdOk() (*string, bool)`

GetResourceAssetIdOk returns a tuple with the ResourceAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAssetId

`func (o *ScopedEntitlementItem) SetResourceAssetId(v string)`

SetResourceAssetId sets ResourceAssetId field to given value.


### GetResourceAsset

`func (o *ScopedEntitlementItem) GetResourceAsset() ScopedEntitlementResourceAsset`

GetResourceAsset returns the ResourceAsset field if non-nil, zero value otherwise.

### GetResourceAssetOk

`func (o *ScopedEntitlementItem) GetResourceAssetOk() (*ScopedEntitlementResourceAsset, bool)`

GetResourceAssetOk returns a tuple with the ResourceAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAsset

`func (o *ScopedEntitlementItem) SetResourceAsset(v ScopedEntitlementResourceAsset)`

SetResourceAsset sets ResourceAsset field to given value.


### GetEntitlements

`func (o *ScopedEntitlementItem) GetEntitlements() []GrantedEntitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *ScopedEntitlementItem) GetEntitlementsOk() (*[]GrantedEntitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *ScopedEntitlementItem) SetEntitlements(v []GrantedEntitlements)`

SetEntitlements sets Entitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


