# MySettingsGetDelegateReadonly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the delegate appointment | 
**Delegate** | [**DelegateAppointmentDelegate**](DelegateAppointmentDelegate.md) |  | 
**StartTime** | Pointer to **time.Time** | The start time of the delegate appointment, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format | [optional] 
**EndTime** | Pointer to **time.Time** | The time when the delegate appointment expires, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format | [optional] 
**Note** | Pointer to **string** | A note that describes the delegate appointment | [optional] 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewMySettingsGetDelegateReadonly

`func NewMySettingsGetDelegateReadonly(id string, delegate DelegateAppointmentDelegate, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *MySettingsGetDelegateReadonly`

NewMySettingsGetDelegateReadonly instantiates a new MySettingsGetDelegateReadonly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySettingsGetDelegateReadonlyWithDefaults

`func NewMySettingsGetDelegateReadonlyWithDefaults() *MySettingsGetDelegateReadonly`

NewMySettingsGetDelegateReadonlyWithDefaults instantiates a new MySettingsGetDelegateReadonly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MySettingsGetDelegateReadonly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MySettingsGetDelegateReadonly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MySettingsGetDelegateReadonly) SetId(v string)`

SetId sets Id field to given value.


### GetDelegate

`func (o *MySettingsGetDelegateReadonly) GetDelegate() DelegateAppointmentDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *MySettingsGetDelegateReadonly) GetDelegateOk() (*DelegateAppointmentDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *MySettingsGetDelegateReadonly) SetDelegate(v DelegateAppointmentDelegate)`

SetDelegate sets Delegate field to given value.


### GetStartTime

`func (o *MySettingsGetDelegateReadonly) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MySettingsGetDelegateReadonly) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MySettingsGetDelegateReadonly) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MySettingsGetDelegateReadonly) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *MySettingsGetDelegateReadonly) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MySettingsGetDelegateReadonly) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MySettingsGetDelegateReadonly) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MySettingsGetDelegateReadonly) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetNote

`func (o *MySettingsGetDelegateReadonly) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *MySettingsGetDelegateReadonly) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *MySettingsGetDelegateReadonly) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *MySettingsGetDelegateReadonly) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MySettingsGetDelegateReadonly) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MySettingsGetDelegateReadonly) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MySettingsGetDelegateReadonly) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *MySettingsGetDelegateReadonly) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MySettingsGetDelegateReadonly) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MySettingsGetDelegateReadonly) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *MySettingsGetDelegateReadonly) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MySettingsGetDelegateReadonly) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MySettingsGetDelegateReadonly) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *MySettingsGetDelegateReadonly) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *MySettingsGetDelegateReadonly) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *MySettingsGetDelegateReadonly) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


