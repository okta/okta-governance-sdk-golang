# RequestConditionFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequesterSettings** | [**RequesterSettingsFullRequesterSettings**](RequesterSettingsFullRequesterSettings.md) |  | 
**AccessScopeSettings** | [**AccessScopeSettingsFullAccessScopeSettings**](AccessScopeSettingsFullAccessScopeSettings.md) |  | 
**AccessDurationSettings** | Pointer to [**AccessDurationSettingsFull**](AccessDurationSettingsFull.md) |  | [optional] 
**ApprovalSequenceId** | Pointer to **string** | If an approval sequence was deleted, then conditions referencing it will become invalid and the approvalSequenceId will not be present. | [optional] 
**Name** | **string** | Writable unique key on create. Modifiable on update. | 
**Description** | Pointer to **string** | Human readable description | [optional] 
**Id** | **string** | The ID of the request condition | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**RequestConditionLinks**](RequestConditionLinks.md) |  | 
**Status** | [**RequestConditionStatus**](RequestConditionStatus.md) |  | 
**Priority** | **int32** | The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority. | 

## Methods

### NewRequestConditionFull

`func NewRequestConditionFull(requesterSettings RequesterSettingsFullRequesterSettings, accessScopeSettings AccessScopeSettingsFullAccessScopeSettings, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestConditionLinks, status RequestConditionStatus, priority int32, ) *RequestConditionFull`

NewRequestConditionFull instantiates a new RequestConditionFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConditionFullWithDefaults

`func NewRequestConditionFullWithDefaults() *RequestConditionFull`

NewRequestConditionFullWithDefaults instantiates a new RequestConditionFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequesterSettings

`func (o *RequestConditionFull) GetRequesterSettings() RequesterSettingsFullRequesterSettings`

GetRequesterSettings returns the RequesterSettings field if non-nil, zero value otherwise.

### GetRequesterSettingsOk

`func (o *RequestConditionFull) GetRequesterSettingsOk() (*RequesterSettingsFullRequesterSettings, bool)`

GetRequesterSettingsOk returns a tuple with the RequesterSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSettings

`func (o *RequestConditionFull) SetRequesterSettings(v RequesterSettingsFullRequesterSettings)`

SetRequesterSettings sets RequesterSettings field to given value.


### GetAccessScopeSettings

`func (o *RequestConditionFull) GetAccessScopeSettings() AccessScopeSettingsFullAccessScopeSettings`

GetAccessScopeSettings returns the AccessScopeSettings field if non-nil, zero value otherwise.

### GetAccessScopeSettingsOk

`func (o *RequestConditionFull) GetAccessScopeSettingsOk() (*AccessScopeSettingsFullAccessScopeSettings, bool)`

GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScopeSettings

`func (o *RequestConditionFull) SetAccessScopeSettings(v AccessScopeSettingsFullAccessScopeSettings)`

SetAccessScopeSettings sets AccessScopeSettings field to given value.


### GetAccessDurationSettings

`func (o *RequestConditionFull) GetAccessDurationSettings() AccessDurationSettingsFull`

GetAccessDurationSettings returns the AccessDurationSettings field if non-nil, zero value otherwise.

### GetAccessDurationSettingsOk

`func (o *RequestConditionFull) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool)`

GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDurationSettings

`func (o *RequestConditionFull) SetAccessDurationSettings(v AccessDurationSettingsFull)`

SetAccessDurationSettings sets AccessDurationSettings field to given value.

### HasAccessDurationSettings

`func (o *RequestConditionFull) HasAccessDurationSettings() bool`

HasAccessDurationSettings returns a boolean if a field has been set.

### GetApprovalSequenceId

`func (o *RequestConditionFull) GetApprovalSequenceId() string`

GetApprovalSequenceId returns the ApprovalSequenceId field if non-nil, zero value otherwise.

### GetApprovalSequenceIdOk

`func (o *RequestConditionFull) GetApprovalSequenceIdOk() (*string, bool)`

GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSequenceId

`func (o *RequestConditionFull) SetApprovalSequenceId(v string)`

SetApprovalSequenceId sets ApprovalSequenceId field to given value.

### HasApprovalSequenceId

`func (o *RequestConditionFull) HasApprovalSequenceId() bool`

HasApprovalSequenceId returns a boolean if a field has been set.

### GetName

`func (o *RequestConditionFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestConditionFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestConditionFull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestConditionFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestConditionFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestConditionFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestConditionFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RequestConditionFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestConditionFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestConditionFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestConditionFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestConditionFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestConditionFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestConditionFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestConditionFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestConditionFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestConditionFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestConditionFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestConditionFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestConditionFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestConditionFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestConditionFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *RequestConditionFull) GetLinks() RequestConditionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestConditionFull) GetLinksOk() (*RequestConditionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestConditionFull) SetLinks(v RequestConditionLinks)`

SetLinks sets Links field to given value.


### GetStatus

`func (o *RequestConditionFull) GetStatus() RequestConditionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestConditionFull) GetStatusOk() (*RequestConditionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestConditionFull) SetStatus(v RequestConditionStatus)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *RequestConditionFull) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestConditionFull) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestConditionFull) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


