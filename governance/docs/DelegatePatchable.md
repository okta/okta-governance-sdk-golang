# DelegatePatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegate** | [**DelegateAppointmentDelegate**](DelegateAppointmentDelegate.md) |  | 
**Note** | Pointer to **string** | A note that describes the delegate appointment | [optional] 
**StartTime** | Pointer to **time.Time** | The start time of the delegate appointment, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format | [optional] 
**EndTime** | Pointer to **time.Time** | The time when the delegate appointment expires, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339.html) date and time format | [optional] 

## Methods

### NewDelegatePatchable

`func NewDelegatePatchable(delegate DelegateAppointmentDelegate, ) *DelegatePatchable`

NewDelegatePatchable instantiates a new DelegatePatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatePatchableWithDefaults

`func NewDelegatePatchableWithDefaults() *DelegatePatchable`

NewDelegatePatchableWithDefaults instantiates a new DelegatePatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegate

`func (o *DelegatePatchable) GetDelegate() DelegateAppointmentDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *DelegatePatchable) GetDelegateOk() (*DelegateAppointmentDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *DelegatePatchable) SetDelegate(v DelegateAppointmentDelegate)`

SetDelegate sets Delegate field to given value.


### GetNote

`func (o *DelegatePatchable) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DelegatePatchable) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DelegatePatchable) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *DelegatePatchable) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStartTime

`func (o *DelegatePatchable) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DelegatePatchable) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DelegatePatchable) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DelegatePatchable) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DelegatePatchable) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DelegatePatchable) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DelegatePatchable) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DelegatePatchable) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


