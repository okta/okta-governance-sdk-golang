# UpdateTaskV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]TaskAssignees**](TaskAssignees.md) | List of task assignees to perform actions on the task. A maximum of 10 assignees can be assigned to a task. | [optional] 

## Methods

### NewUpdateTaskV2Request

`func NewUpdateTaskV2Request() *UpdateTaskV2Request`

NewUpdateTaskV2Request instantiates a new UpdateTaskV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskV2RequestWithDefaults

`func NewUpdateTaskV2RequestWithDefaults() *UpdateTaskV2Request`

NewUpdateTaskV2RequestWithDefaults instantiates a new UpdateTaskV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *UpdateTaskV2Request) GetAssignees() []TaskAssignees`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *UpdateTaskV2Request) GetAssigneesOk() (*[]TaskAssignees, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *UpdateTaskV2Request) SetAssignees(v []TaskAssignees)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *UpdateTaskV2Request) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


