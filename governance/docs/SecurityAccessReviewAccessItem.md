# SecurityAccessReviewAccessItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the access item | 
**Type** | [**SecurityAccessReviewAccessItemType**](SecurityAccessReviewAccessItemType.md) |  | 
**Name** | **string** | The name of the access item | 
**ResourceId** | **string** | The ID of the resource for the access item in a security access review | 
**Severity** | [**SecurityAccessReviewAccessItemSeverity**](SecurityAccessReviewAccessItemSeverity.md) |  | 
**SupportedActions** | [**[]SecurityAccessReviewAccessItemSupportedAction**](SecurityAccessReviewAccessItemSupportedAction.md) |  | 
**RemediationStatus** | Pointer to [**SecurityAccessReviewAccessItemRemediationStatus**](SecurityAccessReviewAccessItemRemediationStatus.md) |  | [optional] 
**ManualRemediationTypes** | Pointer to [**[]SecurityAccessReviewAccessItemManualRemediationType**](SecurityAccessReviewAccessItemManualRemediationType.md) | The reasons for manual remediation | [optional] 
**AppInfo** | Pointer to [**SecurityAccessReviewAccessItemAppInfo**](SecurityAccessReviewAccessItemAppInfo.md) |  | [optional] 
**Summary** | Pointer to [**AiMessage**](AiMessage.md) |  | [optional] 
**SubAccessTypes** | Pointer to [**[]SecurityAccessReviewSubAccessItemType**](SecurityAccessReviewSubAccessItemType.md) | A list of sub-access types for the access item | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | Links related to the access item | [optional] 

## Methods

### NewSecurityAccessReviewAccessItem

`func NewSecurityAccessReviewAccessItem(id string, type_ SecurityAccessReviewAccessItemType, name string, resourceId string, severity SecurityAccessReviewAccessItemSeverity, supportedActions []SecurityAccessReviewAccessItemSupportedAction, ) *SecurityAccessReviewAccessItem`

NewSecurityAccessReviewAccessItem instantiates a new SecurityAccessReviewAccessItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewAccessItemWithDefaults

`func NewSecurityAccessReviewAccessItemWithDefaults() *SecurityAccessReviewAccessItem`

NewSecurityAccessReviewAccessItemWithDefaults instantiates a new SecurityAccessReviewAccessItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewAccessItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewAccessItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewAccessItem) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SecurityAccessReviewAccessItem) GetType() SecurityAccessReviewAccessItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewAccessItem) GetTypeOk() (*SecurityAccessReviewAccessItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewAccessItem) SetType(v SecurityAccessReviewAccessItemType)`

SetType sets Type field to given value.


### GetName

`func (o *SecurityAccessReviewAccessItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityAccessReviewAccessItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityAccessReviewAccessItem) SetName(v string)`

SetName sets Name field to given value.


### GetResourceId

`func (o *SecurityAccessReviewAccessItem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *SecurityAccessReviewAccessItem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *SecurityAccessReviewAccessItem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetSeverity

`func (o *SecurityAccessReviewAccessItem) GetSeverity() SecurityAccessReviewAccessItemSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityAccessReviewAccessItem) GetSeverityOk() (*SecurityAccessReviewAccessItemSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityAccessReviewAccessItem) SetSeverity(v SecurityAccessReviewAccessItemSeverity)`

SetSeverity sets Severity field to given value.


### GetSupportedActions

`func (o *SecurityAccessReviewAccessItem) GetSupportedActions() []SecurityAccessReviewAccessItemSupportedAction`

GetSupportedActions returns the SupportedActions field if non-nil, zero value otherwise.

### GetSupportedActionsOk

`func (o *SecurityAccessReviewAccessItem) GetSupportedActionsOk() (*[]SecurityAccessReviewAccessItemSupportedAction, bool)`

GetSupportedActionsOk returns a tuple with the SupportedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedActions

`func (o *SecurityAccessReviewAccessItem) SetSupportedActions(v []SecurityAccessReviewAccessItemSupportedAction)`

SetSupportedActions sets SupportedActions field to given value.


### GetRemediationStatus

`func (o *SecurityAccessReviewAccessItem) GetRemediationStatus() SecurityAccessReviewAccessItemRemediationStatus`

GetRemediationStatus returns the RemediationStatus field if non-nil, zero value otherwise.

### GetRemediationStatusOk

`func (o *SecurityAccessReviewAccessItem) GetRemediationStatusOk() (*SecurityAccessReviewAccessItemRemediationStatus, bool)`

GetRemediationStatusOk returns a tuple with the RemediationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationStatus

`func (o *SecurityAccessReviewAccessItem) SetRemediationStatus(v SecurityAccessReviewAccessItemRemediationStatus)`

SetRemediationStatus sets RemediationStatus field to given value.

### HasRemediationStatus

`func (o *SecurityAccessReviewAccessItem) HasRemediationStatus() bool`

HasRemediationStatus returns a boolean if a field has been set.

### GetManualRemediationTypes

`func (o *SecurityAccessReviewAccessItem) GetManualRemediationTypes() []SecurityAccessReviewAccessItemManualRemediationType`

GetManualRemediationTypes returns the ManualRemediationTypes field if non-nil, zero value otherwise.

### GetManualRemediationTypesOk

`func (o *SecurityAccessReviewAccessItem) GetManualRemediationTypesOk() (*[]SecurityAccessReviewAccessItemManualRemediationType, bool)`

GetManualRemediationTypesOk returns a tuple with the ManualRemediationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRemediationTypes

`func (o *SecurityAccessReviewAccessItem) SetManualRemediationTypes(v []SecurityAccessReviewAccessItemManualRemediationType)`

SetManualRemediationTypes sets ManualRemediationTypes field to given value.

### HasManualRemediationTypes

`func (o *SecurityAccessReviewAccessItem) HasManualRemediationTypes() bool`

HasManualRemediationTypes returns a boolean if a field has been set.

### GetAppInfo

`func (o *SecurityAccessReviewAccessItem) GetAppInfo() SecurityAccessReviewAccessItemAppInfo`

GetAppInfo returns the AppInfo field if non-nil, zero value otherwise.

### GetAppInfoOk

`func (o *SecurityAccessReviewAccessItem) GetAppInfoOk() (*SecurityAccessReviewAccessItemAppInfo, bool)`

GetAppInfoOk returns a tuple with the AppInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfo

`func (o *SecurityAccessReviewAccessItem) SetAppInfo(v SecurityAccessReviewAccessItemAppInfo)`

SetAppInfo sets AppInfo field to given value.

### HasAppInfo

`func (o *SecurityAccessReviewAccessItem) HasAppInfo() bool`

HasAppInfo returns a boolean if a field has been set.

### GetSummary

`func (o *SecurityAccessReviewAccessItem) GetSummary() AiMessage`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SecurityAccessReviewAccessItem) GetSummaryOk() (*AiMessage, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SecurityAccessReviewAccessItem) SetSummary(v AiMessage)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SecurityAccessReviewAccessItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSubAccessTypes

`func (o *SecurityAccessReviewAccessItem) GetSubAccessTypes() []SecurityAccessReviewSubAccessItemType`

GetSubAccessTypes returns the SubAccessTypes field if non-nil, zero value otherwise.

### GetSubAccessTypesOk

`func (o *SecurityAccessReviewAccessItem) GetSubAccessTypesOk() (*[]SecurityAccessReviewSubAccessItemType, bool)`

GetSubAccessTypesOk returns a tuple with the SubAccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccessTypes

`func (o *SecurityAccessReviewAccessItem) SetSubAccessTypes(v []SecurityAccessReviewSubAccessItemType)`

SetSubAccessTypes sets SubAccessTypes field to given value.

### HasSubAccessTypes

`func (o *SecurityAccessReviewAccessItem) HasSubAccessTypes() bool`

HasSubAccessTypes returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityAccessReviewAccessItem) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReviewAccessItem) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReviewAccessItem) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReviewAccessItem) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


