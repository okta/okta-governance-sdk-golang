# RequestApproval2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ApprovalProviderType**](ApprovalProviderType.md) |  | 
**ProviderName** | **string** | A human readable name of the request approval system, for example, Okta Access Requests or ServiceNow | 
**ProviderDescription** | Pointer to **string** | A description of the request approval system | [optional] 
**ExternalRequestId** | Pointer to **string** | The external request &#x60;id&#x60; from a request approval system, for example, ServiceNow or JIRA | [optional] 
**Status** | Pointer to [**RequestApprovalStatus**](RequestApprovalStatus.md) |  | [optional] 
**Decided** | Pointer to **time.Time** | The date the approval decision is made. | [optional] 
**Decisions** | Pointer to [**[]RequestApprovalDecision**](RequestApprovalDecision.md) | The approval decisions | [optional] 

## Methods

### NewRequestApproval2

`func NewRequestApproval2(type_ ApprovalProviderType, providerName string, ) *RequestApproval2`

NewRequestApproval2 instantiates a new RequestApproval2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestApproval2WithDefaults

`func NewRequestApproval2WithDefaults() *RequestApproval2`

NewRequestApproval2WithDefaults instantiates a new RequestApproval2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestApproval2) GetType() ApprovalProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestApproval2) GetTypeOk() (*ApprovalProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestApproval2) SetType(v ApprovalProviderType)`

SetType sets Type field to given value.


### GetProviderName

`func (o *RequestApproval2) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *RequestApproval2) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *RequestApproval2) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetProviderDescription

`func (o *RequestApproval2) GetProviderDescription() string`

GetProviderDescription returns the ProviderDescription field if non-nil, zero value otherwise.

### GetProviderDescriptionOk

`func (o *RequestApproval2) GetProviderDescriptionOk() (*string, bool)`

GetProviderDescriptionOk returns a tuple with the ProviderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDescription

`func (o *RequestApproval2) SetProviderDescription(v string)`

SetProviderDescription sets ProviderDescription field to given value.

### HasProviderDescription

`func (o *RequestApproval2) HasProviderDescription() bool`

HasProviderDescription returns a boolean if a field has been set.

### GetExternalRequestId

`func (o *RequestApproval2) GetExternalRequestId() string`

GetExternalRequestId returns the ExternalRequestId field if non-nil, zero value otherwise.

### GetExternalRequestIdOk

`func (o *RequestApproval2) GetExternalRequestIdOk() (*string, bool)`

GetExternalRequestIdOk returns a tuple with the ExternalRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRequestId

`func (o *RequestApproval2) SetExternalRequestId(v string)`

SetExternalRequestId sets ExternalRequestId field to given value.

### HasExternalRequestId

`func (o *RequestApproval2) HasExternalRequestId() bool`

HasExternalRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *RequestApproval2) GetStatus() RequestApprovalStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestApproval2) GetStatusOk() (*RequestApprovalStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestApproval2) SetStatus(v RequestApprovalStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestApproval2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDecided

`func (o *RequestApproval2) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *RequestApproval2) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *RequestApproval2) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.

### HasDecided

`func (o *RequestApproval2) HasDecided() bool`

HasDecided returns a boolean if a field has been set.

### GetDecisions

`func (o *RequestApproval2) GetDecisions() []RequestApprovalDecision`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *RequestApproval2) GetDecisionsOk() (*[]RequestApprovalDecision, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *RequestApproval2) SetDecisions(v []RequestApprovalDecision)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *RequestApproval2) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


