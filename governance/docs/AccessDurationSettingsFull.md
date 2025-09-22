# AccessDurationSettingsFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Duration** | **string** | The duration set by the admin for access durations. Use ISO8061 notation for duration values, see https://tc39.es/proposal-temporal/docs/duration.html. You can set up an access duration to a maximum of 72 hours (&#x60;PT72H&#x60;), 90 days (&#x60;90D&#x60;), or 12 weeks (&#x60;P12W&#x60;). For example:    - 24 hours (&#x60;PT24H&#x60;)   - 7 days (&#x60;P7D&#x60;)   - 2 weeks (&#x60;P2W&#x60;)  | 
**MaximumDuration** | **string** | The maximum duration set by the requester for access durations. Use ISO8061 notation for duration values, see https://tc39.es/proposal-temporal/docs/duration.html. The admin sets the maximum duration that can be requested and can&#39;t exceed 72 hours (&#x60;PT72H&#x60;), 90 days (&#x60;90D&#x60;), or 12 weeks (&#x60;P12W&#x60;). For example:    - 24 hours (&#x60;PT24H&#x60;)   - 7 days (&#x60;P7D&#x60;)   - 2 weeks (&#x60;P2W&#x60;)  | 

## Methods

### NewAccessDurationSettingsFull

`func NewAccessDurationSettingsFull(type_ string, duration string, maximumDuration string, ) *AccessDurationSettingsFull`

NewAccessDurationSettingsFull instantiates a new AccessDurationSettingsFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDurationSettingsFullWithDefaults

`func NewAccessDurationSettingsFullWithDefaults() *AccessDurationSettingsFull`

NewAccessDurationSettingsFullWithDefaults instantiates a new AccessDurationSettingsFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessDurationSettingsFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessDurationSettingsFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessDurationSettingsFull) SetType(v string)`

SetType sets Type field to given value.


### GetDuration

`func (o *AccessDurationSettingsFull) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AccessDurationSettingsFull) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AccessDurationSettingsFull) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetMaximumDuration

`func (o *AccessDurationSettingsFull) GetMaximumDuration() string`

GetMaximumDuration returns the MaximumDuration field if non-nil, zero value otherwise.

### GetMaximumDurationOk

`func (o *AccessDurationSettingsFull) GetMaximumDurationOk() (*string, bool)`

GetMaximumDurationOk returns a tuple with the MaximumDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDuration

`func (o *AccessDurationSettingsFull) SetMaximumDuration(v string)`

SetMaximumDuration sets MaximumDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


