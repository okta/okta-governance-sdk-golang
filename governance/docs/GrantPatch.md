# GrantPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**ScheduleSettings** | [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | 

## Methods

### NewGrantPatch

`func NewGrantPatch(id string, scheduleSettings ScheduleSettingsWriteable, ) *GrantPatch`

NewGrantPatch instantiates a new GrantPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantPatchWithDefaults

`func NewGrantPatchWithDefaults() *GrantPatch`

NewGrantPatchWithDefaults instantiates a new GrantPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GrantPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantPatch) SetId(v string)`

SetId sets Id field to given value.


### GetScheduleSettings

`func (o *GrantPatch) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GrantPatch) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GrantPatch) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


