# RequestApprovalPending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of a pending approval | 
**ApprovalId** | **string** | A unique identifier of the approval in the request | 

## Methods

### NewRequestApprovalPending

`func NewRequestApprovalPending(status string, approvalId string, ) *RequestApprovalPending`

NewRequestApprovalPending instantiates a new RequestApprovalPending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestApprovalPendingWithDefaults

`func NewRequestApprovalPendingWithDefaults() *RequestApprovalPending`

NewRequestApprovalPendingWithDefaults instantiates a new RequestApprovalPending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestApprovalPending) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestApprovalPending) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestApprovalPending) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetApprovalId

`func (o *RequestApprovalPending) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *RequestApprovalPending) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *RequestApprovalPending) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


