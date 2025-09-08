# ScheduleSettingsWriteable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **time.Time** | An optional expiration date. If provided, the entitlement bundle grant will be revoked at the specified date and time. | [optional] 
**TimeZone** | Pointer to **string** | The time zone, in IANA format, in which expirationDate was configured | [optional] 

## Methods

### NewScheduleSettingsWriteable

`func NewScheduleSettingsWriteable() *ScheduleSettingsWriteable`

NewScheduleSettingsWriteable instantiates a new ScheduleSettingsWriteable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleSettingsWriteableWithDefaults

`func NewScheduleSettingsWriteableWithDefaults() *ScheduleSettingsWriteable`

NewScheduleSettingsWriteableWithDefaults instantiates a new ScheduleSettingsWriteable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *ScheduleSettingsWriteable) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ScheduleSettingsWriteable) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ScheduleSettingsWriteable) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ScheduleSettingsWriteable) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *ScheduleSettingsWriteable) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ScheduleSettingsWriteable) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ScheduleSettingsWriteable) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ScheduleSettingsWriteable) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


