# AssignedPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **time.Time** | The date on which the principal&#39;s access expires. This property is specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations). | [optional] 
**TimeZone** | Pointer to **string** | The time zone, in IANA format, for the end date of the user access. | [optional] 
**Principal** | Pointer to [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | [optional] 
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]
**CollectionId** | Pointer to **string** | The resource collection &#x60;id&#x60; | [optional] 

## Methods

### NewAssignedPrincipal

`func NewAssignedPrincipal() *AssignedPrincipal`

NewAssignedPrincipal instantiates a new AssignedPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedPrincipalWithDefaults

`func NewAssignedPrincipalWithDefaults() *AssignedPrincipal`

NewAssignedPrincipalWithDefaults instantiates a new AssignedPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *AssignedPrincipal) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AssignedPrincipal) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AssignedPrincipal) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AssignedPrincipal) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *AssignedPrincipal) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AssignedPrincipal) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AssignedPrincipal) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AssignedPrincipal) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetPrincipal

`func (o *AssignedPrincipal) GetPrincipal() TargetPrincipalFull`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AssignedPrincipal) GetPrincipalOk() (*TargetPrincipalFull, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AssignedPrincipal) SetPrincipal(v TargetPrincipalFull)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AssignedPrincipal) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetActor

`func (o *AssignedPrincipal) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AssignedPrincipal) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AssignedPrincipal) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AssignedPrincipal) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCollectionId

`func (o *AssignedPrincipal) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *AssignedPrincipal) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *AssignedPrincipal) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *AssignedPrincipal) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


