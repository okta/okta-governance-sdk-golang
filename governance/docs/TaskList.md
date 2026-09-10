# TaskList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TaskSparse**](TaskSparse.md) | All tasks on the current page | [optional] 
**Links** | Pointer to [**TaskListLinks**](TaskListLinks.md) |  | [optional] 

## Methods

### NewTaskList

`func NewTaskList() *TaskList`

NewTaskList instantiates a new TaskList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskListWithDefaults

`func NewTaskListWithDefaults() *TaskList`

NewTaskListWithDefaults instantiates a new TaskList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaskList) GetData() []TaskSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaskList) GetDataOk() (*[]TaskSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaskList) SetData(v []TaskSparse)`

SetData sets Data field to given value.

### HasData

`func (o *TaskList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *TaskList) GetLinks() TaskListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskList) GetLinksOk() (*TaskListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskList) SetLinks(v TaskListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


