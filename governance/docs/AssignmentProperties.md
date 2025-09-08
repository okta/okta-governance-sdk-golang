# AssignmentProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **time.Time** | The date on which the principal&#39;s access expires. This property is specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations). | [optional] 
**TimeZone** | Pointer to **string** | The time zone, in IANA format, for the end date of the user access. | [optional] 

## Methods

### NewAssignmentProperties

`func NewAssignmentProperties() *AssignmentProperties`

NewAssignmentProperties instantiates a new AssignmentProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentPropertiesWithDefaults

`func NewAssignmentPropertiesWithDefaults() *AssignmentProperties`

NewAssignmentPropertiesWithDefaults instantiates a new AssignmentProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *AssignmentProperties) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AssignmentProperties) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AssignmentProperties) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *AssignmentProperties) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *AssignmentProperties) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AssignmentProperties) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AssignmentProperties) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AssignmentProperties) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


