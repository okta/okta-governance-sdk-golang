# CollectionPrincipalAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | Pointer to **string** | The ORN of the assigned principal | [optional] 
**Actor** | Pointer to **string** | How the assignment was made | [optional] 
**AssignmentType** | Pointer to **string** | Type of assignment | [optional] 
**ExpirationTime** | Pointer to **NullableTime** | When the assignment expires (null for indefinite) | [optional] 

## Methods

### NewCollectionPrincipalAssignment

`func NewCollectionPrincipalAssignment() *CollectionPrincipalAssignment`

NewCollectionPrincipalAssignment instantiates a new CollectionPrincipalAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionPrincipalAssignmentWithDefaults

`func NewCollectionPrincipalAssignmentWithDefaults() *CollectionPrincipalAssignment`

NewCollectionPrincipalAssignmentWithDefaults instantiates a new CollectionPrincipalAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *CollectionPrincipalAssignment) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *CollectionPrincipalAssignment) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *CollectionPrincipalAssignment) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.

### HasPrincipalOrn

`func (o *CollectionPrincipalAssignment) HasPrincipalOrn() bool`

HasPrincipalOrn returns a boolean if a field has been set.

### GetActor

`func (o *CollectionPrincipalAssignment) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CollectionPrincipalAssignment) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CollectionPrincipalAssignment) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CollectionPrincipalAssignment) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAssignmentType

`func (o *CollectionPrincipalAssignment) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *CollectionPrincipalAssignment) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *CollectionPrincipalAssignment) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *CollectionPrincipalAssignment) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetExpirationTime

`func (o *CollectionPrincipalAssignment) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CollectionPrincipalAssignment) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CollectionPrincipalAssignment) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CollectionPrincipalAssignment) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *CollectionPrincipalAssignment) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *CollectionPrincipalAssignment) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


