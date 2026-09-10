# PrincipalEntitlementRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetOrn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**ResourceAsset** | Pointer to [**PrincipalEntitlementRowResourceAsset**](PrincipalEntitlementRowResourceAsset.md) |  | [optional] 
**Entitlements** | [**[]EntitlementWithValues**](EntitlementWithValues.md) | The principal&#39;s complete effective entitlements on this target, calculated on the server side from all grant sources: direct grants, group-inherited grants, and hierarchy-inherited grants. Each item is one entitlement property with only the first page of its effective values returned.  | 

## Methods

### NewPrincipalEntitlementRow

`func NewPrincipalEntitlementRow(targetOrn string, entitlements []EntitlementWithValues, ) *PrincipalEntitlementRow`

NewPrincipalEntitlementRow instantiates a new PrincipalEntitlementRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementRowWithDefaults

`func NewPrincipalEntitlementRowWithDefaults() *PrincipalEntitlementRow`

NewPrincipalEntitlementRowWithDefaults instantiates a new PrincipalEntitlementRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetOrn

`func (o *PrincipalEntitlementRow) GetTargetOrn() string`

GetTargetOrn returns the TargetOrn field if non-nil, zero value otherwise.

### GetTargetOrnOk

`func (o *PrincipalEntitlementRow) GetTargetOrnOk() (*string, bool)`

GetTargetOrnOk returns a tuple with the TargetOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrn

`func (o *PrincipalEntitlementRow) SetTargetOrn(v string)`

SetTargetOrn sets TargetOrn field to given value.


### GetResourceAsset

`func (o *PrincipalEntitlementRow) GetResourceAsset() PrincipalEntitlementRowResourceAsset`

GetResourceAsset returns the ResourceAsset field if non-nil, zero value otherwise.

### GetResourceAssetOk

`func (o *PrincipalEntitlementRow) GetResourceAssetOk() (*PrincipalEntitlementRowResourceAsset, bool)`

GetResourceAssetOk returns a tuple with the ResourceAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAsset

`func (o *PrincipalEntitlementRow) SetResourceAsset(v PrincipalEntitlementRowResourceAsset)`

SetResourceAsset sets ResourceAsset field to given value.

### HasResourceAsset

`func (o *PrincipalEntitlementRow) HasResourceAsset() bool`

HasResourceAsset returns a boolean if a field has been set.

### GetEntitlements

`func (o *PrincipalEntitlementRow) GetEntitlements() []EntitlementWithValues`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *PrincipalEntitlementRow) GetEntitlementsOk() (*[]EntitlementWithValues, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *PrincipalEntitlementRow) SetEntitlements(v []EntitlementWithValues)`

SetEntitlements sets Entitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


