# ReviewerAiAgentConnectedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationServer** | Pointer to [**ReviewerAiAgentConnectedResourceAuthorizationServer**](ReviewerAiAgentConnectedResourceAuthorizationServer.md) |  | [optional] 
**App** | Pointer to [**ReviewerAiAgentConnectedResourceApp**](ReviewerAiAgentConnectedResourceApp.md) |  | [optional] 
**McpServer** | Pointer to [**ReviewerAiAgentConnectedResourceMcpServer**](ReviewerAiAgentConnectedResourceMcpServer.md) |  | [optional] 
**ApiServer** | Pointer to [**ReviewerAiAgentConnectedResourceApiServer**](ReviewerAiAgentConnectedResourceApiServer.md) |  | [optional] 
**ServiceAccount** | Pointer to [**ReviewerAiAgentConnectedResourceServiceAccount**](ReviewerAiAgentConnectedResourceServiceAccount.md) |  | [optional] 
**OpaSecret** | Pointer to [**ReviewerAiAgentConnectedResourceSecret**](ReviewerAiAgentConnectedResourceSecret.md) |  | [optional] 
**Type** | [**AiAgentConnectedResourceType**](AiAgentConnectedResourceType.md) |  | 

## Methods

### NewReviewerAiAgentConnectedResource

`func NewReviewerAiAgentConnectedResource(type_ AiAgentConnectedResourceType, ) *ReviewerAiAgentConnectedResource`

NewReviewerAiAgentConnectedResource instantiates a new ReviewerAiAgentConnectedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerAiAgentConnectedResourceWithDefaults

`func NewReviewerAiAgentConnectedResourceWithDefaults() *ReviewerAiAgentConnectedResource`

NewReviewerAiAgentConnectedResourceWithDefaults instantiates a new ReviewerAiAgentConnectedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationServer

`func (o *ReviewerAiAgentConnectedResource) GetAuthorizationServer() ReviewerAiAgentConnectedResourceAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *ReviewerAiAgentConnectedResource) GetAuthorizationServerOk() (*ReviewerAiAgentConnectedResourceAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *ReviewerAiAgentConnectedResource) SetAuthorizationServer(v ReviewerAiAgentConnectedResourceAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *ReviewerAiAgentConnectedResource) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetApp

`func (o *ReviewerAiAgentConnectedResource) GetApp() ReviewerAiAgentConnectedResourceApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ReviewerAiAgentConnectedResource) GetAppOk() (*ReviewerAiAgentConnectedResourceApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ReviewerAiAgentConnectedResource) SetApp(v ReviewerAiAgentConnectedResourceApp)`

SetApp sets App field to given value.

### HasApp

`func (o *ReviewerAiAgentConnectedResource) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetMcpServer

`func (o *ReviewerAiAgentConnectedResource) GetMcpServer() ReviewerAiAgentConnectedResourceMcpServer`

GetMcpServer returns the McpServer field if non-nil, zero value otherwise.

### GetMcpServerOk

`func (o *ReviewerAiAgentConnectedResource) GetMcpServerOk() (*ReviewerAiAgentConnectedResourceMcpServer, bool)`

GetMcpServerOk returns a tuple with the McpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServer

`func (o *ReviewerAiAgentConnectedResource) SetMcpServer(v ReviewerAiAgentConnectedResourceMcpServer)`

SetMcpServer sets McpServer field to given value.

### HasMcpServer

`func (o *ReviewerAiAgentConnectedResource) HasMcpServer() bool`

HasMcpServer returns a boolean if a field has been set.

### GetApiServer

`func (o *ReviewerAiAgentConnectedResource) GetApiServer() ReviewerAiAgentConnectedResourceApiServer`

GetApiServer returns the ApiServer field if non-nil, zero value otherwise.

### GetApiServerOk

`func (o *ReviewerAiAgentConnectedResource) GetApiServerOk() (*ReviewerAiAgentConnectedResourceApiServer, bool)`

GetApiServerOk returns a tuple with the ApiServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServer

`func (o *ReviewerAiAgentConnectedResource) SetApiServer(v ReviewerAiAgentConnectedResourceApiServer)`

SetApiServer sets ApiServer field to given value.

### HasApiServer

`func (o *ReviewerAiAgentConnectedResource) HasApiServer() bool`

HasApiServer returns a boolean if a field has been set.

### GetServiceAccount

`func (o *ReviewerAiAgentConnectedResource) GetServiceAccount() ReviewerAiAgentConnectedResourceServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ReviewerAiAgentConnectedResource) GetServiceAccountOk() (*ReviewerAiAgentConnectedResourceServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ReviewerAiAgentConnectedResource) SetServiceAccount(v ReviewerAiAgentConnectedResourceServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *ReviewerAiAgentConnectedResource) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetOpaSecret

`func (o *ReviewerAiAgentConnectedResource) GetOpaSecret() ReviewerAiAgentConnectedResourceSecret`

GetOpaSecret returns the OpaSecret field if non-nil, zero value otherwise.

### GetOpaSecretOk

`func (o *ReviewerAiAgentConnectedResource) GetOpaSecretOk() (*ReviewerAiAgentConnectedResourceSecret, bool)`

GetOpaSecretOk returns a tuple with the OpaSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaSecret

`func (o *ReviewerAiAgentConnectedResource) SetOpaSecret(v ReviewerAiAgentConnectedResourceSecret)`

SetOpaSecret sets OpaSecret field to given value.

### HasOpaSecret

`func (o *ReviewerAiAgentConnectedResource) HasOpaSecret() bool`

HasOpaSecret returns a boolean if a field has been set.

### GetType

`func (o *ReviewerAiAgentConnectedResource) GetType() AiAgentConnectedResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReviewerAiAgentConnectedResource) GetTypeOk() (*AiAgentConnectedResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReviewerAiAgentConnectedResource) SetType(v AiAgentConnectedResourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


