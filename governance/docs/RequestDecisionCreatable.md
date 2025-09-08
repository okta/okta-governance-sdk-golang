# RequestDecisionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The Okta user &#x60;id&#x60; | 
**UserEmail** | **string** | E-mail of the user. | 
**UserName** | **string** | Name of the user. | 
**Decision** | [**ApprovalDecisionEnum**](ApprovalDecisionEnum.md) |  | 
**Decided** | **time.Time** | The date the approval decision is made. | 
**OriginalDeciderId** | Pointer to **string** | The Okta user &#x60;id&#x60; | [optional] 
**OriginalDeciderFullName** | Pointer to **string** | Name of the user. | [optional] 
**OriginalDeciderEmail** | Pointer to **string** | E-mail of the user. | [optional] 
**DeciderDelegated** | Pointer to **bool** | Indicates if the decision was made by a delegated decider | [optional] 

## Methods

### NewRequestDecisionCreatable

`func NewRequestDecisionCreatable(userId string, userEmail string, userName string, decision ApprovalDecisionEnum, decided time.Time, ) *RequestDecisionCreatable`

NewRequestDecisionCreatable instantiates a new RequestDecisionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDecisionCreatableWithDefaults

`func NewRequestDecisionCreatableWithDefaults() *RequestDecisionCreatable`

NewRequestDecisionCreatableWithDefaults instantiates a new RequestDecisionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RequestDecisionCreatable) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestDecisionCreatable) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestDecisionCreatable) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserEmail

`func (o *RequestDecisionCreatable) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *RequestDecisionCreatable) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *RequestDecisionCreatable) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserName

`func (o *RequestDecisionCreatable) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RequestDecisionCreatable) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RequestDecisionCreatable) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetDecision

`func (o *RequestDecisionCreatable) GetDecision() ApprovalDecisionEnum`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *RequestDecisionCreatable) GetDecisionOk() (*ApprovalDecisionEnum, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *RequestDecisionCreatable) SetDecision(v ApprovalDecisionEnum)`

SetDecision sets Decision field to given value.


### GetDecided

`func (o *RequestDecisionCreatable) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *RequestDecisionCreatable) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *RequestDecisionCreatable) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.


### GetOriginalDeciderId

`func (o *RequestDecisionCreatable) GetOriginalDeciderId() string`

GetOriginalDeciderId returns the OriginalDeciderId field if non-nil, zero value otherwise.

### GetOriginalDeciderIdOk

`func (o *RequestDecisionCreatable) GetOriginalDeciderIdOk() (*string, bool)`

GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderId

`func (o *RequestDecisionCreatable) SetOriginalDeciderId(v string)`

SetOriginalDeciderId sets OriginalDeciderId field to given value.

### HasOriginalDeciderId

`func (o *RequestDecisionCreatable) HasOriginalDeciderId() bool`

HasOriginalDeciderId returns a boolean if a field has been set.

### GetOriginalDeciderFullName

`func (o *RequestDecisionCreatable) GetOriginalDeciderFullName() string`

GetOriginalDeciderFullName returns the OriginalDeciderFullName field if non-nil, zero value otherwise.

### GetOriginalDeciderFullNameOk

`func (o *RequestDecisionCreatable) GetOriginalDeciderFullNameOk() (*string, bool)`

GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderFullName

`func (o *RequestDecisionCreatable) SetOriginalDeciderFullName(v string)`

SetOriginalDeciderFullName sets OriginalDeciderFullName field to given value.

### HasOriginalDeciderFullName

`func (o *RequestDecisionCreatable) HasOriginalDeciderFullName() bool`

HasOriginalDeciderFullName returns a boolean if a field has been set.

### GetOriginalDeciderEmail

`func (o *RequestDecisionCreatable) GetOriginalDeciderEmail() string`

GetOriginalDeciderEmail returns the OriginalDeciderEmail field if non-nil, zero value otherwise.

### GetOriginalDeciderEmailOk

`func (o *RequestDecisionCreatable) GetOriginalDeciderEmailOk() (*string, bool)`

GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeciderEmail

`func (o *RequestDecisionCreatable) SetOriginalDeciderEmail(v string)`

SetOriginalDeciderEmail sets OriginalDeciderEmail field to given value.

### HasOriginalDeciderEmail

`func (o *RequestDecisionCreatable) HasOriginalDeciderEmail() bool`

HasOriginalDeciderEmail returns a boolean if a field has been set.

### GetDeciderDelegated

`func (o *RequestDecisionCreatable) GetDeciderDelegated() bool`

GetDeciderDelegated returns the DeciderDelegated field if non-nil, zero value otherwise.

### GetDeciderDelegatedOk

`func (o *RequestDecisionCreatable) GetDeciderDelegatedOk() (*bool, bool)`

GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeciderDelegated

`func (o *RequestDecisionCreatable) SetDeciderDelegated(v bool)`

SetDeciderDelegated sets DeciderDelegated field to given value.

### HasDeciderDelegated

`func (o *RequestDecisionCreatable) HasDeciderDelegated() bool`

HasDeciderDelegated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


