# ReviewerAiAgentConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | AI agent connection ID | 
**Name** | Pointer to **string** | Name of the resource associated with this connection | [optional] 
**ConnectedResource** | Pointer to [**ReviewerAiAgentConnectedResource**](ReviewerAiAgentConnectedResource.md) |  | [optional] 

## Methods

### NewReviewerAiAgentConnection

`func NewReviewerAiAgentConnection(id string, ) *ReviewerAiAgentConnection`

NewReviewerAiAgentConnection instantiates a new ReviewerAiAgentConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerAiAgentConnectionWithDefaults

`func NewReviewerAiAgentConnectionWithDefaults() *ReviewerAiAgentConnection`

NewReviewerAiAgentConnectionWithDefaults instantiates a new ReviewerAiAgentConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewerAiAgentConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewerAiAgentConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewerAiAgentConnection) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ReviewerAiAgentConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewerAiAgentConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewerAiAgentConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReviewerAiAgentConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnectedResource

`func (o *ReviewerAiAgentConnection) GetConnectedResource() ReviewerAiAgentConnectedResource`

GetConnectedResource returns the ConnectedResource field if non-nil, zero value otherwise.

### GetConnectedResourceOk

`func (o *ReviewerAiAgentConnection) GetConnectedResourceOk() (*ReviewerAiAgentConnectedResource, bool)`

GetConnectedResourceOk returns a tuple with the ConnectedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedResource

`func (o *ReviewerAiAgentConnection) SetConnectedResource(v ReviewerAiAgentConnectedResource)`

SetConnectedResource sets ConnectedResource field to given value.

### HasConnectedResource

`func (o *ReviewerAiAgentConnection) HasConnectedResource() bool`

HasConnectedResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


