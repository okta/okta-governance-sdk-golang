# RequestConditionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequesterSettings** | [**RequesterSettingsCreatableRequesterSettings**](RequesterSettingsCreatableRequesterSettings.md) |  | 
**AccessScopeSettings** | [**AccessScopeSettingsCreatableAccessScopeSettings**](AccessScopeSettingsCreatableAccessScopeSettings.md) |  | 
**AccessDurationSettings** | Pointer to [**AccessDurationSettingsCreatable**](AccessDurationSettingsCreatable.md) |  | [optional] 
**ApprovalSequenceId** | **string** | The ID of the approval sequence | 
**Priority** | Pointer to **int32** | The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority. | [optional] 
**Name** | **string** | Writable unique key on create. Modifiable on update. | 
**Description** | Pointer to **string** | Human readable description | [optional] 

## Methods

### NewRequestConditionCreatable

`func NewRequestConditionCreatable(requesterSettings RequesterSettingsCreatableRequesterSettings, accessScopeSettings AccessScopeSettingsCreatableAccessScopeSettings, approvalSequenceId string, name string, ) *RequestConditionCreatable`

NewRequestConditionCreatable instantiates a new RequestConditionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConditionCreatableWithDefaults

`func NewRequestConditionCreatableWithDefaults() *RequestConditionCreatable`

NewRequestConditionCreatableWithDefaults instantiates a new RequestConditionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequesterSettings

`func (o *RequestConditionCreatable) GetRequesterSettings() RequesterSettingsCreatableRequesterSettings`

GetRequesterSettings returns the RequesterSettings field if non-nil, zero value otherwise.

### GetRequesterSettingsOk

`func (o *RequestConditionCreatable) GetRequesterSettingsOk() (*RequesterSettingsCreatableRequesterSettings, bool)`

GetRequesterSettingsOk returns a tuple with the RequesterSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSettings

`func (o *RequestConditionCreatable) SetRequesterSettings(v RequesterSettingsCreatableRequesterSettings)`

SetRequesterSettings sets RequesterSettings field to given value.


### GetAccessScopeSettings

`func (o *RequestConditionCreatable) GetAccessScopeSettings() AccessScopeSettingsCreatableAccessScopeSettings`

GetAccessScopeSettings returns the AccessScopeSettings field if non-nil, zero value otherwise.

### GetAccessScopeSettingsOk

`func (o *RequestConditionCreatable) GetAccessScopeSettingsOk() (*AccessScopeSettingsCreatableAccessScopeSettings, bool)`

GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScopeSettings

`func (o *RequestConditionCreatable) SetAccessScopeSettings(v AccessScopeSettingsCreatableAccessScopeSettings)`

SetAccessScopeSettings sets AccessScopeSettings field to given value.


### GetAccessDurationSettings

`func (o *RequestConditionCreatable) GetAccessDurationSettings() AccessDurationSettingsCreatable`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RequestConditionCreatable) GetAccessDurationSettingsOk() (*AccessDurationSettingsCreatable, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RequestConditionCreatable) SetAccessDurationSettings(v AccessDurationSettingsCreatable)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RequestConditionCreatable) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### GetApprovalSequenceId

`func (o *RequestConditionCreatable) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RequestConditionCreatable) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RequestConditionCreatable) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.


### GetPriority

`func (o *RequestConditionCreatable) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestConditionCreatable) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestConditionCreatable) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RequestConditionCreatable) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetName

`func (o *RequestConditionCreatable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestConditionCreatable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestConditionCreatable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestConditionCreatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestConditionCreatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestConditionCreatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestConditionCreatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


