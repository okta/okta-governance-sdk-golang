# GrantPrincipalAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** | The Okta user in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format | 
**ResourceOrn** | **string** | The ORN of the app resource to grant access to | 
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]
**AccessDuration** | Pointer to [**AccessDuration**](AccessDuration.md) |  | [optional] 
**EntitlementData** | Pointer to [**GrantPrincipalEntitlementData**](GrantPrincipalEntitlementData.md) |  | [optional] 

## Methods

### NewGrantPrincipalAccessRequest

`func NewGrantPrincipalAccessRequest(principalOrn string, resourceOrn string, ) *GrantPrincipalAccessRequest`

NewGrantPrincipalAccessRequest instantiates a new GrantPrincipalAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPrincipalAccessRequestWithDefaults

`func NewGrantPrincipalAccessRequestWithDefaults() *GrantPrincipalAccessRequest`

NewGrantPrincipalAccessRequestWithDefaults instantiates a new GrantPrincipalAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *GrantPrincipalAccessRequest) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *GrantPrincipalAccessRequest) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *GrantPrincipalAccessRequest) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetResourceOrn

`func (o *GrantPrincipalAccessRequest) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *GrantPrincipalAccessRequest) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *GrantPrincipalAccessRequest) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetActor

`func (o *GrantPrincipalAccessRequest) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantPrincipalAccessRequest) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantPrincipalAccessRequest) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GrantPrincipalAccessRequest) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAccessDuration

`func (o *GrantPrincipalAccessRequest) GetAccessDuration() AccessDuration`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *GrantPrincipalAccessRequest) GetAccessDurationOk() (*AccessDuration, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *GrantPrincipalAccessRequest) SetAccessDuration(v AccessDuration)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *GrantPrincipalAccessRequest) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.

### GetEntitlementData

`func (o *GrantPrincipalAccessRequest) GetEntitlementData() GrantPrincipalEntitlementData`

GetEntitlementData returns the EntitlementData field if non-nil, zero value otherwise.

### GetEntitlementDataOk

`func (o *GrantPrincipalAccessRequest) GetEntitlementDataOk() (*GrantPrincipalEntitlementData, bool)`

GetEntitlementDataOk returns a tuple with the EntitlementData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementData

`func (o *GrantPrincipalAccessRequest) SetEntitlementData(v GrantPrincipalEntitlementData)`

SetEntitlementData sets EntitlementData field to given value.

### HasEntitlementData

`func (o *GrantPrincipalAccessRequest) HasEntitlementData() bool`

HasEntitlementData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


