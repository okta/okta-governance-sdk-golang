# PrincipalAccessPolicyEnableCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | 
**ResourceOrn** | **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**MergeExistingEntitlements** | Pointer to **bool** | If true, would skip removing bundle grant and create an entitlement grant with delta. By default it is false. | [optional] 

## Methods

### NewPrincipalAccessPolicyEnableCreatable

`func NewPrincipalAccessPolicyEnableCreatable(principalOrn string, resourceOrn string, ) *PrincipalAccessPolicyEnableCreatable`

NewPrincipalAccessPolicyEnableCreatable instantiates a new PrincipalAccessPolicyEnableCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalAccessPolicyEnableCreatableWithDefaults

`func NewPrincipalAccessPolicyEnableCreatableWithDefaults() *PrincipalAccessPolicyEnableCreatable`

NewPrincipalAccessPolicyEnableCreatableWithDefaults instantiates a new PrincipalAccessPolicyEnableCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *PrincipalAccessPolicyEnableCreatable) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *PrincipalAccessPolicyEnableCreatable) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *PrincipalAccessPolicyEnableCreatable) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetResourceOrn

`func (o *PrincipalAccessPolicyEnableCreatable) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *PrincipalAccessPolicyEnableCreatable) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *PrincipalAccessPolicyEnableCreatable) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetMergeExistingEntitlements

`func (o *PrincipalAccessPolicyEnableCreatable) GetMergeExistingEntitlements() bool`

GetMergeExistingEntitlements returns the MergeExistingEntitlements field if non-nil, zero value otherwise.

### GetMergeExistingEntitlementsOk

`func (o *PrincipalAccessPolicyEnableCreatable) GetMergeExistingEntitlementsOk() (*bool, bool)`

GetMergeExistingEntitlementsOk returns a tuple with the MergeExistingEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeExistingEntitlements

`func (o *PrincipalAccessPolicyEnableCreatable) SetMergeExistingEntitlements(v bool)`

SetMergeExistingEntitlements sets MergeExistingEntitlements field to given value.

### HasMergeExistingEntitlements

`func (o *PrincipalAccessPolicyEnableCreatable) HasMergeExistingEntitlements() bool`

HasMergeExistingEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


