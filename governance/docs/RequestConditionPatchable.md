# RequestConditionPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequesterSettings** | Pointer to [**RequesterSettingsCreatableRequesterSettings**](RequesterSettingsCreatableRequesterSettings.md) |  | [optional] 
**AccessScopeSettings** | Pointer to [**AccessScopeSettingsCreatableAccessScopeSettings**](AccessScopeSettingsCreatableAccessScopeSettings.md) |  | [optional] 
**AccessDurationSettings** | Pointer to [**NullableAccessDurationSettingsPatchable**](AccessDurationSettingsPatchable.md) |  | [optional] 
**ApprovalSequenceId** | Pointer to **string** | The ID of the approval sequence | [optional] 
**Priority** | Pointer to **int32** | The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority. | [optional] 
**Name** | Pointer to **string** | Writable unique key on create. Modifiable on update. | [optional] 
**Description** | Pointer to **string** | Human readable description | [optional] 

## Methods

### NewRequestConditionPatchable

`func NewRequestConditionPatchable() *RequestConditionPatchable`

NewRequestConditionPatchable instantiates a new RequestConditionPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConditionPatchableWithDefaults

`func NewRequestConditionPatchableWithDefaults() *RequestConditionPatchable`

NewRequestConditionPatchableWithDefaults instantiates a new RequestConditionPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequesterSettings

`func (o *RequestConditionPatchable) GetRequesterSettings() RequesterSettingsCreatableRequesterSettings`

GetRequesterSettings returns the RequesterSettings field if non-nil, zero value otherwise.

### GetRequesterSettingsOk

`func (o *RequestConditionPatchable) GetRequesterSettingsOk() (*RequesterSettingsCreatableRequesterSettings, bool)`

GetRequesterSettingsOk returns a tuple with the RequesterSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSettings

`func (o *RequestConditionPatchable) SetRequesterSettings(v RequesterSettingsCreatableRequesterSettings)`

SetRequesterSettings sets RequesterSettings field to given value.

### HasRequesterSettings

`func (o *RequestConditionPatchable) HasRequesterSettings() bool`

HasRequesterSettings returns a boolean if a field has been set.

### GetAccessScopeSettings

`func (o *RequestConditionPatchable) GetAccessScopeSettings() AccessScopeSettingsCreatableAccessScopeSettings`

GetAccessScopeSettings returns the AccessScopeSettings field if non-nil, zero value otherwise.

### GetAccessScopeSettingsOk

`func (o *RequestConditionPatchable) GetAccessScopeSettingsOk() (*AccessScopeSettingsCreatableAccessScopeSettings, bool)`

GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScopeSettings

`func (o *RequestConditionPatchable) SetAccessScopeSettings(v AccessScopeSettingsCreatableAccessScopeSettings)`

SetAccessScopeSettings sets AccessScopeSettings field to given value.

### HasAccessScopeSettings

`func (o *RequestConditionPatchable) HasAccessScopeSettings() bool`

HasAccessScopeSettings returns a boolean if a field has been set.

### GetAccessDurationSettings

`func (o *RequestConditionPatchable) GetAccessDurationSettings() AccessDurationSettingsPatchable`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RequestConditionPatchable) GetAccessDurationSettingsOk() (*AccessDurationSettingsPatchable, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RequestConditionPatchable) SetAccessDurationSettings(v AccessDurationSettingsPatchable)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RequestConditionPatchable) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### SetAccessDurationSettingsNil

`func (o *RequestConditionPatchable) SetAccessDurationSettingsNil(b bool)`

 SetAccessDurationSettingsNil sets the value for AccessDurationSettings to be an explicit nil

### UnsetAccessDurationSettings
`func (o *RequestConditionPatchable) UnsetAccessDurationSettings()`

UnsetAccessDurationSettings ensures that no value is present for AccessDurationSettings, not even an explicit nil
### GetApprovalSequenceId

`func (o *RequestConditionPatchable) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RequestConditionPatchable) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RequestConditionPatchable) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.

### HasApprovalSequenceId

`func (o *RequestConditionPatchable) HasApprovalSequenceId() bool`

HasApprovalSequenceId returns a boolean if a field has been set.

### GetPriority

`func (o *RequestConditionPatchable) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestConditionPatchable) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestConditionPatchable) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RequestConditionPatchable) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetName

`func (o *RequestConditionPatchable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestConditionPatchable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestConditionPatchable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestConditionPatchable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RequestConditionPatchable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestConditionPatchable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestConditionPatchable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestConditionPatchable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


