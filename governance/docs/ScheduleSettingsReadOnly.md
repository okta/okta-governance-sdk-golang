# ScheduleSettingsReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScheduleType**](ScheduleType.md) |  | 
**StartDate** | **time.Time** | The date on which the campaign is supposed to start. Accepts date in ISO 8601 format. | 
**DurationInDays** | **float32** | The duration (in days) that the campaign is active. | 
**TimeZone** | **string** | The time zone, in IANA format, for the start date of the campaign. | 
**EndDate** | Pointer to **time.Time** | The date on which the campaign is supposed to end. Date in ISO 8601 format. | [optional] [readonly] 
**Recurrence** | Pointer to [**RecurrenceDefinitionMutable**](RecurrenceDefinitionMutable.md) |  | [optional] 

## Methods

### NewScheduleSettingsReadOnly

`func NewScheduleSettingsReadOnly(type_ ScheduleType, startDate time.Time, durationInDays float32, timeZone string, ) *ScheduleSettingsReadOnly`

NewScheduleSettingsReadOnly instantiates a new ScheduleSettingsReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleSettingsReadOnlyWithDefaults

`func NewScheduleSettingsReadOnlyWithDefaults() *ScheduleSettingsReadOnly`

NewScheduleSettingsReadOnlyWithDefaults instantiates a new ScheduleSettingsReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScheduleSettingsReadOnly) GetType() ScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleSettingsReadOnly) GetTypeOk() (*ScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleSettingsReadOnly) SetType(v ScheduleType)`

SetType sets Type field to given value.


### GetStartDate

`func (o *ScheduleSettingsReadOnly) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ScheduleSettingsReadOnly) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ScheduleSettingsReadOnly) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetDurationInDays

`func (o *ScheduleSettingsReadOnly) GetDurationInDays() float32`

GetDurationInDays returns the DurationInDays field if non-nil, zero value otherwise.

### GetDurationInDaysOk

`func (o *ScheduleSettingsReadOnly) GetDurationInDaysOk() (*float32, bool)`

GetDurationInDaysOk returns a tuple with the DurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInDays

`func (o *ScheduleSettingsReadOnly) SetDurationInDays(v float32)`

SetDurationInDays sets DurationInDays field to given value.


### GetTimeZone

`func (o *ScheduleSettingsReadOnly) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ScheduleSettingsReadOnly) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ScheduleSettingsReadOnly) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetEndDate

`func (o *ScheduleSettingsReadOnly) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ScheduleSettingsReadOnly) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ScheduleSettingsReadOnly) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ScheduleSettingsReadOnly) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRecurrence

`func (o *ScheduleSettingsReadOnly) GetRecurrence() RecurrenceDefinitionMutable`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *ScheduleSettingsReadOnly) GetRecurrenceOk() (*RecurrenceDefinitionMutable, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *ScheduleSettingsReadOnly) SetRecurrence(v RecurrenceDefinitionMutable)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *ScheduleSettingsReadOnly) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


