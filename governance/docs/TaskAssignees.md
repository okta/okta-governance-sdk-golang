# TaskAssignees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The Okta ID | 
**Type** | [**PrincipalType**](PrincipalType.md) |  | 

## Methods

### NewTaskAssignees

`func NewTaskAssignees(externalId string, type_ PrincipalType, ) *TaskAssignees`

NewTaskAssignees instantiates a new TaskAssignees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskAssigneesWithDefaults

`func NewTaskAssigneesWithDefaults() *TaskAssignees`

NewTaskAssigneesWithDefaults instantiates a new TaskAssignees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *TaskAssignees) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TaskAssignees) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TaskAssignees) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *TaskAssignees) GetType() PrincipalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskAssignees) GetTypeOk() (*PrincipalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskAssignees) SetType(v PrincipalType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


