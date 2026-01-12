# RequestApprovalDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeciderId** | **string** | The Okta user &#x60;id&#x60; | 
**DeciderName** | Pointer to **string** | Okta username of the decider | [optional] 
**Decision** | [**ApprovalDecisionEnum**](ApprovalDecisionEnum.md) |  | 
**Decided** | **time.Time** | The date the approval decision is made. | 
**OriginalDeciderId** | Pointer to **string** | The Okta user &#x60;id&#x60; | [optional] 
**OriginalDeciderFullName** | Pointer to **string** | Full name of the original decider | [optional] 
**OriginalDeciderEmail** | Pointer to **string** | Email of the original decider | [optional] 
**DeciderDelegated** | Pointer to **bool** | Indicates if the decision was made by a delegated decider | [optional] 
**DeciderEscalated** | Pointer to **bool** | Indicates if the decision was made by an escalated decider | [optional] 

## Methods

### NewRequestApprovalDecision

`func NewRequestApprovalDecision(deciderId string, decision ApprovalDecisionEnum, decided time.Time, ) *RequestApprovalDecision`

NewRequestApprovalDecision instantiates a new RequestApprovalDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestApprovalDecisionWithDefaults

`func NewRequestApprovalDecisionWithDefaults() *RequestApprovalDecision`

NewRequestApprovalDecisionWithDefaults instantiates a new RequestApprovalDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeciderId

`func (o *RequestApprovalDecision) GetDeciderId() string`

GetDeciderId returns the DeciderId field if non-nil, zero value otherwise.

### GetDeciderIdOk

`func (o *RequestApprovalDecision) GetDeciderIdOk() (*string, bool)`

GetDeciderIdOk returns a tuple with the DeciderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderId

`func (o *RequestApprovalDecision) SetDeciderId(v string)`

SetDeciderId sets DeciderId field to given value.


### GetDeciderName

`func (o *RequestApprovalDecision) GetDeciderName() string`

GetDeciderName returns the DeciderName field if non-nil, zero value otherwise.

### GetDeciderNameOk

`func (o *RequestApprovalDecision) GetDeciderNameOk() (*string, bool)`

GetDeciderNameOk returns a tuple with the DeciderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderName

`func (o *RequestApprovalDecision) SetDeciderName(v string)`

SetDeciderName sets DeciderName field to given value.

### HasDeciderName

`func (o *RequestApprovalDecision) HasDeciderName() bool`

HasDeciderName returns a boolean if a field has been set.

### GetDecision

`func (o *RequestApprovalDecision) GetDecision() ApprovalDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *RequestApprovalDecision) GetDecisionOk() (*ApprovalDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *RequestApprovalDecision) SetDecision(v ApprovalDecisionEnum)`

SetDecision sets Decision field to given value.


### GetDecided

`func (o *RequestApprovalDecision) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *RequestApprovalDecision) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *RequestApprovalDecision) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.


### GetOriginalDeciderId

`func (o *RequestApprovalDecision) GetOriginalDeciderId() string`

GetOriginalDeciderId returns the OriginalDeciderId field if non-nil, zero value otherwise.

### GetOriginalDeciderIdOk

`func (o *RequestApprovalDecision) GetOriginalDeciderIdOk() (*string, bool)`

GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderId

`func (o *RequestApprovalDecision) SetOriginalDeciderId(v string)`

SetOriginalDeciderId sets OriginalDeciderId field to given value.

### HasOriginalDeciderId

`func (o *RequestApprovalDecision) HasOriginalDeciderId() bool`

HasOriginalDeciderId returns a boolean if a field has been set.

### GetOriginalDeciderFullName

`func (o *RequestApprovalDecision) GetOriginalDeciderFullName() string`

GetOriginalDeciderFullName returns the OriginalDeciderFullName field if non-nil, zero value otherwise.

### GetOriginalDeciderFullNameOk

`func (o *RequestApprovalDecision) GetOriginalDeciderFullNameOk() (*string, bool)`

GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderFullName

`func (o *RequestApprovalDecision) SetOriginalDeciderFullName(v string)`

SetOriginalDeciderFullName sets OriginalDeciderFullName field to given value.

### HasOriginalDeciderFullName

`func (o *RequestApprovalDecision) HasOriginalDeciderFullName() bool`

HasOriginalDeciderFullName returns a boolean if a field has been set.

### GetOriginalDeciderEmail

`func (o *RequestApprovalDecision) GetOriginalDeciderEmail() string`

GetOriginalDeciderEmail returns the OriginalDeciderEmail field if non-nil, zero value otherwise.

### GetOriginalDeciderEmailOk

`func (o *RequestApprovalDecision) GetOriginalDeciderEmailOk() (*string, bool)`

GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderEmail

`func (o *RequestApprovalDecision) SetOriginalDeciderEmail(v string)`

SetOriginalDeciderEmail sets OriginalDeciderEmail field to given value.

### HasOriginalDeciderEmail

`func (o *RequestApprovalDecision) HasOriginalDeciderEmail() bool`

HasOriginalDeciderEmail returns a boolean if a field has been set.

### GetDeciderDelegated

`func (o *RequestApprovalDecision) GetDeciderDelegated() bool`

GetDeciderDelegated returns the DeciderDelegated field if non-nil, zero value otherwise.

### GetDeciderDelegatedOk

`func (o *RequestApprovalDecision) GetDeciderDelegatedOk() (*bool, bool)`

GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderDelegated

`func (o *RequestApprovalDecision) SetDeciderDelegated(v bool)`

SetDeciderDelegated sets DeciderDelegated field to given value.

### HasDeciderDelegated

`func (o *RequestApprovalDecision) HasDeciderDelegated() bool`

HasDeciderDelegated returns a boolean if a field has been set.

### GetDeciderEscalated

`func (o *RequestApprovalDecision) GetDeciderEscalated() bool`

GetDeciderEscalated returns the DeciderEscalated field if non-nil, zero value otherwise.

### GetDeciderEscalatedOk

`func (o *RequestApprovalDecision) GetDeciderEscalatedOk() (*bool, bool)`

GetDeciderEscalatedOk returns a tuple with the DeciderEscalated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderEscalated

`func (o *RequestApprovalDecision) SetDeciderEscalated(v bool)`

SetDeciderEscalated sets DeciderEscalated field to given value.

### HasDeciderEscalated

`func (o *RequestApprovalDecision) HasDeciderEscalated() bool`

HasDeciderEscalated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


