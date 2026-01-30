# SecurityAccessReviewSubAccessItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the sub-access item | 
**Name** | **string** | The name of the sub-access item | 
**Type** | [**SecurityAccessReviewSubAccessItemType**](SecurityAccessReviewSubAccessItemType.md) |  | 
**ResourceId** | **string** | The ID of the resource for the sub-access item | 
**Severity** | [**SecurityAccessReviewAccessItemSeverity**](SecurityAccessReviewAccessItemSeverity.md) |  | 
**SupportedActions** | [**[]SecurityAccessReviewAccessItemSupportedAction**](SecurityAccessReviewAccessItemSupportedAction.md) |  | 
**RemediationStatus** | Pointer to [**SecurityAccessReviewAccessItemRemediationStatus**](SecurityAccessReviewAccessItemRemediationStatus.md) |  | [optional] 
**ManualRemediationTypes** | Pointer to [**[]SecurityAccessReviewAccessItemManualRemediationType**](SecurityAccessReviewAccessItemManualRemediationType.md) | The reasons for manual remediation | [optional] 
**EntitlementInfo** | Pointer to [**SecurityAccessReviewSubAccessItemEntitlementInfo**](SecurityAccessReviewSubAccessItemEntitlementInfo.md) |  | [optional] 
**GroupInfo** | Pointer to [**SecurityAccessReviewSubAccessItemGroupInfo**](SecurityAccessReviewSubAccessItemGroupInfo.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewSubAccessItem

`func NewSecurityAccessReviewSubAccessItem(id string, name string, type_ SecurityAccessReviewSubAccessItemType, resourceId string, severity SecurityAccessReviewAccessItemSeverity, supportedActions []SecurityAccessReviewAccessItemSupportedAction, ) *SecurityAccessReviewSubAccessItem`

NewSecurityAccessReviewSubAccessItem instantiates a new SecurityAccessReviewSubAccessItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSubAccessItemWithDefaults

`func NewSecurityAccessReviewSubAccessItemWithDefaults() *SecurityAccessReviewSubAccessItem`

NewSecurityAccessReviewSubAccessItemWithDefaults instantiates a new SecurityAccessReviewSubAccessItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewSubAccessItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewSubAccessItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewSubAccessItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityAccessReviewSubAccessItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityAccessReviewSubAccessItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityAccessReviewSubAccessItem) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SecurityAccessReviewSubAccessItem) GetType() SecurityAccessReviewSubAccessItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewSubAccessItem) GetTypeOk() (*SecurityAccessReviewSubAccessItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewSubAccessItem) SetType(v SecurityAccessReviewSubAccessItemType)`

SetType sets Type field to given value.


### GetResourceId

`func (o *SecurityAccessReviewSubAccessItem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *SecurityAccessReviewSubAccessItem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *SecurityAccessReviewSubAccessItem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetSeverity

`func (o *SecurityAccessReviewSubAccessItem) GetSeverity() SecurityAccessReviewAccessItemSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityAccessReviewSubAccessItem) GetSeverityOk() (*SecurityAccessReviewAccessItemSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityAccessReviewSubAccessItem) SetSeverity(v SecurityAccessReviewAccessItemSeverity)`

SetSeverity sets Severity field to given value.


### GetSupportedActions

`func (o *SecurityAccessReviewSubAccessItem) GetSupportedActions() []SecurityAccessReviewAccessItemSupportedAction`

GetSupportedActions returns the SupportedActions field if non-nil, zero value otherwise.

### GetSupportedActionsOk

`func (o *SecurityAccessReviewSubAccessItem) GetSupportedActionsOk() (*[]SecurityAccessReviewAccessItemSupportedAction, bool)`

GetSupportedActionsOk returns a tuple with the SupportedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedActions

`func (o *SecurityAccessReviewSubAccessItem) SetSupportedActions(v []SecurityAccessReviewAccessItemSupportedAction)`

SetSupportedActions sets SupportedActions field to given value.


### GetRemediationStatus

`func (o *SecurityAccessReviewSubAccessItem) GetRemediationStatus() SecurityAccessReviewAccessItemRemediationStatus`

GetRemediationStatus returns the RemediationStatus field if non-nil, zero value otherwise.

### GetRemediationStatusOk

`func (o *SecurityAccessReviewSubAccessItem) GetRemediationStatusOk() (*SecurityAccessReviewAccessItemRemediationStatus, bool)`

GetRemediationStatusOk returns a tuple with the RemediationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationStatus

`func (o *SecurityAccessReviewSubAccessItem) SetRemediationStatus(v SecurityAccessReviewAccessItemRemediationStatus)`

SetRemediationStatus sets RemediationStatus field to given value.

### HasRemediationStatus

`func (o *SecurityAccessReviewSubAccessItem) HasRemediationStatus() bool`

HasRemediationStatus returns a boolean if a field has been set.

### GetManualRemediationTypes

`func (o *SecurityAccessReviewSubAccessItem) GetManualRemediationTypes() []SecurityAccessReviewAccessItemManualRemediationType`

GetManualRemediationTypes returns the ManualRemediationTypes field if non-nil, zero value otherwise.

### GetManualRemediationTypesOk

`func (o *SecurityAccessReviewSubAccessItem) GetManualRemediationTypesOk() (*[]SecurityAccessReviewAccessItemManualRemediationType, bool)`

GetManualRemediationTypesOk returns a tuple with the ManualRemediationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRemediationTypes

`func (o *SecurityAccessReviewSubAccessItem) SetManualRemediationTypes(v []SecurityAccessReviewAccessItemManualRemediationType)`

SetManualRemediationTypes sets ManualRemediationTypes field to given value.

### HasManualRemediationTypes

`func (o *SecurityAccessReviewSubAccessItem) HasManualRemediationTypes() bool`

HasManualRemediationTypes returns a boolean if a field has been set.

### GetEntitlementInfo

`func (o *SecurityAccessReviewSubAccessItem) GetEntitlementInfo() SecurityAccessReviewSubAccessItemEntitlementInfo`

GetEntitlementInfo returns the EntitlementInfo field if non-nil, zero value otherwise.

### GetEntitlementInfoOk

`func (o *SecurityAccessReviewSubAccessItem) GetEntitlementInfoOk() (*SecurityAccessReviewSubAccessItemEntitlementInfo, bool)`

GetEntitlementInfoOk returns a tuple with the EntitlementInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementInfo

`func (o *SecurityAccessReviewSubAccessItem) SetEntitlementInfo(v SecurityAccessReviewSubAccessItemEntitlementInfo)`

SetEntitlementInfo sets EntitlementInfo field to given value.

### HasEntitlementInfo

`func (o *SecurityAccessReviewSubAccessItem) HasEntitlementInfo() bool`

HasEntitlementInfo returns a boolean if a field has been set.

### GetGroupInfo

`func (o *SecurityAccessReviewSubAccessItem) GetGroupInfo() SecurityAccessReviewSubAccessItemGroupInfo`

GetGroupInfo returns the GroupInfo field if non-nil, zero value otherwise.

### GetGroupInfoOk

`func (o *SecurityAccessReviewSubAccessItem) GetGroupInfoOk() (*SecurityAccessReviewSubAccessItemGroupInfo, bool)`

GetGroupInfoOk returns a tuple with the GroupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInfo

`func (o *SecurityAccessReviewSubAccessItem) SetGroupInfo(v SecurityAccessReviewSubAccessItemGroupInfo)`

SetGroupInfo sets GroupInfo field to given value.

### HasGroupInfo

`func (o *SecurityAccessReviewSubAccessItem) HasGroupInfo() bool`

HasGroupInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


