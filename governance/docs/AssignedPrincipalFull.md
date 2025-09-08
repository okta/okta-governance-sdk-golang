# AssignedPrincipalFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **time.Time** | The date on which the principal&#39;s access expires. This property is specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations). | [optional] 
**TimeZone** | Pointer to **string** | The time zone, in IANA format, for the end date of the user access. | [optional] 
**Principal** | Pointer to [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | [optional] 
**Actor** | Pointer to [**GrantActor**](GrantActor.md) |  | [optional] [default to GRANTACTOR_API]
**CollectionId** | Pointer to **string** | The resource collection &#x60;id&#x60; | [optional] 
**Id** | Pointer to **string** | The assignment &#x60;id&#x60; | [optional] 
**AssignmentType** | Pointer to [**PrincipalAssignmentType**](PrincipalAssignmentType.md) |  | [optional] [default to PRINCIPALASSIGNMENTTYPE_INDIVIDUAL]

## Methods

### NewAssignedPrincipalFull

`func NewAssignedPrincipalFull() *AssignedPrincipalFull`

NewAssignedPrincipalFull instantiates a new AssignedPrincipalFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedPrincipalFullWithDefaults

`func NewAssignedPrincipalFullWithDefaults() *AssignedPrincipalFull`

NewAssignedPrincipalFullWithDefaults instantiates a new AssignedPrincipalFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *AssignedPrincipalFull) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AssignedPrincipalFull) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AssignedPrincipalFull) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AssignedPrincipalFull) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *AssignedPrincipalFull) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AssignedPrincipalFull) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AssignedPrincipalFull) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AssignedPrincipalFull) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetPrincipal

`func (o *AssignedPrincipalFull) GetPrincipal() TargetPrincipalFull`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AssignedPrincipalFull) GetPrincipalOk() (*TargetPrincipalFull, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AssignedPrincipalFull) SetPrincipal(v TargetPrincipalFull)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AssignedPrincipalFull) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetActor

`func (o *AssignedPrincipalFull) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AssignedPrincipalFull) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AssignedPrincipalFull) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AssignedPrincipalFull) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCollectionId

`func (o *AssignedPrincipalFull) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *AssignedPrincipalFull) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *AssignedPrincipalFull) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *AssignedPrincipalFull) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetId

`func (o *AssignedPrincipalFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignedPrincipalFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignedPrincipalFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssignedPrincipalFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignmentType

`func (o *AssignedPrincipalFull) GetAssignmentType() PrincipalAssignmentType`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *AssignedPrincipalFull) GetAssignmentTypeOk() (*PrincipalAssignmentType, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *AssignedPrincipalFull) SetAssignmentType(v PrincipalAssignmentType)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *AssignedPrincipalFull) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


