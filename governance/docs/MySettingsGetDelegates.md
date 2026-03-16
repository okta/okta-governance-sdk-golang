# MySettingsGetDelegates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appointments** | [**[]MySettingsGetDelegateReadonly**](MySettingsGetDelegateReadonly.md) | My delegate appointments | 
**Permissions** | Pointer to **[]string** | My delegate permission settings  | Permission | Description | |------------|-------------| | &#x60;READ&#x60; | I can view my delegates | | &#x60;WRITE&#x60; | I can view and set my own delegates |  | [optional] 

## Methods

### NewMySettingsGetDelegates

`func NewMySettingsGetDelegates(appointments []MySettingsGetDelegateReadonly, ) *MySettingsGetDelegates`

NewMySettingsGetDelegates instantiates a new MySettingsGetDelegates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySettingsGetDelegatesWithDefaults

`func NewMySettingsGetDelegatesWithDefaults() *MySettingsGetDelegates`

NewMySettingsGetDelegatesWithDefaults instantiates a new MySettingsGetDelegates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppointments

`func (o *MySettingsGetDelegates) GetAppointments() []MySettingsGetDelegateReadonly`

GetAppointments returns the Appointments field if non-nil, zero value otherwise.

### GetAppointmentsOk

`func (o *MySettingsGetDelegates) GetAppointmentsOk() (*[]MySettingsGetDelegateReadonly, bool)`

GetAppointmentsOk returns a tuple with the Appointments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppointments

`func (o *MySettingsGetDelegates) SetAppointments(v []MySettingsGetDelegateReadonly)`

SetAppointments sets Appointments field to given value.


### GetPermissions

`func (o *MySettingsGetDelegates) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MySettingsGetDelegates) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MySettingsGetDelegates) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MySettingsGetDelegates) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


