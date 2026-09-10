# RevokeExternalUserEntitlementCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | **string** | The external/source-system identifier of the unmanaged user whose entitlements are being revoked. Provided by the source app at import time. | 
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]
**RevokeOrns** | **[]string** | List of entitlement-value [ORNs](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) to revoke for this external user | 

## Methods

### NewRevokeExternalUserEntitlementCreatable

`func NewRevokeExternalUserEntitlementCreatable(principal string, revokeOrns []string, ) *RevokeExternalUserEntitlementCreatable`

NewRevokeExternalUserEntitlementCreatable instantiates a new RevokeExternalUserEntitlementCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeExternalUserEntitlementCreatableWithDefaults

`func NewRevokeExternalUserEntitlementCreatableWithDefaults() *RevokeExternalUserEntitlementCreatable`

NewRevokeExternalUserEntitlementCreatableWithDefaults instantiates a new RevokeExternalUserEntitlementCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *RevokeExternalUserEntitlementCreatable) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *RevokeExternalUserEntitlementCreatable) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *RevokeExternalUserEntitlementCreatable) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetActor

`func (o *RevokeExternalUserEntitlementCreatable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RevokeExternalUserEntitlementCreatable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RevokeExternalUserEntitlementCreatable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RevokeExternalUserEntitlementCreatable) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetRevokeOrns

`func (o *RevokeExternalUserEntitlementCreatable) GetRevokeOrns() []string`

GetRevokeOrns returns the RevokeOrns field if non-nil, zero value otherwise.

### GetRevokeOrnsOk

`func (o *RevokeExternalUserEntitlementCreatable) GetRevokeOrnsOk() (*[]string, bool)`

GetRevokeOrnsOk returns a tuple with the RevokeOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOrns

`func (o *RevokeExternalUserEntitlementCreatable) SetRevokeOrns(v []string)`

SetRevokeOrns sets RevokeOrns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


