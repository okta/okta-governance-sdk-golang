# TaskSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the task | 
**Assignees** | [**[]TaskAssignees**](TaskAssignees.md) | List of task assignees that perform actions on the task | 
**Status** | [**TaskStatus**](TaskStatus.md) |  | 
**Label** | **string** | Human readable label for the task | 
**RequestId** | **string** | The request ID associated with the task | 
**CreatedAt** | **time.Time** | The ISO 8601 formatted date and time when the object was created | [readonly] 
**UpdatedAt** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**Type** | [**TaskType**](TaskType.md) |  | 
**Value** | Pointer to **NullableString** | The value of the task completion that&#39;s determined by the assignee | [optional] 

## Methods

### NewTaskSparse

`func NewTaskSparse(id string, assignees []TaskAssignees, status TaskStatus, label string, requestId string, createdAt time.Time, updatedAt time.Time, type_ TaskType, ) *TaskSparse`

NewTaskSparse instantiates a new TaskSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSparseWithDefaults

`func NewTaskSparseWithDefaults() *TaskSparse`

NewTaskSparseWithDefaults instantiates a new TaskSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskSparse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskSparse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskSparse) SetId(v string)`

SetId sets Id field to given value.


### GetAssignees

`func (o *TaskSparse) GetAssignees() []TaskAssignees`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *TaskSparse) GetAssigneesOk() (*[]TaskAssignees, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *TaskSparse) SetAssignees(v []TaskAssignees)`

SetAssignees sets Assignees field to given value.


### GetStatus

`func (o *TaskSparse) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskSparse) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskSparse) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetLabel

`func (o *TaskSparse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TaskSparse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TaskSparse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequestId

`func (o *TaskSparse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TaskSparse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TaskSparse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetCreatedAt

`func (o *TaskSparse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskSparse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskSparse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaskSparse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskSparse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskSparse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *TaskSparse) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskSparse) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskSparse) SetType(v TaskType)`

SetType sets Type field to given value.


### GetValue

`func (o *TaskSparse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaskSparse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaskSparse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaskSparse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TaskSparse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TaskSparse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


