# TaskFull

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
**IsEscalated** | Pointer to **bool** | Indicates whether a task is escalated to another user. See [Escalate Tasks](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;csh-escl-task) in the product documentation. | [optional] 
**IsDelegated** | Pointer to **bool** | Indicates whether a task is delegated to another user. See [Governance delegates](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;csh-governance-delegates) in the product documentation. | [optional] 
**OriginalAssigneeId** | Pointer to **NullableString** | ID of the original assignee before delegation or escalation | [optional] 
**CompletedBy** | Pointer to [**TaskCompletedBy**](TaskCompletedBy.md) |  | [optional] 

## Methods

### NewTaskFull

`func NewTaskFull(id string, assignees []TaskAssignees, status TaskStatus, label string, requestId string, createdAt time.Time, updatedAt time.Time, type_ TaskType, ) *TaskFull`

NewTaskFull instantiates a new TaskFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskFullWithDefaults

`func NewTaskFullWithDefaults() *TaskFull`

NewTaskFullWithDefaults instantiates a new TaskFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskFull) SetId(v string)`

SetId sets Id field to given value.


### GetAssignees

`func (o *TaskFull) GetAssignees() []TaskAssignees`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *TaskFull) GetAssigneesOk() (*[]TaskAssignees, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *TaskFull) SetAssignees(v []TaskAssignees)`

SetAssignees sets Assignees field to given value.


### GetStatus

`func (o *TaskFull) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskFull) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskFull) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetLabel

`func (o *TaskFull) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TaskFull) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TaskFull) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequestId

`func (o *TaskFull) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TaskFull) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TaskFull) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetCreatedAt

`func (o *TaskFull) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskFull) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskFull) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaskFull) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskFull) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskFull) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *TaskFull) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskFull) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskFull) SetType(v TaskType)`

SetType sets Type field to given value.


### GetValue

`func (o *TaskFull) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaskFull) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaskFull) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaskFull) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TaskFull) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TaskFull) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEscalated

`func (o *TaskFull) GetIsEscalated() bool`

GetIsEscalated returns the IsEscalated field if non-nil, zero value otherwise.

### GetIsEscalatedOk

`func (o *TaskFull) GetIsEscalatedOk() (*bool, bool)`

GetIsEscalatedOk returns a tuple with the IsEscalated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEscalated

`func (o *TaskFull) SetIsEscalated(v bool)`

SetIsEscalated sets IsEscalated field to given value.

### HasIsEscalated

`func (o *TaskFull) HasIsEscalated() bool`

HasIsEscalated returns a boolean if a field has been set.

### GetIsDelegated

`func (o *TaskFull) GetIsDelegated() bool`

GetIsDelegated returns the IsDelegated field if non-nil, zero value otherwise.

### GetIsDelegatedOk

`func (o *TaskFull) GetIsDelegatedOk() (*bool, bool)`

GetIsDelegatedOk returns a tuple with the IsDelegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelegated

`func (o *TaskFull) SetIsDelegated(v bool)`

SetIsDelegated sets IsDelegated field to given value.

### HasIsDelegated

`func (o *TaskFull) HasIsDelegated() bool`

HasIsDelegated returns a boolean if a field has been set.

### GetOriginalAssigneeId

`func (o *TaskFull) GetOriginalAssigneeId() string`

GetOriginalAssigneeId returns the OriginalAssigneeId field if non-nil, zero value otherwise.

### GetOriginalAssigneeIdOk

`func (o *TaskFull) GetOriginalAssigneeIdOk() (*string, bool)`

GetOriginalAssigneeIdOk returns a tuple with the OriginalAssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAssigneeId

`func (o *TaskFull) SetOriginalAssigneeId(v string)`

SetOriginalAssigneeId sets OriginalAssigneeId field to given value.

### HasOriginalAssigneeId

`func (o *TaskFull) HasOriginalAssigneeId() bool`

HasOriginalAssigneeId returns a boolean if a field has been set.

### SetOriginalAssigneeIdNil

`func (o *TaskFull) SetOriginalAssigneeIdNil(b bool)`

 SetOriginalAssigneeIdNil sets the value for OriginalAssigneeId to be an explicit nil

### UnsetOriginalAssigneeId
`func (o *TaskFull) UnsetOriginalAssigneeId()`

UnsetOriginalAssigneeId ensures that no value is present for OriginalAssigneeId, not even an explicit nil
### GetCompletedBy

`func (o *TaskFull) GetCompletedBy() TaskCompletedBy`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *TaskFull) GetCompletedByOk() (*TaskCompletedBy, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *TaskFull) SetCompletedBy(v TaskCompletedBy)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *TaskFull) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


