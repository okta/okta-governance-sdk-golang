# AccessDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | **time.Time** | The date and time when access expires, in ISO 8601 format | 
**Timezone** | **string** | The timezone for the expiration time, in IANA format | 

## Methods

### NewAccessDuration

`func NewAccessDuration(expirationTime time.Time, timezone string, ) *AccessDuration`

NewAccessDuration instantiates a new AccessDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDurationWithDefaults

`func NewAccessDurationWithDefaults() *AccessDuration`

NewAccessDurationWithDefaults instantiates a new AccessDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *AccessDuration) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AccessDuration) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AccessDuration) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.


### GetTimezone

`func (o *AccessDuration) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AccessDuration) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AccessDuration) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


