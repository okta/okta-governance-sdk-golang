# Grant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantMethod** | Pointer to [**GrantMethod**](GrantMethod.md) |  | [optional] 
**GrantType** | [**GrantType**](GrantType.md) |  | 
**StartTime** | Pointer to **time.Time** | The date on which the user received an access. Date in ISO 8601 format. | [optional] 
**ExpirationTime** | Pointer to **time.Time** | The date on which the user access expires. Date in ISO 8601 format. | [optional] 
**TimeZone** | Pointer to **string** | The time zone, in IANA format, for the end date of the user access. | [optional] 
**Grant** | [**GrantDetails**](GrantDetails.md) |  | 
**Bundle** | Pointer to [**Bundle**](Bundle.md) |  | [optional] 
**Entitlements** | [**[]GrantedEntitlements**](GrantedEntitlements.md) |  | 

## Methods

### NewGrant

`func NewGrant(grantType GrantType, grant GrantDetails, entitlements []GrantedEntitlements, ) *Grant`

NewGrant instantiates a new Grant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantWithDefaults

`func NewGrantWithDefaults() *Grant`

NewGrantWithDefaults instantiates a new Grant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantMethod

`func (o *Grant) GetGrantMethod() GrantMethod`

GetGrantMethod returns the GrantMethod field if non-nil, zero value otherwise.

### GetGrantMethodOk

`func (o *Grant) GetGrantMethodOk() (*GrantMethod, bool)`

GetGrantMethodOk returns a tuple with the GrantMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantMethod

`func (o *Grant) SetGrantMethod(v GrantMethod)`

SetGrantMethod sets GrantMethod field to given value.

### HasGrantMethod

`func (o *Grant) HasGrantMethod() bool`

HasGrantMethod returns a boolean if a field has been set.

### GetGrantType

`func (o *Grant) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Grant) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Grant) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetStartTime

`func (o *Grant) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Grant) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Grant) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Grant) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *Grant) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *Grant) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *Grant) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *Grant) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *Grant) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Grant) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Grant) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Grant) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetGrant

`func (o *Grant) GetGrant() GrantDetails`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *Grant) GetGrantOk() (*GrantDetails, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *Grant) SetGrant(v GrantDetails)`

SetGrant sets Grant field to given value.


### GetBundle

`func (o *Grant) GetBundle() Bundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *Grant) GetBundleOk() (*Bundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *Grant) SetBundle(v Bundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *Grant) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetEntitlements

`func (o *Grant) GetEntitlements() []GrantedEntitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *Grant) GetEntitlementsOk() (*[]GrantedEntitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *Grant) SetEntitlements(v []GrantedEntitlements)`

SetEntitlements sets Entitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


