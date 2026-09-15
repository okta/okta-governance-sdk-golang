# ReviewerAiAgentConnectedResourceServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | Pointer to **string** | Okta Resource Name (ORN) of the connected service account | [optional] 
**Name** | Pointer to **string** | Name of the connected service account | [optional] 
**App** | Pointer to [**ReviewerAiAgentConnectedResourceServiceAccountApp**](ReviewerAiAgentConnectedResourceServiceAccountApp.md) |  | [optional] 

## Methods

### NewReviewerAiAgentConnectedResourceServiceAccount

`func NewReviewerAiAgentConnectedResourceServiceAccount() *ReviewerAiAgentConnectedResourceServiceAccount`

NewReviewerAiAgentConnectedResourceServiceAccount instantiates a new ReviewerAiAgentConnectedResourceServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerAiAgentConnectedResourceServiceAccountWithDefaults

`func NewReviewerAiAgentConnectedResourceServiceAccountWithDefaults() *ReviewerAiAgentConnectedResourceServiceAccount`

NewReviewerAiAgentConnectedResourceServiceAccountWithDefaults instantiates a new ReviewerAiAgentConnectedResourceServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetName

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApp

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetApp() ReviewerAiAgentConnectedResourceServiceAccountApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetAppOk() (*ReviewerAiAgentConnectedResourceServiceAccountApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) SetApp(v ReviewerAiAgentConnectedResourceServiceAccountApp)`

SetApp sets App field to given value.

### HasApp

`func (o *ReviewerAiAgentConnectedResourceServiceAccount) HasApp() bool`

HasApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


