# TaskCompletedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The Okta ID | 
**Type** | [**PrincipalType**](PrincipalType.md) |  | 

## Methods

### NewTaskCompletedBy

`func NewTaskCompletedBy(externalId string, type_ PrincipalType, ) *TaskCompletedBy`

NewTaskCompletedBy instantiates a new TaskCompletedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCompletedByWithDefaults

`func NewTaskCompletedByWithDefaults() *TaskCompletedBy`

NewTaskCompletedByWithDefaults instantiates a new TaskCompletedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *TaskCompletedBy) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TaskCompletedBy) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TaskCompletedBy) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *TaskCompletedBy) GetType() PrincipalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskCompletedBy) GetTypeOk() (*PrincipalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskCompletedBy) SetType(v PrincipalType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


