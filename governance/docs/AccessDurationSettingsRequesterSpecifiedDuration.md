# AccessDurationSettingsRequesterSpecifiedDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**MaximumDuration** | **string** | The maximum duration set by the requester for access durations. Use the [ISO 8601](https://tc39.es/proposal-temporal/docs/duration.html) duration notation for duration values. You can configure up to 365 days (&#x60;P365D&#x60;) or 52 weeks (&#x60;P52W&#x60;) as the maximum access duration of a request. | 

## Methods

### NewAccessDurationSettingsRequesterSpecifiedDuration

`func NewAccessDurationSettingsRequesterSpecifiedDuration(type_ string, maximumDuration string, ) *AccessDurationSettingsRequesterSpecifiedDuration`

NewAccessDurationSettingsRequesterSpecifiedDuration instantiates a new AccessDurationSettingsRequesterSpecifiedDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDurationSettingsRequesterSpecifiedDurationWithDefaults

`func NewAccessDurationSettingsRequesterSpecifiedDurationWithDefaults() *AccessDurationSettingsRequesterSpecifiedDuration`

NewAccessDurationSettingsRequesterSpecifiedDurationWithDefaults instantiates a new AccessDurationSettingsRequesterSpecifiedDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessDurationSettingsRequesterSpecifiedDuration) SetType(v string)`

SetType sets Type field to given value.


### GetMaximumDuration

`func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetMaximumDuration() string`

GetMaximumDuration returns the MaximumDuration field if non-nil, zero value otherwise.

### GetMaximumDurationOk

`func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetMaximumDurationOk() (*string, bool)`

GetMaximumDurationOk returns a tuple with the MaximumDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDuration

`func (o *AccessDurationSettingsRequesterSpecifiedDuration) SetMaximumDuration(v string)`

SetMaximumDuration sets MaximumDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


