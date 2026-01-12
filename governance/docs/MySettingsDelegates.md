# MySettingsDelegates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appointments** | [**[]DelegateReadonly**](DelegateReadonly.md) | My delegate appointments | 
**Permissions** | Pointer to **[]string** | My delegate permission settings  | Permission | Description | |------------|-------------| | &#x60;READ&#x60; | I can view my delegates | | &#x60;WRITE&#x60; | I can view and set my own delegates |  | [optional] 

## Methods

### NewMySettingsDelegates

`func NewMySettingsDelegates(appointments []DelegateReadonly, ) *MySettingsDelegates`

NewMySettingsDelegates instantiates a new MySettingsDelegates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySettingsDelegatesWithDefaults

`func NewMySettingsDelegatesWithDefaults() *MySettingsDelegates`

NewMySettingsDelegatesWithDefaults instantiates a new MySettingsDelegates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppointments

`func (o *MySettingsDelegates) GetAppointments() []DelegateReadonly`

GetAppointments returns the Appointments field if non-nil, zero value otherwise.

### GetAppointmentsOk

`func (o *MySettingsDelegates) GetAppointmentsOk() (*[]DelegateReadonly, bool)`

GetAppointmentsOk returns a tuple with the Appointments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppointments

`func (o *MySettingsDelegates) SetAppointments(v []DelegateReadonly)`

SetAppointments sets Appointments field to given value.


### GetPermissions

`func (o *MySettingsDelegates) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MySettingsDelegates) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MySettingsDelegates) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MySettingsDelegates) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


