# ScheduleSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScheduleType**](ScheduleType.md) |  | 
**StartDate** | **time.Time** | The date on which the campaign is supposed to start. Accepts date in ISO 8601 format. | 
**DurationInDays** | **float32** | The duration (in days) that the campaign is active. | 
**TimeZone** | **string** | The time zone, in IANA format, for the start date of the campaign. | 
**Recurrence** | Pointer to [**RecurrenceDefinitionMutable**](RecurrenceDefinitionMutable.md) |  | [optional] 

## Methods

### NewScheduleSettingsMutable

`func NewScheduleSettingsMutable(type_ ScheduleType, startDate time.Time, durationInDays float32, timeZone string, ) *ScheduleSettingsMutable`

NewScheduleSettingsMutable instantiates a new ScheduleSettingsMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleSettingsMutableWithDefaults

`func NewScheduleSettingsMutableWithDefaults() *ScheduleSettingsMutable`

NewScheduleSettingsMutableWithDefaults instantiates a new ScheduleSettingsMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScheduleSettingsMutable) GetType() ScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleSettingsMutable) GetTypeOk() (*ScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleSettingsMutable) SetType(v ScheduleType)`

SetType sets Type field to given value.


### GetStartDate

`func (o *ScheduleSettingsMutable) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ScheduleSettingsMutable) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ScheduleSettingsMutable) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetDurationInDays

`func (o *ScheduleSettingsMutable) GetDurationInDays() float32`

GetDurationInDays returns the DurationInDays field if non-nil, zero value otherwise.

### GetDurationInDaysOk

`func (o *ScheduleSettingsMutable) GetDurationInDaysOk() (*float32, bool)`

GetDurationInDaysOk returns a tuple with the DurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInDays

`func (o *ScheduleSettingsMutable) SetDurationInDays(v float32)`

SetDurationInDays sets DurationInDays field to given value.


### GetTimeZone

`func (o *ScheduleSettingsMutable) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ScheduleSettingsMutable) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ScheduleSettingsMutable) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetRecurrence

`func (o *ScheduleSettingsMutable) GetRecurrence() RecurrenceDefinitionMutable`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *ScheduleSettingsMutable) GetRecurrenceOk() (*RecurrenceDefinitionMutable, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *ScheduleSettingsMutable) SetRecurrence(v RecurrenceDefinitionMutable)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *ScheduleSettingsMutable) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


