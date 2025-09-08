# RecurrenceDefinitionMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | Recurrence interval specified according to ISO8061 notation for durations. Refer https://tc39.es/proposal-temporal/docs/duration.html for all supported values. Recurring campaigns support - daily, weekly, monthly and yearly interval. One can setup a \&quot;interval\&quot; to max of -   - 182 days (That is &#x60;P182D&#x60;)   - 26 weeks (equivalent of 182 days, that is &#x60;P26W&#x60;)   - 24 months (That is &#x60;P24M&#x60;)   - 2 years (That is &#x60;P2Y&#x60;).  Few examples to better understand, how &#x60;interval&#x60; is used. If the &#x60;startDate&#x60; is &#x60;June 28th 2023&#x60; and the intent is to repeat the campaign -   1) Every 25 days, specify \&quot;P25D\&quot;.   2) Every 4 weeks specify \&quot;P4W\&quot;. In this example the &#x60;startDate&#x60; happens to be a &#x60;Wednesday&#x60;, the campaign would run every &#x60;Wednesday&#x60; of 4th week of the month.   3) Every 3 months specify \&quot;P3M\&quot;. In this example the &#x60;startDate&#x60; is on 28th day. Hence, campaigns repeats every 3 months on 28th day.   4) Every 2 years specify \&quot;P2Y\&quot;. In this example, the &#x60;startDate&#x60; is &#x60;Jun 28&#x60;. Hence repeats every 2 yrs on &#x60;Jun 28th&#x60;.    Also note that, the &#x60;interval&#x60; should not conflict &#x60;durationInDays&#x60;. For eg: If &#x60;durationInDays&#x60; is specified as 21 days, but interval is &#x60;P2W&#x60; (Every 2 weeks), it becomes an invalid request, as duration of the campaign conflicts with the next recurrence of the campaign.  | 
**RepeatOnType** | Pointer to [**RecurrenceRepeatOnType**](RecurrenceRepeatOnType.md) |  | [optional] 
**Ends** | Pointer to **time.Time** | An optional field. Specifies when the recurring schedule can have an end. If not specified, the recurring schedule will run forever. | [optional] 

## Methods

### NewRecurrenceDefinitionMutable

`func NewRecurrenceDefinitionMutable(interval string, ) *RecurrenceDefinitionMutable`

NewRecurrenceDefinitionMutable instantiates a new RecurrenceDefinitionMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceDefinitionMutableWithDefaults

`func NewRecurrenceDefinitionMutableWithDefaults() *RecurrenceDefinitionMutable`

NewRecurrenceDefinitionMutableWithDefaults instantiates a new RecurrenceDefinitionMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *RecurrenceDefinitionMutable) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RecurrenceDefinitionMutable) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RecurrenceDefinitionMutable) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetRepeatOnType

`func (o *RecurrenceDefinitionMutable) GetRepeatOnType() RecurrenceRepeatOnType`

GetRepeatOnType returns the RepeatOnType field if non-nil, zero value otherwise.

### GetRepeatOnTypeOk

`func (o *RecurrenceDefinitionMutable) GetRepeatOnTypeOk() (*RecurrenceRepeatOnType, bool)`

GetRepeatOnTypeOk returns a tuple with the RepeatOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatOnType

`func (o *RecurrenceDefinitionMutable) SetRepeatOnType(v RecurrenceRepeatOnType)`

SetRepeatOnType sets RepeatOnType field to given value.

### HasRepeatOnType

`func (o *RecurrenceDefinitionMutable) HasRepeatOnType() bool`

HasRepeatOnType returns a boolean if a field has been set.

### GetEnds

`func (o *RecurrenceDefinitionMutable) GetEnds() time.Time`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *RecurrenceDefinitionMutable) GetEndsOk() (*time.Time, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *RecurrenceDefinitionMutable) SetEnds(v time.Time)`

SetEnds sets Ends field to given value.

### HasEnds

`func (o *RecurrenceDefinitionMutable) HasEnds() bool`

HasEnds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


