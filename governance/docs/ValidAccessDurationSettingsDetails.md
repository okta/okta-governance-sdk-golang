# ValidAccessDurationSettingsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedTypes** | Pointer to [**[]ValidAccessDurationType**](ValidAccessDurationType.md) | Access duration settings that are eligible to be added to a request condition or risk settings for the specified resource. | [optional] 
**Required** | Pointer to **bool** | Whether &#x60;accessDurationSetting&#x60; must be included in the request conditions or risk settings for the specified resource. | [optional] 
**MaximumDays** | Pointer to **float32** | The minimum and maximum values allowed in &#x60;accessDurationSettings.duration&#x60; expression for a request condition or risk settings, when that expression contains day units. For example: &#x60;P5D&#x60; is valid, and &#x60;P91D&#x60; is not.  | [optional] 
**MaximumHours** | Pointer to **float32** | The minimum and maximum values allowed in &#x60;accessDurationSettings.duration&#x60; expression for a request condition or risk settings, when that expression contains hour units. For example: &#x60;P5H&#x60; is valid, and &#x60;P73H&#x60; is not.  | [optional] 
**MaximumWeeks** | Pointer to **float32** | The minimum and maximum values allowed in &#x60;accessDurationSettings.duration&#x60; expression for a request condition or risk settings, when that expression contains week units. For example: &#x60;P5W&#x60; is valid, and &#x60;P13W&#x60; is not.  | [optional] 

## Methods

### NewValidAccessDurationSettingsDetails

`func NewValidAccessDurationSettingsDetails() *ValidAccessDurationSettingsDetails`

NewValidAccessDurationSettingsDetails instantiates a new ValidAccessDurationSettingsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidAccessDurationSettingsDetailsWithDefaults

`func NewValidAccessDurationSettingsDetailsWithDefaults() *ValidAccessDurationSettingsDetails`

NewValidAccessDurationSettingsDetailsWithDefaults instantiates a new ValidAccessDurationSettingsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedTypes

`func (o *ValidAccessDurationSettingsDetails) GetSupportedTypes() []ValidAccessDurationType`

GetSupportedTypes returns the SupportedTypes field if non-nil, zero value otherwise.

### GetSupportedTypesOk

`func (o *ValidAccessDurationSettingsDetails) GetSupportedTypesOk() (*[]ValidAccessDurationType, bool)`

GetSupportedTypesOk returns a tuple with the SupportedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTypes

`func (o *ValidAccessDurationSettingsDetails) SetSupportedTypes(v []ValidAccessDurationType)`

SetSupportedTypes sets SupportedTypes field to given value.

### HasSupportedTypes

`func (o *ValidAccessDurationSettingsDetails) HasSupportedTypes() bool`

HasSupportedTypes returns a boolean if a field has been set.

### GetRequired

`func (o *ValidAccessDurationSettingsDetails) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ValidAccessDurationSettingsDetails) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ValidAccessDurationSettingsDetails) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ValidAccessDurationSettingsDetails) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMaximumDays

`func (o *ValidAccessDurationSettingsDetails) GetMaximumDays() float32`

GetMaximumDays returns the MaximumDays field if non-nil, zero value otherwise.

### GetMaximumDaysOk

`func (o *ValidAccessDurationSettingsDetails) GetMaximumDaysOk() (*float32, bool)`

GetMaximumDaysOk returns a tuple with the MaximumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDays

`func (o *ValidAccessDurationSettingsDetails) SetMaximumDays(v float32)`

SetMaximumDays sets MaximumDays field to given value.

### HasMaximumDays

`func (o *ValidAccessDurationSettingsDetails) HasMaximumDays() bool`

HasMaximumDays returns a boolean if a field has been set.

### GetMaximumHours

`func (o *ValidAccessDurationSettingsDetails) GetMaximumHours() float32`

GetMaximumHours returns the MaximumHours field if non-nil, zero value otherwise.

### GetMaximumHoursOk

`func (o *ValidAccessDurationSettingsDetails) GetMaximumHoursOk() (*float32, bool)`

GetMaximumHoursOk returns a tuple with the MaximumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHours

`func (o *ValidAccessDurationSettingsDetails) SetMaximumHours(v float32)`

SetMaximumHours sets MaximumHours field to given value.

### HasMaximumHours

`func (o *ValidAccessDurationSettingsDetails) HasMaximumHours() bool`

HasMaximumHours returns a boolean if a field has been set.

### GetMaximumWeeks

`func (o *ValidAccessDurationSettingsDetails) GetMaximumWeeks() float32`

GetMaximumWeeks returns the MaximumWeeks field if non-nil, zero value otherwise.

### GetMaximumWeeksOk

`func (o *ValidAccessDurationSettingsDetails) GetMaximumWeeksOk() (*float32, bool)`

GetMaximumWeeksOk returns a tuple with the MaximumWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumWeeks

`func (o *ValidAccessDurationSettingsDetails) SetMaximumWeeks(v float32)`

SetMaximumWeeks sets MaximumWeeks field to given value.

### HasMaximumWeeks

`func (o *ValidAccessDurationSettingsDetails) HasMaximumWeeks() bool`

HasMaximumWeeks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


