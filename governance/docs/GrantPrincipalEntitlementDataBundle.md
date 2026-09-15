# GrantPrincipalEntitlementDataBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Indicates entitlement bundle assignment | 
**EntitlementBundleOrn** | **string** | The ORN of the entitlement bundle to grant | 
**AccessDuration** | Pointer to [**AccessDuration**](AccessDuration.md) |  | [optional] 

## Methods

### NewGrantPrincipalEntitlementDataBundle

`func NewGrantPrincipalEntitlementDataBundle(type_ string, entitlementBundleOrn string, ) *GrantPrincipalEntitlementDataBundle`

NewGrantPrincipalEntitlementDataBundle instantiates a new GrantPrincipalEntitlementDataBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPrincipalEntitlementDataBundleWithDefaults

`func NewGrantPrincipalEntitlementDataBundleWithDefaults() *GrantPrincipalEntitlementDataBundle`

NewGrantPrincipalEntitlementDataBundleWithDefaults instantiates a new GrantPrincipalEntitlementDataBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GrantPrincipalEntitlementDataBundle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrantPrincipalEntitlementDataBundle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrantPrincipalEntitlementDataBundle) SetType(v string)`

SetType sets Type field to given value.


### GetEntitlementBundleOrn

`func (o *GrantPrincipalEntitlementDataBundle) GetEntitlementBundleOrn() string`

GetEntitlementBundleOrn returns the EntitlementBundleOrn field if non-nil, zero value otherwise.

### GetEntitlementBundleOrnOk

`func (o *GrantPrincipalEntitlementDataBundle) GetEntitlementBundleOrnOk() (*string, bool)`

GetEntitlementBundleOrnOk returns a tuple with the EntitlementBundleOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundleOrn

`func (o *GrantPrincipalEntitlementDataBundle) SetEntitlementBundleOrn(v string)`

SetEntitlementBundleOrn sets EntitlementBundleOrn field to given value.


### GetAccessDuration

`func (o *GrantPrincipalEntitlementDataBundle) GetAccessDuration() AccessDuration`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *GrantPrincipalEntitlementDataBundle) GetAccessDurationOk() (*AccessDuration, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *GrantPrincipalEntitlementDataBundle) SetAccessDuration(v AccessDuration)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *GrantPrincipalEntitlementDataBundle) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


