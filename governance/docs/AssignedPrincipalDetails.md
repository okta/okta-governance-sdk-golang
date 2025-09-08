# AssignedPrincipalDetails

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
**PrincipalProfile** | Pointer to [**PrincipalProfile**](PrincipalProfile.md) |  | [optional] 

## Methods

### NewAssignedPrincipalDetails

`func NewAssignedPrincipalDetails() *AssignedPrincipalDetails`

NewAssignedPrincipalDetails instantiates a new AssignedPrincipalDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedPrincipalDetailsWithDefaults

`func NewAssignedPrincipalDetailsWithDefaults() *AssignedPrincipalDetails`

NewAssignedPrincipalDetailsWithDefaults instantiates a new AssignedPrincipalDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *AssignedPrincipalDetails) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AssignedPrincipalDetails) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AssignedPrincipalDetails) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AssignedPrincipalDetails) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *AssignedPrincipalDetails) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AssignedPrincipalDetails) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AssignedPrincipalDetails) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AssignedPrincipalDetails) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetPrincipal

`func (o *AssignedPrincipalDetails) GetPrincipal() TargetPrincipalFull`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AssignedPrincipalDetails) GetPrincipalOk() (*TargetPrincipalFull, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AssignedPrincipalDetails) SetPrincipal(v TargetPrincipalFull)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *AssignedPrincipalDetails) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetActor

`func (o *AssignedPrincipalDetails) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AssignedPrincipalDetails) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AssignedPrincipalDetails) SetActor(v GrantActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AssignedPrincipalDetails) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCollectionId

`func (o *AssignedPrincipalDetails) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *AssignedPrincipalDetails) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *AssignedPrincipalDetails) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *AssignedPrincipalDetails) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetId

`func (o *AssignedPrincipalDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignedPrincipalDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignedPrincipalDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssignedPrincipalDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignmentType

`func (o *AssignedPrincipalDetails) GetAssignmentType() PrincipalAssignmentType`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *AssignedPrincipalDetails) GetAssignmentTypeOk() (*PrincipalAssignmentType, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *AssignedPrincipalDetails) SetAssignmentType(v PrincipalAssignmentType)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *AssignedPrincipalDetails) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetPrincipalProfile

`func (o *AssignedPrincipalDetails) GetPrincipalProfile() PrincipalProfile`

GetPrincipalProfile returns the PrincipalProfile field if non-nil, zero value otherwise.

### GetPrincipalProfileOk

`func (o *AssignedPrincipalDetails) GetPrincipalProfileOk() (*PrincipalProfile, bool)`

GetPrincipalProfileOk returns a tuple with the PrincipalProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalProfile

`func (o *AssignedPrincipalDetails) SetPrincipalProfile(v PrincipalProfile)`

SetPrincipalProfile sets PrincipalProfile field to given value.

### HasPrincipalProfile

`func (o *AssignedPrincipalDetails) HasPrincipalProfile() bool`

HasPrincipalProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


