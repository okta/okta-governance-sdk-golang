# DelegateAppointment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegator** | [**DelegateAppointmentDelegator**](DelegateAppointmentDelegator.md) |  | 
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

### NewDelegateAppointment

`func NewDelegateAppointment(delegator DelegateAppointmentDelegator, id string, delegate DelegateAppointmentDelegate, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *DelegateAppointment`

NewDelegateAppointment instantiates a new DelegateAppointment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegateAppointmentWithDefaults

`func NewDelegateAppointmentWithDefaults() *DelegateAppointment`

NewDelegateAppointmentWithDefaults instantiates a new DelegateAppointment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegator

`func (o *DelegateAppointment) GetDelegator() DelegateAppointmentDelegator`

GetDelegator returns the Delegator field if non-nil, zero value otherwise.

### GetDelegatorOk

`func (o *DelegateAppointment) GetDelegatorOk() (*DelegateAppointmentDelegator, bool)`

GetDelegatorOk returns a tuple with the Delegator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegator

`func (o *DelegateAppointment) SetDelegator(v DelegateAppointmentDelegator)`

SetDelegator sets Delegator field to given value.


### GetId

`func (o *DelegateAppointment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegateAppointment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegateAppointment) SetId(v string)`

SetId sets Id field to given value.


### GetDelegate

`func (o *DelegateAppointment) GetDelegate() DelegateAppointmentDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *DelegateAppointment) GetDelegateOk() (*DelegateAppointmentDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *DelegateAppointment) SetDelegate(v DelegateAppointmentDelegate)`

SetDelegate sets Delegate field to given value.


### GetStartTime

`func (o *DelegateAppointment) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DelegateAppointment) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DelegateAppointment) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DelegateAppointment) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DelegateAppointment) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DelegateAppointment) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DelegateAppointment) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DelegateAppointment) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetNote

`func (o *DelegateAppointment) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DelegateAppointment) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DelegateAppointment) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *DelegateAppointment) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DelegateAppointment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DelegateAppointment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DelegateAppointment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *DelegateAppointment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DelegateAppointment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DelegateAppointment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *DelegateAppointment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DelegateAppointment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DelegateAppointment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *DelegateAppointment) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *DelegateAppointment) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *DelegateAppointment) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


