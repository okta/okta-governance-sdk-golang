# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the operation | 
**Type** | **string** | The operation type | 
**Status** | [**OperationStatus**](OperationStatus.md) |  | 
**Created** | Pointer to **time.Time** | The ISO 8601 formatted date and time of when the operation was created | [optional] [readonly] 
**Completed** | Pointer to **time.Time** | The ISO 8601 formatted date and time of when the operation completed | [optional] [readonly] 
**TimeElapsedInSeconds** | Pointer to **int32** | Elapsed time since the start of the operation, in seconds | [optional] 
**Links** | [**OperationLinks**](OperationLinks.md) |  | 

## Methods

### NewOperation

`func NewOperation(id string, type_ string, status OperationStatus, links OperationLinks, ) *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Operation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Operation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Operation) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Operation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Operation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Operation) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Operation) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Operation) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Operation) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *Operation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Operation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Operation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Operation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCompleted

`func (o *Operation) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *Operation) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *Operation) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *Operation) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetTimeElapsedInSeconds

`func (o *Operation) GetTimeElapsedInSeconds() int32`

GetTimeElapsedInSeconds returns the TimeElapsedInSeconds field if non-nil, zero value otherwise.

### GetTimeElapsedInSecondsOk

`func (o *Operation) GetTimeElapsedInSecondsOk() (*int32, bool)`

GetTimeElapsedInSecondsOk returns a tuple with the TimeElapsedInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeElapsedInSeconds

`func (o *Operation) SetTimeElapsedInSeconds(v int32)`

SetTimeElapsedInSeconds sets TimeElapsedInSeconds field to given value.

### HasTimeElapsedInSeconds

`func (o *Operation) HasTimeElapsedInSeconds() bool`

HasTimeElapsedInSeconds returns a boolean if a field has been set.

### GetLinks

`func (o *Operation) GetLinks() OperationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Operation) GetLinksOk() (*OperationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Operation) SetLinks(v OperationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


