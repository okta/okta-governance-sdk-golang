# AccessDurationSettingsPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Duration** | **string** | The duration set by the admin for access durations. Use the [ISO 8601](https://tc39.es/proposal-temporal/docs/duration.html) duration notation for duration values. You can configure up to 365 days (&#x60;P365D&#x60;) or 52 weeks (&#x60;P52W&#x60;) as the maximum access duration of a request. | 
**MaximumDuration** | **string** | The maximum duration set by the requester for access durations. Use the [ISO 8601](https://tc39.es/proposal-temporal/docs/duration.html) duration notation for duration values. You can configure up to 365 days (&#x60;P365D&#x60;) or 52 weeks (&#x60;P52W&#x60;) as the maximum access duration of a request. | 

## Methods

### NewAccessDurationSettingsPatchable

`func NewAccessDurationSettingsPatchable(type_ string, duration string, maximumDuration string, ) *AccessDurationSettingsPatchable`

NewAccessDurationSettingsPatchable instantiates a new AccessDurationSettingsPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDurationSettingsPatchableWithDefaults

`func NewAccessDurationSettingsPatchableWithDefaults() *AccessDurationSettingsPatchable`

NewAccessDurationSettingsPatchableWithDefaults instantiates a new AccessDurationSettingsPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessDurationSettingsPatchable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessDurationSettingsPatchable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessDurationSettingsPatchable) SetType(v string)`

SetType sets Type field to given value.


### GetDuration

`func (o *AccessDurationSettingsPatchable) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AccessDurationSettingsPatchable) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AccessDurationSettingsPatchable) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMaximumDuration

`func (o *AccessDurationSettingsPatchable) GetMaximumDuration() string`

GetMaximumDuration returns the MaximumDuration field if non-nil, zero value otherwise.

### GetMaximumDurationOk

`func (o *AccessDurationSettingsPatchable) GetMaximumDurationOk() (*string, bool)`

GetMaximumDurationOk returns a tuple with the MaximumDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDuration

`func (o *AccessDurationSettingsPatchable) SetMaximumDuration(v string)`

SetMaximumDuration sets MaximumDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


