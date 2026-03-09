# RevokePrincipalAccessCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | 
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]
**RevokeOrns** | **[]string** | List of resource [ORNs](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) to revoke access:   * Entitlement value and entitlement bundle resources can be combined in a single request (with a maximum of five resources in a request).   * App resources must be revoked separately (a request can only contain one app ORN). | 

## Methods

### NewRevokePrincipalAccessCreatable

`func NewRevokePrincipalAccessCreatable(principalOrn string, revokeOrns []string, ) *RevokePrincipalAccessCreatable`

NewRevokePrincipalAccessCreatable instantiates a new RevokePrincipalAccessCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokePrincipalAccessCreatableWithDefaults

`func NewRevokePrincipalAccessCreatableWithDefaults() *RevokePrincipalAccessCreatable`

NewRevokePrincipalAccessCreatableWithDefaults instantiates a new RevokePrincipalAccessCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *RevokePrincipalAccessCreatable) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *RevokePrincipalAccessCreatable) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *RevokePrincipalAccessCreatable) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetActor

`func (o *RevokePrincipalAccessCreatable) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RevokePrincipalAccessCreatable) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RevokePrincipalAccessCreatable) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RevokePrincipalAccessCreatable) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetRevokeOrns

`func (o *RevokePrincipalAccessCreatable) GetRevokeOrns() []string`

GetRevokeOrns returns the RevokeOrns field if non-nil, zero value otherwise.

### GetRevokeOrnsOk

`func (o *RevokePrincipalAccessCreatable) GetRevokeOrnsOk() (*[]string, bool)`

GetRevokeOrnsOk returns a tuple with the RevokeOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOrns

`func (o *RevokePrincipalAccessCreatable) SetRevokeOrns(v []string)`

SetRevokeOrns sets RevokeOrns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


