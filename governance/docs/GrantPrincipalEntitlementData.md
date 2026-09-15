# GrantPrincipalEntitlementData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Indicates direct entitlement assignment | 
**Entitlements** | [**[]GrantPrincipalEntitlementGroup**](GrantPrincipalEntitlementGroup.md) | List of entitlement data. Each entry optionally specifies a resource asset scope. Entry without &#x60;resourceAssetOrn&#x60; represent app-level entitlements. | 
**AccessDuration** | Pointer to [**AccessDuration**](AccessDuration.md) |  | [optional] 
**EntitlementBundleOrn** | **string** | The ORN of the entitlement bundle to grant | 

## Methods

### NewGrantPrincipalEntitlementData

`func NewGrantPrincipalEntitlementData(type_ string, entitlements []GrantPrincipalEntitlementGroup, entitlementBundleOrn string, ) *GrantPrincipalEntitlementData`

NewGrantPrincipalEntitlementData instantiates a new GrantPrincipalEntitlementData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPrincipalEntitlementDataWithDefaults

`func NewGrantPrincipalEntitlementDataWithDefaults() *GrantPrincipalEntitlementData`

NewGrantPrincipalEntitlementDataWithDefaults instantiates a new GrantPrincipalEntitlementData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrantPrincipalEntitlementData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrantPrincipalEntitlementData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrantPrincipalEntitlementData) SetType(v string)`

SetType sets Type field to given value.


### GetEntitlements

`func (o *GrantPrincipalEntitlementData) GetEntitlements() []GrantPrincipalEntitlementGroup`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantPrincipalEntitlementData) GetEntitlementsOk() (*[]GrantPrincipalEntitlementGroup, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantPrincipalEntitlementData) SetEntitlements(v []GrantPrincipalEntitlementGroup)`

SetEntitlements sets Entitlements field to given value.


### GetAccessDuration

`func (o *GrantPrincipalEntitlementData) GetAccessDuration() AccessDuration`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *GrantPrincipalEntitlementData) GetAccessDurationOk() (*AccessDuration, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *GrantPrincipalEntitlementData) SetAccessDuration(v AccessDuration)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *GrantPrincipalEntitlementData) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.

### GetEntitlementBundleOrn

`func (o *GrantPrincipalEntitlementData) GetEntitlementBundleOrn() string`

GetEntitlementBundleOrn returns the EntitlementBundleOrn field if non-nil, zero value otherwise.

### GetEntitlementBundleOrnOk

`func (o *GrantPrincipalEntitlementData) GetEntitlementBundleOrnOk() (*string, bool)`

GetEntitlementBundleOrnOk returns a tuple with the EntitlementBundleOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundleOrn

`func (o *GrantPrincipalEntitlementData) SetEntitlementBundleOrn(v string)`

SetEntitlementBundleOrn sets EntitlementBundleOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


