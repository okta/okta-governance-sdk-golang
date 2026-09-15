# ReviewerAiAgentConnectedResourceAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | Pointer to **string** | Okta Resource Name (ORN) of the authorization server | [optional] 
**Name** | Pointer to **string** | Name of the authorization server | [optional] 
**IssuerUrl** | Pointer to **string** | Issuer URL of the connected authorization server | [optional] 
**ResourceIndicator** | Pointer to **string** | Resource indicator used when requesting tokens | [optional] 
**ScopeCondition** | Pointer to [**AuthorizationServerScopeCondition**](AuthorizationServerScopeCondition.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | Connected authorization server scopes that the AI agent has access to | [optional] 

## Methods

### NewReviewerAiAgentConnectedResourceAuthorizationServer

`func NewReviewerAiAgentConnectedResourceAuthorizationServer() *ReviewerAiAgentConnectedResourceAuthorizationServer`

NewReviewerAiAgentConnectedResourceAuthorizationServer instantiates a new ReviewerAiAgentConnectedResourceAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerAiAgentConnectedResourceAuthorizationServerWithDefaults

`func NewReviewerAiAgentConnectedResourceAuthorizationServerWithDefaults() *ReviewerAiAgentConnectedResourceAuthorizationServer`

NewReviewerAiAgentConnectedResourceAuthorizationServerWithDefaults instantiates a new ReviewerAiAgentConnectedResourceAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetName

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetResourceIndicator

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetResourceIndicator() string`

GetResourceIndicator returns the ResourceIndicator field if non-nil, zero value otherwise.

### GetResourceIndicatorOk

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetResourceIndicatorOk() (*string, bool)`

GetResourceIndicatorOk returns a tuple with the ResourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIndicator

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetResourceIndicator(v string)`

SetResourceIndicator sets ResourceIndicator field to given value.

### HasResourceIndicator

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasResourceIndicator() bool`

HasResourceIndicator returns a boolean if a field has been set.

### GetScopeCondition

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopeCondition() AuthorizationServerScopeCondition`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopeConditionOk() (*AuthorizationServerScopeCondition, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetScopeCondition(v AuthorizationServerScopeCondition)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


