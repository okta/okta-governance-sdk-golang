# PrincipalAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Parent** | [**TargetResource**](TargetResource.md) |  | 
**TargetPrincipalOrn** | **string** | The Okta user &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**TargetPrincipal** | [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | 
**ExpirationTime** | Pointer to **time.Time** | The date on which the user access expires. Date in ISO 8601 format. | [optional] 
**TimeZone** | Pointer to **string** | The time zone, in IANA format, for the end date of the user access. | [optional] 
**Base** | Pointer to [**Grant**](Grant.md) |  | [optional] 
**Additional** | Pointer to [**[]Grant**](Grant.md) |  | [optional] 

## Methods

### NewPrincipalAccess

`func NewPrincipalAccess(parentResourceOrn string, parent TargetResource, targetPrincipalOrn string, targetPrincipal TargetPrincipalFull, ) *PrincipalAccess`

NewPrincipalAccess instantiates a new PrincipalAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalAccessWithDefaults

`func NewPrincipalAccessWithDefaults() *PrincipalAccess`

NewPrincipalAccessWithDefaults instantiates a new PrincipalAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentResourceOrn

`func (o *PrincipalAccess) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *PrincipalAccess) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *PrincipalAccess) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.


### GetParent

`func (o *PrincipalAccess) GetParent() TargetResource`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PrincipalAccess) GetParentOk() (*TargetResource, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PrincipalAccess) SetParent(v TargetResource)`

SetParent sets Parent field to given value.


### GetTargetPrincipalOrn

`func (o *PrincipalAccess) GetTargetPrincipalOrn() string`

GetTargetPrincipalOrn returns the TargetPrincipalOrn field if non-nil, zero value otherwise.

### GetTargetPrincipalOrnOk

`func (o *PrincipalAccess) GetTargetPrincipalOrnOk() (*string, bool)`

GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipalOrn

`func (o *PrincipalAccess) SetTargetPrincipalOrn(v string)`

SetTargetPrincipalOrn sets TargetPrincipalOrn field to given value.


### GetTargetPrincipal

`func (o *PrincipalAccess) GetTargetPrincipal() TargetPrincipalFull`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *PrincipalAccess) GetTargetPrincipalOk() (*TargetPrincipalFull, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *PrincipalAccess) SetTargetPrincipal(v TargetPrincipalFull)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetExpirationTime

`func (o *PrincipalAccess) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *PrincipalAccess) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *PrincipalAccess) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *PrincipalAccess) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *PrincipalAccess) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PrincipalAccess) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PrincipalAccess) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PrincipalAccess) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetBase

`func (o *PrincipalAccess) GetBase() Grant`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PrincipalAccess) GetBaseOk() (*Grant, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PrincipalAccess) SetBase(v Grant)`

SetBase sets Base field to given value.

### HasBase

`func (o *PrincipalAccess) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetAdditional

`func (o *PrincipalAccess) GetAdditional() []Grant`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *PrincipalAccess) GetAdditionalOk() (*[]Grant, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *PrincipalAccess) SetAdditional(v []Grant)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *PrincipalAccess) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


