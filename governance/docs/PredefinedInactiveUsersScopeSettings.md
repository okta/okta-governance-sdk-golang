# PredefinedInactiveUsersScopeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InactiveDays** | Pointer to **int32** | The duration the users have not used single sign on (SSO) to access their account within the specific time frame. Minimum 30 days and maximum 365 days are supported. | [optional] 

## Methods

### NewPredefinedInactiveUsersScopeSettings

`func NewPredefinedInactiveUsersScopeSettings() *PredefinedInactiveUsersScopeSettings`

NewPredefinedInactiveUsersScopeSettings instantiates a new PredefinedInactiveUsersScopeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedInactiveUsersScopeSettingsWithDefaults

`func NewPredefinedInactiveUsersScopeSettingsWithDefaults() *PredefinedInactiveUsersScopeSettings`

NewPredefinedInactiveUsersScopeSettingsWithDefaults instantiates a new PredefinedInactiveUsersScopeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInactiveDays

`func (o *PredefinedInactiveUsersScopeSettings) GetInactiveDays() int32`

GetInactiveDays returns the InactiveDays field if non-nil, zero value otherwise.

### GetInactiveDaysOk

`func (o *PredefinedInactiveUsersScopeSettings) GetInactiveDaysOk() (*int32, bool)`

GetInactiveDaysOk returns a tuple with the InactiveDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDays

`func (o *PredefinedInactiveUsersScopeSettings) SetInactiveDays(v int32)`

SetInactiveDays sets InactiveDays field to given value.

### HasInactiveDays

`func (o *PredefinedInactiveUsersScopeSettings) HasInactiveDays() bool`

HasInactiveDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


