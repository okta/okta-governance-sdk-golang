# RequestActionPending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of a pending action | 
**Action** | [**RequestActionEnum**](RequestActionEnum.md) |  | 
**ActionId** | **string** | A unique identifier of the action taken in the request | 

## Methods

### NewRequestActionPending

`func NewRequestActionPending(status string, action RequestActionEnum, actionId string, ) *RequestActionPending`

NewRequestActionPending instantiates a new RequestActionPending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestActionPendingWithDefaults

`func NewRequestActionPendingWithDefaults() *RequestActionPending`

NewRequestActionPendingWithDefaults instantiates a new RequestActionPending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestActionPending) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestActionPending) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestActionPending) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAction

`func (o *RequestActionPending) GetAction() RequestActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestActionPending) GetActionOk() (*RequestActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestActionPending) SetAction(v RequestActionEnum)`

SetAction sets Action field to given value.


### GetActionId

`func (o *RequestActionPending) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *RequestActionPending) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *RequestActionPending) SetActionId(v string)`

SetActionId sets ActionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


