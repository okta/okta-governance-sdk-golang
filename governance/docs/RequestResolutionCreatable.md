# RequestResolutionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalState** | Pointer to [**RequestResolutionTerminalState**](RequestResolutionTerminalState.md) |  | [optional] 
**Decisions** | Pointer to [**[]RequestDecisionCreatable**](RequestDecisionCreatable.md) |  | [optional] 

## Methods

### NewRequestResolutionCreatable

`func NewRequestResolutionCreatable() *RequestResolutionCreatable`

NewRequestResolutionCreatable instantiates a new RequestResolutionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestResolutionCreatableWithDefaults

`func NewRequestResolutionCreatableWithDefaults() *RequestResolutionCreatable`

NewRequestResolutionCreatableWithDefaults instantiates a new RequestResolutionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminalState

`func (o *RequestResolutionCreatable) GetTerminalState() RequestResolutionTerminalState`

GetTerminalState returns the TerminalState field if non-nil, zero value otherwise.

### GetTerminalStateOk

`func (o *RequestResolutionCreatable) GetTerminalStateOk() (*RequestResolutionTerminalState, bool)`

GetTerminalStateOk returns a tuple with the TerminalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalState

`func (o *RequestResolutionCreatable) SetTerminalState(v RequestResolutionTerminalState)`

SetTerminalState sets TerminalState field to given value.

### HasTerminalState

`func (o *RequestResolutionCreatable) HasTerminalState() bool`

HasTerminalState returns a boolean if a field has been set.

### GetDecisions

`func (o *RequestResolutionCreatable) GetDecisions() []RequestDecisionCreatable`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *RequestResolutionCreatable) GetDecisionsOk() (*[]RequestDecisionCreatable, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *RequestResolutionCreatable) SetDecisions(v []RequestDecisionCreatable)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *RequestResolutionCreatable) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


