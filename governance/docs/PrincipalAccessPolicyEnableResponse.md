# PrincipalAccessPolicyEnableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | 
**ResourceOrn** | **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**AccessDuration** | Pointer to [**AssignmentProperties**](AssignmentProperties.md) |  | [optional] 
**EffectiveEntitlements** | [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | 
**EntitlementAccessDetails** | Pointer to [**[]EntitlementAccessDetailsObject**](EntitlementAccessDetailsObject.md) | Entitlement access details | [optional] 

## Methods

### NewPrincipalAccessPolicyEnableResponse

`func NewPrincipalAccessPolicyEnableResponse(principalOrn string, resourceOrn string, effectiveEntitlements []EntitlementFull, ) *PrincipalAccessPolicyEnableResponse`

NewPrincipalAccessPolicyEnableResponse instantiates a new PrincipalAccessPolicyEnableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalAccessPolicyEnableResponseWithDefaults

`func NewPrincipalAccessPolicyEnableResponseWithDefaults() *PrincipalAccessPolicyEnableResponse`

NewPrincipalAccessPolicyEnableResponseWithDefaults instantiates a new PrincipalAccessPolicyEnableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *PrincipalAccessPolicyEnableResponse) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *PrincipalAccessPolicyEnableResponse) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *PrincipalAccessPolicyEnableResponse) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetResourceOrn

`func (o *PrincipalAccessPolicyEnableResponse) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *PrincipalAccessPolicyEnableResponse) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *PrincipalAccessPolicyEnableResponse) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetAccessDuration

`func (o *PrincipalAccessPolicyEnableResponse) GetAccessDuration() AssignmentProperties`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *PrincipalAccessPolicyEnableResponse) GetAccessDurationOk() (*AssignmentProperties, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *PrincipalAccessPolicyEnableResponse) SetAccessDuration(v AssignmentProperties)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *PrincipalAccessPolicyEnableResponse) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.

### GetEffectiveEntitlements

`func (o *PrincipalAccessPolicyEnableResponse) GetEffectiveEntitlements() []EntitlementFull`

GetEffectiveEntitlements returns the EffectiveEntitlements field if non-nil, zero value otherwise.

### GetEffectiveEntitlementsOk

`func (o *PrincipalAccessPolicyEnableResponse) GetEffectiveEntitlementsOk() (*[]EntitlementFull, bool)`

GetEffectiveEntitlementsOk returns a tuple with the EffectiveEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveEntitlements

`func (o *PrincipalAccessPolicyEnableResponse) SetEffectiveEntitlements(v []EntitlementFull)`

SetEffectiveEntitlements sets EffectiveEntitlements field to given value.


### GetEntitlementAccessDetails

`func (o *PrincipalAccessPolicyEnableResponse) GetEntitlementAccessDetails() []EntitlementAccessDetailsObject`

GetEntitlementAccessDetails returns the EntitlementAccessDetails field if non-nil, zero value otherwise.

### GetEntitlementAccessDetailsOk

`func (o *PrincipalAccessPolicyEnableResponse) GetEntitlementAccessDetailsOk() (*[]EntitlementAccessDetailsObject, bool)`

GetEntitlementAccessDetailsOk returns a tuple with the EntitlementAccessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAccessDetails

`func (o *PrincipalAccessPolicyEnableResponse) SetEntitlementAccessDetails(v []EntitlementAccessDetailsObject)`

SetEntitlementAccessDetails sets EntitlementAccessDetails field to given value.

### HasEntitlementAccessDetails

`func (o *PrincipalAccessPolicyEnableResponse) HasEntitlementAccessDetails() bool`

HasEntitlementAccessDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


