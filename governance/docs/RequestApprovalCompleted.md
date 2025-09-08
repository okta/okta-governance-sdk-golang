# RequestApprovalCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of a completed approval | 
**ApproverName** | **string** | Okta username of the user which approved the request | 
**ApprovalId** | **string** | A unique identifier of the approval in the request | 
**ApproverId** | **string** | Okta User.id of the user which approved the request | 
**Decision** | **string** |  | 
**Decided** | **time.Time** | The date the approver made their decision. | 
**OriginalDeciderId** | Pointer to **string** | The Okta user &#x60;id&#x60; | [optional] 
**OriginalDeciderFullName** | Pointer to **string** | Full name of the original decider | [optional] 
**OriginalDeciderEmail** | Pointer to **string** | Email of the original decider | [optional] 
**DeciderDelegated** | Pointer to **bool** | Indicates if the decision was made by a delegated decider | [optional] 
**FieldValues** | [**[]FieldValue**](FieldValue.md) | Values to field prompts provided by the approver at the time of approval.  All approval fields specified in the related request type are represented in the same order as defined in the request type. | 

## Methods

### NewRequestApprovalCompleted

`func NewRequestApprovalCompleted(status string, approverName string, approvalId string, approverId string, decision string, decided time.Time, fieldValues []FieldValue, ) *RequestApprovalCompleted`

NewRequestApprovalCompleted instantiates a new RequestApprovalCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestApprovalCompletedWithDefaults

`func NewRequestApprovalCompletedWithDefaults() *RequestApprovalCompleted`

NewRequestApprovalCompletedWithDefaults instantiates a new RequestApprovalCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestApprovalCompleted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestApprovalCompleted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestApprovalCompleted) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetApproverName

`func (o *RequestApprovalCompleted) GetApproverName() string`

GetApproverName returns the ApproverName field if non-nil, zero value otherwise.

### GetApproverNameOk

`func (o *RequestApprovalCompleted) GetApproverNameOk() (*string, bool)`

GetApproverNameOk returns a tuple with the ApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverName

`func (o *RequestApprovalCompleted) SetApproverName(v string)`

SetApproverName sets ApproverName field to given value.


### GetApprovalId

`func (o *RequestApprovalCompleted) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *RequestApprovalCompleted) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *RequestApprovalCompleted) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.


### GetApproverId

`func (o *RequestApprovalCompleted) GetApproverId() string`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *RequestApprovalCompleted) GetApproverIdOk() (*string, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *RequestApprovalCompleted) SetApproverId(v string)`

SetApproverId sets ApproverId field to given value.


### GetDecision

`func (o *RequestApprovalCompleted) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *RequestApprovalCompleted) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *RequestApprovalCompleted) SetDecision(v string)`

SetDecision sets Decision field to given value.


### GetDecided

`func (o *RequestApprovalCompleted) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *RequestApprovalCompleted) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *RequestApprovalCompleted) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.


### GetOriginalDeciderId

`func (o *RequestApprovalCompleted) GetOriginalDeciderId() string`

GetOriginalDeciderId returns the OriginalDeciderId field if non-nil, zero value otherwise.

### GetOriginalDeciderIdOk

`func (o *RequestApprovalCompleted) GetOriginalDeciderIdOk() (*string, bool)`

GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderId

`func (o *RequestApprovalCompleted) SetOriginalDeciderId(v string)`

SetOriginalDeciderId sets OriginalDeciderId field to given value.

### HasOriginalDeciderId

`func (o *RequestApprovalCompleted) HasOriginalDeciderId() bool`

HasOriginalDeciderId returns a boolean if a field has been set.

### GetOriginalDeciderFullName

`func (o *RequestApprovalCompleted) GetOriginalDeciderFullName() string`

GetOriginalDeciderFullName returns the OriginalDeciderFullName field if non-nil, zero value otherwise.

### GetOriginalDeciderFullNameOk

`func (o *RequestApprovalCompleted) GetOriginalDeciderFullNameOk() (*string, bool)`

GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderFullName

`func (o *RequestApprovalCompleted) SetOriginalDeciderFullName(v string)`

SetOriginalDeciderFullName sets OriginalDeciderFullName field to given value.

### HasOriginalDeciderFullName

`func (o *RequestApprovalCompleted) HasOriginalDeciderFullName() bool`

HasOriginalDeciderFullName returns a boolean if a field has been set.

### GetOriginalDeciderEmail

`func (o *RequestApprovalCompleted) GetOriginalDeciderEmail() string`

GetOriginalDeciderEmail returns the OriginalDeciderEmail field if non-nil, zero value otherwise.

### GetOriginalDeciderEmailOk

`func (o *RequestApprovalCompleted) GetOriginalDeciderEmailOk() (*string, bool)`

GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderEmail

`func (o *RequestApprovalCompleted) SetOriginalDeciderEmail(v string)`

SetOriginalDeciderEmail sets OriginalDeciderEmail field to given value.

### HasOriginalDeciderEmail

`func (o *RequestApprovalCompleted) HasOriginalDeciderEmail() bool`

HasOriginalDeciderEmail returns a boolean if a field has been set.

### GetDeciderDelegated

`func (o *RequestApprovalCompleted) GetDeciderDelegated() bool`

GetDeciderDelegated returns the DeciderDelegated field if non-nil, zero value otherwise.

### GetDeciderDelegatedOk

`func (o *RequestApprovalCompleted) GetDeciderDelegatedOk() (*bool, bool)`

GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderDelegated

`func (o *RequestApprovalCompleted) SetDeciderDelegated(v bool)`

SetDeciderDelegated sets DeciderDelegated field to given value.

### HasDeciderDelegated

`func (o *RequestApprovalCompleted) HasDeciderDelegated() bool`

HasDeciderDelegated returns a boolean if a field has been set.

### GetFieldValues

`func (o *RequestApprovalCompleted) GetFieldValues() []FieldValue`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *RequestApprovalCompleted) GetFieldValuesOk() (*[]FieldValue, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *RequestApprovalCompleted) SetFieldValues(v []FieldValue)`

SetFieldValues sets FieldValues field to given value.


### SetFieldValuesNil

`func (o *RequestApprovalCompleted) SetFieldValuesNil(b bool)`

 SetFieldValuesNil sets the value for FieldValues to be an explicit nil

### UnsetFieldValues
`func (o *RequestApprovalCompleted) UnsetFieldValues()`

UnsetFieldValues ensures that no value is present for FieldValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


